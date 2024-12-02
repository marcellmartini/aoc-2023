package utils

import (
	"log"
	"os"
	"strconv"
)

func CheckError(err error) {
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}

// StrToInt convert s string to s int
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckError(err)

	return i
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
