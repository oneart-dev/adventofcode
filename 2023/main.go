package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/oneart-dev/adventofcode/star1"
	"github.com/oneart-dev/adventofcode/star10"
	"github.com/oneart-dev/adventofcode/star11"
	"github.com/oneart-dev/adventofcode/star12"
	"github.com/oneart-dev/adventofcode/star13"
	"github.com/oneart-dev/adventofcode/star14"
	"github.com/oneart-dev/adventofcode/star15"
	"github.com/oneart-dev/adventofcode/star16"
	"github.com/oneart-dev/adventofcode/star17"
	"github.com/oneart-dev/adventofcode/star2"
	"github.com/oneart-dev/adventofcode/star3"
	"github.com/oneart-dev/adventofcode/star4"
	"github.com/oneart-dev/adventofcode/star5"
	"github.com/oneart-dev/adventofcode/star6"
	"github.com/oneart-dev/adventofcode/star7"
	"github.com/oneart-dev/adventofcode/star8"
	"github.com/oneart-dev/adventofcode/star9"
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
		"1":  star1.Fun,
		"2":  star2.Fun,
		"3":  star3.Fun,
		"4":  star4.Fun,
		"5":  star5.Fun,
		"6":  star6.Fun,
		"7":  star7.Fun,
		"8":  star8.Fun,
		"9":  star9.Fun,
		"10": star10.Fun,
		"11": star11.Fun,
		"12": star12.Fun,
		"13": star13.Fun,
		"14": star14.Fun,
		"15": star15.Fun,
		"16": star16.Fun,
		"17": star17.Fun,
	}

	if _, ok := stars[n]; !ok {
		panic("Star not found")
	}

	return stars[n]
}
