package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getWins(draws, wins []int) int {
	score := 0
	drawMap := make(map[int]bool)
	for _, d := range draws {
		drawMap[d] = true
	}

	for _, w := range wins {
		if drawMap[w] == true {
			score += 1
		}
	}

	return score
}

func stripGameAndId(s string) string {
	i := strings.IndexAny(s, ":")
	s = strings.Trim(s[i+1:], " ")
	return s
}

func parseLine(s string) ([]int, []int) {

	sep := strings.IndexAny(s, "|")
	drawsString, winsString := s[:sep], s[sep+1:]
	drawsString, winsString = strings.Trim(drawsString, " "), strings.Trim(winsString, " ")
	drawsSplit := strings.Fields(drawsString)
	draws := make([]int, len(drawsSplit))
	for i, s := range drawsSplit {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		draws[i] = num
	}

	winsSplit := strings.Fields(winsString)
	wins := make([]int, len(winsSplit))
	for i, s := range winsSplit {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		wins[i] = num
	}

	return draws, wins
}

func main() {
	fmt.Println("Hello there!")
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	score1 := 0
	score2 := 0
	copies := make([]int, 11)
	for scanner.Scan() {
		line := scanner.Text()
		line = stripGameAndId(line)
		draws, wins := parseLine(line)
		numWon := getWins(draws, wins)
		score1 += int(math.Pow(2.0, float64(numWon-1)))
		score2 += 1 + copies[0]

		for i := 1; i < 1+numWon; i++ {
			copies[i] += 1 + copies[0]
		}

		copies = copies[1:]
		copies = append(copies, 0)
	}
	fmt.Println("Score1: ", score1)
	fmt.Println("Score2: ", score2)
}

// 23847 - part 1
