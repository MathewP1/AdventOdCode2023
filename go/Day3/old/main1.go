package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type numInfo struct {
	pos, length int
	value       int
}

type parsedLine struct {
	numbers    []numInfo
	symbolsPos []int
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func extractNumber(s string, i int) (int, int) {
	l := i
	for l < len(s) && isDigit(s[l]) {
		l++
	}

	numS := s[i:l]
	num, err := strconv.Atoi(numS)
	if err != nil {
		panic(err)
	}
	return num, l - i
}

func parseLine(line string) parsedLine {
	p := parsedLine{}
	for i := 0; i < len(line); {
		char := line[i]
		if char == '.' {
			i++
			continue
		}
		if isDigit(char) {
			num, skip := extractNumber(line, i)
			p.numbers = append(p.numbers, numInfo{i, skip, num})
			i += skip
		} else {
			p.symbolsPos = append(p.symbolsPos, i)
			i++
		}

	}
	return p
}

func contains(s []int, val int) bool {
    for _, a := range s {
        if a == val {
            return true
        }
    }
    return false
}
func containsMap(s []int, values map[int]bool) bool {
    for _, a := range s {
        b := values[a]
        if b == true {
            return true
        }
    }
    return false
}
func fillMap(beg, end int) map[int]bool {
    m := make(map[int]bool)
    for i := beg; i < end; i++ {
        m[i] = true
    }
    return m
}

func countLine(line, above, below parsedLine) int {
    score := 0
    for _, nInfo := range line.numbers {
        fmt.Println("Checking nInfo: ", nInfo.value, nInfo.pos, nInfo.length)
        fmt.Println("current lines symbols: ", line.symbolsPos)
        // check neighbours in the same line
        if contains(line.symbolsPos, nInfo.pos-1) || contains(line.symbolsPos, nInfo.pos + nInfo.length + 1) {
            fmt.Println("has neighbour")
            continue
        }
        // check diagonal
        topCheck := fillMap(nInfo.pos - 1, nInfo.pos + nInfo.length + 1)
        bottomCheck := fillMap(nInfo.pos - 1, nInfo.pos + nInfo.length + 1)
        fmt.Println("top line symbols: ", above.symbolsPos)
        fmt.Println("bottom line symbols: ", below.symbolsPos)
        if containsMap(line.symbolsPos, topCheck) || containsMap(line.symbolsPos, bottomCheck) {
            fmt.Println("has top or bottom")
            continue
        }
        score += nInfo.value
    }
    return score
}

func task1(parsedLines []parsedLine) int {
    score := 0
    score += countLine(parsedLines[0], parsedLine{}, parsedLines[1])
    for i := 1; i < len(parsedLines) - 1; i++ {
        
        score += countLine(parsedLines[i], parsedLines[i-1], parsedLines[i+1])


    }
    score += countLine(parsedLines[len(parsedLines)-1], parsedLines[len(parsedLines)-2], parsedLine{})
    
    return score
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var parsedLines []parsedLine
	for scanner.Scan() {
		line := scanner.Text()
		p := parseLine(line)
		parsedLines = append(parsedLines, p)
	}

    fmt.Println(task1(parsedLines))
}
