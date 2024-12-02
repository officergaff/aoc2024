package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func RunDay1(lines []string) {
    l1 := make([]int, len(lines))
    l2 := make([]int, len(lines))

    for i, line := range lines {
        nums := strings.Fields(line)
        l1[i], _ = strconv.Atoi(nums[0])
        l2[i], _ = strconv.Atoi(nums[1])
    }
    fmt.Println(day1Part1(l1, l2))
    fmt.Println(day1Part2(l1, l2))
}

func abs(n int) int {
    if n < 0 { 
        return -n 
    } 
    return n
}

func day1Part1(l1 []int, l2 []int) int {
    difference := 0
    sort.Ints(l1)
    sort.Ints(l2)
    
    for i := range len(l1) {
        difference += abs(l1[i] - l2[i])
    }
    return difference
}

func day1Part2(l1 []int, l2 []int) int {
    counts := make(map[int]int)
    for _, num := range l2 {
        counts[num]++
    }
    score := 0
    for _, num := range l1 {
        score += counts[num] * num
    }
    return score
}
