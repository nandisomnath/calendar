package calendar

// Is the date is the year is leapyear or not
func IsLeap(year int) bool {
	if year%4 == 0 || year%400 == 0 && year%100 != 0 {
		return true
	}
	return false
}


// Returns the weekday of the date
func GetWeekDay(day int, month Month, year int) WeekDay {
	eDay := THURSDAY
	days := SinceEpoch(day, month, year)
	extra := days % 7
	extra = (int(eDay) + extra) % 7
	return WeekDay(extra)
}


// Returns the number of leap years in the range
// from y1 to y2 (exclusive), where y1 and y2 are years.
// This function works for ranges spanning a century change.
func LeapYears(y1, y2 int) int {
	count := 0
	for i := y1; i < y2; i++ {
		if IsLeap(i) {
			count += 1
		}
	}
	return count
}
