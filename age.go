package age

import (
	"time"
)

type EntryDate struct {
	day   int
	month int
	year  int
}

type CurrentDate struct {
	day   int
	month int
	year  int
}

func GetCurrentDate() *CurrentDate {
	date := new(CurrentDate)
	t := time.Now()
	date.day = t.Day()
	date.month = int(t.Month())
	date.year = t.Year()

	return date
}

func getAge(date EntryDate) int {

	return date.year
}
