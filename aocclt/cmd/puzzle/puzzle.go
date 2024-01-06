package puzzle

import (
	"github.com/spf13/cobra"

	"github.com/marcellmartini/aoc-in-go/aocclt/pkg/puzzle"
)

var CmdPuzzle = &cobra.Command{
	Use:   "puzzle [options]",
	Short: "Download puzzle text.",
	Long: `download puzzle text and save it into a txt file in the path 
indicate bye -p. download the day (-d)  and the year (-y)
ex.:

aocclt puzzle -d 05 -y 2023 -p $(pwd)
`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://adventofcode.com/2023/day/1"
		puzzle.GetPuzzle(url)
	},
}
