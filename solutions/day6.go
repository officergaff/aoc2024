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
	day6Part1(graphE)
	day6Part2(graphE)
	graph := process6(strings.TrimSpace(string(data)))
	day6Part1(graph)
	day6Part2(graph)
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

func day6Part1(graph [][]string) [][]int {
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
	arr := [][]int{}
	for k := range visited {
		arr = append(arr, []int{k.Y, k.X})
	}
	return arr
}

type Key2 struct {
	Y   int
	X   int
	dir int
}

func isLoop(graph [][]string) bool {
	currDir := 0
	currPos := findPosition(graph)
	visited := map[Key2]bool{}
	for 0 <= currPos[0] && currPos[0] < len(graph) && 0 <= currPos[1] && currPos[1] < len(graph[0]) {
		_, ok := visited[Key2{currPos[0], currPos[1], currDir}]
		if !ok {
			visited[Key2{currPos[0], currPos[1], currDir}] = true
		} else {
			return true
		}
		nY, nX := getNextCoord(currPos, currDir)
		if !(0 <= nY && nY < len(graph) && 0 <= nX && nX < len(graph[0])) {
			return false
		}
		nCell := graph[nY][nX]
		if nCell == "#" {
			currDir = (currDir + 1) % 4
		} else {
			// very important else statement
			// distinct from the part 1, where we don't really care about the orientation of the guard
			// here we're making sure to record each orientation change in the visited map
			currPos[0], currPos[1] = nY, nX
		}
	}
	return false
}

func day6Part2(graph [][]string) {
	vArr := day6Part1(graph)
	total := 0
	for _, coord := range vArr {
		y, x := coord[0], coord[1]
		if graph[y][x] != "^" {
			graph[y][x] = "#"
			if isLoop(graph) {
				total++
			}
			graph[y][x] = "."
		}
	}
	fmt.Println(total)
}
