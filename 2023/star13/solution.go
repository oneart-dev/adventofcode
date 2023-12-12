package star13

import (
	"slices"
	"strconv"
	"strings"
)

type Race struct {
	Time int
	Dist int
}

var races = make([]Race, 0)

func Fun(input []string) interface{} {
	result := 0

	gameRank := make(map[int]map[string]int)

	for _, r := range input {
		if r == "" {
			continue
		}

		p := strings.Split(r, " ")
		game := p[0]
		num := p[1]

		if game == "" || num == "" {
			panic("should not reach here")
		}

		numInt, _ := strconv.Atoi(num)

		s := calcStrength(game)

		if _, ok := gameRank[s]; !ok {
			gameRank[s] = make(map[string]int)
		}

		gameRank[s][game] = numInt

	}

	sortedGames := make([]int, 0)

	for i := 7; i > 0; i-- {
		if _, ok := gameRank[i]; !ok {
			continue
		}

		v := gameRank[i]

		if len(v) == 1 {
			for s, v := range v {
				sortedGames = append(sortedGames, v)
			}

			continue
		}

		keys := make([]string, 0, len(v))
		for s, _ := range v {
			keys = append(keys, s)
		}

		slices.SortFunc(keys, func(a, b string) int {
			for i := 0; i < len(a); i++ {
				if charStrength[rune(a[i])] > charStrength[rune(b[i])] {
					return -1
				}

				if charStrength[rune(a[i])] < charStrength[rune(b[i])] {
					return 1
				}

				continue
			}

			return 0
		})

		for _, k := range keys {

			sortedGames = append(sortedGames, v[k])
		}

	}

	for i, v := range sortedGames {
		rank := len(sortedGames) - i
		if rank == 0 || v == 0 {
			panic("should not reach here")
		}
		result += rank * v
	}

	return result
}

var charStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func calc(r Race) int {
	n := 0

	for i := 1; i < r.Time; i++ {
		dist := i * (r.Time - i)
		if dist > r.Dist {
			n++
		}
	}

	return n
}

var strength = map[string]int{
	"5":          7,
	"4":          6,
	"full house": 5,
	"3":          4,
	"2pair":      3,
	"pair":       2,
	"high card":  1,
}

func calcStrength(input string) int {
	var cmap = make(map[rune]int, len(input))
	for _, c := range input {
		cmap[c]++
	}

	combinations := map[string]int{
		"3": 0,
		"2": 0,
	}

	for _, v := range cmap {
		switch v {
		case 5:
			return strength["5"]
		case 4:
			return strength["4"]
		case 3:
			combinations["3"]++
		case 2:
			combinations["2"]++
		}
	}

	if combinations["3"] > 0 && combinations["2"] > 0 {
		return strength["full house"]
	}

	if combinations["3"] > 0 {
		return strength["3"]
	}

	if combinations["2"] > 1 {
		return strength["2pair"]
	}

	if combinations["2"] > 0 {
		return strength["pair"]
	}

	return strength["high card"]
}
