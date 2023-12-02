package main

import "fmt"
import "bufio"
import "os"

func parseLine1(s string) int {
	var number int
	firstFound := false
	lastCandidate := 0
	for i := 0; i < len(s); i++ {
		c := int(s[i]) - int('0')
		if c < 0 || c > 10 {
			continue
		}
		if !firstFound {
			number = c * 10
			firstFound = true
		}
		lastCandidate = c
	}
	return number + lastCandidate
}

type digit string

const (
	zero  digit = "zero"
	one   digit = "one"
	two   digit = "two"
	three digit = "three"
	four  digit = "four"
	five  digit = "five"
	six   digit = "six"
	seven digit = "seven"
	eight digit = "eight"
	nine  digit = "nine"
)

var digits []digit = []digit{zero, one, two, three, four, five, six, seven, eight, nine}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func findSubstrDigit(s string) (int, bool) {
	for i, d := range digits {
		if len(string(d)) > len(s) {
			continue
		}
		if string(d) == s[:len(d)] {
			return i, true
		}
	}
	return 0, false
}

func parseLine2(s string) int {
	first := 0
	firstFound := false
	candidate := 0
	fmt.Println("Parsing line: ", s)
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			candidate = int(s[i] - '0')
		} else if d, ok := findSubstrDigit(s[i:]); ok {
			candidate = d
		} else {
			continue
		}
		if !firstFound {
			first = candidate
			firstFound = true
		}
	}
	fmt.Printf("Found first %d and last %d\n", first, candidate)
	return first*10 + candidate
}
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		score += parseLine2(text)
	}
	fmt.Println("Score (method 2): ", score)

}
