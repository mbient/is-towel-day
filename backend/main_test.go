package main

import (
	"testing"
	"time"
)

func TestIsTowelDay(t *testing.T) {
	towelDay := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		currentDate time.Time
		want    bool
	}{
		{time.Date(2024, 5, 25, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2025, 5, 26, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2025, 6, 25, 0, 0, 0, 0, time.UTC), false},
	}

	for _, test := range tests {
		result := IsTowelDay(test.currentDate, towelDay)
		if result != test.want {
			t.Errorf("IsTowelDay(%v, %v) = %v; want %v", test.currentDate, towelDay, result, test.want)
		}
	}
}
