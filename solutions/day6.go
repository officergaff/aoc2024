package solutions

import (
	"fmt"
	"strings"
)

var example6 = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func process6(dat string) [][]string {
	lines := strings.Split(dat, "\n")
	graph := make([][]string, len(lines))
	for i, line := range lines {
		cells := strings.Split(line, "")
		graph[i] = cells
	}
	return graph
}

func RunDay6(data []byte) {
	graphE := process6(example6)
	for _, v := range graphE {
		fmt.Println(v)
	}
	day6Part1(graphE)
	graph := process6(string(data))
	day6Part1(graph)
}

func findPosition(graph [][]string) []int {
	for y, line := range graph {
		for x, v := range line {
			if v == "^" {
				return []int{y, x}
			}
		}
	}
	return []int{-1, -1}
}

func getNextCoord(currPos []int, currDir int) (int, int) {
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	nY, nX := currPos[0]+directions[currDir][0], currPos[1]+directions[currDir][1]
	return nY, nX
}

type Key struct {
	Y int
	X int
}

func day6Part1(graph [][]string) {
	currDir := 0
	currPos := findPosition(graph)
	visited := map[Key]bool{}
	for 0 <= currPos[0] && currPos[0] < len(graph) && 0 <= currPos[1] && currPos[1] < len(graph[0]) {
		visited[Key{currPos[0], currPos[1]}] = true
		nY, nX := getNextCoord(currPos, currDir)
		if !(0 <= nY && nY < len(graph) && 0 <= nX && nX < len(graph[0])) {
			break
		}
		nCell := graph[nY][nX]
		if nCell == "#" {
			currDir = (currDir + 1) % 4
			nY, nX = getNextCoord(currPos, currDir)
		}
		currPos[0], currPos[1] = nY, nX
	}
	fmt.Println(len(visited))
}
