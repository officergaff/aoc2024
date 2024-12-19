package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

var example11 = `125 17`

func RunDay11(data []byte) {
	eStones := strings.Fields(example11)
	for range 6 {
		eStones = blink(eStones)
	}
	fmt.Println(eStones)
	stones := strings.Fields(strings.TrimSpace(string(data)))
	stoneMap := map[string]int{}
	for _, stone := range stones {
		stoneMap[stone] += 1
	}
	for range 25 {
		stoneMap = blinkMemo(stoneMap)
	}
	fmt.Println(sumMap(stoneMap))
	for range 50 {
		stoneMap = blinkMemo(stoneMap)
	}
	fmt.Println(sumMap(stoneMap))
}

func blink(stones []string) []string {
	nStones := []string{}
	for _, stone := range stones {
		if stone == "0" {
			nStones = append(nStones, "1")
		} else if len(stone)%2 == 0 {
			s1, s2 := stone[:len(stone)/2], stone[len(stone)/2:]
			s2n, _ := strconv.Atoi(s2)
			nStones = append(nStones, s1)
			nStones = append(nStones, strconv.Itoa(s2n))
		} else {
			n, _ := strconv.Atoi(stone)
			nStones = append(nStones, strconv.Itoa(n*2024))
		}
	}
	return nStones
}

func sumMap(stoneMap map[string]int) int {
	total := 0
	for _, v := range stoneMap {
		total += v
	}
	return total
}

func blinkMemo(stoneMap map[string]int) map[string]int {
	newMap := map[string]int{}
	for stone, count := range stoneMap {
		if stone == "0" {
			newMap["1"] += count
		} else if len(stone)%2 == 0 {
			s1, s2 := stone[:len(stone)/2], stone[len(stone)/2:]
			s2n, _ := strconv.Atoi(s2)
			s2 = strconv.Itoa(s2n)
			newMap[s1] += count
			newMap[s2] += count
		} else {
			n, _ := strconv.Atoi(stone)
			newMap[strconv.Itoa(n*2024)] += count
		}
	}
	return newMap
}
