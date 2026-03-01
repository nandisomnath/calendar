package calendar



type Date struct {
	day int
	month int
	year int
}


func NewDate(day, month, year int) Date {
	return Date{
		day: day,
		month: month,
		year: year,
	}
}

