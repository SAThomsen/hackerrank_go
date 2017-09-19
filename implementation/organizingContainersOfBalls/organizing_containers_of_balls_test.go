package main

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	testCases := []struct {
		in  matrix
		out bool
	}{
		{[][]int{
			{0, 1},
			{1, 0},
		}, true},
		{[][]int{
			{1, 2},
			{2, 1},
		}, true},
		{[][]int{
			{0, 4, 0},
			{2, 0, 1},
			{1, 0, 2},
		}, true},
		{[][]int{
			{1, 2, 3, 4},
			{2, 1, 4, 3},
			{3, 4, 2, 1},
			{4, 3, 1, 2},
		}, true},
		{[][]int{
			{2, 1},
			{0, 0},
		}, false},
		{[][]int{
			{1, 2, 1, 3},
			{2, 1, 3, 1},
			{1, 3, 2, 1},
			{3, 2, 1, 1},
		}, false},
	}

	for i, t := range testCases {
		if t.in.canBecomePure() != t.out {
			fmt.Println("Failed: ", i)
		} else {
			fmt.Println("Passed: ", i)
		}
	}
}
