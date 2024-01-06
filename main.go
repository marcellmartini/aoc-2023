package main

import (
	"fmt"

	"github.com/marcellmartini/aoc-in-go/challenges/2023/day01"
	"github.com/marcellmartini/aoc-in-go/challenges/2023/day02"
	"github.com/marcellmartini/aoc-in-go/challenges/2023/day03"
	"github.com/marcellmartini/aoc-in-go/challenges/2023/day04"
)

func main() {
	fmt.Printf("Day 01:\n")
	fmt.Printf("\tPart 01, SumOfCalibrations: %d\n", day01.SumOfCalibrations(day01.Input))
	fmt.Printf(
		"\tPart 02, SumOfCalibrationWithWord: %d\n",
		day01.SumOfCalibrationWithWord(day01.Input),
	)

	fmt.Printf("Day 02:\n")
	fmt.Printf("\tPart 01, SumValidGamesIds: %d\n", day02.SumValidGamesIds(day02.Input))
	fmt.Printf("\tPart 02, SumFewestGamePossible: %d\n", day02.SumFewestGamePossible(day02.Input))

	fmt.Printf("Day 03:\n")
	fmt.Printf("\tPart 01, SumOfAllPartNumbers: %d\n", day03.SumOfAllPartNumbers(day03.Input))
	fmt.Printf("\tPart 01, SumAllGearRations: %d\n", day03.SumAllGearRations(day03.Input))

	fmt.Printf("Day 04:\n")
	fmt.Printf("\tPart 01, SumOfScratchcards: %d\n", day04.SumOfScratchcards(day04.Input))
	fmt.Printf("\tPart 02, SumOfScratchcardsTotal: %d\n", day04.SumOfScratchcardsTotal(day04.Input))
}
