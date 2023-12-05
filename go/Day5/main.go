package main

import (
	"bufio"
	"fmt"
	"os"
)

type gardenMap struct {
	m map[int]int
}

func getMaps() map[string]*gardenMap {
	return map[string]*gardenMap{
		"seed-to-soil":            &gardenMap{},
		"soil-to-fertilizer":      &gardenMap{},
		"fetilizer-to-water":      &gardenMap{},
		"water-to-light":          &gardenMap{},
		"light-to-temperature":    &gardenMap{},
		"temperature-to-humidity": &gardenMap{},
		"humidity-to-location":    &gardenMap{},
	}

}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			fmt.Println("empty")
		}
	}
}
