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
	Part1 = iota // Part1 is the index for the solution array that solves Part 1 of the puzzle
	Part2        // Part2 is the index for the solution array that solves Part 2 of the puzzle
)

// SolutionFunc is a type that represents a function to solve the puzzle
type SolutionFunc func([]string) int

// Solution is an array of functions that solve the puzzle.
// The first element solves Part 1.
// The second element solves Part 2.
type Solution []SolutionFunc

// Input is a map that stores input data for the puzzle
type Input map[string][]string

// Answer is a map that stores the expected answers for the puzzle
type Answer map[string][]int

// Puzzle is the main structure that stores the puzzle's inputs, answers, and solutions
type Puzzle struct {
	PWD       string   // Working directory where the files are located
	Inputs    Input    // Input data loaded from files
	Answers   Answer   // Expected answers loaded from files
	Solutions Solution // Array of functions to solve the puzzle
}

// PuzzleBuilder is a helper structure to construct a Puzzle object
type PuzzleBuilder struct {
	Puzzle Puzzle // Instance of the puzzle being configured
}

// NewBuilder creates and returns a new PuzzleBuilder instance with default settings
func NewBuilder() PuzzleBuilder {
	return PuzzleBuilder{
		Puzzle: Puzzle{
			PWD:       "/",        // Default directory is the root
			Inputs:    Input{},    // Initializes an empty input map
			Answers:   Answer{},   // Initializes an empty answer map
			Solutions: Solution{}, // Initializes an empty solution list
		},
	}
}

// Build returns the configured Puzzle
func (pb PuzzleBuilder) Build() Puzzle {
	return pb.Puzzle
}

// ConfigureSolutions adds solution functions to the Puzzle
func (pb PuzzleBuilder) ConfigureSolutions(sf ...SolutionFunc) PuzzleBuilder {
	pb.Puzzle.Solutions = append(pb.Puzzle.Solutions, sf...)
	return pb
}

// ConfigurePWD sets the working directory for the Puzzle
func (pb PuzzleBuilder) ConfigurePWD(pwd string) PuzzleBuilder {
	pb.Puzzle.PWD = filepath.Dir(pwd)
	return pb
}

// LoadFiles loads default and additional files into the Puzzle
func (pb PuzzleBuilder) LoadFiles(files ...string) PuzzleBuilder {
	// List of default files to be loaded
	defaultFiles := []string{"input", "input_example", "answer", "answer_example"}
	defaultFiles = append(defaultFiles, files...) // Adds extra files specified by the user

	// Loads the content of each file
	for _, file := range defaultFiles {
		pb.fetchFile(file)
	}
	return pb
}

// fetchFile reads the file content and processes it based on its type (input or answer)
// TODO: Refactor to use io.Reader instead of directly reading from files.
// This would allow greater flexibility for testing and working with different data sources.
func (pb PuzzleBuilder) fetchFile(fileName string) {
	content, err := pb.readFile(fileName) // Reads the file
	utils.CheckError(err)

	switch {
	// If the file is an input file
	case strings.Contains(fileName, "input"):
		pb.Puzzle.Inputs[fileName] = content

	// If the file is an answer file
	case strings.Contains(fileName, "answer"):
		for _, answer := range content {
			// Converts the answer to an integer and adds it to the map
			i, err := strconv.Atoi(answer)
			utils.CheckError(err)
			pb.Puzzle.Answers[fileName] = append(pb.Puzzle.Answers[fileName], i)
		}
	}
}

// readFile reads the content of a file line by line
func (pb PuzzleBuilder) readFile(fileName string) ([]string, error) {
	var file *os.File
	var xs []string

	// Constructs the full file path
	filePath := filepath.Join(pb.Puzzle.PWD, fileName+".txt")

	// Checks if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		// If the file does not exist, creates it
		file, err = os.Create(filePath)
		utils.CheckError(err)
		defer file.Close()

		// Writes default values for answer files
		if strings.Contains(fileName, "answer") {
			fmt.Fprintf(file, "%d\n%d", -1, -1)
		}
	}

	// Opens the file for reading
	file, err = os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	// Reads the file line by line and adds it to the slice xs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xs = append(xs, scanner.Text())
	}
	return xs, nil
}
