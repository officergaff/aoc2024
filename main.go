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
	data, err := os.ReadFile("inputs/day" + day + ".txt")
	if err != nil {
		fmt.Println(err)
	}

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
	case "6":
		solutions.RunDay6(data)
	case "7":
		solutions.RunDay7(data)
	case "8":
		solutions.RunDay8(data)
	case "9":
		solutions.RunDay9(data)
	case "10":
		solutions.RunDay10(data)
	case "11":
		solutions.RunDay11(data)
	case "12":
		solutions.RunDay12(data)
	}
}
