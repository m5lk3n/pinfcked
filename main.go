package main

import (
	"errors"
	"fmt"
	"log"
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

func printError(err error) {
	if err != nil {
		log.Printf("%v\n", err)
	}
}

func printCombinations(c [4][3]int) {
	for b, a := range c {
		fmt.Printf("%d ", b)
		for i := 0; i < len(a); i++ {
			fmt.Printf("%v", a[i])
		}
		fmt.Println()
	}

	combinations := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					fmt.Printf("%v%v%v%v\n", c[0][i], c[1][j], c[2][k], c[3][l])
					combinations++
				}
			}
		}
	}
	fmt.Println(combinations)
}

func main() {
	a1, _ := getAdjacents(0)
	a2, _ := getAdjacents(3)
	a3, _ := getAdjacents(5)
	a4, _ := getAdjacents(8)
	c := [4][3]int{
		a1,
		a2,
		a3,
		a4,
	}
	printCombinations(c)

	_, err := getAdjacents(-1)
	printError(err)
	_, err = getAdjacents(11)
	printError(err)
}
