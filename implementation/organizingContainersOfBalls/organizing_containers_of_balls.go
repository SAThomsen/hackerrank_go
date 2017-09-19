package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	matrixes, err := readInput(io)
	if err != nil {
		panic(err)
	}

	for _, m := range matrixes {
		possible := m.canBecomePure()
		if possible {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}

type matrix [][]int

func (m matrix) canBecomePure() bool {
	conSum := make([]int, len(m))
	typeSum := make([]int, len(m))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			conSum[i] += m[j][i]
			typeSum[j] += m[j][i]
		}
	}

	sort.Ints(conSum)
	sort.Ints(typeSum)

	for k := 0; k < len(conSum); k++ {
		if conSum[k] != typeSum[k] {
			return false
		}
	}

	return true
}

func newMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	return m
}

func readInput(r io.Reader) ([]matrix, error) {
	var q int
	_, err := fmt.Fscan(r, &q)
	if err != nil {
		return nil, err
	}

	ms := make([]matrix, q)
	var n int
	for i := 0; i < q; i++ {
		_, err := fmt.Fscan(r, &n)
		if err != nil {
			return nil, err
		}
		m := newMatrix(n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				_, err := fmt.Fscan(r, &m[j][k])
				if err != nil {
					return nil, err
				}
			}
		}
		ms[i] = m
	}
	return ms, nil
}
