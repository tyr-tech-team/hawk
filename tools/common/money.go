package common

import "strconv"

func FormatMoney(m int64) string {

	// s := strconv.Itoa(int(m))
	s := strconv.FormatInt(m, 10)
	str := ""
	if m == 0 {
		return ""
	}
	x := 0
	for i := len(s) - 1; i >= 0; i-- {
		if x != 0 && (x%3) == 0 {
			str = "," + str
			x = 0
		}
		x++
		str = string(s[i]) + str
	}

	return str

}
