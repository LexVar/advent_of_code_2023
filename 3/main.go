package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	// -1 means it's a symbol
	number     int
	lineNumber int
	startIndex int
	endIndex   int
}

func getPotentialNumbers(line string, lineNumber int) []Number {
	buffer := ""
	numbers := []Number{}
	currentNumber := Number{lineNumber: lineNumber}
	for i, char := range line {
		if unicode.IsDigit(char) {
			if buffer == "" {
				currentNumber.startIndex = i
			}
			buffer += string(char)
		} else {
			if buffer != "" {
				currentNumber.number, _ = strconv.Atoi(buffer)
				currentNumber.endIndex = i - 1
				numbers = append(numbers, currentNumber)
				// New number, reset buffer
				buffer = ""
				currentNumber = Number{lineNumber: lineNumber}
			}
			if char != '.' {
				numbers = append(numbers, Number{number: -1, lineNumber: lineNumber, startIndex: i, endIndex: i})
			}
		}
	}
	if buffer != "" {
		currentNumber.number, _ = strconv.Atoi(buffer)
		currentNumber.endIndex = len(line) - 1
		numbers = append(numbers, currentNumber)
	}
	return numbers
}

func checkForSymbolInLine(symbols []Number, startIndex int, endIndex int) bool {
	endIndex++
	startIndex--
	for _, symbol := range symbols {
		if symbol.number == -1 && symbol.startIndex >= startIndex && symbol.endIndex <= endIndex {
			return true
		}
	}
	return false
}

func isPartNumber(schema [][]Number, number Number) bool {
	if number.number == -1 {
		return false
	}
	if number.lineNumber == 0 {
		return (checkForSymbolInLine(schema[number.lineNumber+1], number.startIndex, number.endIndex) ||
			checkForSymbolInLine(schema[number.lineNumber], number.startIndex, number.endIndex))
	} else {
		return (checkForSymbolInLine(schema[number.lineNumber-1], number.startIndex, number.endIndex) ||
			checkForSymbolInLine(schema[number.lineNumber+1], number.startIndex, number.endIndex) ||
			checkForSymbolInLine(schema[number.lineNumber], number.startIndex, number.endIndex))
	}
}

func searchForNumbers(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	schema := [][]Number{}
	for line := 0; scanner.Scan(); line++ {
		schema = append(schema, getPotentialNumbers(scanner.Text(), line))
	}
	schema = append(schema, []Number{})
	for _, row := range schema {
		for _, potentialNumber := range row {
			if isPartNumber(schema, potentialNumber) {
				total += potentialNumber.number
			}
		}
	}
	fmt.Println(total)
}

func main() {
	filename := flag.String("filename", "input_example", "Filename with puzzle input")
	flag.Parse()
	searchForNumbers(*filename)
}
