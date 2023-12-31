package day03

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

type numbers struct {
	number string
	row    int
	col    int
	len    int
}

func SumOfAllPartNumbers(schematic []string) int {
	numberList := getAllValidNumbers(schematic)

	return sumValidNumbers(numberList)
}

func getAllValidNumbers(schematic *[]string) []numbers {
	var buffer string
	resp := []numbers{}
	lenght := 0

	buff := bytes.NewBufferString(buffer)

	for i, line := range *schematic {
		for j, r := range line {
			if unicode.IsDigit(r) {
				buff.WriteRune(r)
				lenght++

				// if number is near of the end of the line
				if j+1 == len(line) && unicode.IsDigit(rune(line[j])) {
					appendValidNumber(schematic, &resp, buff, &lenght, i, j)

					// all other numbers position
				} else if !unicode.IsDigit(rune(line[j+1])) {
					appendValidNumber(schematic, &resp, buff, &lenght, i, j)
				}
			}
		}
	}
	return resp
}

func appendValidNumber(
	schematic *[]string,
	xn *[]numbers,
	b *bytes.Buffer,
	length *int,
	i int,
	j int,
) {
	n, ok := scanSurroundingsNumbers(
		schematic,
		numbers{
			b.String(),
			i,
			j - (*length - 1),
			*length,
		})

	if ok {
		*xn = append(*xn, n)
	}

	b.Reset()
	*length = 0
}

// Locations relative to i, j
//
// -1,-1 -1, 0 -1, 1
//  0,-1  i, j  0, 1
//  1,-1  1, 0  1, 1
var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func scanSurroundingsNumbers(schematic *[]string, number numbers) (numbers, bool) {
	rows, cols := len(*schematic), len((*schematic)[0])
	r, err := regexp.Compile("[^\\d.]")
	if err != nil {
		panic(err)
	}

	// in the lenght of the number ...
	for l := 0; l < number.len; l++ {
		// looking for in all direction ...
		for _, d := range directions {
			newRow, newCol := number.row+d[0], number.col+d[1]+l
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				// if have a symbol ...
				if r.MatchString(string((*schematic)[newRow][newCol])) {
					// return the number.
					return number, true
				}
			}
		}
	}

	// if not have a symbol
	return numbers{}, false
}

func sumValidNumbers(xn []numbers) int {
	result := 0
	for _, n := range xn {
		ni, err := strconv.Atoi(n.number)
		if err != nil {
			panic(err)
		}
		result += ni
	}

	return result
}
