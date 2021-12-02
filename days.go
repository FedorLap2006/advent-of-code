package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveSecond(input string) string {
	commands := strings.Split(input, "\n")
	var x, depth, aim int64
	for _, command := range commands {
		sectors := strings.Split(command, " ")
		amount, _ := strconv.ParseInt(sectors[1], 10, 16)

		switch sectors[0] {
		case "up":
			aim -= amount
		case "down":
			aim += amount
		case "forward":
			x += amount
			depth += aim * amount
		}
	}
	return fmt.Sprintf("%d %d %d", x, depth, x*depth)
}

func SolveFirst(input string) string {
	measurements := strings.Split(input, "\n")
	var prev, count int
	for i := 0; i < len(measurements)-2; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			v, _ := strconv.Atoi(measurements[i+j])
			sum += v
		}
		if i > 0 && sum > prev {
			count++
		}
		prev = sum
	}

	return fmt.Sprintf("%d", count)
}
