package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello, World!")
	csvFileName := flag.String("csv", "problems.csv", "csv file of problems & answer")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file: %s\n", *csvFileName))
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv file")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problems #%d: %s = \n", i+1, p.ques)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.ans {
			// fmt.Print("correct \n")
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d .\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem {
			ques: line[0],
			ans: strings.TrimSpace(line[1]),
		}
	}
	return ret 
}

type problem struct {
	ques string
	ans string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}