package shared

import "strconv"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Atoi(s string) int {
	intValue, _ := strconv.Atoi(s)
	return intValue
}
