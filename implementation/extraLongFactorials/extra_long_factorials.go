package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	var n int64
	fmt.Fscan(io, &n)

	fac := factorial(n)

	fmt.Println(fac)
}

func factorial(n int64) *big.Int {
	nB := big.NewInt(n)
	f := big.NewInt(1)
	oneB := big.NewInt(1)
	for 0 <= nB.Cmp(oneB) {
		f.Mul(f, nB)
		nB.Sub(nB, oneB)
	}

	return f
}
