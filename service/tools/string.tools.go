package tools

import (
	"log"
	"strconv"
)

func IntToString(n int64) (s string, err error) {
	s = strconv.Itoa(int(n))

	return s, err
}

func StringToInt64(s string) (n int64) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		return n
	}

	n = int64(i)

	return n
}
