package clock

import (
	"encoding/json"
	"time"

	"github.com/aodin/date"
)

type Time struct {
	time.Time
}

type Date struct {
	date.Date
}

func NewDate(year int, month time.Month, day int) Date {
	d := date.New(year, month, day)
	return Date{Date: d}
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// 28-1-2008
	layout := "2006-01-02"
	time, err := time.Parse(layout, value)
	d.Date = date.FromTime(time)
	return err
}

func (d Date) MarshalSchema() string {
	if d.Equal(time.Time{}) {
		return ""
	}
	return d.Format("2006-01-02")
}

type Currency string

type CurrencyValue struct {
	Cents    int      `json:"cents"`
	Currency Currency `json:"currency"`
}
