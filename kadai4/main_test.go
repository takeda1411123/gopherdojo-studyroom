package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func SetTime(t time.Time) {
	timeNow = func() time.Time {
		return t
	}
}

func TestGetHandlerSanganichi(t *testing.T) {
	cases := []struct {
		name   string
		date   time.Time
		output Date
	}{
		{name: "1-1-2021", date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), output: Date{Year: 2021, Month: time.Month(1), Day: 1, Result: "大吉"}},
		{name: "1-2-2021", date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local), output: Date{Year: 2021, Month: time.Month(1), Day: 2, Result: "大吉"}},
		{name: "1-3-2021", date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local), output: Date{Year: 2021, Month: time.Month(1), Day: 3, Result: "大吉"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			SetTime(c.date)
			// Arrange
			reqBody := bytes.NewBufferString("request body")
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get", reqBody)

			// レスポンスを受け止める*httptest.ResponseRecorder
			got := httptest.NewRecorder()

			// Act
			getHandler(got, req)
			result, err := ioutil.ReadAll(got.Body)
			if err != nil {
				t.Error(err)
			}
			jsonBytes := ([]byte)(result)
			data := new(Date)
			if err := json.Unmarshal(jsonBytes, data); err != nil {
				t.Error(err)
			}

			if got.Code != http.StatusOK {
				t.Errorf("want %d, but %d", 200, got.Code)
			}

			if !reflect.DeepEqual(c.output, *data) {
				t.Errorf("invalid result. testCase:%#v, actual:%v", c.output, *data)
			}

		})
	}
}

func TestGetHandlerNotSanganichi(t *testing.T) {
	cases := []struct {
		name   string
		date   time.Time
		kind   [6]string
		output Date
	}{
		{name: "1-4-2021", date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local), kind: [6]string{"吉", "吉", "吉", "吉", "吉", "吉"}, output: Date{Year: 2021, Month: time.Month(1), Day: 4, Result: "吉"}},
		{name: "12-31-2021", date: time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), kind: [6]string{"凶", "凶", "凶", "凶", "凶", "凶"}, output: Date{Year: 2021, Month: time.Month(12), Day: 31, Result: "凶"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			SetTime(c.date)
			sample = c.kind
			// Arrange
			reqBody := bytes.NewBufferString("request body")
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get", reqBody)

			// レスポンスを受け止める*httptest.ResponseRecorder
			got := httptest.NewRecorder()

			// Act
			getHandler(got, req)
			result, err := ioutil.ReadAll(got.Body)
			if err != nil {
				t.Error(err)
			}
			jsonBytes := ([]byte)(result)
			data := new(Date)
			if err := json.Unmarshal(jsonBytes, data); err != nil {
				t.Error(err)
			}

			if got.Code != http.StatusOK {
				t.Errorf("want %d, but %d", 200, got.Code)
			}

			if !reflect.DeepEqual(c.output, *data) {
				t.Errorf("invalid result. testCase:%#v, actual:%v", c.output, *data)
			}

		})
	}
}
