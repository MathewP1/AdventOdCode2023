package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

func getWinsOptimized(draws, wins []int) int {
	score := 0
    drawMap := [100]int{}
	for _, d := range draws {
		drawMap[d] = 1 
	}

	for _, w := range wins {
		if drawMap[w] == 1 {
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

func Task1(lines []string) int {
    score1 := 0
    for _, line := range lines {
        line =stripGameAndId(line)
        draws, wins := parseLine(line)
        numWon := getWins(draws, wins)
        score1 += int(math.Pow(2.0, float64(numWon-1)))
    }

    return score1
}

func Task1Optimized(lines []string) int {
    score1 := 0
    for _, line := range lines {
        line =stripGameAndId(line)
        draws, wins := parseLine(line)
        numWon := getWinsOptimized(draws, wins)
        score1 += int(math.Pow(2.0, float64(numWon-1)))
    }

    return score1
}

func Task2(lines []string) int {
	score2 := 0
	copies := make([]int, 11)
    for _, line := range lines {
		line = stripGameAndId(line)
		draws, wins := parseLine(line)
		numWon := getWins(draws, wins)
		score2 += 1 + copies[0]

		for i := 1; i < 1+numWon; i++ {
			copies[i] += 1 + copies[0]
		}

		copies = copies[1:]
		copies = append(copies, 0)
	}

    return score2
}



func main() {
	fmt.Println("Hello there!")
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
    lines := []string{}
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    
    score1 := Task1(lines)
    score2 := Task2(lines)


	fmt.Println("Score1: ", score1)
	fmt.Println("Score2: ", score2)
    

    fmt.Println("Starting benchmarks...")
    counts := []int{1, 100, 1000}

    for _, c := range counts {
        sc := 0
        start := time.Now()
        for i := 0; i < c; i++ {
           sc += Task1(lines) 
        }
        duration := time.Since(start)
        fmt.Printf("%d times took %dms\n", c, duration.Milliseconds())

    }
    for _, c := range counts {
        sc := 0
        start := time.Now()
        for i := 0; i < c; i++ {
            sc += Task1Optimized(lines) 
        }
        duration := time.Since(start)
        fmt.Printf("%d times took %dms\n", c, duration.Milliseconds())

    }

}

// 23847 - part 1
