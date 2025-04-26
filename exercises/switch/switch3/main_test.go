// switch3
// Make me compile!

package main_test

import (
	"errors"
	"testing"
)

func weekDay(day int) (string, error) {
	// Return the day of the week based on the
	// integer. Use a switch case to satisfy all test cases below
	switch {
	case day == 0:
		return "Sunday", nil
	case day == 1:
		return "Monday", nil
	case day == 2:
		return "Tuesday", nil
	case day == 3:
		return "Wednesday", nil
	case day == 4:
		return "Thursday", nil
	case day == 5:
		return "Friday", nil
	case day == 6:
		return "Saturday", nil
	case day == 7:
		return "Sunday", nil
	}
	return "", errors.New("day was greater than 6; limits of the days of the week")
}

func TestWeekDay(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{input: 0, want: "Sunday"},
		{input: 1, want: "Monday"},
		{input: 2, want: "Tuesday"},
		{input: 3, want: "Wednesday"},
		{input: 4, want: "Thursday"},
		{input: 5, want: "Friday"},
		{input: 6, want: "Saturday"},
	}

	for _, tc := range tests {
		day, err := weekDay(tc.input)
		if err != nil {
			t.Errorf("Got a too high/low integer")
		}
		if day != tc.want {
			t.Errorf("expected %s but got %s", tc.want, day)
		}
	}
}
