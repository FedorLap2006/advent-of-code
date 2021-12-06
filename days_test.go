package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThird(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	output := SolveThird(input)

	assert.Equal(t, strconv.Itoa(9*22), output)
}

func TestThirdP2(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	output := SolveThirdP2(input)

	assert.Equal(t, strconv.Itoa(23*10), output)
}

func TestFourth(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`
	t.Run("pt1", func(t *testing.T) {
		output := SolveFourth(input)
		assert.Equal(t, strconv.Itoa(188*24), output)
	})
	t.Run("pt2", func(t *testing.T) {
		output := SolveFourthP2(input)
		assert.Equal(t, strconv.Itoa(148*13), output)
	})
}

func TestDay5(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, "5", SolveDay5(input))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, "12", SolveDay5P2(input))
	})
}

func TestDay6(t *testing.T) {
	input := `3,4,3,1,2`

	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, "5934", SolveDay6(input))
	})
	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, "26984457539", SolveDay6P2(input))
	})
}
