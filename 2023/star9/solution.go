package star9

import (
	"strconv"
	"strings"
)

type DistSourceMap struct {
	Source int
	Dist   int
	Length int
}

var mapsOrder = make([]string, 0)
var seeds = make([]int, 0)
var iterations = make(map[string][]DistSourceMap, 0)

func Fun(input []string) interface{} {
	result := 0

	preProcess(input)

	for _, seed := range seeds {
		r := calc(seed)
		if r < result || result == 0 {
			result = r
		}
	}

	return result
}

func calc(seed int) int {
	n := seed
	for _, v := range mapsOrder {
		for _, dsm := range iterations[v] {
			if dsm.Source <= n && n <= dsm.Source+dsm.Length {
				n = dsm.Dist + (n - dsm.Source)
				break
			}
		}
	}

	return n
}

func preProcess(input []string) {
	mapIndex := ""
	for i, line := range input {
		if line == "" {
			mapIndex = ""
			continue
		}

		if i == 0 {
			p := strings.Split(strings.Split(line, ": ")[1], " ")
			for _, v := range p {
				nseed, _ := strconv.Atoi(v)
				seeds = append(seeds, nseed)
			}

			continue
		}

		if strings.Contains(line, "map") {
			mapIndex = strings.Split(line, " ")[0]
			mapsOrder = append(mapsOrder, mapIndex)
			continue
		}

		p := strings.Split(line, " ")
		source, _ := strconv.Atoi(p[1])
		dist, _ := strconv.Atoi(p[0])
		length, _ := strconv.Atoi(p[2])
		dsm := DistSourceMap{
			Source: source,
			Dist:   dist,
			Length: length,
		}

		iterations[mapIndex] = append(iterations[mapIndex], dsm)
	}
}
