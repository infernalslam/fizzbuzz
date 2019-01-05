package cap

import (
	"strconv"
)

// Cap must
func Cap(n1, n2, n3, n4 int) string {
	text := ""
	if n1 == 1 {
		text += strconv.Itoa(n2)
		text += operation(n3)
		text += numberic(n4)
	} else if n1 == 2 {
		text += numberic(n2)
		text += operation(n3)
		text += strconv.Itoa(n4)
	}
	return text
}

func numberic(n int) string {
	numberic := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return numberic[n-1]
}

func operation(n3 int) string {
	text := ""
	if n3 == 1 {
		text += " + "
	} else if n3 == 2 {
		text += " - "
	} else if n3 == 3 {
		text += " / "
	} else if n3 == 4 {
		text += " * "
	}
	return text
}
