package day01

import "strings"

var digit = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func SumOfCalibrations(input []string) int {
	value := 0

	for _, line := range input {
		value += getLeftRight(line)
	}
	return value
}

func getLeftRight(line string) int {
	l, r := 0, 0
	i, j := 0, len(line)-1
	ok := false

	for i <= j {
		// Loop forward to find the first digit (Left)
		if l, ok = digit[line[i]]; !ok {
			i += 1
		}

		// Loop backward to find the last digit (Right)
		if r, ok = digit[line[j]]; !ok {
			j -= 1
		}

		// Stop the loop if find both left and right digits 
		if l != 0 && r != 0 {
			break
		}
	}

	return l*10 + r
}

var digitWithWord = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func SumOfCalibrationWithWord(input []string) int {
	value := 0
	for _, line := range input {
		value += getLeftRightWithWord(line)
	}
	return value
}

func hasWordNumber(s string) (int, bool) {
	for k, v := range digitWithWord{
		if strings.Contains(s, k) {
			return v, true
		}
	}
	return 0, false
}

func getLeftRightWithWord(line string) int {
	l, r := 0, 0
	i, j := 0, len(line)-1
	ok := false

	for i <= j {
		// Loop forward to find the first digit (Left)
		if d, hasWorld := hasWordNumber(line[:i]); hasWorld {
			l = d
		} else if l, ok = digit[line[i]]; !ok {
			i += 1
		}

		// Loop backward to find the last digit (Right)
		if d, hasWorld := hasWordNumber(line[j:]); hasWorld {
			r = d
		} else if r, ok = digit[line[j]]; !ok {
			j -= 1
		}

		// Stop the loop if find both left and right digits 
		if l != 0 && r != 0 {
			break
		}
	}
	return l*10 + r
}
