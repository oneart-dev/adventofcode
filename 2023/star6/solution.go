package star6

import (
	"fmt"
	"strconv"
)

// Skipped graph structure and used map instead
func Fun(input []string) interface{} {
	result := 0
	for i, line := range input {
		prev := ""
		next := ""
		if i > 0 {
			prev = input[i-1]
		}

		if i < len(input)-1 {
			next = input[i+1]
		}

		calc(prev, line, next, i)
	}

	// TODO: dont like second loop here
	for _, v := range engine {
		if len(v) == 2 {
			result += v[0] * v[1]
			continue
		}
	}

	return result
}

var engine = map[string][]int{}

func calc(prev, input, next string, rowIndex int) {

	number := []rune{}
	foundSymbol := ""

	for i, char := range input {

		if !isNumber(char) {
			if len(number) > 0 && foundSymbol != "" {
				nint, _ := strconv.Atoi(string(number))
				engine[foundSymbol] = append(engine[foundSymbol], nint)
			}

			number = []rune{}
			foundSymbol = ""
			continue
		}

		number = append(number, char)

		if foundSymbol == "" {
			if s := checkAroundForSymbol(prev, input, next, i, rowIndex); s != "" {

				foundSymbol = s
			}
		}

	}

	if len(number) > 0 && foundSymbol != "" {
		nint, _ := strconv.Atoi(string(number))
		engine[foundSymbol] = append(engine[foundSymbol], nint)
	}

}

// Ugly but works
func checkAroundForSymbol(prev, input, next string, i int, rowI int) string {

	if i > 0 && isExistAndSymbol(input, i-1) {
		return fmt.Sprintf("%d-%d", rowI, i-1)
	}

	if i < len(input)-1 && isExistAndSymbol(input, i+1) {
		return fmt.Sprintf("%d-%d", rowI, i+1)
	}

	if prev != "" {
		if isExistAndSymbol(prev, i-1) {
			return fmt.Sprintf("%d-%d", rowI-1, i-1)
		}

		if isExistAndSymbol(prev, i+1) {
			return fmt.Sprintf("%d-%d", rowI-1, i+1)
		}

		if isExistAndSymbol(prev, i) {
			return fmt.Sprintf("%d-%d", rowI-1, i)
		}
	}

	if next != "" {
		if isExistAndSymbol(next, i-1) {
			return fmt.Sprintf("%d-%d", rowI+1, i-1)
		}

		if isExistAndSymbol(next, i+1) {
			return fmt.Sprintf("%d-%d", rowI+1, i+1)
		}

		if isExistAndSymbol(next, i) {
			return fmt.Sprintf("%d-%d", rowI+1, i)
		}
	}

	return ""
}

func isExistAndNumber(input string, index int) bool {
	if index < 0 || index >= len(input) {
		return false
	}

	r := []rune(input)
	return isNumber(r[index])
}

func isExistAndSymbol(input string, index int) bool {
	if index < 0 || index >= len(input) {
		return false
	}

	r := []rune(input)
	return !isNumber(r[index]) && r[index] != '.'
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}
