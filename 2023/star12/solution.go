package star12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	Time int
	Dist int
}

var races = make([]Race, 0)

func Fun(input []string) interface{} {
	result := 1

	space := regexp.MustCompile(`\s+`)
	t := space.ReplaceAllString(input[0], " ")
	d := space.ReplaceAllString(input[1], " ")
	time := strings.Split(strings.Split(t, ": ")[1], " ")
	fmt.Println(time)
	dist := strings.Split(strings.Split(d, ": ")[1], " ")
	fmt.Println(dist)
	for i := 0; i < len(time); i++ {
		if time[i] == "" {
			continue
		}
		ntime, _ := strconv.Atoi(time[i])
		ndist, _ := strconv.Atoi(dist[i])
		races = append(races, Race{
			Time: ntime,
			Dist: ndist,
		})
		fmt.Printf("Time: %d, Dist: %d\n", ntime, ndist)
	}

	for _, r := range races {
		result *= calc(r)
	}

	fmt.Printf("Result: %d\n", result)

	return result
}

func calc(r Race) int {
	n := 0

	fmt.Printf("Time: %d, Dist: %d\n", r.Time, r.Dist)

	for i := 1; i < r.Time; i++ {
		dist := i * (r.Time - i)
		if dist > r.Dist {
			fmt.Println(dist)
			n++
		}
	}

	return n
}
