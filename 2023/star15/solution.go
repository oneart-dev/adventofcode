package star15

import (
	"strings"
)

func Fun(input []string) interface{} {
	dir := ""
	path := make(map[string][]string, len(input))

	for i, r := range input {
		if r == "" {
			continue
		}

		if i == 0 {
			dir = r
			continue
		}

		p := strings.Split(r, " = (")
		row := p[0]
		parts := strings.Split(strings.TrimSuffix(p[1], ")"), ", ")

		if _, ok := path[row]; ok {
			panic("should not reach here")
		}

		path[row] = parts
	}

	i := 0
	steps := 0
	newDir := ""
	pos := "AAA"
	for pos != "ZZZ" {
		if i == len(dir) {
			i = 0
		}

		newDir += string(dir[i])

		r := rune(dir[i])
		if r == rune('R') {
			pos = path[pos][1]
		} else if r == 'L' {
			pos = path[pos][0]
		} else {
			panic("should not reach here")
		}

		i++
		steps++
	}

	return steps
}
