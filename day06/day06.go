package day06

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
	races := parseInput(in)

	waysToWin := util.Map(races, func(r race) int { return r.waysToWin() })
	value := 1
	for _, v := range waysToWin {
		if v != 0 {
			value *= v
		}
	}

	return strconv.Itoa(value)
}

// N2 computes the results for Ex2 on the given input-file
func N2(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range in {
		in[i] = strings.Replace(in[i], " ", "", -1)
		in[i] = strings.Replace(in[i], ":", " ", -1)
	}
	races := parseInput(in)

	waysToWin := util.Map(races, func(r race) int { return r.waysToWin() })
	value := 1
	for _, v := range waysToWin {
		if v != 0 {
			value *= v
		}
	}

	return strconv.Itoa(value)
}

type race struct {
	time     int
	distance int
}

func (r *race) waysToWin() int {
	// Using the quadratic formula to solve equation
	// 0 = -x² + time * x - distance

	a := float64(-1)
	b := float64(r.time)
	c := -1.0 * float64(r.distance)
	x1 := (-1.0*b + math.Sqrt(math.Pow(b, 2)-4.0*a*c)) / (a * 2.0)
	x2 := (-1.0*b - math.Sqrt(math.Pow(b, 2)-4.0*a*c)) / (a * 2.0)

	if math.Mod(x1, 1.0) == 0 {
		x1 += 0.5
	}

	if math.Mod(x2, 1.0) == 0 {
		x2 -= 0.5
	}

	x1 = math.Ceil(x1)
	x2 = math.Floor(x2)

	return int(x2-x1) + 1
}

func parseInput(lines []string) (r []race) {
	// Parse Time
	split := util.Filter(strings.Split(lines[0], " "), func(s string) bool { return s != "" })
	for _, n := range split[1:] {
		t, _ := strconv.Atoi(n)
		r = append(r, race{time: t})
	}
	split = util.Filter(strings.Split(lines[1], " "), func(s string) bool { return s != "" })
	for i, n := range split[1:] {
		t, _ := strconv.Atoi(n)
		r[i].distance = t
	}
	return
}
