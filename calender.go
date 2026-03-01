package calendar

// Is the date is the year is leapyear or not
func IsLeap(year int) bool {

	if year%4 == 0 || year%400 == 0 && year%100 != 0 {
		return true
	}

	return false
}


func weekday()  {
	
}