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
    solutions.Run(day)
}
