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
