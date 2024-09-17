package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdjacents(t *testing.T) {
	tests := []struct {
		index    int
		expected [3]int
		err      error
	}{
		{0, [3]int{9, 0, 1}, nil},
		{3, [3]int{2, 3, 4}, nil},
		{5, [3]int{4, 5, 6}, nil},
		{8, [3]int{7, 8, 9}, nil},
		{9, [3]int{8, 9, 0}, nil},
		{-1, [3]int{}, errors.New("index must not be negative")},
		{11, [3]int{}, errors.New("index must not be greater than 10")},
	}

	for _, test := range tests {
		result, err := getAdjacents(test.index)
		if test.err != nil {
			assert.EqualError(t, err, test.err.Error())
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, test.expected, result)
	}
}
