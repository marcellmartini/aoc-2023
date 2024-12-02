package day0X

import (
	"runtime"

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
		value := 0
		//
		// ADD YOUR CODE HERE
		//
		return value
	}
}

func part2() puzzle.SolutionFunc {
	return func(input []string) int {
		value := 0
		//
		// ADD YOUR CODE HERE
		//
		return value
	}
}

//
// ADD YOUR CODE HERE
//
