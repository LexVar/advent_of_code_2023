package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Mapping struct {
	source      int
	destination int
	valueCount  int
}

var mappings = [7][]Mapping{}

func seedToLocation(seed int) int {
	for _, mapping := range mappings {
		for _, m := range mapping {
			if seed >= m.source && seed < m.source+m.valueCount {
				seed = m.destination + (seed - m.source)
				break
			}
		}
	}
	return seed
}

func part1(seeds []int) []int {
	locations := make([]int, 0)
	for _, seed := range seeds {
		locations = append(locations, seedToLocation(seed))
	}
	return locations
}

var wg sync.WaitGroup

func calculateSeedRange(seed int, length int, ch chan int) {
	defer wg.Done()
	closestLocation := math.MaxInt
	for i := 0; i < length; i++ {
		closestLocation = min(seedToLocation(seed+i), closestLocation)
	}
	ch <- closestLocation
}

func part2(seeds []int) int {
	closestLocation := math.MaxInt
	channel := make(chan int, 20)
	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go calculateSeedRange(seeds[i], seeds[i+1], channel)
	}
	wg.Wait()
	close(channel)
	for location := range channel {
		closestLocation = min(location, closestLocation)
	}

	return closestLocation
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
	seeds := make([]int, 0)
	for _, seed := range strings.Split(strings.Split(scanner.Text(), ": ")[1], " ") {
		seed, _ := strconv.Atoi(seed)
		seeds = append(seeds, seed)
	}
	i := 0
	for scanner.Scan(); scanner.Scan(); i++ {
		scanner.Scan()
		for ; scanner.Text() != ""; scanner.Scan() {
			mappingValues := strings.Split(scanner.Text(), " ")
			destination, _ := strconv.Atoi(mappingValues[0])
			source, _ := strconv.Atoi(mappingValues[1])
			valueCount, _ := strconv.Atoi(mappingValues[2])
			mappings[i] = append(mappings[i], Mapping{source: source, destination: destination, valueCount: valueCount})
		}
	}

	fmt.Println("part1", slices.Min(part1(seeds)))
	fmt.Println("part2", part2(seeds))
}

func main() {
	filename := flag.String("filename", "input_example", "Filename with puzzle input")
	flag.Parse()

	start(*filename)
}
