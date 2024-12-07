package solutions

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var example = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

type Node struct {
	outNodes []int
	inDegree int
}

func RunDay5(data []byte) {
	partsE := strings.Split(example, "\n\n")
	rulesE := strings.Split(partsE[0], "\n")
	updatesE := strings.Split(partsE[1], "\n")
	fmt.Println(day5part1(rulesE, updatesE))

	parts := strings.Split(string(data), "\n\n")
	rules := strings.Split(parts[0], "\n")
	updates := strings.Split(parts[1], "\n")
	fmt.Println(day5part1(rules, updates))
}

// Build a topological order only with the numbers involved in the updates
// This will avoid circular graphs
func buildTopoMap(rules []string, updateNums []int) map[int]int {
	// Topological sorting
	// Start with building a topological graph
	graph := map[int]*Node{}
	for _, rule := range rules {
		ordering := strings.Split(rule, "|")
		first, _ := strconv.Atoi(ordering[0])
		second, _ := strconv.Atoi(ordering[1])
		if slices.Contains(updateNums, first) && slices.Contains(updateNums, second) {
			fNode, ok := graph[first]
			if !ok {
				fNode = &Node{}
				graph[first] = fNode
			}
			fNode.outNodes = append(fNode.outNodes, second)
			sNode, ok := graph[second]
			if !ok {
				sNode = &Node{}
				graph[second] = sNode
			}
			sNode.inDegree++
		}
	}
	// Figure out the order with a graph traversal
	queue := []int{}
	for k, v := range graph {
		if v.inDegree == 0 {
			queue = append(queue, k)
		}
	}
	order := []int{}
	for len(queue) > 0 {
		nKey := queue[0]
		queue = queue[1:]
		order = append(order, nKey)
		for _, nxt := range graph[nKey].outNodes {
			graph[nxt].inDegree--
			if graph[nxt].inDegree == 0 {
				queue = append(queue, nxt)
			}
		}
	}
	orderMap := map[int]int{}
	for i, num := range order {
		orderMap[num] = i
	}
	return orderMap
}

func day5part1(rules []string, updates []string) int {
	// Convert the updates into their ordering value
	total := 0
	for _, u := range updates {
		uFields := strings.Split(u, ",")
		nums := []int{}
		for _, num := range uFields {
			num, _ := strconv.Atoi(num)
			nums = append(nums, num)
		}
		orderMap := buildTopoMap(rules, nums)
		ord := []int{}
		for _, num := range nums {
			v, ok := orderMap[num]
			if ok {
				ord = append(ord, v)
			}
		}
		if sort.IntsAreSorted(ord) {
			mid := len(nums) / 2
			total += nums[mid]
		}
	}
	return total
}
