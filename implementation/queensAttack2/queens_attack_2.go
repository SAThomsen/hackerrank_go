package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	n, _, queenPos, obstacles := readInput(io)
	g := newGame(n, queenPos)
	sum := g.calculateQueenMoves(&obstacles)
	fmt.Println(sum)
}

func readInput(io io.Reader) (int, int, point, []point) {
	// game size
	var n int
	fmt.Fscan(io, &n)
	// number of obstacles
	var k int
	fmt.Fscan(io, &k)

	var qx int
	var qy int
	fmt.Fscan(io, &qx)
	fmt.Fscan(io, &qy)
	queenPos := point{x: qx, y: qy}

	obstacles := make([]point, k)
	var ox, oy int
	for i := 0; i < k; i++ {
		fmt.Fscan(io, &ox)
		fmt.Fscan(io, &oy)
		obstacles[i] = point{x: ox, y: oy}
	}

	return n, k, queenPos, obstacles
}

type point struct {
	x int
	y int
}

func (p1 *point) distance(p2 *point) int {
	diffX := float64(p2.x - p1.x)
	diffY := float64(p2.y - p1.y)
	distance := math.Sqrt(math.Pow(diffX, 2.0) + math.Pow(diffY, 2.0))

	return int(distance)
}

type game struct {
	queenPos        point
	dimension       int
	distanceToQueen [8]int
}

func newGame(n int, q point) game {
	g := game{dimension: n, queenPos: q}
	g.initDistanceToQueen()
	return g
}

func (g *game) onBoard(p *point) bool {
	d := g.dimension
	return !(p.x < 1 || p.y < 1 || d < p.x || d < p.y)
}

func (g *game) initDistanceToQueen() {
	d, q := g.dimension, g.queenPos

	g.distanceToQueen[0] = intMin(q.y, q.x) - 1
	g.distanceToQueen[1] = q.x - 1
	g.distanceToQueen[2] = intMin(d-q.y, intAbs(1-q.x))
	g.distanceToQueen[3] = d - q.y
	g.distanceToQueen[4] = d - intMax(q.y, q.x)
	g.distanceToQueen[5] = d - q.x
	g.distanceToQueen[6] = intMin(d-q.x, q.y-1)
	g.distanceToQueen[7] = q.y - 1
}

func (g *game) getCurrentDistanceInDirection(p *point) *int {
	q := g.queenPos
	x, y := p.x-q.x, p.y-q.y

	if x < 0 && y < 0 && x == y {
		return &g.distanceToQueen[0]
	}

	if x < 0 && y == 0 {
		return &g.distanceToQueen[1]
	}

	if x < 0 && y > 0 && intAbs(x) == intAbs(y) {
		return &g.distanceToQueen[2]
	}

	if x == 0 && y > 0 {
		return &g.distanceToQueen[3]
	}

	if x > 0 && y > 0 && x == y {
		return &g.distanceToQueen[4]
	}

	if x > 0 && y == 0 {
		return &g.distanceToQueen[5]
	}

	if x > 0 && y < 0 && intAbs(x) == intAbs(y) {
		return &g.distanceToQueen[6]
	}

	if x == 0 && y < 0 {
		return &g.distanceToQueen[7]
	}

	return nil
}

func (g *game) calculateQueenMoves(obstacles *[]point) int {
	for _, obs := range *obstacles {
		if g.onBoard(&obs) {
			curDistance := g.getCurrentDistanceInDirection(&obs)
			if curDistance == nil {
				continue
			}

			obsDistance := g.queenPos.distance(&obs) - 1
			if *curDistance > obsDistance {
				*curDistance = obsDistance
			}
		}
	}

	sum := 0
	for i := 0; i < 8; i++ {
		sum += g.distanceToQueen[i]
	}

	return sum
}

func (g *game) printDistanceToQueen() {
	fmt.Printf("%d\t%d\t%d\n", g.distanceToQueen[2], g.distanceToQueen[3], g.distanceToQueen[4])
	fmt.Printf("%d\tQ\t%d\n", g.distanceToQueen[1], g.distanceToQueen[5])
	fmt.Printf("%d\t%d\t%d\n", g.distanceToQueen[0], g.distanceToQueen[7], g.distanceToQueen[6])
}

func intAbs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func intMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
