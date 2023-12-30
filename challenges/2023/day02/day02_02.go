package day02

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
