package star8

import (
	"strconv"
	"strings"
)

func Fun(input []string) interface{} {
	result := 0
	var cardWinningCount = make(map[int]int, len(input))
	for i, line := range input {
		if line == "" {
			continue
		}

		c := calc(line)
		cardWinningCount[i] += 1
		if c > 0 {
			for j := i + 1; j <= i+c; j++ {
				cardWinningCount[j] += 1 * cardWinningCount[i]
			}
		}
	}

	for _, v := range cardWinningCount {
		result += v
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
			winCount += 1
		}
	}

	return winCount
}
