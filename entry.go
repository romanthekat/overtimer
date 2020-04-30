package main

import "time"

type entryType string

const (
	overtime entryType = "overtime"
	spending entryType = "spending"
)

type entry struct {
	EntryType entryType `json:"entry_type"`
	StartTime time.Time `json:"start_time"`
}

type finishedEntry struct {
	entry
	EndTime time.Time `json:"end_time"`
}
