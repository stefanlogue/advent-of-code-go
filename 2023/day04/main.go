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

var cardNumberRegex = regexp.MustCompile(`(\d+):`)

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

func getNumbersFromString(s string) (numbers map[int]bool) {
	strArray := strings.Split(s, " ")
	numbers = make(map[int]bool)
	for _, str := range strArray {
		if len(str) > 0 && unicode.IsDigit(rune(str[0])) {
			numbers[cast.ToInt(str)] = true
		}
	}
	return numbers
}

func intersection(a, b map[int]bool) (c map[int]bool) {
	c = make(map[int]bool)
	if len(a) > len(b) {
		a, b = b, a
	}
	for key := range a {
		if b[key] {
			c[key] = true
		}
	}
	return c
}

func part1(input string) int {
	sum := 0
	cards := strings.Split(input, "\n")
	for _, card := range cards {
		numbers := strings.Split(card, ": ")[1]
		bothSetsOfNums := strings.Split(numbers, " | ")
		winningNumbers := getNumbersFromString(bothSetsOfNums[0])
		myNumbers := getNumbersFromString(bothSetsOfNums[1])
		matchingNumbers := intersection(winningNumbers, myNumbers)
		if len(matchingNumbers) == 0 {
			continue
		}
		points := 0
		for i := 0; i < len(matchingNumbers); i++ {
			if points == 0 {
				points++
			} else {
				points *= 2
			}
		}
		sum += points
	}
	return sum
}

type Card struct {
	number  int
	copies  int
	matches int
}

func part2(input string) int {
	sum := 0
	var cards []Card
	cardStrings := strings.Split(input, "\n")
	for _, cs := range cardStrings {
		cardNumber := cast.ToInt(cardNumberRegex.FindStringSubmatch(cs)[1])
		numbers := strings.Split(cs, ": ")[1]
		bothSetsOfNums := strings.Split(numbers, " | ")
		winners := getNumbersFromString(bothSetsOfNums[0])
		myNumbers := getNumbersFromString(bothSetsOfNums[1])
		matchingNumbers := intersection(winners, myNumbers)
		matches := len(matchingNumbers)
		card := Card{
			number:  cardNumber,
			copies:  1,
			matches: matches,
		}
		cards = append(cards, card)
	}
	for i, card := range cards {
		for j := 0; j < card.matches; j++ {
			cards[i+j+1].copies += card.copies
		}
		sum += card.copies
	}
	return sum
}
