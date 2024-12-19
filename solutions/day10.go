package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

var example10 = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

type Summit struct {
	y int
	x int
}

func exploreTrail(cHeight int, r int, c int, tmap [][]int, summits map[Summit]bool) {
	if cHeight == 9 {
		summits[Summit{r, c}] = true
		return
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range directions {
		y, x := r+d[0], c+d[1]
		if 0 <= y && y < len(tmap) && 0 <= x && x < len(tmap[0]) {
			nHeight := tmap[y][x]
			if cHeight == nHeight-1 {
				exploreTrail(nHeight, y, x, tmap, summits)
			}
		}
	}
}

func exploreDistinctTrail(cHeight int, r int, c int, tmap [][]int) int {
	if cHeight == 9 {
		return 1
	}
	total := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range directions {
		y, x := r+d[0], c+d[1]
		if 0 <= y && y < len(tmap) && 0 <= x && x < len(tmap[0]) {
			nHeight := tmap[y][x]
			if cHeight == nHeight-1 {
				total += exploreDistinctTrail(nHeight, y, x, tmap)
			}
		}
	}
	return total
}

func day10Part2(tmap [][]int) int {
	total := 0
	for j, row := range tmap {
		for i, height := range row {
			if height == 0 {
				total += exploreDistinctTrail(0, j, i, tmap)
			}
		}
	}
	return total
}

func day10Part1(tmap [][]int) int {
	total := 0
	for j, row := range tmap {
		for i, height := range row {
			if height == 0 {
				summits := map[Summit]bool{}
				exploreTrail(0, j, i, tmap, summits)
				total += len(summits)
			}
		}
	}
	return total
}

func RunDay10(data []byte) {
	eLines := strings.Split(example10, "\n")
	eMap := [][]int{}
	for _, line := range eLines {
		row := []int{}
		for _, s := range strings.Split(line, "") {
			n, _ := strconv.Atoi(s)
			row = append(row, n)
		}
		eMap = append(eMap, row)
	}
	e1 := day10Part1(eMap)
	e2 := day10Part2(eMap)
	fmt.Println(e1)
	fmt.Println(e2)

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	tmap := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, s := range strings.Split(line, "") {
			n, _ := strconv.Atoi(s)
			row = append(row, n)
		}
		tmap = append(tmap, row)
	}
	part1 := day10Part1(tmap)
	part2 := day10Part2(tmap)
	fmt.Println(part1)
	fmt.Println(part2)
}
