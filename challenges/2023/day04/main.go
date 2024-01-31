package day04

import (
	"math"
	"regexp"
	"runtime"
	"strings"

	"github.com/marcellmartini/aoc-in-go/puzzle"
)

var _, fp, _, _ = runtime.Caller(0)

var Puzzle = puzzle.NewBuilder().
	ConfigurePWD(fp).
	ConfigureSolutions(part1(), part2()).
	LoadFiles().
	Build()

func part1() puzzle.SolutionFunc {
	return func(input []string) int {
		sc := lineToCards(&input)
		return sumOfPoints(&sc)
	}
}

func part2() puzzle.SolutionFunc {
	return func(input []string) int {
		deck := lineToCards(&input)
		return sumOfPointsAndCopies(&deck)
	}
}

type card struct {
	id              string
	wNumbers        map[string]bool
	HaveNumbers     []string
	countWinNumbers int
	copies          int
}

type deck struct {
	cards      []card
	worthPoint int
}

func sumOfPointsAndCopies2(sc deck) int {
	for _, n := range sc.cards[0].HaveNumbers {
		if _, ok := sc.cards[0].wNumbers[n]; ok {
			sc.cards[0].countWinNumbers++
		}
	}

	// Add the total of copies ...
	for j := 1; j <= sc.cards[0].copies; j++ {
		for k := 1; k <= sc.cards[0].countWinNumbers; k++ {
			sc.cards[0+k].copies++
		}
	}

	if len(sc.cards) == 1 {
		return sc.cards[0].copies
	}

	return sc.cards[0].copies + sumOfPointsAndCopies2(
		deck{
			sc.cards[1:],
			0,
		},
	)
}

func sumOfPointsAndCopies(sc *deck) int {
	// In All deck of scratchcards ...
	for i, card := range sc.cards {
		// Count the number of wins ...
		for _, n := range card.HaveNumbers {
			if _, ok := card.wNumbers[n]; ok {
				sc.cards[i].countWinNumbers++
			}
		}

		// Add the total of copies ...
		for j := 1; j <= sc.cards[i].copies; j++ {
			for k := 1; k <= sc.cards[i].countWinNumbers; k++ {
				sc.cards[i+k].copies++
			}
		}

		// sum the total of copies ...
		sc.worthPoint += sc.cards[i].copies
	}

	// and finaly return the total of scratchcards.
	return sc.worthPoint
}

func sumOfPoints(sc *deck) int {
	for k, card := range sc.cards {
		for _, n := range card.HaveNumbers {
			if _, ok := card.wNumbers[n]; ok {
				sc.cards[k].countWinNumbers++
			}
		}
	}

	for k := range sc.cards {
		if sc.cards[k].countWinNumbers != 0 {
			sc.worthPoint += int(math.Pow(2, float64(sc.cards[k].countWinNumbers-1)))
		}
	}
	return sc.worthPoint
}

func lineToCards(sc *[]string) deck {
	xCard := []card{}
	space := regexp.MustCompile(`\s+`)

	for _, line := range *sc {
		xLine := strings.Split(line, ":")

		// xLine[0] - Card id
		ci := space.ReplaceAllString(xLine[0], " ")
		cardId := strings.Split(ci, " ")

		// xLine[1] - numbers
		n := space.ReplaceAllString(xLine[1], " ")
		xNumbers := strings.Split(n, "|")

		// xNumbers[0] - Winning numbers
		wn := strings.TrimSpace(xNumbers[0])
		xwn := strings.Split(wn, " ")
		mwn := map[string]bool{}

		for _, k := range xwn {
			mwn[k] = true
		}

		// xNumbers[1] - Have numbers
		hn := strings.TrimSpace(xNumbers[1])
		xhn := strings.Split(hn, " ")

		xCard = append(xCard, card{cardId[1], mwn, xhn, 0, 1})
	}

	return deck{xCard, 0}
}
