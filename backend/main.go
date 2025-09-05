package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "time"
)

type TowelDayResponse struct {
  IsTowelDay   bool   `json:"is_towel_day"`
  TowelDay     string `json:"towel_day"`
  CurrentDate  string `json:"current_date"`
  DaysUntil    int    `json:"days_until"`
  Message      string `json:"message"`
}

func isTowelDayHandler(w http.ResponseWriter, r *http.Request) {
  currentDate := time.Now().UTC()
  towelDay := time.Date(currentDate.Year(), time.May, 25, 0, 0, 0, 0, time.UTC)

  isTowelDay := currentDate.Year() == towelDay.Year() && currentDate.Month() == towelDay.Month() && currentDate.Day() == towelDay.Day()

  var daysUntil int
  var message string

  if isTowelDay {
    daysUntil = 0
    message = "Today is Towel Day! Don't forget to bring your towel."
  } else {
    if currentDate.After(towelDay) {
      towelDay = towelDay.AddDate(1, 0, 0)
    }
    daysUntil = int(towelDay.Sub(currentDate).Hours() / 24)
    message = fmt.Sprintf("There are %d days until Towel Day.", daysUntil)
  }

  response := TowelDayResponse{
    IsTowelDay:   isTowelDay,
    TowelDay:     towelDay.Format("January 2"),
    CurrentDate:  currentDate.Format("January 2"),
    DaysUntil:    daysUntil,
    Message:      message,
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func main() {
  http.HandleFunc("/is-towel-day", isTowelDayHandler)
  http.ListenAndServe(":8080", nil)
}
