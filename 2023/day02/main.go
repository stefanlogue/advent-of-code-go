package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/stefanlogue/advent-of-code-go/cast"
	"github.com/stefanlogue/advent-of-code-go/util"
)

//go:embed input.txt
var input string

var (
	redRegex   = regexp.MustCompile(`(\d+) red`)
	greenRegex = regexp.MustCompile(`(\d+) green`)
	blueRegex  = regexp.MustCompile(`(\d+) blue`)
)

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func numberOfCubes(input string) (red int, green int, blue int) {
	var numRed int
	var numGreen int
	var numBlue int
	if redStrings := redRegex.FindStringSubmatch(input); len(redStrings) > 0 {
		numRed = cast.ToInt(redStrings[1])
	}
	if greenStrings := greenRegex.FindStringSubmatch(input); len(greenStrings) > 0 {
		numGreen = cast.ToInt(greenStrings[1])
	}
	if blueStrings := blueRegex.FindStringSubmatch(input); len(blueStrings) > 0 {
		numBlue = cast.ToInt(blueStrings[1])
	}
	return numRed, numGreen, numBlue
}

func part1(input string) int {
	maxRed, maxGreen, maxBlue := 12, 13, 14
	games := strings.Split(input, "\n")
	sum := 0
	for id, game := range games {
		checks := strings.Split(game, "; ")
		gameIsValid := true
		for _, check := range checks {
			numRed, numGreen, numBlue := numberOfCubes(check)
			if numRed+numGreen+numBlue == 0 {
				gameIsValid = false
			}
			if numRed > maxRed || numGreen > maxGreen || numBlue > maxBlue {
				gameIsValid = false
				break
			}
		}
		if gameIsValid {
			sum += id + 1
		}
	}

	return sum
}

func part2(input string) int {
	games := strings.Split(input, "\n")
	sum := 0
	for _, game := range games {
		minRed, minGreen, minBlue := 0, 0, 0
		for _, check := range strings.Split(game, "; ") {
			numRed, numGreen, numBlue := numberOfCubes(check)
			if numRed > minRed {
				minRed = numRed
			}
			if numBlue > minBlue {
				minBlue = numBlue
			}
			if numGreen > minGreen {
				minGreen = numGreen
			}
		}
		power := minRed * minGreen * minBlue
		sum += power
	}
	return sum
}
