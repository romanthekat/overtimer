package main

import (
	"fmt"
	"time"
)

type entryType string

const (
	overtime entryType = "overtime"
	spending entryType = "spending"
)

type entry struct {
	EntryType entryType `json:"entry_type"`
	StartTime time.Time `json:"start_time"`
}

func (e entry) String() string {
	return fmt.Sprintf("Active %v since %v", e.EntryType, e.StartTime)
}

type finishedEntry struct {
	entry
	EndTime time.Time `json:"end_time"`
}

func (f finishedEntry) String() string {
	return fmt.Sprintf("{%v: %v - %v}", f.EntryType, f.StartTime, f.EndTime)
}

func (f finishedEntry) getDuration() time.Duration {
	return f.EndTime.Sub(f.StartTime)
}
