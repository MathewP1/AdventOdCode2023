package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// input is list of times alloted and best ever distances in this time
// in each race you must go farther that distance in input
// the longer you hold down the button, the faster the boat will go
// each second is +1 mm/s

func getPossibleWins(t, s int64) int64 {
    wins := int64(0)
    for x := int64(1); x < t; x++ {
        if s < x* (t-x) {
            wins+=1
        }
    }
    return wins 
}

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(f)
    lines := [2]string{}
    for i := 0; scanner.Scan(); i++ {
        lines[i] = scanner.Text()
    }
    fmt.Println(lines)
    timeLine := lines[0]
    distanceLine := lines[1]

    timeLine = timeLine[strings.IndexAny(timeLine, ":")+1:]
    distanceLine = distanceLine[strings.IndexAny(distanceLine, ":")+1:]

    timeLine = strings.TrimSpace(timeLine)
    distanceLine = strings.TrimSpace(distanceLine)
    
    readLine := func(l string) []int {
        split := strings.Split(l, " ")
        ret := make([]int, 0, 0)
        for _, s := range split {
            if s == "" || s == " " {
                continue
            }

            val, err := strconv.Atoi(s)
            if err != nil {
                panic(err)
            }
            ret = append(ret, val)

        }
        return ret 
    }
    
    times := readLine(timeLine)
    distances := readLine(distanceLine)
    
    mult := int64(1)
    for i:=0; i < len(times); i++ {
        possibleWins := getPossibleWins(int64(times[i]), int64(distances[i]))
        fmt.Printf("Possible wins with t=%d and s=%d: %d\n", times[i], distances[i], possibleWins)
        mult *= possibleWins
    }
    fmt.Println("Possible numbers of wins multipied: ", mult)


    timeString := fmt.Sprintf("%d%d%d%d", times[0], times[1], times[2], times[3])
    distanceString := fmt.Sprintf("%d%d%d%d", distances[0], distances[1], distances[2], distances[3])

    time, err := strconv.Atoi(timeString)
    if err != nil {
        panic(err)
    }
    distance, err := strconv.Atoi(distanceString)
    if err != nil {
        panic(err)
    }

    fmt.Println(time, distance)

    fmt.Println("Possible wins: ", getPossibleWins(int64(time), int64(distance)))

}
