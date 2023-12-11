package day11

import (
	"daniel/aoc2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// N1 computes the results for Ex1 on the given input-file
func N1(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	g := build(in, 2)
	sum := 0
	for i, g1 := range g {
		for j := i + 1; j < len(g); j++ {
			sum += g1.dist(g[j])
		}
	}
	return strconv.Itoa(sum)
}

// N2 computes the results for Ex2 on the given input-file
func N2(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	g := build(in, 1000000)
	sum := 0
	for i, g1 := range g {
		for j := i + 1; j < len(g); j++ {
			sum += g1.dist(g[j])
		}
	}
	return strconv.Itoa(sum)
}

type galaxy struct {
	index int
	x     int
	y     int
}

func (g *galaxy) dist(g2 galaxy) int {
	return int(math.Abs(float64(g.x-g2.x))) + int(math.Abs(float64(g.y-g2.y)))
}

func build(lines []string, expansion int) (g []galaxy) {
	index := 1
	xOffset, yOffset := offsets(lines, expansion)
	for x, l := range lines {
		for y, c := range l {
			if c == '#' {
				g = append(g, galaxy{index: index, x: x + xOffset[x], y: y + yOffset[y]})
				index++
			}
		}
	}

	return
}

func offsets(in []string, expansion int) (xOffset []int, yOffset []int) {
	xOffset = make([]int, len(in))
	yOffset = make([]int, len(in[0]))
	emptyLine := strings.Repeat(".", len(in[0]))
	for i, l := range in {
		if l == emptyLine {
			for j := i; j < len(xOffset); j++ {
				xOffset[j] += expansion - 1
			}
		}
	}
	// add vertical lines
	for i := 0; i < len(in[0]); i++ {
		dots := true
		for _, l := range in {
			if l[i] != '.' {
				dots = false
			}
		}
		if dots {
			for j := i; j < len(in[0]); j++ {
				yOffset[j] += expansion - 1
			}
		}
	}
	return
}
