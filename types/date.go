package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Date is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".
type Date struct {
	time.Time
}

var (
	_ json.Marshaler   = &Date{}
	_ json.Unmarshaler = &Date{}
	_ fmt.Stringer     = &Date{}
)

// NewDate returns an instance of Date from a time.Time.
func NewDate(t time.Time) *Date {
	d := DateFromTime(t)
	return &d
}

// DateFromTime returns a Date from a time.Time.
func DateFromTime(t time.Time) Date {
	return Date{t}
}

// NewDateFromString returns an instance of Date from a string formatted as "2006-01-02".
func NewDateFromString(str string) (*Date, error) {
	d, err := DateFromString(str)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// DateFromString returns a Date from a string formatted as "2006-01-02".
func DateFromString(str string) (Date, error) {
	var d Date
	var err error

	d.Time, err = time.Parse("2006-01-02", str)
	return d, err
}

// MustNewDateFromString returns an instance of Date from a string formatted as "2006-01-02" or panics.
// Avoid using this function in production code.
func MustNewDateFromString(str string) *Date {
	d := MustDateFromString(str)
	return &d
}

// MustDateFromString returns a Date from a string formatted as "2006-01-02" or panics.
// Avoid using this function in production code.
func MustDateFromString(str string) Date {
	d, err := DateFromString(str)
	if err != nil {
		panic(err)
	}
	return d
}

func (d Date) GetTime() time.Time {
	return d.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format("2006-01-02"))), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var err error

	str := string(data)
	str = strings.Trim(str, `"`)

	d.Time, err = time.Parse("2006-01-02", str)
	return err
}

func (d Date) String() string {
	return d.Time.Format("2006-01-02")
}
