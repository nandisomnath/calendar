package calendar

type Month int

const (
	JANUARY Month = iota + 1
	FEBRUARY
	MARCH
	APRIL
	MAY
	JUNE
	JULY
	AUGUST
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER
)

// Returns the days number
// year - is need to determine number of days for february
func getDayNumber(month Month, year int) int {

	switch month {
	case JANUARY:
		return 31
	case FEBRUARY:
		if IsLeap(year) {
			return 29
		}
		return 28
	case MARCH:
		return 31
	case APRIL:
		return 30
	case MAY:
		return 31
	case JUNE:
		return 30
	case JULY:
		return 31
	case AUGUST:
		return 31
	case SEPTEMBER:
		return 30
	case OCTOBER:
		return 31
	case NOVEMBER:
		return 30
	case DECEMBER:
		return 31
	}
	return -1
}
