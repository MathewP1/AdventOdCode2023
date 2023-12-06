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

type valRange struct {
	srcStart, desStart, length int64
}

func (v *valRange) in(value int64) bool {
	return value >= v.srcStart && value < v.srcStart+v.length
}

func getValue(value int64, ranges []valRange) int64 {
	for _, r := range ranges {
		if !r.in(value) {
			continue
		}
		s := value - r.srcStart
		return r.desStart + s
	}

	return value
}

type almanac struct {
	rangesMap map[mapKey][]valRange
	seeds     []int64
}
type almanac2 struct {
	rangesMap map[mapKey][]valRange
	seeds     []valRange
}


type mapKey string

const (
	seedToSoil            mapKey = "seed-to-soil"
	soilToFertilizer      mapKey = "soil-to-fertilizer"
	fertilizerToWater     mapKey = "fertilizer-to-water"
	waterToLight          mapKey = "water-to-light"
	lightToTemperature    mapKey = "light-to-temperature"
	temperatureToHumidity mapKey = "temperature-to-humidity"
	humidityToLocation    mapKey = "humidity-to-location"
	invalid               mapKey = "invalid"
)

func newAlmanac() *almanac {
	return &almanac{
		rangesMap: map[mapKey][]valRange{
			seedToSoil:            {},
			soilToFertilizer:      {},
			fertilizerToWater:     {},
			waterToLight:          {},
			lightToTemperature:    {},
			temperatureToHumidity: {},
			humidityToLocation:    {},
		},
	}
}

func (a *almanac) parseSeeds(s string) {
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
	a.seeds = make([]int64, 0, len(split))
	for _, v := range split {
        val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		a.seeds = append(a.seeds, val)
	}
}

func (a *almanac2) parseSeeds(s string) {
    s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
    if len(split) % 2 != 0 {
        panic("Seeds odd")
    }
    a.seeds = make([]valRange, 0, len(split) / 2)
    for i := 0; i < len(split); i += 2 {
        v1, err := strconv.ParseInt(split[i], 10, 64)
        if err != nil {
        panic(err)
        }
        v2, err := strconv.ParseInt(split[i+1], 10, 64)
        if err != nil {
        panic(err)
        }
        fmt.Printf("Adding range from %d to %d\n", v1, v2)
        a.seeds = append(a.seeds, valRange{srcStart : int64(v1), length: int64(v2)})
    }
}

func (a *almanac) parseMap(s string, key mapKey) {
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
    destinationStart ,err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		panic(err)
	}
    sourceStart, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		panic(err)
	}
    count, err := strconv.ParseInt(split[2], 10, 64)
	if err != nil {
		panic(err)
	}

	currentRanges := a.rangesMap[key]
	currentRanges = append(currentRanges, valRange{srcStart: sourceStart, desStart: destinationStart, length: count})
	a.rangesMap[key] = currentRanges
}

func (a *almanac) getLocation(seedNum int64) int64 {
    value := seedNum
	steps := []mapKey{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation}
    test := make([]int64, 0, len(steps) + 1)
    test = append(test, value)
    for _, step := range steps {
        m := a.rangesMap[step]
        value = getValue(value, m)
        test = append(test, value)
    }
   // fmt.Println("Steps: ", test)
	return value
}

func (al *almanac) findLowestInRange(seedNum valRange, lowestChan chan int64) {
    lowest := int64(math.MaxInt64)
    var i int64
    for i=0; i < seedNum.length; i++ {
        v := al.getLocation(i + seedNum.srcStart)
        if v < 4917125 {
            fmt.Println("LOWER!!! ", v)
        }
        if v < lowest {
            lowest = v
        }
    }
    lowestChan <- lowest
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	al := newAlmanac()
    al2 := &almanac2{}

	currentMapKey := invalid

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			currentMapKey = invalid
			continue
		}

		if currentMapKey != invalid {
			al.parseMap(line, currentMapKey)
			continue
		}

		firstSpace := strings.IndexAny(line, " ")
		word := line[:firstSpace]

		if word == "seeds:" {
			al.parseSeeds(line[firstSpace:])
            al2.parseSeeds(line[firstSpace:]) 
			continue
		}

		key := mapKey(word)
		if _, b := al.rangesMap[key]; !b {
			panic(fmt.Sprintf("Key: %s doesn't exist!", string(key)))
		}
		currentMapKey = key

	}
    al2.rangesMap = al.rangesMap


	fmt.Println("Calculating lowest...")

	lowest := int64(math.MaxInt64)
	for _, seed := range al.seeds {
		loc := al.getLocation(seed)
		if loc < lowest {
			lowest = loc
		}
	}

	fmt.Println("Lowest location num: ", lowest)

    fmt.Println("Max int32: ", math.MaxInt32)
    fmt.Println("Max int64: ", math.MaxInt64)
    fmt.Println("Calculation lowest for seed ranges...")
    start := time.Now()
    lowestChan := make(chan int64)
    for _, seed := range al2.seeds {
        fmt.Println("Range :", seed.srcStart, seed.length)
        go al.findLowestInRange(seed, lowestChan)
    }

    lowest = math.MaxInt64
    for i := 0; i < len(al2.seeds); i++ {
        result := <-lowestChan
        fmt.Println("got value from channel: ", result)
        if result < lowest {
            lowest = result
        }
    }
    close(lowestChan)
    duration := time.Since(start).Milliseconds()
    fmt.Println("Took: ", duration)

    fmt.Println("Lowest location for seed ranges: ", lowest)

}
