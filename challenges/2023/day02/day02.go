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

func SumValidGamesIds(games []string) int {
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

func SumFewestGamePossible(games []string) int {
	result := 0
	red, green, blue := 0, 0, 0

	for _, g := range games {
		red, green, blue = fewestGamePossible(g)
		result += red * green * blue
	}
	return result
}

func fewestGamePossible(g string) (mr, mg, mb int) {
	game := lineToGame(g)

	for _, s := range game.sets {
		if s.red > mr && s.red != 0 {
			mr = s.red
		}
		if s.green > mg && s.green != 0 {
			mg = s.green
		}
		if s.blue > mb && s.blue != 0 {
			mb = s.blue
		}
	}
	return
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
		r, g, b := getSetOFValues(colors)
		sets = append(sets, set{r, g, b})
	}

	return game{
		gameId,
		sets,
	}
}

func getSetOFValues(parts []string) (int, int, int) {
	colors := map[string]int{"red": 0, "green": 0, "blue": 0}

	for _, part := range parts {
		p := strings.TrimSpace(part)
		color := strings.Split(p, " ")

		// color[0]: value of color
		val, err := strconv.Atoi(color[0])
		if err != nil {
			panic(err)
		}

		// color[1]: color that have a value
		colors[color[1]] = val
	}

	return colors["red"], colors["green"], colors["blue"]
}
