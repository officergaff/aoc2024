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

func getAntiNodes1(a Point, b Point) []Point {
	// Creating a vector from a->b
	yDist, xDist := b.Y-a.Y, b.X-a.X
	return []Point{{a.Y - yDist, a.X - xDist}, {b.Y + yDist, b.X + xDist}}
}

func getAntiNodes2(a Point, b Point, maxRow int, maxCol int) []Point {
	yDist, xDist := b.Y-a.Y, b.X-a.X
	nodes := []Point{}
	for i := 0; ; i-- {
		node := Point{a.Y - (yDist * i), a.X - (xDist * i)}
		if !(0 <= node.Y && node.Y < maxRow && 0 <= node.X && node.X < maxCol) {
			break
		}
		nodes = append(nodes, node)
	}
	for i := 1; ; i++ {
		node := Point{a.Y + (yDist * i), a.X + (xDist * i)}
		if !(0 <= node.Y && node.Y < maxRow && 0 <= node.X && node.X < maxCol) {
			break
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func addNodes(nodeMap map[Point]bool, nodes []Point, maxRow int, maxCol int) {
	for _, node := range nodes {
		if 0 <= node.Y && node.Y < maxRow && 0 <= node.X && node.X < maxCol {
			nodeMap[node] = true
		}
	}
}

func getSignalCoords(graph [][]string) map[string][]Point {
	signalCoords := map[string][]Point{}
	for j, line := range graph {
		for i, cell := range line {
			if cell == "." {
				continue
			}
			signalCoords[cell] = append(signalCoords[cell], Point{j, i})
		}
	}
	return signalCoords
}

func day8Part1(graph [][]string) int {
	maxRow, maxCol := len(graph), len(graph[0])
	antiNodes := map[Point]bool{}
	signalCoords := getSignalCoords(graph)
	for _, antennas := range signalCoords {
		for curr, antenna := range antennas {
			for idx, other := range antennas {
				if curr == idx {
					continue
				}
				nodes := getAntiNodes1(antenna, other)
				addNodes(antiNodes, nodes, maxRow, maxCol)
			}
		}
	}
	return len(antiNodes)
}

func day8Part2(graph [][]string) int {
	maxRow, maxCol := len(graph), len(graph[0])
	antiNodes := map[Point]bool{}
	signalCoords := getSignalCoords(graph)
	for _, antennas := range signalCoords {
		for curr, antenna := range antennas {
			for idx, other := range antennas {
				if curr == idx {
					continue
				}
				nodes := getAntiNodes2(antenna, other, maxRow, maxCol)
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
	part2 := day8Part2(graph)
	fmt.Println(part2)
}
