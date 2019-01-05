package fizzbuzz

import (
	"strconv"
)

// Say must return string
func Say(n int) string {
	if n == 3 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
