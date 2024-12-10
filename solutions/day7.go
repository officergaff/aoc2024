package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

var example7 = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func backtrack(goal int, nums []int, currNum int) bool {
	if len(nums) == 0 {
		if goal == currNum {
			return true
		} else {
			return false
		}
	}
	return backtrack(goal, nums[1:], currNum*nums[0]) || backtrack(goal, nums[1:], currNum+nums[0])
}

func concat(num1 int, num2 int) int {
	conv := strconv.Itoa(num1) + strconv.Itoa(num2)
	vonc, _ := strconv.Atoi(conv)
	return vonc
}

func backtrack1(goal int, nums []int, currNum int) bool {
	if len(nums) == 0 {
		if goal == currNum {
			return true
		} else {
			return false
		}
	}
	return backtrack1(goal, nums[1:], currNum*nums[0]) || backtrack1(goal, nums[1:], currNum+nums[0]) || backtrack1(goal, nums[1:], concat(currNum, nums[0]))
}

func day7Part1(lines []string) int {
	total := 0
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		goal, _ := strconv.Atoi(fields[0])
		nums := []int{}
		for _, num := range strings.Fields(fields[1]) {
			num, _ := strconv.Atoi(num)
			nums = append(nums, num)
		}
		if backtrack(goal, nums[1:], nums[0]) {
			total += goal
		}
	}
	return total
}

func day7Part2(lines []string) int {
	total := 0
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		goal, _ := strconv.Atoi(fields[0])
		nums := []int{}
		for _, num := range strings.Fields(fields[1]) {
			num, _ := strconv.Atoi(num)
			nums = append(nums, num)
		}
		if backtrack1(goal, nums[1:], nums[0]) {
			total += goal
		}
	}
	return total
}

func RunDay7(data []byte) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	// linesE := strings.Split(example7, "\n")
	part1 := day7Part1(lines)
	fmt.Println(part1)
	part2 := day7Part2(lines)
	fmt.Println(part2)
}
