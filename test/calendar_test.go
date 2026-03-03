package test

import (
	"fmt"
	"testing"

	"github.com/nandisomnath/calendar"
)



func TestGetWeekDay(t *testing.T)  {
	weekday := calendar.GetWeekDay(3, calendar.MARCH, 2026)
	if weekday != calendar.TUESDAY {
		t.Fatalf("not equal weekday %d\n", weekday)
	}
}


func TestSinceEpoch(t *testing.T)  {
	days := calendar.SinceEpoch(3, calendar.MARCH, 2026)
	if 20515 != days {
		t.Fatalf("required: %d, found: %d\n", 20515, days)
	}
}


func TestIsLeap(t *testing.T)  {
	for i := 1970; i <= 2026; i++ {
		if calendar.IsLeap(i) {
			fmt.Printf("%d, ", i)
		}
	}
	fmt.Println()
}