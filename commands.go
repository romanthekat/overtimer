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

	activeDuration := time.Now().Sub(app.ActiveEntry.StartTime)
	if app.ActiveEntry.EntryType == overtime {
		result += activeDuration
	} else if app.ActiveEntry.EntryType == spending {
		result -= activeDuration
	}

	result = result.Round(time.Second)
	if result > 0 {
		return result, hasOvertime
	} else {
		return result, hasDebt
	}
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}