package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func part1(winningNumbers []string, myNumbers []string) int {
	total := 0
	for _, winningNumber := range winningNumbers {
		winningNumber = strings.TrimSpace(winningNumber)
		for _, myNumber := range myNumbers {
			myNumber = strings.TrimSpace(myNumber)
			if myNumber != "" && winningNumber == myNumber {
				if total == 0 {
					total = 1
				} else {
					total *= 2
				}
			}
		}
	}
	return total
}

type ScratchCard struct {
	copies          int
	matchingNumbers int
}

func countMatchingNumbers(winningNumbers []string, myNumbers []string) int {
	matchingNumbers := 0
	for _, winningNumber := range winningNumbers {
		winningNumber = strings.TrimSpace(winningNumber)
		for _, myNumber := range myNumbers {
			myNumber = strings.TrimSpace(myNumber)
			if myNumber != "" && winningNumber == myNumber {
				matchingNumbers++
			}
		}
	}
	return matchingNumbers
}

func calculateScratchCards(scratchCards []ScratchCard) int {
	totalScratchcards := 0
	for card, scratchCard := range scratchCards {
		for i := 1; i <= scratchCard.matchingNumbers; i++ {
			scratchCards[card+i].copies += scratchCard.copies
		}
		totalScratchcards += scratchCard.copies
	}
	return totalScratchcards
}

func start(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scratchCards := make([]ScratchCard, 0)
	total := 0
	// Builds the schema matrix with number objects
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		numbers := strings.Split(strings.Split(scanner.Text(), ": ")[1], " | ")
		winningNumbers, myNumbers := strings.Split(numbers[0], " "), strings.Split(numbers[1], " ")
		total += part1(winningNumbers, myNumbers)
		scratchCards = append(scratchCards, ScratchCard{copies: 1, matchingNumbers: countMatchingNumbers(winningNumbers, myNumbers)})
	}
	fmt.Println("part 1", total)
	fmt.Println("part 2", calculateScratchCards(scratchCards))
}

func main() {
	filename := flag.String("filename", "input_example", "Filename with puzzle input")
	flag.Parse()

	start(*filename)
}
