package main

import (
	"fmt"

	"github.com/marcellmartini/aoc-in-go/challenges/2023/day01"
	"github.com/marcellmartini/aoc-in-go/challenges/2023/day02"
	"github.com/marcellmartini/aoc-in-go/challenges/2023/day03"
	"github.com/marcellmartini/aoc-in-go/challenges/2023/day04"
	"github.com/marcellmartini/aoc-in-go/puzzle"
)

func main() {
	fmt.Printf("Day 01:\n")
	fmt.Printf(
		"\tPart 01: %d\n",
		day01.Puzzle.Solutions[puzzle.Part1](day01.Puzzle.Inputs["input"]),
	)
	fmt.Printf(
		"\tPart 02: %d\n",
		day01.Puzzle.Solutions[puzzle.Part2](day01.Puzzle.Inputs["input"]),
	)

	fmt.Printf("Day 02:\n")
	fmt.Printf(
		"\tPart 01: %d\n",
		day02.Puzzle.Solutions[puzzle.Part1](day02.Puzzle.Inputs["input"]),
	)
	fmt.Printf(
		"\tPart 02: %d\n",
		day02.Puzzle.Solutions[puzzle.Part2](day02.Puzzle.Inputs["input"]),
	)

	fmt.Printf("Day 03:\n")
	fmt.Printf("\tPart 01, SumOfAllPartNumbers: %d\n", day03.SumOfAllPartNumbers(day03.Input))
	fmt.Printf("\tPart 01, SumAllGearRations: %d\n", day03.SumAllGearRations(day03.Input))

	fmt.Printf("Day 04:\n")
	fmt.Printf("\tPart 01, SumOfScratchcards: %d\n", day04.SumOfScratchcards(day04.Input))
	fmt.Printf("\tPart 02, SumOfScratchcardsTotal: %d\n", day04.SumOfScratchcardsTotal(day04.Input))
}
