package day01

import (
	"daniel/aoc2023/util"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func N1(in []string) int {
	n := getNumbers(in)
	return util.Sum(n)

}

func N2(in []string) int {
	var newIn []string
	for _, line := range in {
		newIn = append(newIn, parseLine(line))
	}
	n := getNumbers(newIn)
	return util.Sum(n)
}

func getNumbers(in []string) (out []int) {
	for _, line := range in {
		first := -1
		last := -1
		for _, letter := range []rune(line) {
			if unicode.IsDigit(letter) {
				if first == -1 {
					first = int(letter) - 48
					last = int(letter) - 48
				} else {
					last = int(letter) - 48
				}
			}
		}
		out = append(out, 10*first+last)
	}
	return
}

func parseLine(line string) string {
	lline := line
	returnline := ""
	for lline != "" {
		position := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := range position {
			position[i] = strings.Index(lline, util.ToText(i+1))
		}
		index := slices.Index(position, 0)
		if index != -1 {
			returnline += strconv.Itoa(index + 1)
		} else if unicode.IsDigit([]rune(lline)[0]) {
			returnline = returnline + lline[:1]
		}
		lline = lline[1:]
	}
	return returnline
}
