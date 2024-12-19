package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

var example9 = `2333133121414131402`

func isNumeric(r string) bool {
	_, err := strconv.Atoi(r)
	return err == nil
}

func calculateScore(diskmap []string) int {
	total := 0
	for idx, r := range diskmap {
		num, err := strconv.Atoi(r)
		if err == nil {
			total += (num * idx)
		}
	}
	return total
}

func buildDisk(input string) []string {
	nums := []int{}
	for _, v := range strings.Split(input, "") {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}
	diskmap := []string{}
	for idx, num := range nums {
		if idx%2 == 0 {
			for range num {
				str := strconv.Itoa(idx / 2)
				diskmap = append(diskmap, str)
			}
		} else {
			for range num {
				diskmap = append(diskmap, ".")
			}
		}
	}
	return diskmap
}

func day9Part1(input string) int {
	diskmap := buildDisk(input)
	left, right := 0, len(diskmap)-1
	for left < right {
		if isNumeric(diskmap[left]) {
			left++
			continue
		}
		if !isNumeric(diskmap[right]) {
			right--
			continue
		}
		swapSpaces(diskmap[right], left, right, diskmap)
	}
	return calculateScore(diskmap)
}

func checkSpaceLength(ptr int, diskmap []string) int {
	length := 0
	for {
		if isNumeric(diskmap[ptr]) {
			return length
		}
		length++
		ptr++
	}
}

func findSpace(filesize int, right int, diskmap []string) int {
	ptr := 0
	for ptr < right {
		if isNumeric(diskmap[ptr]) {
			ptr++
			continue
		}
		spaceLength := checkSpaceLength(ptr, diskmap)
		if spaceLength >= filesize {
			return ptr
		}
		ptr += spaceLength
	}
	return -1
}

func checkDataLength(data string, ptr int, diskmap []string) int {
	length := 0
	for ptr >= 0 {
		if diskmap[ptr] != data {
			return length
		}
		length++
		ptr--
	}
	return length
}

func swapSpaces(data string, left int, right int, diskmap []string) {
	for left < right {
		if diskmap[right] != data {
			return
		}
		diskmap[left], diskmap[right] = diskmap[right], diskmap[left]
		left++
		right--
	}
}

func day9Part2(input string) int {
	diskmap := buildDisk(input)
	right := len(diskmap) - 1
	for 0 <= right {
		if !isNumeric(diskmap[right]) {
			right--
			continue
		}
		dataLength := checkDataLength(diskmap[right], right, diskmap)
		// Find free space index
		space := findSpace(dataLength, right, diskmap)
		if space != -1 {
			swapSpaces(diskmap[right], space, right, diskmap)
		} else {
			right -= dataLength
		}
	}
	return calculateScore(diskmap)
}

func RunDay9(data []byte) {
	input := strings.TrimSpace(string(data))
	part1 := day9Part1(input)
	fmt.Println(part1)
	part2 := day9Part2(input)
	fmt.Println(part2)
}
