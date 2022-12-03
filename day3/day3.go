package day3

import (
	"bufio"
	"os"
)

func Answer(input *os.File) (int, int) {
	scanner := bufio.NewScanner(input)

	prioritySum := 0
	badgePrioritySum := 0
	lineCache := make([]string, 0, 3)
	for scanner.Scan() {
		line := scanner.Text()
		lineCache = append(lineCache, line)
		prioritySum += calculatePriority(findCommonItem(line))
		if len(lineCache) == 3 {
			badgePrioritySum += calculatePriority(findBadge(lineCache[0], lineCache[1], lineCache[2]))
			lineCache = lineCache[0:0]
		}
	}

	return prioritySum, badgePrioritySum
}

func calculatePriority(item byte) int {
	if item > 96 {
		return int(item - 96)
	}
	return int(item - 38)
}

func findCommonItem(sack string) byte {
	left, right := sack[:len(sack)/2], sack[len(sack)/2:]
	seenLeft := make(map[byte]bool, len(left))
	for i := 0; i < len(left); i++ {
		seenLeft[left[i]] = true
	}
	for i := 0; i < len(right); i++ {
		if seenLeft[right[i]] {
			return right[i]
		}
	}

	return 0
}

func findBadge(first, second, third string) byte {
	seenFirst := make(map[byte]bool, len(first))
	seenSecond := make(map[byte]bool, len(second))
	for i := 0; i < len(first); i++ {
		seenFirst[first[i]] = true
	}
	for i := 0; i < len(second); i++ {
		if seenFirst[second[i]] {
			seenSecond[second[i]] = true
		}
	}
	for i := 0; i < len(third); i++ {
		if seenSecond[third[i]] {
			return third[i]
		}
	}

	return 0
}
