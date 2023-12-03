package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/oneart-dev/adventofcode/star1"
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
	}

	return stars[n]
}
