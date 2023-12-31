package day01_test

import (
	"day01"
	"testing"
)

func TestSumOfCalibrations(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect int
	}{
		{
			desc:   "Just one entry",
			input:  []string{"1abc2"},
			expect: 12,
		},
		{
			desc:   "Just one entry 2",
			input:  []string{"12abc2"},
			expect: 12,
		},
		{
			desc:   "An array entry",
			input:  []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
			expect: 142,
		},
		{
			desc:   "All input",
			input:  day01.Input,
			expect: 55477,
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
			desc:   "Just one entry",
			input:  []string{"two1nine"},
			expect: 29,
		},
		{
			desc:   "An array entry",
			input:  []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"},
			expect: 281,
		},
		{
			desc:   "All input",
			input:  day01.Input,
			expect: 54431,
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
