package main

import (
	"bufio"
	"daniel/aoc2023/day1"
	"daniel/aoc2023/day2"
	"daniel/aoc2023/day3"
	"daniel/aoc2023/day4"
	"daniel/aoc2023/day5"
	"daniel/aoc2023/day6"
	"daniel/aoc2023/day7"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strconv"
	"strings"
	"time"
)

const Day = 0

func main() {
	localDay := Day
	if localDay == -1 {
		fmt.Println("Please select the day:")
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			return
		}

		// remove the delimiter from the string
		input = strings.TrimSuffix(input, "\n")
		localDay, err = strconv.Atoi(input)
		if err != nil {
			localDay = -1
		}
	}
	switch localDay {
	case 0:
		runDay1()
		runDay2()
		runDay3()
		runDay4()
		runDay5()
		runDay6()
		runDay7()
	case 1:
		runDay1()
	case 2:
		runDay2()
	case 3:
		runDay3()
	case 4:
		runDay4()
	case 5:
		runDay5()
	case 6:
		runDay6()
	case 7:
		runDay7()
	default:
		fmt.Println("Day does not exist")
	}
	printRuns()
}

func runDay1() {
	runAndPrint(day1.N1, []string{"day1/test1.txt", "day1/input.txt"}, "1", "1")
	runAndPrint(day1.N2, []string{"day1/test2.txt", "day1/input.txt"}, "1", "2")
}

func runDay2() {
	runAndPrint(day2.N1, []string{"day2/test1.txt", "day2/input.txt"}, "2", "1")
	runAndPrint(day2.N2, []string{"day2/test2.txt", "day2/input.txt"}, "2", "2")
}

func runDay3() {
	runAndPrint(day3.N1, []string{"day3/test1.txt", "day3/input.txt"}, "3", "1")
	runAndPrint(day3.N2, []string{"day3/test2.txt", "day3/input.txt"}, "3", "2")
}

func runDay4() {
	runAndPrint(day4.N1, []string{"day4/test1.txt", "day4/input.txt"}, "4", "1")
	runAndPrint(day4.N2, []string{"day4/test2.txt", "day4/input.txt"}, "4", "2")
}

func runDay5() {
	runAndPrint(day5.N1, []string{"day5/test.txt", "day5/input.txt"}, "5", "1")
	runAndPrint(day5.N2, []string{"day5/test.txt", "day5/input.txt"}, "5", "2")
}

func runDay6() {
	runAndPrint(day6.N1, []string{"day6/test.txt", "day6/input.txt"}, "6", "1")
	runAndPrint(day6.N2, []string{"day6/test.txt", "day6/input.txt"}, "6", "2")
}

func runDay7() {
	runAndPrint(day7.N1, []string{"day7/test.txt", "day7/input.txt"}, "7", "1")
	runAndPrint(day7.N2, []string{"day7/test.txt", "day7/input.txt"}, "7", "2")
}

type run struct {
	day      string
	ex       string
	input    string
	result   string
	duration time.Duration
}

var runs []run

func runAndPrint(f func(string) string, params []string, day string, ex string) {
	for _, p := range params {
		start := time.Now()
		out := f(p)
		elapsed := time.Since(start)
		runs = append(runs, run{
			day,
			ex,
			p,
			out,
			elapsed,
		})
	}
}

func printRuns() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Advent of Code")
	t.AppendHeader(table.Row{"Day", "Part", "Input", "Result", "Duration"})
	for _, r := range runs {
		t.AppendRows([]table.Row{
			{r.day, r.ex, r.input, r.result, r.duration},
		})
	}
	t.SetStyle(table.StyleColoredDark)
	t.Style().Title.Align = text.AlignCenter
	t.Render()
}
