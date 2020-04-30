package main

import (
	"fmt"
	"time"
)

type totalType string

const (
	hasOvertime totalType = "has overtime"
	hasDebt     totalType = "has debt"
)

func (app *App) calculateTotal() (time.Duration, totalType) {
	result := time.Duration(0)

	for _, entry := range app.FinishedEntries {
		if entry.EntryType == overtime {
			result += entry.getDuration()
		} else if entry.EntryType == spending {
			result -= entry.getDuration()
		}
	}

	activeEntry := app.ActiveEntry
	if activeEntry != nil {
		activeDuration := time.Now().Sub(activeEntry.StartTime)
		if activeEntry.EntryType == overtime {
			result += activeDuration
		} else if activeEntry.EntryType == spending {
			result -= activeDuration
		}
	}

	result = result.Round(time.Second)
	if result > 0 {
		return result, hasOvertime
	} else {
		return result, hasDebt
	}
}

func (app *App) start() bool {
	if app.ActiveEntry == nil {
		app.ActiveEntry = newEntry(overtime, time.Now())
		return true
	}

	return false
}

func (app *App) stop() (entryType, error) {
	finishedEntry := app.finishActive()
	if finishedEntry != nil {
		return finishedEntry.EntryType, nil
	} else {
		return "", fmt.Errorf("no active entry found - can't perform stop")
	}
}

func (app *App) finishActive() *finishedEntry {
	activeEntry := app.ActiveEntry
	if activeEntry != nil {
		finishedEntry := newFinishedEntry(activeEntry.EntryType, activeEntry.StartTime, time.Now())
		app.FinishedEntries = append(app.FinishedEntries, *finishedEntry)
		app.ActiveEntry = nil
		return finishedEntry
	} else {
		return nil
	}
}

func (app *App) spend() bool {
	if app.ActiveEntry != nil && app.ActiveEntry.EntryType == spending {
		return false
	}

	app.finishActive()

	app.ActiveEntry = newEntry(spending, time.Now())
	return true
}
