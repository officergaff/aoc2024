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

func RunDay9(data []byte) {
	nums := []int{}
	input := strings.TrimSpace(string(data))
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
	fmt.Println(total)
}
