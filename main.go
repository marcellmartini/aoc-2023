package main

import (
	"day02"
	"fmt"
)

func main() {
	fmt.Printf("Day 02:\n")
	fmt.Printf("\tPart 01, SumValidGamesIds: %d\n", day02.SumValidGamesIds(day02.Input))
	fmt.Printf("\tPart 02, SumFewestGamePossible: %d\n", day02.SumFewestGamePossible(day02.Input))
}
