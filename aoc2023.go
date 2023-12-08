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
	"daniel/aoc2023/day8"
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
		runDay8()
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
	case 8:
		runDay8()
	default:
		fmt.Println("Day does not exist")
	}
	printRuns()
}

func runDay1() {
	f := getFilesIn("day1/")
	runAndPrint(day1.N1, f, "1", "1")
	runAndPrint(day1.N2, f, "1", "2")
}

func runDay2() {
	f := getFilesIn("day2/")
	runAndPrint(day2.N1, f, "2", "1")
	runAndPrint(day2.N2, f, "2", "2")
}

func runDay3() {
	f := getFilesIn("day3/")
	runAndPrint(day3.N1, f, "3", "1")
	runAndPrint(day3.N2, f, "3", "2")
}

func runDay4() {
	f := getFilesIn("day4/")
	runAndPrint(day4.N1, f, "4", "1")
	runAndPrint(day4.N2, f, "4", "2")
}

func runDay5() {
	f := getFilesIn("day5/")
	runAndPrint(day5.N1, f, "5", "1")
	runAndPrint(day5.N2, f, "5", "2")
}

func runDay6() {
	f := getFilesIn("day6/")
	runAndPrint(day6.N1, f, "6", "1")
	runAndPrint(day6.N2, f, "6", "2")
}

func runDay7() {
	f := getFilesIn("day7/")
	runAndPrint(day7.N1, f, "7", "1")
	runAndPrint(day7.N2, f, "7", "2")
}

func runDay8() {
	f1 := getFilesIn("day8/inputPart1/")
	f2 := getFilesIn("day8/inputPart2/")
	runAndPrint(day8.N1, f1, "8", "1")
	runAndPrint(day8.N2, f2, "8", "2")
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

func getFilesIn(s string) (out []string) {
	f, _ := os.Open(s)
	files, _ := f.Readdir(0)
	for _, f := range files {
		if strings.Contains(f.Name(), ".txt") {
			out = append(out, s+f.Name())
		}
	}
	return
}
