package main

import (
	"testing"
	"time"
)

func TestApp_calculateTotal(t *testing.T) {
	type fields struct {
		ActiveEntry     Entry
		FinishedEntries []FinishedEntry
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
				ActiveEntry: *newEntry(spending, time.Now().Add(-10*time.Second)),
				FinishedEntries: []FinishedEntry{
					*newFinishedEntry(overtime, time.Now().Add(-30*time.Second), time.Now()),
					*newFinishedEntry(overtime, time.Now().Add(-70*time.Second), time.Now()),
				}},
			want:  90 * time.Second,
			want1: hasOvertime,
		},
		{
			name: "debt",
			fields: fields{
				FinishedEntries: []FinishedEntry{
					*newFinishedEntry(overtime, time.Now().Add(-30*time.Second), time.Now()),
					*newFinishedEntry(spending, time.Now().Add(-70*time.Second), time.Now()),
				}},
			want:  -40 * time.Second,
			want1: hasDebt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				ActiveEntry:     &tt.fields.ActiveEntry,
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

func TestApp_StartStop(t *testing.T) {
	//given
	app := NewApp(nil, nil, []FinishedEntry{})

	//when
	app.start()
	entryType, err := app.stop()

	//then
	if err != nil {
		t.Error(err)
	}

	if entryType != overtime {
		t.Errorf("app start/stop Entry type, got %v, want %v", entryType, overtime)
	}

	if len(app.FinishedEntries) != 1 {
		t.Errorf("app start/stop entries count, got %d, want 1", len(app.FinishedEntries))
	}
}

func TestApp_SpendStop(t *testing.T) {
	//given
	app := NewApp(nil, nil, []FinishedEntry{})

	//when
	app.spend()
	entryType, err := app.stop()

	//then
	if err != nil {
		t.Error(err)
	}

	if entryType != spending {
		t.Errorf("app spend/stop Entry type, got %v, want %v", entryType, spending)
	}

	if len(app.FinishedEntries) != 1 {
		t.Errorf("app spend/stop entries count, got %d, want 1", len(app.FinishedEntries))
	}
}
