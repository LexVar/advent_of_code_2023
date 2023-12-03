package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var limits = map[string]int{
	"red":   12,
	"blue":  14,
	"green": 13,
}

func isGameValid(hands []string) bool {
	for _, hand := range hands {
		cubes := strings.Split(hand, ",")
		for _, cube := range cubes {
			elements := strings.Split(strings.TrimSpace(cube), " ")
			color := elements[1]
			count, _ := strconv.Atoi(elements[0])
			if count > limits[color] {
				return false
			}
		}
	}
	return true
}

func parseGamePart2(hands []string) int {
	fewest := map[string]float64{
		"red":   0.0,
		"blue":  0.0,
		"green": 0.0,
	}
	for _, hand := range hands {
		cubes := strings.Split(hand, ",")
		for _, cube := range cubes {
			elements := strings.Split(strings.TrimSpace(cube), " ")
			color := elements[1]
			count, _ := strconv.Atoi(elements[0])
			fewest[color] = math.Max(fewest[color], float64(count))
		}
	}
	return int(fewest["red"] * fewest["blue"] * fewest["green"])
}

func parseGames(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total, total2 := 0, 0
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), ":")
		game, err := strconv.Atoi(strings.TrimPrefix(splits[0], "Game "))
		if err != nil {
			fmt.Println("Error converting game number:", err)
			continue
		}
		if isGameValid(strings.Split(splits[1], ";")) {
			total += game
		}
		total2 += parseGamePart2(strings.Split(splits[1], ";"))
	}
	fmt.Println(total, total2)
}

func main() {
	filename := flag.String("filename", "input_example", "Filename with puzzle input")
	flag.Parse()
	parseGames(*filename)
}
