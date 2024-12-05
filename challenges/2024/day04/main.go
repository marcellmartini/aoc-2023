package day04

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
		return countXmas(input)
	}
}

func part2() puzzle.SolutionFunc {
	return func(input []string) int {
		return countX_mas(input)
	}
}

//   0 1 2 3
// 0 X . M .
// 1 . A . .
// 2 X . M .
// 3 . . . .

func countXmas(input []string) int {
	answer := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if j < len(input[0])-3 && input[i][j] == 'X' && input[i][j+1] == 'M' &&
				input[i][j+2] == 'A' &&
				input[i][j+3] == 'S' {
				answer++
			}
			if j >= 3 && input[i][j] == 'X' && input[i][j-1] == 'M' && input[i][j-2] == 'A' &&
				input[i][j-3] == 'S' {
				answer++
			}
			if i < len(input[0])-3 && input[i][j] == 'X' && input[i+1][j] == 'M' &&
				input[i+2][j] == 'A' &&
				input[i+3][j] == 'S' {
				answer++
			}
			if i >= 3 && input[i][j] == 'X' && input[i-1][j] == 'M' && input[i-2][j] == 'A' &&
				input[i-3][j] == 'S' {
				answer++
			}
			if i < len(input[0])-3 && j < len(input[0])-3 && input[i][j] == 'X' &&
				input[i+1][j+1] == 'M' &&
				input[i+2][j+2] == 'A' &&
				input[i+3][j+3] == 'S' {
				answer++
			}
			if i >= 3 && j >= 3 && input[i][j] == 'X' &&
				input[i-1][j-1] == 'M' &&
				input[i-2][j-2] == 'A' &&
				input[i-3][j-3] == 'S' {
				answer++
			}
			if i >= 3 && j < len(input[0])-3 && input[i][j] == 'X' &&
				input[i-1][j+1] == 'M' &&
				input[i-2][j+2] == 'A' &&
				input[i-3][j+3] == 'S' {
				answer++
			}
			if i < len(input[0])-3 && j >= 3 && input[i][j] == 'X' &&
				input[i+1][j-1] == 'M' &&
				input[i+2][j-2] == 'A' &&
				input[i+3][j-3] == 'S' {
				answer++
			}
		}
	}
	return answer
}

func countX_mas(input []string) int {
	answer := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if i < len(input)-2 && j < len(input)-2 &&
				input[i][j] == 'M' &&
				input[i][j+2] == 'M' &&
				input[i+1][j+1] == 'A' &&
				input[i+2][j] == 'S' &&
				input[i+2][j+2] == 'S' {
				answer++
			}
			if i < len(input)-2 && j < len(input)-2 &&
				input[i][j] == 'S' &&
				input[i][j+2] == 'M' &&
				input[i+1][j+1] == 'A' &&
				input[i+2][j] == 'S' &&
				input[i+2][j+2] == 'M' {
				answer++
			}
			if i < len(input)-2 && j < len(input)-2 &&
				input[i][j] == 'S' &&
				input[i][j+2] == 'S' &&
				input[i+1][j+1] == 'A' &&
				input[i+2][j] == 'M' &&
				input[i+2][j+2] == 'M' {
				answer++
			}
			if i < len(input)-2 && j < len(input)-2 &&
				input[i][j] == 'M' &&
				input[i][j+2] == 'S' &&
				input[i+1][j+1] == 'A' &&
				input[i+2][j] == 'M' &&
				input[i+2][j+2] == 'S' {
				answer++
			}

		}
	}
	return answer
}
