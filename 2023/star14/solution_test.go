package star14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	cases := []struct {
		input  string
		output int
	}{
		{"", 1},
		{"AKQT9", 1},
		{"AAQT9", 2},
		{"AAQQ9", 3},
		{"AAAQT", 4},
		{"AAAKK", 5},
		{"AAAAQ", 6},
		{"AAAAA", 7},
		{"AAAAJ", 7},
		{"AAAJJ", 7},
		{"AAAJT", 6},
		{"AAJJT", 6},
		{"AKQJT", 2},
		{"AKQJJ", 4},
		{"AAQQJ", 5},
		{"AAQQJ", 5},
	}

	for _, c := range cases {
		v := calcStrength(c.input)
		assert.Equal(t, c.output, v, "input: %s wrong output: %d != %d", c.input, c.output, v)
	}
}
