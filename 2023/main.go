package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/oneart-dev/adventofcode/star1"
	"github.com/oneart-dev/adventofcode/star2"
	"github.com/oneart-dev/adventofcode/star3"
	"github.com/oneart-dev/adventofcode/star4"
)

var name string

func main() {
	// parse string flag
	flag.StringVar(&name, "n", "1", "start number")
	flag.Parse()

	input, err := loadFile(fmt.Sprintf("star%s/input", name))
	if err != nil {
		panic(err)
	}

	// call function
	result := starsGenerator(name)(input)
	fmt.Println(result)
}

func loadFile(path string) ([]string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(f), "\n"), nil
}

type fun func(input []string) interface{}

func starsGenerator(n string) fun {
	stars := map[string]fun{
		"1": star1.Fun,
		"2": star2.Fun,
		"3": star3.Fun,
		"4": star4.Fun,
	}

	if _, ok := stars[n]; !ok {
		panic("Star not found")
	}

	return stars[n]
}
