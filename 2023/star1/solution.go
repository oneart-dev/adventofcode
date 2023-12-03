package star1

import (
	"strconv"
)

func Fun(input []string) interface{} {

	result := 0

	for _, line := range input {
		num := []rune{}
		for _, char := range line {
			if char >= '0' && char <= '9' {
				switch len(num) {
				case 0, 1:
					num = append(num, char)
				case 2:
					num[1] = char
				}
			}
		}

		switch len(num) {
		case 0:
			continue
		case 1:
			num = append(num, num[0])
		}
		nint, _ := strconv.Atoi(string(num))
		result += nint
	}

	return result
}
