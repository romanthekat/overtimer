package main

import (
	"fmt"
	"time"
)

type entryType string

type settings struct {
	WorkStartHour int `json:"work_start_hour"`
	WorkEndHour   int `json:"work_end_hour"`
}

const (
	overtime entryType = "overtime"
	spending entryType = "spending"
)

type entry struct {
	EntryType entryType `json:"entry_type"`
	StartTime time.Time `json:"start_time"`
}

func newEntry(entryType entryType, startTime time.Time) *entry {
	return &entry{EntryType: entryType, StartTime: startTime}
}

func (e entry) String() string {
	return fmt.Sprintf("Active %v since %v", e.EntryType, e.StartTime)
}

type finishedEntry struct {
	entry
	EndTime time.Time `json:"end_time"`
}

func newFinishedEntry(entryType entryType, startTime time.Time, endTime time.Time) *finishedEntry {
	return &finishedEntry{entry: *newEntry(entryType, startTime), EndTime: endTime}
}

func (f finishedEntry) String() string {
	return fmt.Sprintf("{%v: %v - %v}", f.EntryType, f.StartTime, f.EndTime)
}

func (f finishedEntry) getDuration() time.Duration {
	return f.EndTime.Sub(f.StartTime)
}
