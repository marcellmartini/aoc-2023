package day01

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
