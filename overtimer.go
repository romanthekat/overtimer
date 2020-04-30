package main

import (
	"fmt"
	"log"
)

type App struct {
	ActiveEntry     entry           `json:"active_entry,omitempty"`
	FinishedEntries []finishedEntry `json:"finished_entries"`
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

	case stop:

	case spend:

	case status:
		fmt.Println(app)
	}

	err = app.save()
	if err != nil {
		log.Fatal(err)
	}
}
