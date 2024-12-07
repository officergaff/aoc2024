package solutions

import (
	"fmt"
	"strings"
)

var example4 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func RunDay4(data []byte) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	exampleLines := strings.Split(example4, "\n")
	exampleChars := make([][]string, len(exampleLines))
	for i, line := range exampleLines {
		exampleChars[i] = strings.Split(line, "")
	}

	fmt.Println(day4Part1(exampleChars))
	fmt.Println(day4Part2(exampleChars))
	chars := make([][]string, len(lines))
	for i, line := range lines {
		chars[i] = strings.Split(line, "")
	}
	fmt.Println(day4Part1(chars))
	fmt.Println(day4Part2(chars))
}

func findXmases(i int, j int, chars [][]string) int {
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {-1, -1}, {1, -1}, {1, 1}}
	words := make([]string, 8)
	for dIdx := range directions {
		for k := range 4 {
			dy, dx := i+k*directions[dIdx][0], j+k*directions[dIdx][1]
			if 0 <= dx && dx < len(chars[0]) && 0 <= dy && dy < len(chars) {
				words[dIdx] += chars[dy][dx]
			}
		}
	}
	count := 0
	for _, word := range words {
		if word == "XMAS" {
			count++
		}
	}
	return count
}

func findMases(i int, j int, chars [][]string) bool {
	directions := [][]int{{-1, 1}, {-1, -1}}
	words := make([]string, 2)
	for dIdx := range directions {
		for k := -1; k < 2; k++ {
			dy, dx := i+k*directions[dIdx][0], j+k*directions[dIdx][1]
			if 0 <= dx && dx < len(chars[0]) && 0 <= dy && dy < len(chars) {
				words[dIdx] += chars[dy][dx]
			}
		}
	}
	for _, word := range words {
		if word != "MAS" && word != "SAM" {
			return false
		}
	}
	return true
}

func day4Part1(chars [][]string) int {
	total := 0
	for i := range chars {
		for j := range chars[0] {
			if chars[i][j] == "X" {
				total += findXmases(i, j, chars)
			}
		}
	}
	return total
}

func day4Part2(chars [][]string) int {
	total := 0
	for i := range chars {
		for j := range chars[0] {
			if chars[i][j] == "A" {
				if findMases(i, j, chars) {
					total++
				}
			}
		}
	}
	return total
}
