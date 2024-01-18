package day04_test

import (
	"testing"

	"github.com/marcellmartini/aoc-in-go/challenges/2023/day04"
)

func TestSumOfScratchcards(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test Input",
			input:  day04.InputTest,
			expect: day04.AnswerTest,
		},
		{
			desc:   "Full Input",
			input:  day04.Input,
			expect: day04.Answer,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := day04.SumOfScratchcards(tC.input)
			if got != tC.expect {
				t.Errorf("SumOfScratchcardsWorth(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}

func TestSumOfScratchcardsAndCopies(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test Input",
			input:  day04.InputTest,
			expect: day04.AnswerTest,
		},
		{
			desc:   "Full Input",
			input:  day04.Input,
			expect: day04.Answer,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := day04.SumOfScratchcardsTotal(tC.input)
			if got != tC.expect {
				t.Errorf("SumOfScratchcardsWorth(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}
