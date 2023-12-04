package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/stefanlogue/advent-of-code-go/cast"
	"github.com/stefanlogue/advent-of-code-go/util"
)

//go:embed input.txt
var input string

var digitRegex = regexp.MustCompile("[0-9]")

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

func part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		digits := getDigitsFromLine(line)
		sum += digits
	}
	return sum
}

func getDigitsFromLine(line string) int {
	matches := digitRegex.FindAllString(line, -1)
	if len(matches) == 0 {
		return 0
	}
	first := matches[0]
	last := matches[len(matches)-1]
	combined := first + last
	asInt := cast.ToInt(combined)
	return asInt
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		digits := []string{}
		for i, char := range line {
			if unicode.IsDigit(char) && cast.ToInt(string(char)) != 0 && cast.ToInt(string(char)) < 10 {
				digits = append(digits, string(char))
			}
			for j, value := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], value) {
					digits = append(digits, fmt.Sprint(j+1))
				}
			}
		}
		combined := string(digits[0]) + string(digits[len(digits)-1])
		score := cast.ToInt(combined)
		sum += score
	}
	return sum
}
