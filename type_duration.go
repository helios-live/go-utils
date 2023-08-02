package utils

import (
	"database/sql/driver"
	"errors"
	"time"
)

// Duration represents a duration
type Duration struct {
	time.Duration
}

// MarshalJSON .
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}

// UnmarshalJSON .
func (d *Duration) UnmarshalText(b []byte) error {

	t, err := time.ParseDuration(string(b))
	if err != nil {
		return err
	}
	d.Duration = t
	return nil
}

// UnmarshalJSON .
func (d *Duration) UnmarshalJSON(b []byte) error {
	return d.UnmarshalText(b[1 : len(b)-1])
}

// Adds parses d2 and adds the duration to d
func (d *Duration) Adds(d2 string) {
	d2d, _ := ParseDuration(d2)
	d.Add(d2d)
}

// Add returns d2 added to d
func (d *Duration) Add(d2 Duration) {
	d.Duration = time.Duration(d.Duration.Nanoseconds() + d2.Duration.Nanoseconds())
}

// ParseDuration returns a Duration struct parsed from a string in format 10h2m3s
func ParseDuration(str string) (Duration, error) {
	d, err := time.ParseDuration(str)
	if err != nil {
		return Duration{0}, err
	}
	return Duration{d}, nil
}

// SubZero returns true if the underlying duration is negative
func (d *Duration) SubZero() bool {
	return d.Duration < 0
}

func (d *Duration) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil {
		// set the value of the pointer yne to YesNoEnum(false)
		*d = Duration{0}
		return nil
	}
	if bv, err := driver.String.ConvertValue(value); err == nil {
		// if this is a bool type
		if v, ok := bv.(string); ok {

			r, err := ParseDuration(v)
			if err != nil {
				return err
			}
			// set the value of the pointer yne to YesNoEnum(v)
			*d = r
			return nil
		}
	}
	// otherwise, return an error
	return errors.New("failed to scan Duration")
}
func (d *Duration) Value() (driver.Value, error) {
	if d.Duration == 0 {
		return "0", nil
	}
	return d.String(), nil
}
