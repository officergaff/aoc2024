package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunDay3(lines []string) {
	fullstring := strings.Join(lines, "\n")
	part1 := day3Part1(fullstring)
	part2 := day3Part2(fullstring)
	fmt.Println(part1)
	fmt.Println(part2)
}

func day3Part1(fullstring string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(fullstring, -1)
	sum := 0
	for _, expr := range matches {
		a, _ := strconv.Atoi(expr[1])
		b, _ := strconv.Atoi(expr[2])
		sum += a * b
	}
	return sum
}

func day3Part2(fullstring string) int {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(fullstring, -1)
	isEnabled := true
	sum := 0
	for _, expr := range matches {
		if expr[0] == "do()" {
			isEnabled = true
			continue
		}
		if expr[0] == "don't()" {
			isEnabled = false
			continue
		}
		if isEnabled {
			a, _ := strconv.Atoi(expr[1])
			b, _ := strconv.Atoi(expr[2])
			sum += a * b
		}
	}
	return sum
}
