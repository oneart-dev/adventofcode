package star5

import (
	"fmt"
	"strconv"
)

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

		result += calc(prev, line, next)
	}

	return result
}

var engine = [][]rune{}

func calc(prev, input, next string) int {

	number := []rune{}
	foundSymbol := false
	sum := 0

	fmt.Println(input)

	for i, char := range input {
		if !isNumber(char) {
			if len(number) > 0 && foundSymbol {
				nint, _ := strconv.Atoi(string(number))
				sum += nint
			}

			number = []rune{}
			foundSymbol = false
			continue
		}

		number = append(number, char)

		fmt.Println("number", string(number))

		if !foundSymbol {
			foundSymbol = checkAroundForSymbol(prev, input, next, i)
		}

		fmt.Println("NOT Found")

	}

	if len(number) > 0 && foundSymbol {
		nint, _ := strconv.Atoi(string(number))
		sum += nint
	}

	return sum
}

func checkAroundForSymbol(prev, input, next string, i int) bool {
	if (i > 0 && isExistAndSymbol(input, i-1)) ||
		(i < len(input)-1 && isExistAndSymbol(input, i+1)) {
		return true
	}

	if prev != "" &&
		(isExistAndSymbol(prev, i-1) || isExistAndSymbol(prev, i+1) || isExistAndSymbol(prev, i)) {
		return true
	}

	if next != "" &&
		(isExistAndSymbol(next, i-1) || isExistAndSymbol(next, i+1) || isExistAndSymbol(next, i)) {
		return true
	}

	return false
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
