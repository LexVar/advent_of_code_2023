package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countDifferentTimes(time int, distance int) int {
	timeButtonIsHeld := 1
	for timeButtonIsHeld*(time-timeButtonIsHeld) <= distance {
		timeButtonIsHeld++
	}
	return time - 2*timeButtonIsHeld + 1
}

func part1(times []int, distances []int) int {
	product := 1
	for i := 0; i < len(times); i++ {
		product *= countDifferentTimes(times[i], distances[i])
	}
	return product
}

func start(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line1 := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
	scanner.Scan()
	line2 := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")

	times := make([]int, len(line1))
	distances := make([]int, len(line2))
	for i := 0; i < len(times); i++ {
		times[i], _ = strconv.Atoi(line1[i])
		distances[i], _ = strconv.Atoi(line2[i])
	}

	fmt.Println("Part 1:", part1(times, distances))
}

func main() {
	filename := flag.String("filename", "input_example", "Filename with puzzle input")
	flag.Parse()

	start(*filename)
}
