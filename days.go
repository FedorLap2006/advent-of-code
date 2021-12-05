package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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
func SolveThird(input string) string {
	diagnostics := strings.Split(input, "\n")

	gamma, epsilon := "", ""
	for i := 0; i < len(diagnostics[0]); i++ {
		zero, one := 0, 0
		for _, d := range diagnostics {
			if d[i] == '0' {
				zero++
			} else if d[i] == '1' {
				one++
			}
		}
		if one >= zero {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammav, _ := strconv.ParseUint(gamma, 2, len(diagnostics[0]))
	epsilonv, _ := strconv.ParseUint(epsilon, 2, len(diagnostics[0]))
	return strconv.Itoa(int(gammav) * int(epsilonv))
}
func SolveThirdP2(input string) string {
	diagnostics := strings.Split(input, "\n")

	o2diagnostics := strings.Split(input, "\n")
	co2diagnostics := strings.Split(input, "\n")
	var o2, co2 string

	for i := 0; i < len(diagnostics[0]); i++ {
		var o2crit, co2crit byte
		zero, one := 0, 0
		for _, d := range o2diagnostics {
			if d[i] == '0' {
				zero++
			} else if d[i] == '1' {
				one++
			}
		}
		if one >= zero {
			o2crit = '1'
		} else {
			o2crit = '0'
		}
		zero, one = 0, 0
		for _, d := range co2diagnostics {
			if d[i] == '0' {
				zero++
			} else if d[i] == '1' {
				one++
			}
		}
		if one >= zero {
			co2crit = '0'
		} else {
			co2crit = '1'
		}

		var filteredO2, filteredCO2 []string
		for _, d := range o2diagnostics {
			if d[i] == o2crit {
				filteredO2 = append(filteredO2, d)
			}
		}
		for _, d := range co2diagnostics {
			if d[i] == co2crit {
				filteredCO2 = append(filteredCO2, d)
			}
		}

		o2diagnostics, co2diagnostics = filteredO2, filteredCO2
		if o2 == "" && len(o2diagnostics) == 1 {
			o2 = o2diagnostics[0]
		}
		if co2 == "" && len(co2diagnostics) == 1 {
			co2 = co2diagnostics[0]
		}

		if o2 != "" && co2 != "" {
			break
		}
	}
	o2v, _ := strconv.ParseUint(o2, 2, len(diagnostics[0]))
	co2v, _ := strconv.ParseUint(co2, 2, len(diagnostics[0]))
	return strconv.FormatUint(o2v*co2v, 10)
}

func SolveFourth(input string) string {
	data := strings.Split(input, "\n")
	var numbers []int
	for _, sv := range strings.Split(data[0], ",") {
		v, _ := strconv.Atoi(sv)
		numbers = append(numbers, v)
	}
	data = data[1:]

	boards := make([][]int, len(data)/6)
	for idx := 0; idx < len(data)/6; idx++ {
		boards[idx] = make([]int, 0, 5*5)
		for row := 0; row < 5; row++ {
			for _, value := range strings.Fields(data[1+idx*6+row]) {
				parsed, _ := strconv.Atoi(value)
				boards[idx] = append(boards[idx], parsed)
			}
		}
	}
	marked := make([]int32, len(boards))
	winner, winnerNum := 0, numbers[0]
	for _, num := range numbers {
		for idx, board := range boards {
			for boardIdx, value := range board {
				if value == num {
					marked[idx] |= 1 << boardIdx
				}
			}
			for i := 0; i < 5; i++ {
				if marked[idx]>>(i*5)&0x1f == 0x1f || marked[idx]&(0x108421<<i) == (0x108421<<i) {
					winner = idx
					winnerNum = num
					goto calcScore
				}
			}
		}
	}
calcScore:
	score := 0
	for i := 0; i < 5*5; i++ {
		if marked[winner]&(1<<i) == 0 {
			score += boards[winner][i]
		}
	}
	score *= winnerNum
	return fmt.Sprintf("%d", score)
}

func SolveFourthP2(input string) string {
	data := strings.Split(input, "\n")
	var numbers []int
	for _, sv := range strings.Split(data[0], ",") {
		v, _ := strconv.Atoi(sv)
		numbers = append(numbers, v)
	}
	data = data[1:]

	boards := make([][]int, len(data)/6)
	for idx := 0; idx < len(data)/6; idx++ {
		boards[idx] = make([]int, 0, 5*5)
		for row := 0; row < 5; row++ {
			for _, value := range strings.Fields(data[1+idx*6+row]) {
				parsed, _ := strconv.Atoi(strings.TrimSpace(value))
				boards[idx] = append(boards[idx], parsed)
			}
		}
	}

	marked := make([]int32, len(boards))
	winners, winnerNum := make([]int, 0, len(boards)), numbers[0]
	for _, num := range numbers {
		for idx, board := range boards {
			for _, v := range winners {
				if idx == v {
					goto nextIteration
				}
			}

			for boardIdx, value := range board {
				if value == num {
					marked[idx] |= 1 << boardIdx
				}
			}
			for i := 0; i < 5; i++ {
				if marked[idx]>>(i*5)&0x1f == 0x1f || marked[idx]&(0x108421<<i) == (0x108421<<i) {
					winners = append(winners, idx)
					winnerNum = num
					if len(winners) == len(boards) {
						goto calcScore
					}
					break
				}
			}
		nextIteration:
		}
	}
calcScore:
	score := 0
	for i := 0; i < 5*5; i++ {
		if marked[winners[len(winners)-1]]&(1<<i) == 0 {
			score += boards[winners[len(winners)-1]][i]
		}
	}
	return fmt.Sprintf("%d", score*winnerNum)
}

const spaceDim = 1000

func printSpace(space [][spaceDim]int) {
	for i := 0; i < spaceDim; i++ {
		for j := 0; j < spaceDim; j++ {
			if space[i][j] == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(space[i][j])
		}
		fmt.Println()
	}
}

func coverVent(space [][spaceDim]int, x, y, x2, y2 int, diagonals bool) {
	xMax, xMin := int(math.Max(float64(x), float64(x2))), int(math.Min(float64(x), float64(x2)))
	yMax, yMin := int(math.Max(float64(y), float64(y2))), int(math.Min(float64(y), float64(y2)))

	if diagonals && xMax-xMin == yMax-yMin {
		reverse := (x > x2 || y > y2) && !(x > x2 && y > y2)
		startX, startY := xMin, yMin
		if reverse {
			startX, startY = xMin, yMax
		}
		for i := 0; i <= yMax-yMin; i++ {
			for j := 0; j <= xMax-xMin; j++ {
				if i == j {
					if reverse {
						space[startY-i][startX+j]++
					} else {
						space[startY+i][startX+j]++
					}
				}
			}
		}

	} else if y == y2 {
		for i := xMin; i <= xMax; i++ {
			space[y][i]++
		}
	} else if x == x2 {
		for i := yMin; i <= yMax; i++ {
			space[i][x]++
		}
	}
}

func SolveDay5(input string) string {
	var space [][spaceDim]int = make([][spaceDim]int, spaceDim)
	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, " -> ")
		var pairs [2][2]int
		for i := 0; i < 2; i++ {
			for idx, v := range strings.Split(coords[i], ",") {
				pairs[i][idx], _ = strconv.Atoi(v)
			}
		}
		coverVent(space, pairs[0][0], pairs[0][1], pairs[1][0], pairs[1][1], false)
	}
	count := 0
	for i := 0; i < spaceDim; i++ {
		for j := 0; j < spaceDim; j++ {
			if space[i][j] >= 2 {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func SolveDay5P2(input string) string {
	var space [][spaceDim]int = make([][spaceDim]int, spaceDim)
	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, " -> ")
		var pairs [2][2]int
		for i := 0; i < 2; i++ {
			for idx, v := range strings.Split(coords[i], ",") {
				pairs[i][idx], _ = strconv.Atoi(v)
			}
		}
		coverVent(space, pairs[0][0], pairs[0][1], pairs[1][0], pairs[1][1], true)
	}
	// coverVent(space, 3, 3, 1, 1, true) // false
	// coverVent(space, 1, 1, 3, 3, true) // false
	// coverVent(space, 9, 7, 7, 9, true) // true
	// coverVent(space, 7, 9, 9, 7, true) // true
	// if x > x2 {
	// 	for i := 0; i <= y-y2; i++ {
	// 		fmt.Println("khm")
	// 		for j := 0; j <= x-x2; j++ {
	// 			if i == j {
	// 				space[y+i][x-j]++
	// 			}
	// 		}
	// 	}

	// } else {
	// 	for i := 0; i <= y2-y; i++ {
	// 		for j := 0; j <= x2-x; j++ {
	// 			hidx := x + j
	// 			if i == j {
	// 				space[y+i][hidx]++
	// 			}
	// 		}
	// 	}

	// }
	// space[7][9]++
	// space[8][8]++
	// space[9][7]++
	/*
		for i := 0; i <= 9-7; i++ {
			fmt.Println("khm")
			for j := 0; j <= 9-7; j++ {
				if i == j {
					space[7+i][9-j]++
				}
			}
		}

	*/
	// printSpace(space)
	count := 0
	for i := 0; i < spaceDim; i++ {
		for j := 0; j < spaceDim; j++ {
			if space[i][j] >= 2 {
				count++
			}
		}
	}

	return strconv.Itoa(count)

}
