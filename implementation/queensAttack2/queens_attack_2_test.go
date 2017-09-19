package main

import (
	"testing"
)

func TestCalculateMoves(t *testing.T) {
	testCases := []struct {
		n         int
		k         int
		queenX    int
		queenY    int
		obstacles [][]int
		out       int
	}{
		{4, 0, 4, 4, [][]int{}, 9},
		{4, 0, 4, 4, [][]int{{3, 3}}, 6},
		{5, 3, 4, 3, [][]int{{5, 5}, {4, 2}, {2, 3}}, 10},
		{100000, 0, 4187, 5068, [][]int{}, 308369},
	}

	for _, tc := range testCases {
		g := game{length: tc.n}
		g.initializeBoard(tc.obstacles)

		r := g.calculateMoves(tc.queenX, tc.queenY)
		if r != tc.out {
			t.Errorf("Failed. Expected: %d, Got: %d", tc.out, r)
		}
	}
}
