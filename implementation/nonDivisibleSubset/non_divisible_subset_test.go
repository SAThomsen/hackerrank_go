package main

import "testing"
import "fmt"

func TestNonDivisibleSubset(t *testing.T) {
	testCases := []struct {
		in  []int
		k   int
		out int
	}{
		{[]int{1, 7, 2, 4}, 3, 3},
	}

	for _, tc := range testCases {
		tr := nonDivisibleSubset(tc.in, tc.k)
		fmt.Println(tr)
		if tr != tc.out {
			t.Error("Failed!")
		}
	}
}
