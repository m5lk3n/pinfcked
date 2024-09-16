package main

import (
	"errors"
	"fmt"
	"log"
)

var numbers = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func getAdjacents(i int) ([3]int, error) {
	if i < 0 {
		return [3]int{}, errors.New("index must not be negative")
	}
	l := len(numbers)
	if i > l {
		return [3]int{}, errors.New(fmt.Sprintf("index must not be greater than %d", l))
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

func printError(err error) {
	if err != nil {
		log.Printf("%v\n", err)

	}
}

func main() {
	for i := 0; i < 10; i++ {
		res, _ := getAdjacents(i)
		fmt.Println(res)
	}

	_, err := getAdjacents(-1)
	printError(err)
	_, err = getAdjacents(11)
	printError(err)
}
