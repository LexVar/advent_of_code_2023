package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CamelHand struct {
	bid      int
	hand     string
	typeHand int
	jokers   int
}

var cards = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

/*
"five of a kind" - 7
"four of a kind" - 6
"fullhouse" - 5
"three of a kind" - 4
"two pair" - 3
"one pair" - 2
"high card" - 1
*/
func getHandType(hand CamelHand) int {
	cardCount := make(map[rune]int)
	pairs, triples, fours, fives := 0, 0, 0, 0
	for _, card := range hand.hand {
		cardCount[card] = strings.Count(hand.hand, string(card))
	}
	for _, count := range cardCount {
		if count == 5 {
			fives++
		} else if count == 4 {
			fours++
		} else if count == 2 {
			pairs++
		} else if count == 3 {
			triples++
		}
	}
	if fives == 1 {
		return 7
	} else if fours == 1 {
		return 6
	} else if triples == 1 && pairs == 1 {
		return 5
	} else if triples == 1 {
		return 4
	} else if pairs == 2 {
		return 3
	} else if pairs == 1 {
		return 2
	} else {
		return 1
	}
}

type ByType []CamelHand

func (a ByType) Len() int           { return len(a) }
func (a ByType) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByType) Less(i, j int) bool { return a[i].typeHand < a[j].typeHand }

type ByCard []CamelHand

func (a ByCard) Len() int      { return len(a) }
func (a ByCard) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCard) Less(i, j int) bool {
	for k := 0; k < len(a[i].hand); k++ {
		if cards[a[i].hand[k]] == cards[a[j].hand[k]] {
			continue
		}
		return cards[a[i].hand[k]] < cards[a[j].hand[k]]
	}
	return false
}

func part1(hands []CamelHand) int {
	totalWinnings := 0
	sort.Sort(ByType(hands))
	t, startIndex := hands[0].typeHand, 0
	for i := 0; i < len(hands); i++ {
		if hands[i].typeHand != t {
			sort.Sort(ByCard(hands[startIndex:i]))
			startIndex = i
			t = hands[i].typeHand
		}
	}
	for ind, hand := range hands {
		totalWinnings += hand.bid * (ind + 1)
	}
	return totalWinnings
}

// -------------------- Part 2 --------------------------------
var cards2 = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 0,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func getHandTypeJoker(hand CamelHand) int {
	if hand.typeHand < 1 {
		hand.typeHand = getHandType(hand)
	}
	if hand.jokers <= 0 {
		return hand.typeHand
	} else if hand.jokers == 1 {
		switch hand.typeHand {
		case 1, 6:
			hand.typeHand++
		case 2, 3, 4:
			hand.typeHand += 2
		}
	} else if hand.jokers == 2 {
		switch hand.typeHand {
		case 2:
			hand.typeHand = 4
		case 3:
			hand.typeHand = 6
		case 5:
			hand.typeHand = 7
		}
	} else if hand.jokers == 3 {
		switch hand.typeHand {
		case 4:
			hand.typeHand = 6
		case 5:
			hand.typeHand = 7
		}
	} else if hand.jokers == 4 && hand.typeHand == 6 {
		hand.typeHand = 7
	}
	return hand.typeHand
}

type ByCardJoker []CamelHand

func (a ByCardJoker) Len() int      { return len(a) }
func (a ByCardJoker) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCardJoker) Less(i, j int) bool {
	for k := 0; k < len(a[i].hand); k++ {
		if cards2[a[i].hand[k]] == cards2[a[j].hand[k]] {
			continue
		}
		return cards2[a[i].hand[k]] < cards2[a[j].hand[k]]
	}
	return false
}

func part2(hands []CamelHand) int {
	totalWinnings := 0
	sort.Sort(ByType(hands))
	t, startIndex := hands[0].typeHand, 0
	for i := 0; i < len(hands); i++ {
		if hands[i].typeHand != t {
			sort.Sort(ByCardJoker(hands[startIndex:i]))
			startIndex = i
			t = hands[i].typeHand
		}
	}
	sort.Sort(ByCardJoker(hands[startIndex:]))
	for ind, hand := range hands {
		totalWinnings += hand.bid * (ind + 1)
	}
	return totalWinnings
}

func start(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	hands := make([]CamelHand, 0)
	hands2 := make([]CamelHand, 0)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(line[1])
		// part 1
		newHand := CamelHand{hand: line[0], bid: bid, typeHand: 0}
		newHand.typeHand = getHandType(newHand)
		hands = append(hands, newHand)

		// part 2
		newHand.jokers = strings.Count(newHand.hand, "J")
		newHand.typeHand = getHandTypeJoker(newHand)
		hands2 = append(hands2, newHand)
	}

	fmt.Println("Part 1:", part1(hands))
	fmt.Println("Part 2:", part2(hands2))
}

func main() {
	filename := flag.String("filename", "input_example", "Filename with puzzle input")
	flag.Parse()

	start(*filename)
}
