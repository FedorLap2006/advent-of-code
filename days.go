package main

import (
	"container/list"
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

type ventSpace [spaceDim][spaceDim]int

func (space *ventSpace) mapVent(x, y, x2, y2 int, diagonal bool) {
	xMax, xMin := int(math.Max(float64(x), float64(x2))), int(math.Min(float64(x), float64(x2)))
	yMax, yMin := int(math.Max(float64(y), float64(y2))), int(math.Min(float64(y), float64(y2)))

	if diagonal && xMax-xMin == yMax-yMin {
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

func (space *ventSpace) countSafe(minIntersections int) (count int) {
	for i := 0; i < spaceDim; i++ {
		for j := 0; j < spaceDim; j++ {
			if space[i][j] >= minIntersections {
				count++
			}
		}
	}
	return
}
func parseVentPoints(input string) (points [][2][2]int) {
	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, " -> ")
		var pairs [2][2]int
		for i := 0; i < 2; i++ {
			for idx, v := range strings.Split(coords[i], ",") {
				pairs[i][idx], _ = strconv.Atoi(v)
			}
		}
		points = append(points, pairs)
	}
	return
}

func SolveDay5(input string) string {
	space := new(ventSpace)
	points := parseVentPoints(input)
	for _, p := range points {
		space.mapVent(p[0][0], p[0][1], p[1][0], p[1][1], false)
	}
	return strconv.Itoa(space.countSafe(2))
}

func SolveDay5P2(input string) string {
	space := new(ventSpace)
	points := parseVentPoints(input)
	for _, p := range points {
		space.mapVent(p[0][0], p[0][1], p[1][0], p[1][1], true)
	}

	return strconv.Itoa(space.countSafe(2))
}

type Lanternfish int

func (t *Lanternfish) Reproduce() (Lanternfish, bool) {
	if *t < 0 {
		*t = 6
		return 9, true
	}
	return 0, false
}

func (t *Lanternfish) Tick() {
	*t--
}

func printAllFish(fish *list.List) {
	count := [11]int{}
	for e := fish.Front(); e != nil; e = e.Next() {
		count[int(e.Value.(Lanternfish))]++
	}
	fmt.Println(count)
}

func SolveDay6(input string) string {
	var fish = list.New()
	for _, v := range strings.Split(input, ",") {
		age, _ := strconv.Atoi(v)
		fish.PushBack(Lanternfish(age))
	}
	// fmt.Printf("Initial state: ")
	// printAllFish(fish)

	for i := 0; i < 80; i++ {
		for e := fish.Front(); e != nil; e = e.Next() {
			current := e.Value.(Lanternfish)
			current.Tick()
			if newFish, ready := current.Reproduce(); ready {
				fish.PushBack(newFish)
			}
			e.Value = current
		}
		printAllFish(fish)
		// fmt.Printf("After %d days: ", i+1)
		// printAllFish(fish)
	}

	return fmt.Sprintf("%d", fish.Len())
}

func SolveDay6P2(input string) string {
	var fish [8 + 1]int
	for _, v := range strings.Split(input, ",") {
		timer, _ := strconv.Atoi(v)
		fish[timer]++
	}

	for i := 0; i < 256; i++ {
		var zero int = fish[0]
		for timer, count := range fish[1:] {
			fish[timer] = count
		}
		fish[6] += zero
		fish[8] = zero
	}
	count := 0
	for _, v := range fish {
		count += v
	}
	return fmt.Sprintf("%d", count)
}

type CrabSubmarine struct {
	Position int
}

func (c CrabSubmarine) MovingCost(position int) int {
	n := int(math.Abs(float64(c.Position - position)))
	return n * (n + 1) / 2
}

func SolveDay7(input string) string {
	positions := []int{}
	for _, posRaw := range strings.Split(input, ",") {
		position, _ := strconv.Atoi(posRaw)
		positions = append(positions, position)
	}

	min := math.MaxInt32
	for i, v := range positions {
		sum := 0
		for j, other := range positions {
			if j == i {
				continue
			}
			// fmt.Printf("Move from %d to %d: %d fuel\n", v, other, int(math.Abs(float64(other-v))))
			sum += int(math.Abs(float64(other - v)))
		}
		// fmt.Printf("Sum (based on %d): %d\n", v, sum)
		if sum < min {
			// fmt.Println("Found minimum")
			min = sum
		}
	}

	return fmt.Sprintf("%d", min)
}

func SolveDay7P2(input string) string {
	positions := []CrabSubmarine{}
	maxPosition := 0
	for _, posRaw := range strings.Split(input, ",") {
		position, _ := strconv.Atoi(posRaw)
		if position > maxPosition {
			maxPosition = position
		}
		positions = append(positions, CrabSubmarine{Position: position})
	}
	min := math.MaxInt32
	for v := 0; v <= maxPosition*2; v++ {
		sum := 0
		for _, other := range positions {
			// fmt.Printf("Move from %d to %d: %d fuel\n", v, other, other.MovingCost(v))
			sum += other.MovingCost(v)
		}
		// fmt.Printf("Sum (based on %d): %d\n", v, sum)
		if sum < min {
			// fmt.Println("Found minimum")
			min = sum
		}
	}

	return fmt.Sprintf("%d", min)
}
