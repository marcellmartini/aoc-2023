package day03

import (
	"bytes"
	"regexp"
	"runtime"
	"strconv"
	"unicode"

	"github.com/marcellmartini/aoc-in-go/puzzle"
)

var _, fp, _, _ = runtime.Caller(0)

var Puzzle = puzzle.NewBuilder().
	ConfigurePWD(fp).
	ConfigureSolutions(part1(), part2()).
	LoadFiles().
	Build()

func part1() puzzle.SolutionFunc {
	return func(input []string) int {
		numberList, _ := getAllValidNumbersAndGearlist(&input)

		return sumValidNumbers(numberList)
	}
}

func part2() puzzle.SolutionFunc {
	return func(input []string) int {
		result := 0

		numberList, gearList := getAllValidNumbersAndGearlist(&input)

		for _, g := range gearList {
			if ok, num := isGearAdjacentTwoNumbers(g, numberList); ok {
				n1, err := strconv.Atoi(num[0].number)
				if err != nil {
					panic(err)
				}

				n2, err := strconv.Atoi(num[1].number)
				if err != nil {
					panic(err)
				}

				result += n1 * n2
			}
		}

		return result
	}
}

type numbers struct {
	number string
	row    int
	col    int
	length int
}

type gear struct {
	row int
	col int
}

func isGearAdjacentTwoNumbers(g gear, numberList []numbers) (bool, []numbers) {
	nearbyCount := 0
	nl := []numbers{}

	// Check if the gear is adjacent number ...
	for _, num := range numberList {
		if isGearAdjacentToNumber(g, &num) {
			nearbyCount++
			nl = append(nl, num)
		}
	}

	// and return true if is adjacent to exactly a two numbers
	if nearbyCount == 2 {
		return true, nl
	}

	// or false if not .
	return false, []numbers{}
}

func isGearAdjacentToNumber(g gear, num *numbers) bool {
	// Check if a gear is adjacent to a number based on directions
	for _, dir := range directions {
		newGear := gearDeslocated(g, dir)
		if newGear.row == num.row &&
			num.col <= newGear.col &&
			newGear.col < num.col+num.length {
			return true
		}
	}
	return false
}

func gearDeslocated(g gear, direction []int) gear {
	// Change gear location based in the direction
	return gear{
		g.row + direction[0],
		g.col + direction[1],
	}
}

func getAllValidNumbersAndGearlist(schematic *[]string) ([]numbers, []gear) {
	var buffer string
	numberList := []numbers{}
	gearList := []gear{}
	lenght := 0

	buff := bytes.NewBufferString(buffer)

	for i, line := range *schematic {
		for j, r := range line {
			if unicode.IsDigit(r) {
				buff.WriteRune(r)
				lenght++

				// if number is near of the end of the line
				if j+1 == len(line) && unicode.IsDigit(rune(line[j])) {
					appendValidNumber(schematic, &numberList, buff, &lenght, i, j)

					// all other numbers position
				} else if !unicode.IsDigit(rune(line[j+1])) {
					appendValidNumber(schematic, &numberList, buff, &lenght, i, j)
				}
			}

			if r == '*' {
				gearList = append(gearList, gear{i, j})
			}

		}
	}
	return numberList, gearList
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
//
//	0,-1  i, j  0, 1
//	1,-1  1, 0  1, 1
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
	for l := 0; l < number.length; l++ {
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
