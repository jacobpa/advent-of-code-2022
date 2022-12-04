package day4

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var (
	numberExpression = regexp.MustCompile("\\d+")
)

func Answer(input *os.File) (int,int) {
	scanner := bufio.NewScanner(input)

	countEncapsulates := 0
	countOverlaps := 0
	for scanner.Scan() {
		line := scanner.Text()
		if encapsulates(line) {
			countEncapsulates++
		}
		if overlaps(line) {
			countOverlaps++
		}
	}

	return countEncapsulates, countOverlaps
}

func encapsulates(ranges string) bool {
	numbers := rangesToIntSlice(ranges)

	if numbers[0] <= numbers[2] && numbers[1] >= numbers[3] {
		return true
	}
	if numbers[2] <= numbers[0] && numbers[3] >= numbers[1] {
		return true
	}

	return false
}

func overlaps(ranges string) bool {
	numbers := rangesToIntSlice(ranges)

	if numbers[0] < numbers[2] && numbers[1] < numbers[2] {
		return false
	}
	if numbers[0] > numbers[3] && numbers[1] > numbers[3] {
		return false
	}

	return true
}

func rangesToIntSlice(input string) []int {
	strings := numberExpression.FindAllString(input, 4)
	if len(strings) != 4 {
		panic("Unexpected input for parsing range")
	}

	numbers := make([]int, 4)
	for i, number := range strings {
		numbers[i], _ = strconv.Atoi(number)
	}

	return numbers
}
