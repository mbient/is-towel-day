package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIsTowelDay(t *testing.T) {
	towelDay := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		currentDate time.Time
		want        bool
	}{
		{time.Date(2024, 5, 25, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2025, 5, 26, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2025, 6, 25, 0, 0, 0, 0, time.UTC), false},
	}

	for _, test := range tests {
		result := IsEventToday(test.currentDate, towelDay)
		if result != test.want {
			t.Errorf("IsTowelDay(%v, %v) = %v; want %v", test.currentDate, towelDay, result, test.want)
		}
	}
}

func TestTodayIsEvent(t *testing.T) {
	towelDay := time.Now().UTC()

	tests := []struct {
		currentDate time.Time
		want        bool
	}{
		{towelDay, true},
		{time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC), false},
	}

	for _, test := range tests {
		result := IsEventToday(test.currentDate, towelDay)
		if result != test.want {
			t.Errorf("IsTowelDay(%v, %v) = %v; want %v", test.currentDate, towelDay, result, test.want)
		}
	}
}

func TestCalculateDaysUntil(t *testing.T) {
	tests := []struct {
		currentDate time.Time
		eventDate   time.Time
		want        int
	}{
		{time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC), time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC), 1},
		{time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC), time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC), 0},
		{time.Date(2025, 5, 26, 0, 0, 0, 0, time.UTC), time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC), -1},
	}

	for _, test := range tests {
		result := CalculateDaysUntil(test.currentDate, test.eventDate)
		if result != test.want {
			t.Errorf("CalculateDaysUntil(%v, %v) = %v; want %v", test.currentDate, test.eventDate, result, test.want)
		}
	}
}

func TestGenerateMessage(t *testing.T) {
	tests := []struct {
		isEvent bool
		days    int
		want    string
	}{
		{true, 0, "Today is Towel Day! Don't forget to bring your towel."},
		{false, 5, "There are 5 days until Towel Day."},
	}

	for _, test := range tests {
		result := GenerateMessage(test.isEvent, test.days)
		if result != test.want {
			t.Errorf("GenerateMessage(%v, %d) = %v; want %v", test.isEvent, test.days, result, test.want)
		}
	}
}

func TestIsEventHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/is-towel-day", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(isEventHandler)

	handler.ServeHTTP(rr, req)

	// Check server codes
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check JSON-Response
	var response EventResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	// Check response
	currentDate := time.Now().UTC()
	towelDay := time.Date(currentDate.Year(), time.May, 25, 0, 0, 0, 0, time.UTC)
	isEvent := IsEventToday(currentDate, towelDay)

	if response.IsEvent != isEvent {
		t.Errorf("Expected IsEvent to be %v, got %v", isEvent, response.IsEvent)
	}
}
