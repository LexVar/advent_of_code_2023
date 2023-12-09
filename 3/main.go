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
	// < 0 means it's a symbol
	// -2 means it's *
	// -1 means it's something else
	number       int
	lineNumber   int
	startIndex   int
	endIndex     int
	isPartNumber bool
}

func getPotentialNumbers(line string, lineNumber int) []Number {
	buffer := ""
	numbers := []Number{}
	currentNumber := Number{isPartNumber: false, lineNumber: lineNumber}
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
				currentNumber = Number{isPartNumber: false, lineNumber: lineNumber}
			}
			if char != '.' {
				if char == '*' {
					numbers = append(numbers, Number{isPartNumber: false, number: -2, lineNumber: lineNumber, startIndex: i, endIndex: i})
				} else {
					numbers = append(numbers, Number{isPartNumber: false, number: -1, lineNumber: lineNumber, startIndex: i, endIndex: i})
				}
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
		if symbol.number < 0 && symbol.startIndex >= startIndex && symbol.endIndex <= endIndex {
			return true
		}
	}
	return false
}

func isPartNumber(schema [][]Number, number Number) bool {
	if number.number < 0 {
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

func overlap(startIndex1 int, endIndex1 int, startIndex2 int, endIndex2 int) bool {
	return (startIndex1 >= startIndex2 && startIndex1 <= endIndex2) ||
		(endIndex1 >= startIndex2 && endIndex1 <= endIndex2) ||
		(startIndex1 <= startIndex2 && endIndex1 >= endIndex2)
}

func getPartNumbers(symbols []Number, startIndex int, endIndex int) []Number {
	endIndex++
	startIndex--
	partNumbers := []Number{}
	for _, symbol := range symbols {
		if symbol.number > 0 && symbol.isPartNumber && overlap(symbol.startIndex, symbol.endIndex, startIndex, endIndex) {
			partNumbers = append(partNumbers, symbol)
		}
	}
	return partNumbers
}

func gearRatio(schema [][]Number, number Number) int {
	if number.number != -2 {
		return 0
	}
	partNumbers := []Number{}
	if number.lineNumber == 0 {
		partNumbers = append(getPartNumbers(schema[number.lineNumber+1], number.startIndex, number.endIndex),
			getPartNumbers(schema[number.lineNumber], number.startIndex, number.endIndex)...)
	} else {
		partNumbers = append(append(getPartNumbers(schema[number.lineNumber-1], number.startIndex, number.endIndex),
			getPartNumbers(schema[number.lineNumber+1], number.startIndex, number.endIndex)...),
			getPartNumbers(schema[number.lineNumber], number.startIndex, number.endIndex)...)
	}
	if len(partNumbers) == 2 {
		return partNumbers[0].number * partNumbers[1].number
	} else {
		return 0
	}
}

func part1(filename string) (int, [][]Number) {
	schema := [][]Number{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, schema
	}
	defer file.Close()
	// Builds the schema matrix with number objects
	for scanner, line := bufio.NewScanner(file), 0; scanner.Scan(); line++ {
		schema = append(schema, getPotentialNumbers(scanner.Text(), line))
	}
	schema = append(schema, []Number{})
	sum := 0
	for i, row := range schema {
		for j, potentialNumber := range row {
			if isPartNumber(schema, potentialNumber) {
				sum += potentialNumber.number
				potentialNumber.isPartNumber = true
				schema[i][j].isPartNumber = true
			}
		}
	}
	return sum, schema
}

func part2(schema [][]Number) int {
	total := 0
	for _, row := range schema {
		for _, potentialNumber := range row {
			if ratio := gearRatio(schema, potentialNumber); ratio > 0 {
				total += ratio
			}
		}
	}
	return total
}

func main() {
	filename := flag.String("filename", "input_example", "Filename with puzzle input")
	flag.Parse()

	result, schema := part1(*filename)
	fmt.Println(result)

	fmt.Println(part2(schema))
}
