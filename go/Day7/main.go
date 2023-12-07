package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var strengths [13]byte = [13]byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var strengthsJ [13]byte = [13]byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

var strengthMap map[byte]uint8
var strengthMapJ map[byte]uint8

type handType uint8

const (
	fiveOfaKind  handType = 0
	fourOfaKind  handType = 1
	fullHouse    handType = 2
	threeOfaKind handType = 3
	twoPair      handType = 4
	onePair      handType = 5
	highCard     handType = 6
)

type hand struct {
	bid          int
	handOriginal [5]byte
	hand         [5]byte
	t            handType
}


type handJoker struct {
    bid int
    handOriginal [5]byte
    hand [5]byte
    t handType
}


func findBestType(sortedHand [5]byte) handType {
    repeats := [3]uint8{}
    repeats[0] = 1
    repI := 0
    for i := 1; i < len(sortedHand); i++ {
        if sortedHand[i] == sortedHand[i-1] {
            repeats[repI] += 1
            continue
        }

        if repeats[repI] != 1 {
            repI += 1
            repeats[repI] = 1
            continue
        }
    }

    t := highCard 
    if repeats[0] == 5 {
        t = fiveOfaKind
    } else if repeats[0] == 4 || repeats[1] == 4 {
        t = fourOfaKind
    } else if (repeats[0] == 3 && repeats[1] == 2) || (repeats[0] == 2 && repeats[1] == 3) {
        t = fullHouse
    } else if repeats[0] == 3 || repeats[1] == 3 {
        t = threeOfaKind
    } else if repeats[0] == 2 && repeats[1] == 2 {
        t = twoPair
    } else if repeats[0] == 2 || repeats[1] == 2 {
        t = onePair
    } else {
        t = highCard
    }
    return t
}

func (h *handJoker) assignHandType() {
    copy(h.handOriginal[:], h.hand[:])
    sort.Slice(h.hand[:], func(i, j int) bool {
        v1 := h.hand[i]
        v2 := h.hand[j]
        return strengthMapJ[v1] > strengthMapJ[v2]
    })
    jokerCount := uint8(0)
    for i:= len(h.hand)-1; i >=0; i-- {
        if h.hand[i] == 'J' {
            jokerCount+=1
        }
    }
    if jokerCount == 0 {
        h.t = findBestType(h.hand)   
        return
    }

    fmt.Println("has jokers: ", jokerCount)
    repeats := [3]uint8{}
    repeats[0] = 1
    repI := 0
    for i := 1; i < len(h.hand) - int(jokerCount); i++ {
        if h.hand[i] == h.hand[i-1] {
            repeats[repI] += 1
            continue
        }

        if repeats[repI] != 1 {
            repI += 1
            repeats[repI] = 1
            continue
        }
    }

    higherRepeats := repeats[0]
    lowerRepeats := repeats[1]
    if repeats[1] > repeats[0] {
        higherRepeats = repeats[1]
        lowerRepeats = repeats[0]
    }



    t := onePair
    sumHigher := higherRepeats + jokerCount
    sumLower := lowerRepeats + jokerCount
    fmt.Println(string(h.hand[:]))
    //fmt.Printf("jokers: %d, hiRepeats: %d, lowRepeats: %d, sumHigher: %d, sumLower: %d\n", jokerCount, higherRepeats, lowerRepeats, sumHigher, sumLower)
    if sumHigher == 5 || jokerCount == 5{
        t = fiveOfaKind
    } else if sumHigher == 4 {
        t = fourOfaKind
    } else if repeats[0] == 2 && repeats[1] == 2 && jokerCount == 1{
        t = fullHouse
    fmt.Printf("jokers: %d, hiRepeats: %d, lowRepeats: %d, sumHigher: %d, sumLower: %d\n", jokerCount, higherRepeats, lowerRepeats, sumHigher, sumLower)
    } else if sumHigher == 3 || sumLower == 3 {
        t = threeOfaKind
    } else if (sumHigher == 2 && lowerRepeats == 2) || (sumLower == 2 && higherRepeats == 2) {
        t = twoPair
    } else {
        t = onePair
    }
    h.t =  t

}

func (h *hand) assignHandType() {
	// TODO: opt idea - bit masks?
	copy(h.handOriginal[:], h.hand[:])
	sort.Slice(h.hand[:], func(i, j int) bool {
		v1 := h.hand[i]
		v2 := h.hand[j]
		return strengthMap[v1] > strengthMap[v2]
	})

    h.t = findBestType(h.hand)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	strengthMap = make(map[byte]uint8, 13)
	for i, s := range strengths {
		strengthMap[s] = uint8(13 - i)
	}
    strengthMapJ = make(map[byte]uint8, 13)
    for i, s := range strengthsJ {
        strengthMapJ[s] = uint8(13 - i)
    }

	hands := make([]hand, 0, 100)
    handsJoker:= make([]handJoker, 0, 100)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		cards := lineSplit[0]
		if len(cards) != 5 {
			panic("Len of cards not 5")
		}
		bid, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, hand{bid: bid})
		copy(hands[i].hand[:], cards)
        handsJoker = append(handsJoker, handJoker{bid: bid})
        copy(handsJoker[i].hand[:], cards)
		hands[i].assignHandType()
        handsJoker[i].assignHandType()
	}


	sort.Slice(hands, func(i, j int) bool {
		if hands[i].t == hands[j].t {
			for h := 0; h < len(hands[i].hand); h++ {
				if hands[i].handOriginal[h] == hands[j].handOriginal[h] {
					continue
				}
				return strengthMap[hands[i].handOriginal[h]] > strengthMap[hands[j].handOriginal[h]]
			}
		}
		return hands[i].t < hands[j].t
	})

    

	winnings := uint64(0)
	for i, h := range hands {
		rank := uint64(len(hands) - i)
		fmt.Println(i, rank, string(h.handOriginal[:]), string(h.hand[:]), h.bid)
		winnings += rank * uint64(h.bid)
	}

	fmt.Println(winnings)
    

    fmt.Println("---- Part2 -----")

    sort.Slice(handsJoker, func(i, j int) bool {
		if handsJoker[i].t == handsJoker[j].t {
			for h := 0; h < len(handsJoker[i].hand); h++ {
				if handsJoker[i].handOriginal[h] == handsJoker[j].handOriginal[h] {
					continue
				}
				return strengthMapJ[handsJoker[i].handOriginal[h]] > strengthMapJ[handsJoker[j].handOriginal[h]]
			}
		}

		return handsJoker[i].t < handsJoker[j].t
	})

    winnings = uint64(0)
    for i, h := range handsJoker {
        rank := uint64(len(hands) - i)
        fmt.Println(i, rank, string(h.handOriginal[:]), string(h.hand[:]), h.t, h.bid)
        winnings += rank * uint64(h.bid)
    }

    fmt.Println(winnings)


}
