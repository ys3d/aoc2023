package main

import (
	"bufio"
	"daniel/aoc2023/day1"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Day = 0

func main() {
	fmt.Println("AOC 2023")
	localDay := Day
	if localDay == 0 {
		fmt.Println("Please select the day:")
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")
		localDay, err = strconv.Atoi(input)
		if err != nil {
			localDay = -1
		}
	}
	switch localDay {
	case 1:
		fmt.Println("######### Day 1 Ex 1 ##########")
		out := day1.N1("day1/test1.txt")
		fmt.Println("--> Solution for Test: ", out)
		out = day1.N1("day1/input.txt")
		fmt.Println("--> Solution for Input: ", out)

		fmt.Println("######### Day 1 Ex 2 ##########")
		out = day1.N2("day1/test2.txt")
		fmt.Println("--> Solution for Test: ", out)
		out = day1.N2("day1/input.txt")
		fmt.Println("--> Solution for Input: ", out)
	default:
		fmt.Println("Day does not exist")
	}

}
