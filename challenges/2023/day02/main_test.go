package day02_test

import (
	"testing"

	"github.com/marcellmartini/aoc-in-go/challenges/2023/day02"
)

func TestSumValidGames(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test input",
			input:  day02.InputTest,
			expect: day02.AnswerTest,
		},
		{
			desc:   "Full Input",
			input:  day02.Input,
			expect: day02.Answer,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := day02.SumValidGamesIds(tC.input)
			if got != tC.expect {
				t.Errorf("SumValidGamesIds(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}

func TestSumFewestGamePossible(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test input",
			input:  day02.InputTest,
			expect: day02.AnswerTest,
		},
		{
			desc:   "Full Input",
			input:  day02.Input,
			expect: day02.Answer,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := day02.SumFewestGamePossible(tC.input)
			if got != tC.expect {
				t.Errorf("SumFewestGamePossible(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}
