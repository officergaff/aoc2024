package solutions

import (
	"fmt"
	"strings"
)

var example12 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func RunDay12(data []byte) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	regionMap := [][]string{}
	for _, line := range lines {
		regionMap = append(regionMap, strings.Split(line, ""))
	}
	part1 := day12Part1(regionMap)
	fmt.Println(part1)
}

type Plot struct {
	y int
	x int
}

func day12Part1(regionMap [][]string) int {
	visited := map[Plot]bool{}
	cost := 0
	for j, row := range regionMap {
		for i := range row {
			if !visited[Plot{j, i}] {
				area, perimeter := exploreRegion(j, i, regionMap, visited)
				cost += area * perimeter
			}
		}
	}
	return cost
}

func exploreRegion(r int, c int, regionMap [][]string, visited map[Plot]bool) (int, int) {
	plot := Plot{r, c}
	plotType := regionMap[r][c]
	visited[plot] = true
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	area, perimeter := 1, 0
	for _, d := range directions {
		dy, dx := r+d[0], c+d[1]
		if 0 <= dy && dy < len(regionMap) && 0 <= dx && dx < len(regionMap[0]) {
			if regionMap[dy][dx] == plotType {
				if visited[Plot{dy, dx}] == false {
					a, p := exploreRegion(dy, dx, regionMap, visited)
					area += a
					perimeter += p
				}
			} else {
				perimeter++
			}
		} else {
			perimeter++
		}
	}
	return area, perimeter
}
