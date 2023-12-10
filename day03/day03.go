package day03

import (
	"daniel/aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func N1(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	emptyLine := strings.Repeat(".", len(in[0]))
	in = append([]string{emptyLine}, in...)
	in = append(in, emptyLine)

	var numbers []int
	for i := range in {
		if i > 0 && i < (len(in)-1) {
			numbers = append(numbers, findNumber(in[i-1:i+2])...)
		}
	}
	//fmt.Println(numbers)

	return strconv.Itoa(util.Sum(numbers))

}

func N2(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	emptyLine := strings.Repeat(".", len(in[0]))
	in = append([]string{emptyLine}, in...)
	in = append(in, emptyLine)

	for i := range in {
		in[i] = "." + in[i] + "."
	}

	var ratios []int

	for i := range in {
		index := strings.Index(in[i], "*")
		for index >= 0 {
			numbers := getNumbersAround(in, i, index)
			if len(numbers) == 2 {
				ratios = append(ratios, numbers[0]*numbers[1])
			}
			in[i] = strings.Replace(in[i], "*", ".", 1)
			index = strings.Index(in[i], "*")
		}
	}

	return strconv.Itoa(util.Sum(ratios))
}

func findNumber(lines []string) []int {
	var numbers []int
	var number int
	var found bool
	for i := range lines[1] {
		digit := []rune(lines[1][i : i+1])[0]
		isDigit := unicode.IsDigit(digit)
		if isDigit {
			number = number*10 + int(digit-'0')
		}
		enable0, _ := regexp.MatchString("([^0-9.])", lines[0][i:i+1])
		enable1, _ := regexp.MatchString("([^0-9.])", lines[1][i:i+1])
		enable2, _ := regexp.MatchString("([^0-9.])", lines[2][i:i+1])
		enable := enable0 || enable1 || enable2
		if enable {
			found = true
		}
		if !isDigit {
			if number != 0 && found {
				numbers = append(numbers, number)
			}
			number = 0
		}

		if isDigit && found && (i == len(lines[1])-1) {
			numbers = append(numbers, number)
		}

		if !isDigit && found && !enable {
			found = false
		}
	}
	return numbers
}

func getNumbersAround(lines []string, starX int, starY int) (numbers []int) {
	t := lines[starX-1]
	m := lines[starX]
	b := lines[starX+1]

	numbers = append(numbers, numbersAround(t, starY)...)
	numbers = append(numbers, numbersAround(m, starY)...)
	numbers = append(numbers, numbersAround(b, starY)...)
	return
}

func numbersAround(s string, i int) (numbers []int) {
	if unicode.IsDigit([]rune(s[i : i+1])[0]) {
		// one number case
		start := i - 1
		for unicode.IsDigit([]rune(s[start : start+1])[0]) {
			start -= 1
		}
		start += 1
		part := s[start:]
		number := 0
		for unicode.IsDigit([]rune(part[:1])[0]) {
			number = number*10 + int([]rune(part[:1])[0]-'0')
			part = part[1:]
		}
		numbers = append(numbers, number)
		return
	}

	// two number case
	if unicode.IsDigit([]rune(s[i+1 : i+2])[0]) {
		part := s[i+1:]
		number := 0
		for unicode.IsDigit([]rune(part[:1])[0]) {
			number = number*10 + int([]rune(part[:1])[0]-'0')
			part = part[1:]
		}
		numbers = append(numbers, number)
	}
	if unicode.IsDigit([]rune(s[i-1 : i])[0]) {
		part := s[:i]
		factor := 1
		number := 0
		for unicode.IsDigit([]rune(part[len(part)-1:])[0]) {
			number += factor * int([]rune(part[len(part)-1:])[0]-'0')
			factor *= 10
			part = part[:len(part)-1]
		}
		numbers = append(numbers, number)
	}
	return
}
