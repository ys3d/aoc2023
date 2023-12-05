package day4

import (
	"daniel/aoc2023/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func N1(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var games []game

	for _, l := range in {
		g := game{}
		g.parseFromLine(l)
		games = append(games, g)
	}

	return strconv.Itoa(util.Sum(util.Map(games, func(g game) int { return g.worth() })))
}

func N2(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var games []game

	for _, l := range in {
		g := game{}
		g.parseFromLine(l)
		games = append(games, g)
	}
	i := len(games) - 1
	for i >= 0 {
		n := games[i].matchingNumbers()
		for n > 0 {
			if i+n < len(games) {
				games[i].cards += games[i+n].cards
				n -= 1
			}
		}
		i--
	}

	return strconv.Itoa(util.Sum(util.Map(games, func(g game) int { return g.cards })))
}

type game struct {
	index          int
	winningNumbers []int
	numbers        []int
	cards          int
}

func (g *game) parseFromLine(line string) {
	g.index, _ = strconv.Atoi(util.Last(strings.Split(strings.Split(line, ": ")[0], " ")))
	winningNumbersS := strings.TrimSpace(strings.Split(strings.Split(line, " | ")[0], ": ")[1])
	for _, n := range strings.Split(winningNumbersS, " ") {
		if n == "" {
			continue
		}
		number, _ := strconv.Atoi(n)
		g.winningNumbers = append(g.winningNumbers, number)
	}

	numbersS := strings.TrimSpace(strings.Split(line, " | ")[1])
	for _, n := range strings.Split(numbersS, " ") {
		if n == "" {
			continue
		}
		number, _ := strconv.Atoi(strings.TrimSpace(n))
		g.numbers = append(g.numbers, number)
	}
	g.cards = 1
}

func (g *game) matchingNumbers() (m int) {
	for _, n := range g.numbers {
		if slices.Contains(g.winningNumbers, n) {
			m += 1
		}
	}
	return
}

func (g *game) worth() (w int) {
	m := g.matchingNumbers()
	for m > 0 {
		if w == 0 {
			w = 1
		} else {
			w *= 2
		}
		m -= 1
	}
	return
}
