package main

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		in  int64
		out big.Int
	}{
		{1, *big.NewInt(1)},
		{2, *big.NewInt(2)},
		{3, *big.NewInt(6)},
		{4, *big.NewInt(24)},
	}

	for _, tc := range testCases {
		if tc.out.Cmp(factorial(tc.in)) != 0 {
			t.Error("The factorial failed")
		}
	}

}
