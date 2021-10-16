package main

/**
 * 月份天数
 * @param year: a number year
 * @param month: a number month
 * @return: Given the year and the month, return the number of days of the month.
 */
func getTheMonthDays(year int, month int) int {
	if month == 2 {
		if isLeapYear(year) {
			return 29
		} else {
			return 28
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	} else {
		return 31
	}
}

func isLeapYear(n int) bool {
	if (n%4 == 0 && n%100 != 0) || n%400 == 0 {
		return true
	}
	return false
}
