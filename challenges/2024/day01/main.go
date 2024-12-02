package day0X

import (
	"runtime"
	"sort"
	"strings"

	"github.com/marcellmartini/aoc-in-go/puzzle"
	"github.com/marcellmartini/aoc-in-go/utils"
)

var _, fp, _, _ = runtime.Caller(0)

var Puzzle = puzzle.NewBuilder().
	ConfigurePWD(fp).
	ConfigureSolutions(part1(), part2()).
	LoadFiles().
	Build()

func part1() puzzle.SolutionFunc {
	return func(input []string) int {
		return absSum(input)
	}
}

func part2() puzzle.SolutionFunc {
	return func(input []string) int {
		return multiRepNumber(input)
	}
}

func absSum(lines []string) (result int) {
	xl, xr := extractLeftRightSlicers(lines)

	sort.Ints(xl)
	sort.Ints(xr)

	for i := 0; i < len(xl); i++ {
		result += utils.Abs(xl[i] - xr[i])
	}

	return
}

func extractLeftRightSlicers(lines []string) (xl, xr []int) {
	for _, line := range lines {
		s := strings.Fields(line)
		xl = append(xl, utils.StrToInt(s[0]))
		xr = append(xr, utils.StrToInt(s[1]))
	}

	return
}

func multiRepNumber(lines []string) (result int) {
	xl, xr := extractLeftRightSlicers(lines)

	mapr := mapRepNumber(xr)

	for i := 0; i < len(xl); i++ {
		result += xl[i] * mapr[xl[i]]
	}

	return
}

func mapRepNumber(xr []int) map[int]int {
	result := map[int]int{}

	for _, v := range xr {
		result[v]++
	}

	return result
}
