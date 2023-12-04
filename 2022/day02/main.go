package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

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

func part1(input string) int {
	score := 0
	games := strings.Split(input, "\n")
	for _, game := range games {
		moves := strings.Split(game, " ")
		if len(moves) != 2 {
			break
		}
		var player1Move string
		switch moves[0] {
		case "A":
			player1Move = "rock"
		case "B":
			player1Move = "paper"
		case "C":
			player1Move = "scissors"
		}
		var player2Move string
		switch moves[1] {
		case "X":
			player2Move = "rock"
			score += 1
		case "Y":
			player2Move = "paper"
			score += 2
		case "Z":
			player2Move = "scissors"
			score += 3
		}

		if player1Move == player2Move {
			score += 3
		} else if player1Move == "rock" && player2Move == "scissors" {
			continue
		} else if player1Move == "paper" && player2Move == "rock" {
			continue
		} else if player1Move == "scissors" && player2Move == "paper" {
			continue
		} else {
			score += 6
		}

	}

	return score
}

func part2(input string) int {
	score := 0
	games := strings.Split(input, "\n")
	for _, game := range games {
		moves := strings.Split(game, " ")
		if len(moves) != 2 {
			break
		}
		switch moves[1] {
		case "X":
			switch moves[0] {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}
		case "Y":
			switch moves[0] {
			case "A":
				score += 4
			case "B":
				score += 5
			case "C":
				score += 6
			}
		case "Z":
			switch moves[0] {
			case "A":
				score += 8
			case "B":
				score += 9
			case "C":
				score += 7
			}
		}
	}
	return score
}
