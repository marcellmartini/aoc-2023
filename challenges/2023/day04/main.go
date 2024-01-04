package day04

import (
	"math"
	"regexp"
	"strings"
)

type card struct {
	id              string
	wNumbers        map[string]bool
	HaveNumbers     []string
	countWinNumbers int
}

type deck struct {
	cards      []card
	worthPoint int
}

func SumOfScratchcards(scratchcards []string) int {
	sc := lineToCards(&scratchcards)
	return sumOfPoints(&sc)
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

		xCard = append(xCard, card{cardId[1], mwn, xhn, 0})
	}

	return deck{xCard, 0}
}
