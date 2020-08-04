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
	finishedEntry, err := app.finishActive()
	if err != nil {
		return "", fmt.Errorf("no active Entry found - can't perform stop")
	}

	return finishedEntry.EntryType, nil
}

func (app *App) finishActive() (*FinishedEntry, error) {
	activeEntry := app.ActiveEntry
	if activeEntry != nil {
		if activeEntry.EntryType == lunching { //TODO migrate this logic on type?
			return app.finishLunching(activeEntry), nil
		} else {
			return app.finishGenericEntry(activeEntry), nil
		}
	} else {
		return nil, fmt.Errorf("no active entry found")
	}
}

func (app *App) finishGenericEntry(activeEntry *Entry) *FinishedEntry {
	finishedEntry := app.addEntry(activeEntry.EntryType, activeEntry.StartTime, time.Now())
	app.ActiveEntry = nil
	return finishedEntry
}

func (app *App) finishLunching(activeEntry *Entry) *FinishedEntry {
	activeEntry = app.ActiveEntry
	app.ActiveEntry = nil

	now := time.Now()

	delta := now.Sub(activeEntry.StartTime)
	lunchEndTime := activeEntry.StartTime.Add(1 * time.Hour)

	if delta.Minutes() > 60 {
		return app.addEntry(spending, lunchEndTime, now)
	} else {
		return app.addEntry(overtime, now, lunchEndTime)
	}
}

func (app *App) addEntry(t EntryType, startTime time.Time, endTime time.Time) *FinishedEntry {
	finishedEntry := newFinishedEntry(t, startTime, endTime)
	app.FinishedEntries = append(app.FinishedEntries, *finishedEntry)
	return finishedEntry
}

func (app *App) spend() (bool, error) {
	if app.ActiveEntry != nil && app.ActiveEntry.EntryType == spending {
		return false, nil
	}

	app.finishActive()

	app.ActiveEntry = newEntry(spending, time.Now())
	return true, nil
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
		return fmt.Sprintf("overtime till %s added\n%s", startTime, app), nil
	case t.After(startTime) && t.Before(endTime) && t.Sub(startTime) < endTime.Sub(t):
		app.addEntry(spending, startTime, t)
		return fmt.Sprintf("spending from %s added\n%s", startTime, app), nil
	case t.After(startTime) && t.Before(endTime) && t.Sub(startTime) > endTime.Sub(t):
		app.addEntry(spending, t, endTime)
		return fmt.Sprintf("spending from %s to %s added\n%s", t, endTime, app), nil
	case t.After(endTime):
		app.addEntry(overtime, endTime, t)
		return fmt.Sprintf("overtime from %s added\n%s", endTime, app), nil
	default:
		return "nothing performed", nil
	}
}

func newDate(t time.Time, hour int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, t.Location())
}

func (app *App) lunch() bool {
	if app.ActiveEntry == nil {
		app.ActiveEntry = newEntry(lunching, time.Now())
		return true
	}

	return false
}
