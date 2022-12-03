package main

import (
	"adventofcode2022/day2"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Must provide path to input file")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Could not open file at %v\n", os.Args[1])
	}
	defer file.Close()
	// partOne, partTwo := day1.Answer(file)
	partOne, partTwo := day2.Answer(file)
	fmt.Printf("The answers are:\nPart One: %v\nPart Two: %v\n", partOne, partTwo)
}
