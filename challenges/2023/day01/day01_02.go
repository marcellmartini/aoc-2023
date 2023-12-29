package day01

import (
	"fmt"
	"strings"
)

var digit_02 = map[string]int{
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

func AnserDay01_02(input []string) int {
	value := 0
	for _, line := range input {
		value += getLeftRight_02(line)
	}
	return value
}

func hasWordNumber(s string) (int, bool) {
	for k, v := range digit_02 {
		if strings.Contains(s, k) {
			return v, true
		}
	}
	return 0, false
}

func getLeftRight_02(line string) int {
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

	fmt.Printf("line %v, number %v\n", line, l*10+r)
	return l*10 + r
}
