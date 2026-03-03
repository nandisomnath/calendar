package calendar

// Is the date is the year is leapyear or not
func IsLeap(year int) bool {

	if year%4 == 0 || year%400 == 0 && year%100 != 0 {
		return true
	}

	return false
}

func GetWeekDay(day int, month Month, year int) WeekDay {
	eDay := THURSDAY
	days := SinceEpoch(day, month, year)
	extra := days % 7
	extra = (int(eDay) + extra) % 7
	return WeekDay(extra)
}
