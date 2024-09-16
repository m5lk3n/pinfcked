package main

import (
	"errors"
	"fmt"
	"log"
)

var numbers = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func getAdjacents(i int) error {
	if i < 0 {
		return errors.New("index must not be negative")
	}
	l := len(numbers)
	if i > l {
		return errors.New(fmt.Sprintf("index must not be greater than %d\n", i))
	}

	h := i - 1
	if i == 0 {
		h = l - 1
	}
	fmt.Printf("%d\n", numbers[h])
	fmt.Printf("%d*\n", i)
	j := i + 1
	if j > l {
		j = 0
	}
	fmt.Printf("%d\n", numbers[j])

	return nil
}

func printError(err error) {
	if err != nil {
		log.Printf("%v\n", err)

	}
}

func main() {
	printError(getAdjacents(1))
	printError(getAdjacents(0))

	printError(getAdjacents(-1))
	printError(getAdjacents(11))
}
