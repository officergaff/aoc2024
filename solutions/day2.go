package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func RunDay2(data []byte) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	part1 := day2Part1(lines)
	part2 := day2Part2(lines)
	fmt.Println(part1)
	fmt.Println(part2)
}

func toNums(line string) []int {
	nums := []int{}
	for _, num := range strings.Fields(line) {
		if n, err := strconv.Atoi(num); err == nil {
			nums = append(nums, n)
		}
	}
	return nums
}

func checkSafe(nums []int) bool {
	curr := nums[0]
	next := nums[1]
	isDecreasing := curr > next
	for i := range len(nums) - 1 {
		curr = nums[i]
		next = nums[i+1]
		if isDecreasing != (curr > next) {
			return false
		} else {
			diff := abs(curr - next)
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}

func day2Part1(lines []string) int {
	safecount := 0
	for _, line := range lines {
		isSafe := checkSafe(toNums(line))
		if isSafe {
			safecount++
		}
	}
	return safecount
}

func day2Part2(lines []string) int {
	safecount := 0
	for _, line := range lines {
		nums := toNums(line)
		isSafe := checkSafe(nums)
		if isSafe {
			safecount++
		} else {
			for i := range nums {
				modNums := make([]int, len(nums)-1)
				copy(modNums, nums[:i])
				copy(modNums[i:], nums[i+1:])
				if checkSafe(modNums) {
					isSafe = true
				}
			}
			if isSafe {
				safecount++
			}
		}
	}
	return safecount
}
