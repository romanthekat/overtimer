package main

import (
	"fmt"
	"log"
	"time"
)

type App struct {
	Settings        *settings       `json:"settings"`
	ActiveEntry     *entry          `json:"active_entry,omitempty"`
	FinishedEntries []finishedEntry `json:"finished_entries"`
}

func NewApp(currentSettings *settings, activeEntry *entry, finishedEntries []finishedEntry) *App {
	return &App{Settings: currentSettings, ActiveEntry: activeEntry, FinishedEntries: finishedEntries}
}

func NewAppDefault() *App {
	return &App{Settings: &settings{
		WorkStartHour: 10,
		WorkEndHour:   19,
	}}
}

func (app *App) String() string {
	total, totalType := app.calculateTotal()
	if totalType == hasOvertime {
		return fmt.Sprintf("overtime: %v", total)
	} else {
		return fmt.Sprintf("debt: %v", total)
	}
}

func main() {
	command, err := readCommand()
	if err != nil {
		log.Fatal(err)
	}

	app, err := getApp()
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case start:
		started := app.start()
		if started {
			fmt.Println("overtime started at", nowTimeFormatted())
		} else {
			fmt.Println("overtime is already in progress")
		}
	case stop:
		entryType, err := app.stop()
		if err != nil {
			log.Fatal("error occurred during stopping: ", err)
		}
		fmt.Printf("%v stopped at %v", entryType, nowTimeFormatted())
	case spend:
		started := app.spend()
		if started {
			fmt.Println("time spending started at", nowTimeFormatted())
		} else {
			fmt.Println("spending is already in progress")
		}
	case routine:
		result, err := app.routine()
		if err != nil {
			log.Fatal("error occurred during stopping: ", err)
		}
		fmt.Printf("%s", result)
	case status:
		fmt.Println(app)
	}

	err = app.save()
	if err != nil {
		log.Fatal(err)
	}
}

func nowTimeFormatted() string {
	return time.Now().Format(time.RFC3339)
}
