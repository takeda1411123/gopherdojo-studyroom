package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Init
var timeNow = func() time.Time { return time.Now() }
var sample = [6]string{"大吉", "吉", "中吉", "小吉", "末吉", "凶"}

// Date : declare struct
type Date struct {
	Year   int        `json:"year"`
	Month  time.Month `json:"month"`
	Day    int        `json:"day"`
	Result string     `json:"result"`
}

func main() {
	http.HandleFunc("/get", getHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// getHandler : do omikuji
func getHandler(w http.ResponseWriter, r *http.Request) {
	// GET
	t := timeNow()
	today := Date{
		Year:   t.Year(),
		Month:  t.Month(),
		Day:    t.Day(),
		Result: "",
	}

	resultSample := sample
	rand.Seed(time.Now().UnixNano())

	if today.Month.String() == "January" && today.Day <= 3 {
		today.Result = "大吉"
	} else {
		today.Result = resultSample[rand.Intn(6)]
	}

	json.NewEncoder(w).Encode(today)
}
