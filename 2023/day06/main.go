package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/stefanlogue/advent-of-code-go/cast"
	"github.com/stefanlogue/advent-of-code-go/util"
)

//go:embed input.txt
var input string

var getNumberFromLineRegex = regexp.MustCompile(`\d+`)

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

type Race struct {
	time   int
	record int
}

func calculateBounds(time, record int) (int, int) {
	lower := (float64(time) - (math.Sqrt(math.Pow(float64(time), 2) - float64(4*(record))))) / 2
	adjustedLower := int(math.Floor(lower) + 1)
	upper := (float64(time) + (math.Sqrt(math.Pow(float64(time), 2) - float64(4*(record))))) / 2
	adjustedUpper := int(math.Ceil(upper) - 1)
	return adjustedLower, adjustedUpper
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	times := getNumberFromLineRegex.FindAllString(lines[0], -1)
	recordDistances := getNumberFromLineRegex.FindAllString(lines[1], -1)
	var races []Race
	for i := range times {
		races = append(races, Race{
			time:   cast.ToInt(times[i]),
			record: cast.ToInt(recordDistances[i]),
		})
	}
	var waysToWin []int
	for _, race := range races {
		lower, upper := calculateBounds(race.time, race.record)
		waysToWin = append(waysToWin, int(upper-lower+1))
	}
	answer := 1
	for _, way := range waysToWin {
		answer *= way
	}
	return answer
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	times := getNumberFromLineRegex.FindAllString(lines[0], -1)
	recordDistances := getNumberFromLineRegex.FindAllString(lines[1], -1)
	var timeStr string
	var recordStr string
	for _, t := range times {
		timeStr += t
	}
	for _, r := range recordDistances {
		recordStr += r
	}
	time := cast.ToInt(timeStr)
	record := cast.ToInt(recordStr)
	lower, upper := calculateBounds(time, record)
	waysToWin := upper - lower + 1
	return waysToWin
}
