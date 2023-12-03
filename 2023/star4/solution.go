package star4

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

	rolledColors := map[string]int{
		"red":   -1,
		"green": -1,
		"blue":  -1,
	}

	games := strings.Split(p[1], ";")
	for _, game := range games {
		rolls := strings.Split(game, ",")
		for _, roll := range rolls {
			t := strings.Split(strings.Trim(roll, " "), " ")
			rolledColor := t[1]
			rolledNumber, err := strconv.Atoi(strings.Split(t[0], " ")[0])
			if err != nil {
				log.Panicf("Error converting num to int: %s", err)
			}

			if rolledNumber > rolledColors[rolledColor] {
				rolledColors[rolledColor] = rolledNumber
			}
		}
	}

	v := rolledColors["red"] * rolledColors["green"] * rolledColors["blue"]
	if v < 0 {
		return v * -1
	}

	return v
}
