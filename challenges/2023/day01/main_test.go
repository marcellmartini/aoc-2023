package day01_test

import (
	"testing"

	"github.com/marcellmartini/aoc-in-go/challenges/2023/day01"
)

func TestSumOfCalibrations(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test input",
			input:  day01.InputTest,
			expect: day01.AnswerTest,
		},
		{
			desc:   "Full input",
			input:  day01.Input,
			expect: day01.Answer,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := day01.SumOfCalibrations(tC.input)
			if got != tC.expect {
				t.Errorf("SumOfCalibrations(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}

func TestSumOfCalibrationWithWord(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Test input",
			input:  day01.InputTest2,
			expect: day01.AnswerTest,
		},
		{
			desc:   "Full input",
			input:  day01.Input,
			expect: day01.Answer,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := day01.SumOfCalibrationWithWord(tC.input)
			if got != tC.expect {
				t.Errorf("SumOfCalibrationWithWord(): exprect %v, got %v", tC.expect, got)
			}
		})
	}
}
