package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/stefanlogue/advent-of-code-go/cast"
	"github.com/stefanlogue/advent-of-code-go/util"
)

//go:embed input.txt
var input string

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

type Elf struct {
	number   int
	calories int
}

func populateCaloriesSlice(lines []string) []int {
	var caloriesSlice []int
	calories := 0
	for _, line := range lines {
		if line == "\n" || line == "" {
			caloriesSlice = append(caloriesSlice, calories)
			calories = 0
			continue
		} else {
			calories += cast.ToInt(line)
		}
	}
	return caloriesSlice
}

func part1(input string) int {
	caloriesSlice := populateCaloriesSlice(strings.Split(input, "\n"))

	return slices.Max(caloriesSlice)
}

func part2(input string) int {
	caloriesSlice := populateCaloriesSlice(strings.Split(input, "\n"))
	slices.Sort(caloriesSlice)
	sum := 0
	for i := 0; i < 3; i++ {
		sum += caloriesSlice[len(caloriesSlice)-1-i]
	}
	return sum
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}
