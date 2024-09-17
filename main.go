package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var numbers = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // int because later operations are applied, like len(), which return an int

func getAdjacents(i int) ([3]int, error) {
	if i < 0 {
		return [3]int{}, errors.New("index must not be negative")
	}
	l := len(numbers)
	if i > l {
		return [3]int{}, fmt.Errorf("index must not be greater than %d", l)
	}

	h := i - 1
	if i == 0 {
		h = l - 1
	}
	j := i + 1
	if j >= l {
		j = 0
	}

	return [...]int{numbers[h], i, numbers[j]}, nil
}

func printCombinations(c [4][3]int) {
	n := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					fmt.Printf("%v%v%v%v\n", c[0][i], c[1][j], c[2][k], c[3][l])
					n++
				}
			}
		}
	}
	log.Debugf("%d combinations", n)
}

func initDebug() { // don't make this init() because it would fail the tests
	var debug bool
	flag.BoolVar(&debug, "debug", false, "enable/disable debug mode")
	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}
}

func getDigit(b byte) int {
	s := string(b)
	d, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid digit: %s\n", s)
	}

	return d
}

func main() {
	initDebug()

	pin := flag.Arg(0)
	if len(pin) != 4 {
		log.Fatalln("Suspected PIN must be 4 digits long, where each digit is a number between 0 and 9")
	}
	p0 := getDigit(pin[0])
	p1 := getDigit(pin[1])
	p2 := getDigit(pin[2])
	p3 := getDigit(pin[3])

	r1, _ := getAdjacents(p0)
	r2, _ := getAdjacents(p1)
	r3, _ := getAdjacents(p2)
	r4, _ := getAdjacents(p3)
	c := [4][3]int{
		r1,
		r2,
		r3,
		r4,
	}
	printCombinations(c)
}
