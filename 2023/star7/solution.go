package star7

import (
	"strconv"
	"strings"
)

// Skipped graph structure and used map instead
func Fun(input []string) interface{} {
	result := 0
	for _, line := range input {
		result += calc(line)
	}

	return result
}

func calc(line string) int {
	winCount := 0
	if line == "" {
		return 0
	}

	p := strings.Split(line, ":")
	if len(p) < 2 {
		return 0
	}

	p2 := strings.Split(p[1], " | ")
	if len(p2) < 2 {
		return 0
	}

	winningNumbersString := strings.Split(p2[0], " ")
	myNumbersString := strings.Split(p2[1], " ")

	winningNumbers := make(map[int]bool, len(winningNumbersString))

	for _, v := range winningNumbersString {
		if v == "" {
			continue
		}

		key, _ := strconv.Atoi(v)
		winningNumbers[key] = true
	}

	for _, v := range myNumbersString {
		if v == "" {
			continue
		}

		vv, _ := strconv.Atoi(v)
		if _, ok := winningNumbers[vv]; ok {
			if winCount == 0 {
				winCount = 1
			} else {
				winCount *= 2
			}
		}
	}

	return winCount
}
