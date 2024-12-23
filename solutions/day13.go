package solutions

import (
	"fmt"
	"math"
	"strings"
)

var example13 = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

/*
Solved with Hrishi:
a * ax + b * bx = X
a * ay + b * by = Y
*/
func RunDay13(data []byte) {
	machines := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	part1 := 0
	part2 := 0
	for _, machine := range machines {
		var ax, ay, bx, by, X, Y float64
		fmt.Sscanf(machine, "Button A: X+%v, Y+%v\nButton B: X+%v, Y+%v\nPrize: X=%v, Y=%v", &ax, &ay, &bx, &by, &X, &Y)
		part1 += solve(ax, ay, bx, by, X, Y)
		part2 += solve(ax, ay, bx, by, X+10000000000000, Y+10000000000000)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func solve(ax, ay, bx, by, X, Y float64) int {
	delta := 0.01
	b := (X*ay - ax*Y) / (bx*ay - ax*by)
	b_round := math.Round(b)
	if math.Abs(b-b_round) > delta {
		return 0
	}
	a := Y/ay - by/ay*b_round
	a_round := math.Round(a)
	if math.Abs(a-a_round) > delta {
		return 0
	}
	return int(3*a_round + b_round)
}
