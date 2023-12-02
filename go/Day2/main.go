package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 3 cubes: red, green, blue
// random number of cubes of each color in a bag
// each game random cubes are taken out, this is this task's input
// Game id: x1 blue, y1 green, z1 red; ...

// Task: Count the sum of ids of games that are possible, given that in the bag
// there are: 12 red, 13 green and 14 blue cubes.
type handful struct {
	red, green, blue int
}

type game struct {
	id       int
	handfuls []handful
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func extractNum(s string) (int, string) {
	if s[0] == ' ' {
		s = s[1:]
	}
	end := strings.IndexAny(s, " ")
	num, err := strconv.Atoi(s[:end])
	if err != nil {
		panic(err)
	}
	return num, s[end+1:]
}

func parseLine(s string) game {
	s = s[5:] // skip "Game "
	semicolonId := strings.IndexAny(s, ":")
	var g game
	gameId, err := strconv.Atoi(string(s[:semicolonId]))

	if err != nil {
		panic(err)
	}

	g.id = gameId

	s = s[semicolonId+1:] // skip " "
	semicolonSplit := strings.Split(s, ";")
	handfuls := make([]handful, len(semicolonSplit))
	for i, draw := range semicolonSplit {
		commaSplit := strings.Split(draw, ",")
		for _, color := range commaSplit {
			num, color := extractNum(color)
			switch color {
			case "green":
				handfuls[i].green = num
				break
			case "blue":
				handfuls[i].blue = num
				break
			case "red":
				handfuls[i].red = num
				break
			default:
				panic(fmt.Sprintf("Unexpected default case: %s", color))
			}
		}
	}
	g.handfuls = handfuls
	return g
}

func colorsValid(g game, rLimit, gLimit, bLimit int) bool {
	for _, hand := range g.handfuls {
		if hand.red > rLimit || hand.green > gLimit || hand.blue > bLimit {
			return false
		}
	}
	return true
}

func minimumRequired(g game) (red, green, blue int) {
	red, green, blue = 0, 0, 0
	for _, hand := range g.handfuls {
		if hand.red > red {
			red = hand.red
		}
		if hand.green > green {
			green = hand.green
		}
		if hand.blue > blue {
			blue = hand.blue
		}
	}
	return
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	scanner := bufio.NewScanner(f)
	score1, score2 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parseLine(line)
		if colorsValid(game, redLimit, greenLimit, blueLimit) {
			score1 += game.id
		}
		r, g, b := minimumRequired(game)
		score2 += r * g * b
	}

	fmt.Println(score1, score2)
}
