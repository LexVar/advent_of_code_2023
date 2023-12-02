package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

func findStringDigit(s string) int {
	numberStringToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	if val, ok := numberStringToInt[s]; ok {
		return val
	} else {
		return 0
	}
}

func findFirstDigit(s string) int {
	buffer := "     "
	for _, element := range s {
		// Keep sliding buffer updated
		buffer = buffer[1:] + string(element)
		if unicode.IsDigit(element) {
			return int(element - 48)
		} else if val := findStringDigit(buffer) + findStringDigit(buffer[1:]) + findStringDigit(buffer[2:]); val != 0 {
			return val
		}
	}
	return -1
}

func findLastDigit(s string) int {
	buffer := "     "
	for i, _ := range s {
		i = len(s) - i - 1
		// Keep sliding buffer updated
		buffer = string(s[i]) + buffer[:4]

		if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - 48)
		} else if val := findStringDigit(buffer) + findStringDigit(buffer[:4]) + findStringDigit(buffer[:3]); val != 0 {
			return val
		}
	}
	return -1
}

func part1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		total += 10*findFirstDigit(scanner.Text()) + findLastDigit(scanner.Text())
	}
	fmt.Println(total)
}

func main() {
	filename := flag.String("filename", "input_example", "Filename with puzzle input")
	flag.Parse()
	part1(*filename)
}
