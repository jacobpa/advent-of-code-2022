package main

import (
	"aoc2022/day1"
	"aoc2022/day2"
	"aoc2022/day3"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	OK int = iota
	INVALID_ARGUMENT
	NO_SOLUTION
	INPUT_READ_ERROR
)

var funcMapping = map[string]func(*os.File) (int, int){
	"1": day1.Answer,
	"2": day2.Answer,
	"3": day3.Answer,
}

func usage(name string) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("USAGE: %v day input\n", filepath.Base(name)))
	builder.WriteString(fmt.Sprintln("\tday\tday of challenge (1-25)"))
	builder.WriteString(fmt.Sprintln("\tinput\tpath to input file"))
	return builder.String()
}

func main() {
	if len(os.Args) != 3 {
		fmt.Print(usage(os.Args[0]))
		os.Exit(INVALID_ARGUMENT)
	}

	if _, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid day parameter: %v\n", os.Args[1])
		fmt.Print(usage(os.Args[0]))
		os.Exit(INVALID_ARGUMENT)
	}

	answerFunc := getAnswerFunc(os.Args[1])
	if answerFunc == nil {
		fmt.Fprintf(os.Stderr, "Could not find challenge implementaiton for day %v\n", os.Args[1])
		os.Exit(NO_SOLUTION)
	}

	file, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open file at %v\n", os.Args[2])
		os.Exit(INPUT_READ_ERROR)
	}
	defer file.Close()

	partOne, partTwo := answerFunc(file)
	fmt.Printf("The answers are:\nPart One: %v\nPart Two: %v\n", partOne, partTwo)
}

func getAnswerFunc(day string) func(*os.File) (int, int) {
	if answerFunc, ok := funcMapping[day]; ok {
		return answerFunc
	}
	return nil
}
