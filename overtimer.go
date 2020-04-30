package main

import "time"

type entryType string

const (
	overtime entryType = "overtime"
	spending entryType = "spending"
)

type entry struct {
	entryType entryType
	startTime time.Time
}

type finishedEntry struct {
	entry

	endTime time.Time
}

func main() {

}
