package main

import "testing"

func Test_readCommandLineArg(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    commandLineArg
		wantErr bool
	}{
		{"start", []string{"start"}, start, false},
		{"stop", []string{"stop"}, stop, false},
		{"unknown", []string{"unknown"}, status, true},
		{"empty", []string{""}, status, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCommandLineArg(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCommandLineArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readCommandLineArg() got = %v, want %v", got, tt.want)
			}
		})
	}
}
