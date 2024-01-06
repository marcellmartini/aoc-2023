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
			"Sample Input 1",
			[]string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			4361,
		},
		{
			desc: "Sample Input 2",
			input: []string{
				"...*.../..",
				"...35...63",
			},
			expect: 98,
		},
		{
			desc:   "Full Input",
			input:  day03.Input,
			expect: 540025,
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
			"Sample Input 1",
			[]string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			467835,
		},
		{
			desc:   "Full Input",
			input:  day03.Input,
			expect: 84584891,
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
