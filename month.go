package calendar



type Month int



const (
	JANUARY Month = iota
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
// INFO: This is a protopye function, NOT FOR USE
func getDayNumber(month Month) int {
	switch month {
	case JANUARY:
		return 31
	case FEBRUARY:
		return -1 // This is unkown only from month
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

