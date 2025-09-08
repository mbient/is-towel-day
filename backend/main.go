package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EventResponse struct {
	IsEvent     bool   `json:"is_event"`
	EventDate   string `json:"event_date"`
	CurrentDate string `json:"current_date"`
	DaysUntil   int    `json:"days_until"`
	Message     string `json:"message"`
}

func IsEventToday(currentDate, eventDate time.Time) bool {
	return currentDate.Month() == eventDate.Month() && currentDate.Day() == eventDate.Day()
}

func CalculateDaysUntil(currentDate, eventDate time.Time) int {
	return int(eventDate.Sub(currentDate).Hours() / 24)
}

func GenerateMessage(isEvent bool, days int) string {
	var message string
	if isEvent {
		message = "Today is Towel Day! Don't forget to bring your towel."
	} else {
		message = fmt.Sprintf("There are %d days until Towel Day.", days)
	}
	return message
}

func isEventHandler(w http.ResponseWriter, r *http.Request) {
	currentDate := time.Now().UTC()
	towelDay := time.Date(currentDate.Year(), time.May, 25, 0, 0, 0, 0, time.UTC)
	isEvent := IsEventToday(currentDate, towelDay)

	var daysUntil int
	var message string

	if isEvent {
		daysUntil = 0
		message = GenerateMessage(isEvent, daysUntil)
	} else {
		if currentDate.After(towelDay) {
			towelDay = towelDay.AddDate(1, 0, 0)
		}
		daysUntil = CalculateDaysUntil(currentDate, towelDay)
		message = GenerateMessage(isEvent, daysUntil)
	}

	response := EventResponse{
		IsEvent:     isEvent,
		EventDate:   towelDay.Format("January 2"),
		CurrentDate: currentDate.Format("January 2"),
		DaysUntil:   daysUntil,
		Message:     message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/is-towel-day", isEventHandler)
	http.ListenAndServe(":8080", nil)
}
