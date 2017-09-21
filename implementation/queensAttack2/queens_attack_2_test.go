package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func TestCalculateMoves(t *testing.T) {
	testCases := []struct {
		n         int
		k         int
		queenPos  point
		obstacles []point
		out       int
	}{
		{4, 0, point{4, 4}, []point{}, 9},
		{4, 0, point{4, 4}, []point{{3, 3}}, 6},
		{5, 3, point{4, 3}, []point{{5, 5}, {4, 2}, {2, 3}}, 10},
		{100000, 0, point{4187, 5068}, []point{}, 308369},
	}

	for _, tc := range testCases {
		g := newGame(tc.n, tc.queenPos)
		sum := g.calculateQueenMoves(&tc.obstacles)
		if sum != tc.out {
			t.Errorf("Failed. Expected: %d, Got: %d", tc.out, sum)
		}
	}
}

func TestDistance(t *testing.T) {
	testCases := []struct {
		a   point
		b   point
		out int
	}{
		{point{4, 4}, point{1, 1}, 3},
	}

	for _, tc := range testCases {
		d := tc.a.distance(&tc.b)
		if d == tc.out {
			t.Errorf("Distance was calculated to be: %d expected %d", d, tc.out)
		}
	}
}

func TestMain(t *testing.T) {
	testCases := []struct {
		inFile  string
		outFile string
	}{
		{"test_input_6.txt", "test_output_6.txt"},
		{"test_input_18.txt", "test_output_18.txt"},
		{"test_input_19.txt", "test_output_19.txt"},
	}

	for _, tc := range testCases {
		subproc := exec.Command("go run queens_attack_2")

		inputBuf, _ := ioutil.ReadFile(tc.inFile)
		input := string(inputBuf)
		subproc.Stdin = strings.NewReader(input)
		output, _ := subproc.Output()

		outputBuf, _ := ioutil.ReadFile(tc.outFile)
		expectedOutput := string(outputBuf)
		if string(output) == expectedOutput {
			t.Errorf("Wanted: %v, Got: %v", string(output), expectedOutput)
		}
		subproc.Wait()
	}
}
