package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type CustomDate struct {
	time.Time
}

func GetWeekRange(day CustomDate) CustomDate {
	return CustomDate{time.Date(day.Year(), day.Month(), day.Day()+7, 0, 0, 0, 0, day.Location())}
}

func GetMonthRange(day CustomDate) CustomDate {
	return CustomDate{time.Date(day.Year(), day.Month()+1, 0, 0, 0, 0, 0, day.Location())}
}

func GetDate(date string) (CustomDate, error) {
	t, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return CustomDate{}, errors.New("wrong date template, use YYYY-MM-DD")
	}
	return CustomDate{t}, nil
}

func (d *CustomDate) UnmarshalJSON(b []byte) error {
	timeStr := strings.ReplaceAll(string(b), `"`, "")
	t, err := time.Parse(time.DateOnly, timeStr)
	if err != nil {
		return fmt.Errorf("unmarshal failed: %v", err)
	}

	*d = CustomDate{t}
	return nil
}

func (d *CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.Format(time.DateOnly))), nil
}

func (d CustomDate) Value() (driver.Value, error) {
	return d.Format(time.DateOnly), nil
}

func (d *CustomDate) Scan(src interface{}) (err error) {
	var t time.Time
	switch src.(type) {
	case string:
		t, err = time.Parse(time.DateOnly, src.(string))
	case []byte:
		t, err = time.Parse(time.DateOnly, string(src.([]byte)))
	}
	if err != nil {
		return fmt.Errorf("scan failed: %v", err)
	}
	*d = CustomDate{t}
	return nil
}
