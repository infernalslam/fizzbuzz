package cap

import (
	"strconv"
)

// Cap must
func Cap(n1, n2, n3, n4 int) string {
	text := ""
	numberic := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	if n1 == 1 {
		text += strconv.Itoa(n2)
		if n3 == 1 {
			text += " + "
		} else if n3 == 2 {
			text += " - "
		} else if n3 == 3 {
			text += " / "
		} else if n3 == 4 {
			text += " * "
		}
		text += numberic[n4-1]
	} else if n1 == 2 {
		text += numberic[n2-1]
		if n3 == 1 {
			text += " + "
		} else if n3 == 2 {
			text += " - "
		} else if n3 == 3 {
			text += " / "
		} else if n3 == 4 {
			text += " * "
		}
		text += strconv.Itoa(n4)
	}
	return text
}
