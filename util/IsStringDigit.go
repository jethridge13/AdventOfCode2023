package util

import "strconv"

func IsStringDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
