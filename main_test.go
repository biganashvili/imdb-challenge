package main

import (
	"testing"
)

func Test_getPlot(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "correct id",
			args:    args{ID: "tt0036443"},
			want:    "Third Reich's Nazi propaganda epic about a heroic fictional German officer on board of the RMS Titanic. On its maiden voyage in April 1912, the supposedly unsinkable ship hits an iceberg in the Atlantic Ocean and starts to go down.",
			wantErr: false,
		},
		{
			name:    "wrong id",
			args:    args{ID: "tp0036443"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPlot(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPlot() = %v, want %v", got, tt.want)
			}
		})
	}
}
