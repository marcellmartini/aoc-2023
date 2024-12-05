package utils

import (
	"log"
	"os"
	"sort"
	"strconv"
)

// CheckError logs the error and exits the program if the error is not nil.
// This is a utility function to handle unexpected errors.
func CheckError(err error) {
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}

// StrToInt converts a string to an integer.
// It will exit the program if the string cannot be converted.
// TODO: Consider returning an error instead of exiting directly.
// This would provide more flexibility for error handling by the caller.
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckError(err)
	return i
}

// Abs returns the absolute value of the input integer x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// RemoveIndex removes the element at index i from a sort.IntSlice.
// It creates a new slice without the element at index i and returns it.
// This does not modify the original slice.
func RemoveIndex(s sort.IntSlice, i int) sort.IntSlice {
	// Create a new slice to hold the result
	ret := make(sort.IntSlice, 0)

	// Append elements before index i
	ret = append(ret, s[:i]...)

	// Append elements after index i
	return append(ret, s[i+1:]...)
}
