package day10

import (
	"daniel/aoc2023/util"
	"github.com/dominikbraun/graph"
	"strings"
)

// N1 computes the results for Ex1 on the given input-file
func N1(in []string) int {
	in = append([]string{strings.Repeat(".", len(in[0]))}, in...)
	in = append(in, strings.Repeat(".", len(in[0])))
	for i := 0; i < len(in); i++ {
		in[i] = "." + in[i] + "."
	}
	g := buildGraphOfPipes(in)
	x, y := findStart(in)
	d := 0
	_ = graph.DFS(g, vertexIndex(x, y, len(in[0])), func(value int) bool {
		d++
		return d != 1 && value == vertexIndex(x, y, len(in[0]))
	})
	return d / 2
}

// N2 computes the results for Ex2 on the given input-file
func N2(in []string) int {
	in = append([]string{strings.Repeat(".", len(in[0]))}, in...)
	in = append(in, strings.Repeat(".", len(in[0])))
	for i := 0; i < len(in); i++ {
		in[i] = "." + in[i] + "."
	}
	in = cleanup(in)
	var extendedIn []string
	for _, l := range in {
		extendedIn = append(extendedIn, toMultiLine(l)...)
	}
	var field [][]int
	for x, l := range extendedIn {
		field = append(field, make([]int, len(l)))
		for y, c := range l {
			if c != '.' {
				field[x][y] = 1
			}
		}
	}
	g := buildGraphOfDots(extendedIn)
	_ = graph.DFS(g, vertexIndex(0, 0, len(extendedIn[0])), func(value int) bool {
		x, y := vertexCoordinates(value, len(extendedIn[0]))
		field[x][y] = 1
		return false
	})
	count := 0
	for x, l := range field {
		for y, c := range l {
			if c == 0 && x%3 == 1 && y%3 == 1 {
				count++
			}
		}
	}
	return count
}

func cleanup(lines []string) (out []string) {
	for _, l := range lines {
		out = append(out, strings.Repeat(".", len(l)))
	}
	g := buildGraphOfPipes(lines)
	xStart, yStart := findStart(lines)
	_ = graph.DFS(g, vertexIndex(xStart, yStart, len(lines[0])), func(value int) bool {
		x, y := vertexCoordinates(value, len(lines[0]))
		out[x] = util.ReplaceAtIndex(out[x], rune(lines[x][y]), y)
		return false
	})
	return
}

func buildGraphOfPipes(lines []string) graph.Graph[int, int] {
	maxY := len(lines[0])
	g := graph.New(graph.IntHash)

	for i, l := range lines {
		for j, c := range l {
			index := vertexIndex(i, j, maxY)
			switch c {
			case '|', 'L':
				_ = g.AddVertex(index)
				switch lines[i-1][j] {
				case '|', '7', 'S', 'F':
					_ = g.AddEdge(index, vertexIndex(i-1, j, maxY))
				}
			case '-', '7':
				_ = g.AddVertex(index)
				switch lines[i][j-1] {
				case '-', 'S', 'F', 'L':
					_ = g.AddEdge(index, vertexIndex(i, j-1, maxY))
				}

			case 'J', 'S':
				_ = g.AddVertex(index)
				switch lines[i][j-1] {
				case '-', 'S', 'F', 'L':
					_ = g.AddEdge(index, vertexIndex(i, j-1, maxY))
				}
				switch lines[i-1][j] {
				case '|', '7', 'S', 'F':
					_ = g.AddEdge(index, vertexIndex(i-1, j, maxY))
				}
			case 'F':
				_ = g.AddVertex(index)
			}
		}
	}
	return g
}

func buildGraphOfDots(lines []string) graph.Graph[int, int] {
	maxY := len(lines[0])
	g := graph.New(graph.IntHash)

	for i, l := range lines {
		for j, c := range l {
			index := vertexIndex(i, j, maxY)
			if c == '.' {
				_ = g.AddVertex(index)
				if j-1 >= 0 && lines[i][j-1] == '.' {
					_ = g.AddEdge(index, vertexIndex(i, j-1, maxY))
				}
				if i-1 >= 0 {
					if j-1 > 0 && lines[i-1][j-1] == '.' {
						_ = g.AddEdge(index, vertexIndex(i-1, j-1, maxY))
					}
					if lines[i-1][j] == '.' {
						_ = g.AddEdge(index, vertexIndex(i-1, j, maxY))
					}
					if j+1 < len(l) && lines[i-1][j+1] == '.' {
						_ = g.AddEdge(index, vertexIndex(i-1, j+1, maxY))
					}
				}
			}
		}
	}
	return g
}

func toMultiLine(l string) (out []string) {
	out = make([]string, 3)
	for _, c := range l {
		switch c {
		case '.':
			out[0] += "..."
			out[1] += "..."
			out[2] += "..."
		case '|':
			out[0] += ".x."
			out[1] += ".x."
			out[2] += ".x."
		case '-':
			out[0] += "..."
			out[1] += "xxx"
			out[2] += "..."
		case 'L':
			out[0] += ".X."
			out[1] += ".xx"
			out[2] += "..."
		case 'J':
			out[0] += ".x."
			out[1] += "xx."
			out[2] += "..."
		case '7':
			out[0] += "..."
			out[1] += "xx."
			out[2] += ".x."
		case 'F':
			out[0] += "..."
			out[1] += ".xx"
			out[2] += ".x."
		case 'S':
			out[0] += ".x."
			out[1] += "xxx"
			out[2] += ".x."
		}
	}
	return
}

func vertexIndex(x int, y int, maxY int) int {
	return x*maxY + y
}

func vertexCoordinates(index int, maxY int) (int, int) {
	return index / maxY, index % maxY
}

func findStart(lines []string) (int, int) {
	for i, l := range lines {
		j := strings.Index(l, "S")
		if j != -1 {
			return i, j
		}
	}
	return -1, -1
}
