package main

import "fmt"
import "os"
import "bufio"
import "strconv"

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

// return true if has adjacent symbols
func checkNum(start, end int, above, current, below string) bool {
	if start-1 < 0 {
		start = 0
	} else {
		start = start - 1
	}
	if above != "" {
		for i := start; i < end+1 && i < len(above); i++ {
			if !isDigit(above[i]) && above[i] != '.' {
				return true
			}
		}
	}
	if !isDigit(current[start]) && current[start] != '.' {
		return true
	}
	if end < len(current) && !isDigit(current[end]) && current[end] != '.' {
		return true
	}

	if below != "" {
		for i := start; i < end+1 && i < len(below); i++ {
			if !isDigit(below[i]) && below[i] != '.' {
				return true
			}
		}
	}

	return false
}

func findAdjacent(pos int, s string) []int {
    var ret []int
    for i := 0; i < len(s); {
        if !isDigit(s[i]) {
            i++
            continue
        }
        num, skip := extractNumber(s, i)
        for j := i; j < i + skip; j++ {
            if j == pos - 1 || j == pos || j == pos +1 {
                fmt.Println("adding: ", num)
                ret = append(ret, num)
                break
                }
        }
        i += skip
    
    }
    return ret 
}

func findAdjacentNeighbours(pos int, s string) []int {
    var ret []int
    for i := 0; i < len(s); {
        if !isDigit(s[i]) {
            i++
            continue
        }
        num, skip := extractNumber(s, i)
        fmt.Printf("Extracted %d, at pos %d of size %d, symbol at: %d\n", num, i, skip, pos)
       if i == pos + 1 || i + skip - 1 == pos -1 {
        fmt.Println("Adding side: ", num)
            ret = append(ret, num)
        }
        i += skip
    }
    return ret
     
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)

	}
	score := 0

	for i, line := range lines {
		above, below := "", ""
		for j := 0; j < len(line); {
			if !isDigit(line[j]) {
				j++
				continue
			}
			num, length := extractNumber(line, j)
			if i != 0 {
				above = lines[i-1]
			}
			if i != len(lines)-1 {
				below = lines[i+1]
			}
			if checkNum(j, j+length, above, line, below) {
				score += num
			}
			j += length

		}

	}
    fmt.Println("Score: ", score)

    score = 0
    for i, line := range lines {
        for j := 1; j < len(line) - 1; j++ {
            if line[j] != '*' {
                continue
            }
            above := lines[i-1]
            adjAbove := findAdjacent(j, above)
            adjNeighbour := findAdjacentNeighbours(j, line)
            below := lines[i+1]
            adjBelow := findAdjacent(j, below)
            adjAbove = append(adjAbove, adjNeighbour...)
            adjAbove = append(adjAbove, adjBelow...)
            if len(adjAbove) != 2 {
                continue
            }
            score += adjAbove[0] * adjAbove[1]

        }
    }
    fmt.Println("Score: ", score)

}
