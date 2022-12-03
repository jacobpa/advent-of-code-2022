package day1

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func Answer(input *os.File) (int, int) {
	elves := countElfCalories(input)
	return elves[len(elves)-1], elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
}

func countElfCalories(input *os.File) []int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	elves := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Encountered non-number token, creating new elf")
			elves = append(elves, sum)
			sum = 0
		}
		sum += cal
	}

	if sum != 0 {
		elves = append(elves, sum)
	}

	sort.Ints(elves)
	return elves
}
