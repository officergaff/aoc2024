package main

import (
	"aoc2024/solutions"
	"fmt"
	"os"
	"strings"
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
    lines := strings.Split(strings.TrimSpace(string(data)), "\n")

    switch day {
    case "1": solutions.RunDay1(lines)
case "2": solutions.RunDay2(lines)
    }
}
