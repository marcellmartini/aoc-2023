package day05

import (
	"math"
	"runtime"
	"strconv"
	"strings"

	"github.com/marcellmartini/aoc-in-go/puzzle"
	. "github.com/marcellmartini/aoc-in-go/utils"
)

var _, fp, _, _ = runtime.Caller(0)

var Puzzle = puzzle.NewBuilder().
	ConfigurePWD(fp).
	ConfigureSolutions(part1(), part2()).
	LoadFiles().
	Build()

func part1() puzzle.SolutionFunc {
	return func(input []string) int {
		a := Almanac{}
		return a.lowestLocation(&input, puzzle.Part1)
	}
}

// TODO: find better solution
func part2() puzzle.SolutionFunc {
	return func(input []string) int {
		return 0
	}
}

type Translate struct {
	Dest   int
	Orig   int
	Length int
}

type Map []Translate

type Almanac struct {
	Seeds []int
	Maps  []Map
}

func (a *Almanac) lowestLocation(input *[]string, part int) int {
	a.loadAlmanac(input, part)
	return a.result()
}

func (a *Almanac) loadAlmanac(input *[]string, part int) {
	splitLine := strings.Split((*input)[0], ":")

	for _, v := range strings.Fields(splitLine[1]) {
		i, err := strconv.Atoi(v)
		CheckError(err)
		a.Seeds = append(a.Seeds, i)
	}

	*input = (*input)[1:]

	a.loadMaps(input)
}

func (a *Almanac) loadMaps(input *[]string) {
	i := -1
	a.Maps = make([]Map, 7)

	for _, line := range *input {
		if line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			i++
			continue
		}

		fields := strings.Fields(line)

		dest, err := strconv.Atoi(fields[0])
		CheckError(err)

		orig, err := strconv.Atoi(fields[1])
		CheckError(err)

		length, err := strconv.Atoi(fields[2])
		CheckError(err)

		a.Maps[i] = append(a.Maps[i], Translate{
			Dest:   dest,
			Orig:   orig,
			Length: length,
		})
	}
}

func (a *Almanac) result() int {
	result := math.MaxInt

	for _, s := range a.Seeds {
		for _, Map := range a.Maps {
			for _, t := range Map {
				if s >= t.Orig && s < t.Orig+t.Length {
					s = t.Dest + (s - t.Orig)
					break
				}
			}
		}

		if s < result {
			result = s
		}
	}

	return result
}
