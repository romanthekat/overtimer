package main

import (
	"fmt"
	"log"
)

type App struct {
	ActiveEntry     entry           `json:"active_entry,omitempty"`
	FinishedEntries []finishedEntry `json:"finished_entries"`
}

func main() {
	app, err := getApp()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("app: %+v", app)
}
