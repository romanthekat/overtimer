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

func (app *App) stop() (EntryType, error) {
	finishedEntry := app.finishActive()
	if finishedEntry != nil {
		return finishedEntry.EntryType, nil
	} else {
		return "", fmt.Errorf("no active Entry found - can't perform stop")
	}
}

func (app *App) finishActive() *FinishedEntry {
	activeEntry := app.ActiveEntry
	if activeEntry != nil {
		finishedEntry := app.addEntry(activeEntry.EntryType, activeEntry.StartTime, time.Now())
		app.ActiveEntry = nil
		return finishedEntry
	} else {
		return nil
	}
}

func (app *App) addEntry(t EntryType, startTime time.Time, endTime time.Time) *FinishedEntry {
	finishedEntry := newFinishedEntry(t, startTime, endTime)
	app.FinishedEntries = append(app.FinishedEntries, *finishedEntry)
	return finishedEntry
}

func (app *App) spend() bool {
	if app.ActiveEntry != nil && app.ActiveEntry.EntryType == spending {
		return false
	}

	app.finishActive()

	app.ActiveEntry = newEntry(spending, time.Now())
	return true
}

func (app *App) routine() (string, error) {
	return app.routineAt(time.Now())
}

func (app *App) routineAt(t time.Time) (string, error) {
	if app.ActiveEntry != nil {
		return "", fmt.Errorf("active Entry exists, routine cannot be performed automatically")
	}

	startTime := newDate(t, app.Settings.WorkStartHour)
	endTime := newDate(t, app.Settings.WorkEndHour)

	switch {
	case t.Before(startTime):
		app.addEntry(overtime, t, startTime)
		return fmt.Sprintf("overtime till %s added", startTime), nil
	case t.After(startTime) && t.Before(endTime):
		app.addEntry(spending, startTime, t)
		return fmt.Sprintf("spending from %s added", startTime), nil
	case t.After(endTime):
		app.addEntry(overtime, endTime, t)
		return fmt.Sprintf("overtime from %s added", endTime), nil
	default:
		return "nothing performed", nil
	}
}

func newDate(t time.Time, hour int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, t.Location())
}
