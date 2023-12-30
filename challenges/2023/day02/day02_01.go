package day02

import (
	"strconv"
	"strings"
)

type set struct {
	red   int
	green int
	blue  int
}

type bag set

type game struct {
	id   int
	sets []set
}

func AnserDay02_01(games []string) int {
	bag := bag{12, 13, 14}

	result := 0
	for _, g := range games {
		if ok, gameId := isValidGame(g, bag); ok {
			result += gameId
		}
	}
	return result
}

func isValidGame(g string, b bag) (bool, int) {
	game := lineToGame(g)
	valid := true

	for _, set := range game.sets {
		if set.red > b.red ||
			set.green > b.green ||
			set.blue > b.blue {
			valid = false
		}
	}
	return valid, game.id
}

func lineToGame(line string) game {
	// xLine[0] - Game id
	xLine := strings.Split(line, ":")
	gameId, err := strconv.Atoi(strings.SplitAfter(xLine[0], " ")[1])
	if err != nil {
		panic(err)
	}

	// xLine[1] - sets
	xLineSets := strings.Split(xLine[1], ";")
	sets := []set{}

	for _, lineSet := range xLineSets {
		colors := strings.Split(lineSet, ",")
		r, g, b := getSetValues(colors)
		sets = append(sets, set{r, g, b})
	}

	return game{
		gameId,
		sets,
	}
}

func getSetValues(parts []string) (int, int, int) {
	r, g, b := 0, 0, 0

	for _, part := range parts {
		p := strings.TrimSpace(part)
		color := strings.Split(p, " ")

		val, err := strconv.Atoi(color[0])
		if err != nil {
			panic(err)
		}

		switch color[1] {
		case "red":
			r = val
		case "green":
			g = val
		case "blue":
			b = val
		}
	}

	return r, g, b
}
