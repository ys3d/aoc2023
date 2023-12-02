package day2

import (
	"daniel/aoc2023/util"
	"fmt"
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
		g.fromLine(l)
		games = append(games, g)
	}
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	sum := 0
	for _, g := range games {
		if g.red <= maxRed && g.green <= maxGreen && g.blue <= maxBlue {
			sum += g.index
		}
	}
	return strconv.Itoa(sum)

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
		g.fromLine(l)
		games = append(games, g)
	}

	sum := 0
	for _, g := range games {
		sum += g.red * g.green * g.blue
	}
	return strconv.Itoa(sum)
}

type game struct {
	index int
	red   int
	green int
	blue  int
}

func (g *game) fromLine(line string) {
	g.index, _ = strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
	rounds := strings.Split(strings.Split(line, ": ")[1], "; ")
	for _, r := range rounds {
		red, _ := strconv.Atoi(util.Last(strings.Split(strings.Split(r, " red")[0], " ")))
		green, _ := strconv.Atoi(util.Last(strings.Split(strings.Split(r, " green")[0], " ")))
		blue, _ := strconv.Atoi(util.Last(strings.Split(strings.Split(r, " blue")[0], " ")))

		g.red = max(red, g.red)
		g.green = max(green, g.green)
		g.blue = max(blue, g.blue)

	}
}
