package puzzle

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/marcellmartini/aoc-in-go/utils"
)

const (
	Part1 = iota // Part1 is the first element of Solution array
	Part2        // Part2 is the second element of Solution array
)

// SolutionFunc is the function to solve the puzzle
type SolutionFunc func([]string) int

// Solution is an array of solutions that solve the puzzle.
// the first element solve the puzzle Part 1.
// the second element solve the puzzle Part 2.
type Solution []SolutionFunc

type Input map[string][]string

type Answer map[string][]int

type Puzzle struct {
	PWD       string
	Inputs    Input
	Answers   Answer
	Solutions Solution
}

type PuzzleBuilder struct {
	Puzzle Puzzle
}

func NewBuilder() PuzzleBuilder {
	return PuzzleBuilder{
		Puzzle: Puzzle{
			PWD:       "/",
			Inputs:    Input{},
			Answers:   Answer{},
			Solutions: Solution{},
		},
	}
}

func (pb PuzzleBuilder) Build() Puzzle {
	return pb.Puzzle
}

func (pb PuzzleBuilder) ConfigureSolutions(sf ...SolutionFunc) PuzzleBuilder {
	pb.Puzzle.Solutions = append(pb.Puzzle.Solutions, sf...)
	return pb
}

func (pb PuzzleBuilder) ConfigurePWD(pwd string) PuzzleBuilder {
	pb.Puzzle.PWD = filepath.Dir(pwd)
	return pb
}

func (pb PuzzleBuilder) LoadFiles(files ...string) PuzzleBuilder {
	defaultFiles := []string{"input", "input_example", "answer", "answer_example"}
	defaultFiles = append(defaultFiles, files...)

	for _, file := range defaultFiles {
		pb.fetchFile(file)
	}
	return pb
}

// TODO can be change to io.Reader
func (pb PuzzleBuilder) fetchFile(fileName string) {
	content, err := pb.readFile(fileName)
	utils.CheckError(err)

	switch {
	case strings.Contains(fileName, "input"):
		pb.Puzzle.Inputs[fileName] = content

	case strings.Contains(fileName, "answer"):
		for _, answer := range content {
			i, err := strconv.Atoi(answer)
			utils.CheckError(err)
			pb.Puzzle.Answers[fileName] = append(pb.Puzzle.Answers[fileName], i)
		}
	}
}

func (pb PuzzleBuilder) readFile(fileName string) ([]string, error) {
	var file *os.File
	var xs []string

	filePath := filepath.Join(pb.Puzzle.PWD, fileName+".txt")

	_, err := os.Stat(filePath)
	if err != nil {
		file, err = os.Create(filePath)
		utils.CheckError(err)
		defer file.Close()

		if strings.Contains(fileName, "answer") {
			fmt.Fprintf(file, "%d\n%d", -1, -1)
			return []string{"-1", "-1"}, nil
		}
	}

	file, err = os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xs = append(xs, scanner.Text())
	}
	return xs, nil
}
