package fizzbuzz

import (
	"strconv"
)

// Say must return string
func Say(n int) string {
	if n == 3 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
