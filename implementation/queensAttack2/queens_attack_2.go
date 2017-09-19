package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	n, _, qx, qy, obs := readInput(io)

	g := game{length: n}
	g.initializeBoard(obs)

	count := g.calculateMoves(qx, qy)

	fmt.Println(count)
}

func readInput(io io.Reader) (int, int, int, int, [][]int) {
	// game size
	var n int
	fmt.Fscan(io, &n)
	//number of obstacles
	var k int
	fmt.Fscan(io, &k)

	var qx, qy int
	fmt.Fscan(io, &qx)
	fmt.Fscan(io, &qy)

	obstacles := make([][]int, k)
	var ox, oy int
	for i := 0; i < k; i++ {
		fmt.Fscan(io, &ox)
		fmt.Fscan(io, &oy)
		obstacles[i] = []int{ox, oy}
	}

	return n, k, qx, qy, obstacles
}

type game struct {
	board  map[int]map[int]bool
	length int
}

func (g *game) crossedEdge(x int, y int) bool {
	l := g.length
	return x < 1 || y < 1 || l < x || l < y
}

func (g *game) isObstacle(x int, y int) bool {
	my, ok := g.board[x]
	if !ok {
		return false
	}
	b, ok := my[y]
	if !ok {
		return false
	}
	return b
}

func (g *game) initializeBoard(obstacles [][]int) {
	bx := make(map[int]map[int]bool)
	for _, obs := range obstacles {
		by := make(map[int]bool)
		by[obs[1]] = true
		bx[obs[0]] = by
	}
	g.board = bx
}

func (g *game) calculateMoves(qx int, qy int) int {
	count := 0
	sq := []int{-1, 0, 1}
	for _, ix := range sq {
		for _, iy := range sq {
			moves := g.movesInDirection(ix, iy, qx, qy)
			count += moves
		}
	}
	return count
}

func (g *game) movesInDirection(ix int, iy int, qx int, qy int) int {
	if ix == 0 && iy == 0 {
		return 0
	}
	count := 0
	for true {
		qx -= ix
		qy -= iy
		if g.crossedEdge(qx, qy) {
			break
		}
		if g.isObstacle(qx, qy) {
			break
		}
		count++
	}

	return count
}

func (g *game) printObstacles() {
	for x, m := range g.board {
		for y, b := range m {
			fmt.Printf("x: %d, y: %d val: %t", x, y, b)
		}
	}
}
