package day03

import (
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/marcellmartini/aoc-in-go/puzzle"
)

var _, fp, _, _ = runtime.Caller(0)

var Puzzle = puzzle.NewBuilder().
	ConfigurePWD(fp).
	ConfigureSolutions(part1(), part2()).
	LoadFiles("input_example2").
	Build()

func part1() puzzle.SolutionFunc {
	return func(input []string) int {
		return multiSomeNumbers(input, pattern)
	}
}

func part2() puzzle.SolutionFunc {
	return func(input []string) int {
		return multiSomeNumbers2(input, pattern)
	}
}

var pattern = regexp.MustCompile(`^mul\((?P<left>\d{1,3}),(?P<right>\d{1,3})\)`)

func multiSomeNumbers(input []string, pattern *regexp.Regexp) int {
	value := 0

	for _, v := range input {
		for i := 0; i < len(v); i++ {
			opa := pattern.FindStringSubmatch(v[i:])
			value += calculate(opa, pattern, &i)
		}
	}
	return value
}

func multiSomeNumbers2(input []string, pattern *regexp.Regexp) int {
	enable := true
	value := 0

	for _, v := range input {
		for i := 0; i < len(v); i++ {
			if strings.HasPrefix(v[i:], "do()") {
				enable = true
				i += 3
				continue
			}

			if strings.HasPrefix(v[i:], "don't()") {
				enable = false
				i += 6
				continue
			}

			opa := pattern.FindStringSubmatch(v[i:])
			if enable {
				value += calculate(opa, pattern, &i)
			}

		}
	}
	return value
}

func calculate(s []string, pattern *regexp.Regexp, i *int) int {
	if s != nil {
		il := pattern.SubexpIndex("left")
		ir := pattern.SubexpIndex("right")

		l, _ := strconv.Atoi(s[il])
		r, _ := strconv.Atoi(s[ir])

		*i += len(s[0]) - 1
		return l * r
	}
	return 0
}
