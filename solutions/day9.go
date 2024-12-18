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
		diskmap[left], diskmap[right] = diskmap[right], diskmap[left]
		left++
		right--
	}
	total := 0
	for idx, r := range diskmap {
		num, err := strconv.Atoi(r)
		if err == nil {
			total += (num * idx)
		}
	}
	return total
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

func findSpaces(filesize int, left int, right int, diskmap []string) map[int]int {
	spaces := map[int]int{}
	for left < right {
		if isNumeric(diskmap[left]) {
			left++
			continue
		}
		spaceLength := checkSpaceLength(left, diskmap)
		if spaceLength >= filesize {
			spaces[left] = spaceLength
		}
		left += spaceLength
	}
	return spaces
}

func checkDataLength(data string, ptr int, diskmap []string) int {
	length := 0
	for {
		if diskmap[ptr] != data {
			return length
		}
		length++
		ptr--
	}
}

func swapSpaces(data string, left int, right int, diskmap []string) {
	for {
		if diskmap[right] != data {
			return
		}
		diskmap[left], diskmap[right] = diskmap[right], diskmap[left]
		left++
		right--
	}
}

/*
Todo: instead of a map, find the earliest free space available, return nil if there is no such thing
*/
func day9Part2(input string) int {
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
		dataLength := checkDataLength(diskmap[right], right, diskmap)
		spaces := findSpaces(dataLength, left, right, diskmap)
		swapped := false
		fmt.Printf("datalength of: %v, spaces available: %v, diskmap state: %v\n", dataLength, spaces, diskmap)
		for k, v := range spaces {
			if dataLength <= v && !swapped {
				swapSpaces(diskmap[right], k, right, diskmap)
				right -= dataLength
				swapped = true
			}
		}
		if !swapped {
			right -= dataLength
		}
	}
	total := 0
	for idx, r := range diskmap {
		num, err := strconv.Atoi(r)
		if err == nil {
			total += (num * idx)
		}
	}
	fmt.Println(diskmap)
	return total
}

func RunDay9(data []byte) {
	// input := strings.TrimSpace(string(data))
	// part1 := day9Part1(input)
	// fmt.Println(part1)
	part2 := day9Part2(example9)
	fmt.Println(part2)
}
