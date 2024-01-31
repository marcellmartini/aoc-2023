package day01

import (
	"testing"

	"github.com/marcellmartini/aoc-in-go/puzzle"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test input",
			input:  Puzzle.Inputs["input_example"],
			expect: Puzzle.Answers["answer_example"][puzzle.Part1],
		},
		{
			desc:   "Full input",
			input:  Puzzle.Inputs["input"],
			expect: Puzzle.Answers["answer"][puzzle.Part1],
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Puzzle.Solutions[puzzle.Part1](tC.input)
			if got != tC.expect {
				t.Errorf("Part1(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test input",
			input:  Puzzle.Inputs["input_example2"],
			expect: Puzzle.Answers["answer_example"][puzzle.Part2],
		},
		{
			desc:   "Full input",
			input:  Puzzle.Inputs["input"],
			expect: Puzzle.Answers["answer"][puzzle.Part2],
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Puzzle.Solutions[puzzle.Part2](tC.input)
			if got != tC.expect {
				t.Errorf("Part2(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}
