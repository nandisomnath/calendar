package calendar


// Returns the day number since epoch
// For now it only works if year >= epoachYear
// TODO: do some testing please
func SinceEpoch(day int, month Month, year int) int {

	epochYear := 1970
	epochMonth := JANUARY
	epochDay := 1


	days := 0

	for i := epochYear; i <= year; i++ {
		
		if IsLeap(i) {
			days += 366
		} else {
			days += 365
		}
	}


	for i := epochMonth; i <= month; i++ {
		days += getDayNumber(i, year)
	}

	days += (day - epochDay)

	return days
}
