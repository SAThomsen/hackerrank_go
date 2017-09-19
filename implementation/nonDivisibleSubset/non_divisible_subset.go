package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(io, &n)

	var k int
	fmt.Fscan(io, &k)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(io, &s[i])
	}

	c := nonDivisibleSubset(s, k)

	fmt.Println(c)
}

func nonDivisibleSubset(s []int, k int) int {

	counts := make([]int, k)
	for _, i := range s {
		counts[i%k]++
	}
	fmt.Println(counts)

	count := min(counts[0], 1)

	for i := 1; i < k-i; i++ {
		count += max(counts[i], counts[k-i])
	}

	if k%2 == 0 {
		count++
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
