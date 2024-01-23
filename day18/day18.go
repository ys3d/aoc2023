package day18

import (
	"daniel/aoc2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// N1 computes the results for Ex1 on the given input-file
func N1(in []string) int {
	edges := parse(in)
	return area(edges)
}

// N2 computes the results for Ex2 on the given input-file
func N2(in []string) int {
	points := parsePoints(in)
	fmt.Println(points)
	return areaPoint(points)
}

type edge struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type point struct {
	x int
	y int
}

func toPoints(edges []edge) (points []point) {
	for _, e := range edges {
		points = append(points, point{
			x: e.x2,
			y: e.y2,
		})
	}
	return
}

func areaPoint(points []point) (result int) {
	for i := 1; i < len(points); i++ {
		a1 := points[i-1].x * points[i].y
		a2 := points[i-1].y * points[i].x
		result = result + a1 - a2
	}
	if result < 0 {
		result *= -1
	}
	result /= 2
	return
}

func area(edges []edge) (result int) {
	maxX := util.Reduce(util.Map(edges, func(e edge) int {
		return max(e.x1, e.x2)
	}), math.MinInt, func(i int, j int) int {
		return max(i, j)
	})
	maxY := util.Reduce(util.Map(edges, func(e edge) int {
		return max(e.y1, e.y2)
	}), math.MinInt, func(i int, j int) int {
		return max(i, j)
	})

	edges = util.Map(edges, func(e edge) edge {
		e.x1++
		e.x2++
		e.y1++
		e.y2++
		return e
	})

	field := make([][]rune, maxX+3)
	for i := range field {
		field[i] = make([]rune, maxY+3)
		for j := range field[i] {
			field[i][j] = 'O'
		}
	}
	for _, p := range edges {
		xSmall := min(p.x1, p.x2)
		xBig := max(p.x1, p.x2)
		ySmall := min(p.y1, p.y2)
		yBig := max(p.y1, p.y2)
		for x := xSmall; x <= xBig; x++ {
			for y := ySmall; y <= yBig; y++ {
				field[x][y] = 'A'
			}
		}
	}
	field[0][0] = 'F'
	found := true
	for found {
		found = false
		for x, l := range field {
			for y, c := range l {
				if c == 'O' {
					if x != 0 && field[x-1][y] == 'F' {
						field[x][y] = 'F'
						found = true
					} else if x != len(field)-1 && field[x+1][y] == 'F' {
						field[x][y] = 'F'
						found = true
					} else if y != 0 && field[x][y-1] == 'F' {
						field[x][y] = 'F'
						found = true
					} else if y != len(field[x])-1 && field[x][y+1] == 'F' {
						field[x][y] = 'F'
						found = true
					}
				}
			}
		}
	}
	fmt.Print("a\n")
	numberField := util.Map(field, func(l []rune) []int {
		return util.Map(l, func(c rune) int {
			if c != 'F' {
				return 1
			}
			return 0
		})
	})

	return util.Sum(util.Map(numberField, func(l []int) int {
		return util.Sum(l)
	}))
}

func parsePoints(lines []string) (points []point) {
	prev := point{x: 0, y: 0}
	prevDir := util.None
	points = append(points, prev)
	inner := false
	isOuterCornerChange := func(prevDir int, curDir int) (bool, func(point) point) {
		if (prevDir == util.Up && curDir == util.Right) ||
			(prevDir == util.Down && curDir == util.Left) ||
			(prevDir == util.Left && curDir == util.Up) ||
			(prevDir == util.Right && curDir == util.Down) {
			switch prevDir {
			case util.Up:
				points[len(points)-1].x--
				return false, func(p point) point {
					p.x--
					return p
				}
			case util.Down:
				points[len(points)-1].x++
				return false, func(p point) point {
					p.x++
					return p
				}
			case util.Left:
				points[len(points)-1].y--
				return false, func(p point) point {
					p.y--
					return p
				}
			case util.Right:
				points[len(points)-1].y++
				return false, func(p point) point {
					p.y++
					return p
				}
			}
		}
		return prevDir != util.None, func(p point) point {
			return p
		}
	}
	for _, l := range lines {
		newP := prev
		split := strings.Split(l, " ")
		n, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "U":
			newP.x -= n
			innerNew, f := isOuterCornerChange(prevDir, util.Up)
			newP = f(newP)
			if inner && innerNew {
				newP.x++
			}
			inner = innerNew
			prevDir = util.Up
		case "D":
			newP.x += n
			innerNew, f := isOuterCornerChange(prevDir, util.Down)
			newP = f(newP)
			if inner && innerNew {
				newP.x--
			}
			inner = innerNew
			prevDir = util.Down
		case "L":
			newP.y -= n
			innerNew, f := isOuterCornerChange(prevDir, util.Left)
			newP = f(newP)
			if inner && innerNew {
				newP.y++
			}
			inner = innerNew
			prevDir = util.Left
		case "R":
			newP.y += n
			innerNew, f := isOuterCornerChange(prevDir, util.Right)
			newP = f(newP)
			if inner && innerNew {
				newP.y--
			}
			inner = innerNew
			prevDir = util.Right
		}
		prev = newP
		points = append(points, newP)
	}
	return
}

func parse(lines []string) (edges []edge) {
	startE := edge{x1: 0, y1: 0, x2: 0, y2: 0}
	prev := startE
	edges = append(edges, startE)
	for _, l := range lines {
		newE := prev
		newE.x1 = newE.x2
		newE.y1 = newE.y2
		split := strings.Split(l, " ")
		n, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "U":
			newE.x2 -= n
		case "D":
			newE.x2 += n
		case "L":
			newE.y2 -= n
		case "R":
			newE.y2 += n
		}
		prev = newE
		edges = append(edges, newE)
	}
	edges = makePositive(edges)
	return
}

func parseExchange(lines []string) (edges []edge) {
	startE := edge{x1: 0, y1: 0, x2: 0, y2: 0}
	prev := startE
	edges = append(edges, startE)
	for _, l := range lines {
		newE := prev
		newE.x1 = newE.x2
		newE.y1 = newE.y2
		c := strings.Split(l, " ")[2]
		c = c[2 : len(c)-1]
		n, _ := strconv.ParseInt(c[:len(c)-1], 16, 0)
		switch c[len(c)-1 : len(c)] {
		case "3":
			newE.x2 -= int(n)
		case "1":
			newE.x2 += int(n)
		case "2":
			newE.y2 -= int(n)
		case "0":
			newE.y2 += int(n)
		}
		prev = newE
		edges = append(edges, newE)
	}

	return
}

func makePositive(edges []edge) []edge {
	minX := util.Reduce(util.Map(edges, func(e edge) int {
		return min(e.x1, e.x2)
	}), math.MaxInt, func(i int, j int) int {
		return min(i, j)
	})
	minY := util.Reduce(util.Map(edges, func(e edge) int {
		return min(e.y1, e.y2)
	}), math.MaxInt, func(i int, j int) int {
		return min(i, j)
	})

	if minX < 0 {
		minX *= -1
		for i := range edges {
			edges[i].x1 += minX
			edges[i].x2 += minX
		}
	}
	if minY < 0 {
		minY *= -1
		for i := range edges {
			edges[i].y1 += minY
			edges[i].y2 += minY
		}
	}
	return edges
}
