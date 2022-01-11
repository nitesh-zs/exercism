package exercism

// IsLeap checks if a given year is leap year or not
func isLeap(year int) bool {
	if year%100 == 0 && year%400 != 0 {
		return false
	} else if year%4 != 0 {
		return false
	}
	return true
}
