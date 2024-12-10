package solutions

import (
	"fmt"
	"strings"
)

var example8 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

type Point struct {
	Y int
	X int
}

func getAntiNodes(a Point, b Point) []Point {
	yDist, xDist := a.Y-b.Y, a.X-b.X
	return []Point{{a.Y + yDist, a.X + xDist}, {b.Y - yDist, b.X - xDist}}
}

func addNodes(nodeMap map[Point]bool, nodes []Point, maxRow int, maxCol int) {
	for _, node := range nodes {
		if 0 <= node.Y && node.Y < maxRow && 0 <= node.X && node.X < maxCol {
			nodeMap[node] = true
		}
	}
}

func day8Part1(graph [][]string) int {
	maxRow, maxCol := len(graph), len(graph[0])
	signalCoords := map[string][]Point{}
	for j, line := range graph {
		for i, cell := range line {
			if cell == "." {
				continue
			}
			signalCoords[cell] = append(signalCoords[cell], Point{j, i})
		}
	}
	antiNodes := map[Point]bool{}
	for _, antennas := range signalCoords {
		for curr, antenna := range antennas {
			for idx, other := range antennas {
				if curr == idx {
					continue
				}
				nodes := getAntiNodes(antenna, other)
				addNodes(antiNodes, nodes, maxRow, maxCol)
			}
		}
	}
	return len(antiNodes)
}

func RunDay8(data []byte) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	graph := [][]string{}
	for _, line := range lines {
		graph = append(graph, strings.Split(line, ""))
	}
	part1 := day8Part1(graph)
	fmt.Println(part1)
}
