package utils

import (
	"log"
	"strconv"
)

func StringToInt(intStr string) int {
	num, err := strconv.Atoi(intStr)
	if err != nil {
		log.Fatal("$1 was failed to convert to integer type", intStr)
	}
	return num
}
