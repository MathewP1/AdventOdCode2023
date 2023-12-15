package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseExpectedGroups(groupsString string) []int {
	split := strings.Split(groupsString, ",")
	out := make([]int, 0, len(split))
	for _, s := range split {
		value, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		out = append(out, value)
	}
	return out
}

type Params struct {
	input  []byte
	i      int // index in input
	groups []int
	g      int // index in groups
}

func CountPossibilities(params Params) int {
	input := params.input
	i := params.i
	groups := params.groups
	g := params.g
    fmt.Printf("Params: input: %s, i: %d, groups: %v, g: %d\n", input, i, groups, g)
	if i == len(input) {
		// base case reached, that means this string is valid
		return 1
    } else if(strings.IndexAny(string(input[i:]), "#?") == -1) {
        fmt.Println("No need to check more!")
        return 0
    }
    // TODO: else if no more # in string


	if input[i] == '?' {
		count := 0
		input[i] = '#'
		count += CountPossibilitiesCached(Params{input, i, groups, g})
		input[i] = '.'
		count += CountPossibilitiesCached(Params{input, i, groups, g})
		return count
	}

	if input[i] == '#' {
		// TODO: implement check
	} else if input[i] == '.' {
		// TODO: implement check

	} else {
		panic("unxpected input")
	}
	return CountPossibilitiesCached(Params{input, i + 1, groups, g})
}

func CountPossibilitiesCached(params Params) int {
	return CountPossibilities(params)
}

func main() {
	fmt.Println("Hello there!")

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scores := make(chan int)
	i := 0
	for ; scanner.Scan(); i++ {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		damagedInput := lineSplit[0]
		expectedGroups := parseExpectedGroups(lineSplit[1])
		fmt.Println("Damaged input: ", damagedInput)
		fmt.Println("Groups: ", expectedGroups)
		go func() {
			scores <- CountPossibilitiesCached(Params{[]byte(damagedInput), 0, expectedGroups, 0})
		}()
	}

	overall_score := 0
	for j := 0; j < i; j++ {
		overall_score += <-scores
	}
	close(scores)
	fmt.Println("Score: ", overall_score)

}
