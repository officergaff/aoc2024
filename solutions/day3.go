package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunDay3(lines []string) {
	fullstring := strings.Join(lines, "\n")
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(fullstring, -1)
	fmt.Println(matches)
	sum := 0
	for _, expr := range matches {
		a, _ := strconv.Atoi(expr[1])
		b, _ := strconv.Atoi(expr[2])
		sum += a * b
	}
	fmt.Println(sum)
}
