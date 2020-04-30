package main

import (
	"testing"
	"time"
)

func TestApp_calculateTotal(t *testing.T) {
	type fields struct {
		ActiveEntry     entry
		FinishedEntries []finishedEntry
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
		want1  totalType
	}{
		{
			name: "overtime",
			fields: fields{
				ActiveEntry: entry{
					EntryType: spending,
					StartTime: time.Now().Add(-10 * time.Second),
				},
				FinishedEntries: []finishedEntry{
					{
						entry: entry{
							EntryType: overtime,
							StartTime: time.Now().Add(-30 * time.Second),
						},
						EndTime: time.Now(),
					},
					{
						entry: entry{
							EntryType: overtime,
							StartTime: time.Now().Add(-70 * time.Second),
						},
						EndTime: time.Now(),
					},
				}},
			want:  90 * time.Second,
			want1: hasOvertime,
		},
		{
			name: "debt",
			fields: fields{
				ActiveEntry: entry{
					EntryType: overtime,
					StartTime: time.Now().Add(-10 * time.Second),
				},
				FinishedEntries: []finishedEntry{
					{
						entry: entry{
							EntryType: spending,
							StartTime: time.Now().Add(-30 * time.Second),
						},
						EndTime: time.Now(),
					},
					{
						entry: entry{
							EntryType: spending,
							StartTime: time.Now().Add(-70 * time.Second),
						},
						EndTime: time.Now(),
					},
				}},
			want:  -90 * time.Second,
			want1: hasDebt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				ActiveEntry:     tt.fields.ActiveEntry,
				FinishedEntries: tt.fields.FinishedEntries,
			}
			got, got1 := app.calculateTotal()
			if got != tt.want {
				t.Errorf("calculateTotal() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculateTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
