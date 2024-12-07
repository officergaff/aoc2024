package main

import (
	"fmt"
	"os"

	"aoc2024/solutions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]
	Run(day)
}

func Run(day string) {
	data, _ := os.ReadFile("inputs/day" + day + ".txt")

	switch day {
	case "1":
		solutions.RunDay1(data)
	case "2":
		solutions.RunDay2(data)
	case "3":
		solutions.RunDay3(data)
	case "4":
		solutions.RunDay4(data)
	case "5":
		solutions.RunDay5(data)
	}
}
