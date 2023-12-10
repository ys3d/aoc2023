package main

import (
	"bufio"
	"daniel/aoc2023/day01"
	"daniel/aoc2023/day02"
	"daniel/aoc2023/day03"
	"daniel/aoc2023/day04"
	"daniel/aoc2023/day05"
	"daniel/aoc2023/day06"
	"daniel/aoc2023/day07"
	"daniel/aoc2023/day08"
	"daniel/aoc2023/day09"
	"daniel/aoc2023/day10"
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
		runDay9()
		runDay10()
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
	case 9:
		runDay9()
	case 10:
		runDay10()
	default:
		fmt.Println("Day does not exist")
	}
	printRuns()
}

func runDay1() {
	f := getFilesIn("day01/")
	runAndPrint(day01.N1, f, "1", "1")
	runAndPrint(day01.N2, f, "1", "2")
}

func runDay2() {
	f := getFilesIn("day02/")
	runAndPrint(day02.N1, f, "2", "1")
	runAndPrint(day02.N2, f, "2", "2")
}

func runDay3() {
	f := getFilesIn("day03/")
	runAndPrint(day03.N1, f, "3", "1")
	runAndPrint(day03.N2, f, "3", "2")
}

func runDay4() {
	f := getFilesIn("day04/")
	runAndPrint(day04.N1, f, "4", "1")
	runAndPrint(day04.N2, f, "4", "2")
}

func runDay5() {
	f := getFilesIn("day05/")
	runAndPrint(day05.N1, f, "5", "1")
	runAndPrint(day05.N2, f, "5", "2")
}

func runDay6() {
	f := getFilesIn("day06/")
	runAndPrint(day06.N1, f, "6", "1")
	runAndPrint(day06.N2, f, "6", "2")
}

func runDay7() {
	f := getFilesIn("day07/")
	runAndPrint(day07.N1, f, "7", "1")
	runAndPrint(day07.N2, f, "7", "2")
}

func runDay8() {
	f1 := getFilesIn("day08/inputPart1/")
	f2 := getFilesIn("day08/inputPart2/")
	runAndPrint(day08.N1, f1, "8", "1")
	runAndPrint(day08.N2, f2, "8", "2")
}

func runDay9() {
	f := getFilesIn("day09/")
	runAndPrint(day09.N1, f, "9", "1")
	runAndPrint(day09.N2, f, "9", "2")
}

func runDay10() {
	f := getFilesIn("day10/")
	runAndPrint(day10.N1, f, "10", "1")
	runAndPrint(day10.N2, f, "10", "2")
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
