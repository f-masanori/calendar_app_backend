package utils

import (
	"log"
	"strconv"
)

func StrToInt(s string) int {
	toInt, err := strconv.Atoi(s)
	if err != nil {
		log.Panicln(err)
		return -1
	}
	return toInt
}
