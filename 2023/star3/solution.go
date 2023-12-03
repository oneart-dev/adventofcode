package star3

import (
	"log"
	"strconv"
	"strings"
)

var rules = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Fun(input []string) interface{} {
	result := 0
	for _, line := range input {
		result += calc(line)
	}

	return result
}

func calc(input string) int {
	if input == "" {
		return 0
	}

	p := strings.Split(input, ":")
	if len(p) < 1 {
		return 0
	}

	games := strings.Split(p[1], ";")
	for _, game := range games {

		rolls := strings.Split(game, ",")
		for _, roll := range rolls {

			t := strings.Split(strings.Trim(roll, " "), " ")

			num, err := strconv.Atoi(strings.Split(t[0], " ")[0])
			if err != nil {
				log.Panicf("Error converting num to int: %s", err)
			}

			if num > rules[t[1]] {
				return 0
			}

		}

	}

	gameID, err := strconv.Atoi(strings.Split(p[0], " ")[1])
	if err != nil {
		log.Panicf("Error converting gameID to int: %s", err)
	}

	return gameID
}
