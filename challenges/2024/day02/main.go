package day02

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
		return numSafeReports(input)
	}
}

func part2() puzzle.SolutionFunc {
	return func(input []string) int {
		return numSafeReportsWithTolerate(input)
	}
}

type safeFunc func(xi sort.IntSlice) bool

func numSafeReports(lines []string) int {
	reports := extractINTSlicer(lines, len(lines))

	return countSafeReports(reports, isSafe)
}

func extractINTSlicer(lines []string, len int) []sort.IntSlice {
	xi := make([]sort.IntSlice, len)

	for i, line := range lines {
		for _, field := range strings.Fields(line) {
			xi[i] = append(xi[i], utils.StrToInt(field))
		}
	}
	return xi
}

func countSafeReports(report []sort.IntSlice, sf safeFunc) int {
	result := 0

	for _, r := range report {
		if sf(r) {
			result++
		}
	}
	return result
}

func isSafe(r sort.IntSlice) bool {
	inc_or_desc := sort.IsSorted(r) || sort.IsSorted(sort.Reverse(r))

	valid := true
	for i := 0; i < r.Len()-1; i++ {
		diff := utils.Abs(r[i] - r[i+1])

		if diff < 1 || diff > 3 {
			valid = false
		}
	}
	return inc_or_desc && valid
}

func numSafeReportsWithTolerate(lines []string) int {
	reports := extractINTSlicer(lines, len(lines))

	return countSafeReports(reports, isSafeWithTolerate)
}

func isSafeWithTolerate(r sort.IntSlice) bool {
	result := false

	for j := 0; j < r.Len(); j++ {
		ret := utils.RemoveIndex(r, j)
		if isSafe(ret) {
			result = true
		}
	}
	return result
}
