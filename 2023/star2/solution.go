package star2

import (
	"fmt"
	"strconv"
	"strings"
)

var stringDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func calc(input string, direction int) string {

	ss := ""
	i := 0

	if len(input) == 0 {
		return ""
	}

	if direction == -1 {
		i = len(input) - 1
	}

	for i < len(input) {
		if input[i] >= '0' && input[i] <= '9' {
			return string(input[i])
		}

		if direction > 0 {
			ss += string(input[i])
		} else {
			ss = string(input[i]) + ss
		}

		for stringDigit, stringNum := range stringDigits {
			if strings.Contains(ss, stringDigit) {
				return stringNum
			}
		}

		i += direction

	}

	return ""
}

func Fun(input []string) interface{} {
	result := 0
	for _, line := range input {
		nint, _ := strconv.Atoi(fmt.Sprintf("%s%s", calc(line, 1), calc(line, -1)))
		result += nint
	}

	return result
}
