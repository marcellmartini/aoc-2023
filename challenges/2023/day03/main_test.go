package day03_test

import (
	"testing"

	"github.com/marcellmartini/aoc-in-go/challenges/2023/day03"
)

func TestSumOfAllPartNumbers(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test Input",
			input:  day03.InputTest,
			expect: day03.AnswerTest,
		},
		{
			desc:   "Full Input",
			input:  day03.Input,
			expect: day03.Answer,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := day03.SumOfAllPartNumbers(tC.input)
			if got != tC.expect {
				t.Errorf("part1(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}

func TestSumAllGearRations(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test Input",
			input:  day03.InputTest,
			expect: day03.AnswerTest,
		},
		{
			desc:   "Full Input",
			input:  day03.Input,
			expect: day03.Answer,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := day03.SumAllGearRations(tC.input)
			if got != tC.expect {
				t.Errorf("part1(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}
