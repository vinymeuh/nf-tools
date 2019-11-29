package main

import (
	"testing"
	"time"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		path     string
		datetime time.Time
		want     string
	}{
		{"DSC_1234.jpg", time.Date(2019, 11, 28, 12, 4, 5, 0, time.UTC), "./2019-11-28/2019-11-28_120405_DSC_1234.jpg"},
		{"./Import/DSC_1234.jpg", time.Date(2019, 11, 28, 12, 4, 5, 0, time.UTC), "Import/2019-11-28/2019-11-28_120405_DSC_1234.jpg"},
		{"./Import/2019-11-28_120405_DSC_1234.jpg", time.Date(2019, 11, 28, 12, 4, 5, 0, time.UTC), "Import/2019-11-28/2019-11-28_120405_DSC_1234.jpg"},
		{"./Import/2019-11-28_120405.jpg", time.Date(2019, 11, 28, 12, 4, 5, 0, time.UTC), "Import/2019-11-28/2019-11-28_120405.jpg"},
	}

	for i, tc := range tests {
		got := newName(tc.path, tc.datetime)
		if got != tc.want {
			t.Fatalf("Test n°%d: got=%s, want=%s", i+1, got, tc.want)
		}
	}

}
