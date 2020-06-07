package main

import (
	"fmt"
	"time"
)

type EntryType string

const (
	overtime EntryType = "overtime"
	spending EntryType = "spending"
)

type Entry struct {
	EntryType EntryType `json:"entry_type"`
	StartTime time.Time `json:"start_time"`
}

func newEntry(entryType EntryType, startTime time.Time) *Entry {
	return &Entry{EntryType: entryType, StartTime: startTime}
}

func (e Entry) String() string {
	return fmt.Sprintf("Active %v since %v", e.EntryType, e.StartTime)
}

type FinishedEntry struct {
	Entry
	EndTime time.Time `json:"end_time"`
}

func newFinishedEntry(entryType EntryType, startTime time.Time, endTime time.Time) *FinishedEntry {
	return &FinishedEntry{Entry: *newEntry(entryType, startTime), EndTime: endTime}
}

func (f FinishedEntry) String() string {
	return fmt.Sprintf("{%v: %v - %v}", f.EntryType, f.StartTime, f.EndTime)
}

func (f FinishedEntry) getDuration() time.Duration {
	return f.EndTime.Sub(f.StartTime)
}

type Settings struct {
	WorkStartHour int `json:"work_start_hour"`
	WorkEndHour   int `json:"work_end_hour"`
}
