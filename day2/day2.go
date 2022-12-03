package day2

import (
	"bufio"
	"os"
)

var (
	// A,X == Rock
	// B,Y == Paper
	// C,Z == Scissors
	points = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	losePlays = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
		"X": "B",
		"Y": "C",
		"Z": "A",
	}
	winPlays = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
	drawPlays = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
)

func Answer(input *os.File) (int, int) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	partOneScore := 0
	partTwoScore := 0
	for scanner.Scan() {
		opponentInput := scanner.Text()
		scanner.Scan() // Know based on input we can scan twice per line
		guideInput := scanner.Text()
		partOneScore += caluclatePlayerScore(opponentInput, guideInput)
		partTwoScore += caluclatePlayerScore(opponentInput, calculateDesiredHand(opponentInput, guideInput))
	}

	return partOneScore, partTwoScore
}

func caluclatePlayerScore(opponent, player string) int {
	outcome := 0
	if drawPlays[player] == opponent {
		outcome = 3
	} else if winPlays[player] == opponent {
		outcome = 6
	}

	return outcome + points[player]
}

func calculateDesiredHand(opponent, outcome string) string {
	if outcome == "X" {
		return losePlays[opponent]
	} else if outcome == "Y" {
		return drawPlays[opponent]
	} else {
		// Implicit "Z" case, no other valid inputs
		return winPlays[opponent]
	}
}
