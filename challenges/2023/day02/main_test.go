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
			desc: "Sample input",
			input: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expect: 8,
		},
		{
			desc:   "Full Input",
			input:  day02.Input,
			expect: 2348,
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
			desc: "Sample input",
			input: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expect: 2286,
		},
		{
			desc:   "Full Input",
			input:  day02.Input,
			expect: 76008,
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
