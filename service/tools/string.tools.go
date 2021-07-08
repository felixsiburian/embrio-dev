package tools

import "strconv"

func IntToString(n int64) (s string, err error) {
	s = strconv.Itoa(int(n))

	return s, err
}
