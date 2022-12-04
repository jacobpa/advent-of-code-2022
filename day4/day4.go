package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	builder = strings.Builder{}
)

func Answer(input *os.File) (int, int) {
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
	numbers := make([]int, 0, 4)

	for i := 0; true; i++ {
		if input[i] > 47 && input[i] < 58 {
			builder.WriteByte(input[i])
		} else {
			converted, _ := strconv.Atoi(builder.String())
			numbers = append(numbers, converted)
			builder.Reset()
		}
	}

	converted, _ := strconv.Atoi(builder.String())
	numbers = append(numbers, converted)
	builder.Reset()

	if len(numbers) != 4 {
		panic("Unexpected input for parsing range")
	}

	return numbers
}
