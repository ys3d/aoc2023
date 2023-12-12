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
	"daniel/aoc2023/day11"
	"daniel/aoc2023/day12"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strconv"
	"strings"
	"time"
)

const Day = 0

var jobs []job

//var runs []run

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
		runDay11()
		runDay12()
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
	case 11:
		runDay11()
	case 12:
		runDay12()
	default:
		fmt.Println("Day does not exist")
	}
	printRuns(runScheduled())
}

func runDay1() {
	f := getFilesIn("day01/")
	schedule(day01.N1, f, 1, 1)
	schedule(day01.N2, f, 1, 1)
}

func runDay2() {
	f := getFilesIn("day02/")
	schedule(day02.N1, f, 2, 1)
	schedule(day02.N2, f, 2, 2)
}

func runDay3() {
	f := getFilesIn("day03/")
	schedule(day03.N1, f, 3, 1)
	schedule(day03.N2, f, 3, 2)
}

func runDay4() {
	f := getFilesIn("day04/")
	schedule(day04.N1, f, 4, 1)
	schedule(day04.N2, f, 4, 2)
}

func runDay5() {
	f := getFilesIn("day05/")
	schedule(day05.N1, f, 5, 1)
	schedule(day05.N2, f, 5, 2)
}

func runDay6() {
	f := getFilesIn("day06/")
	schedule(day06.N1, f, 6, 1)
	schedule(day06.N2, f, 6, 2)
}

func runDay7() {
	f := getFilesIn("day07/")
	schedule(day07.N1, f, 7, 1)
	schedule(day07.N2, f, 7, 2)
}

func runDay8() {
	f1 := getFilesIn("day08/inputPart1/")
	f2 := getFilesIn("day08/inputPart2/")
	schedule(day08.N1, f1, 8, 1)
	schedule(day08.N2, f2, 8, 2)
}

func runDay9() {
	f := getFilesIn("day09/")
	schedule(day09.N1, f, 9, 1)
	schedule(day09.N2, f, 9, 2)
}

func runDay10() {
	f := getFilesIn("day10/")
	schedule(day10.N1, f, 10, 1)
	schedule(day10.N2, f, 10, 2)
}

func runDay11() {
	f := getFilesIn("day11/")
	schedule(day11.N1, f, 11, 1)
	schedule(day11.N2, f, 11, 2)
}

func runDay12() {
	f := getFilesIn("day12/")
	schedule(day12.N1, f, 12, 1)
	schedule(day12.N2, f, 12, 2)
}

type job struct {
	f    func(string) string
	file string
	day  int
	ex   int
}

func (j *job) run() run {
	start := time.Now()
	out := j.f(j.file)
	elapsed := time.Since(start)
	return run{
		strconv.Itoa(j.day),
		strconv.Itoa(j.ex),
		j.file,
		out,
		elapsed,
	}
}

type run struct {
	day      string
	ex       string
	input    string
	result   string
	duration time.Duration
}

func runScheduled() (runs []run) {
	for _, j := range jobs {
		runs = append(runs, j.run())
	}
	return
}

func schedule(f func(string) string, files []string, day int, ex int) {
	for _, file := range files {
		j := job{
			f:    f,
			file: file,
			day:  day,
			ex:   ex,
		}
		jobs = append(jobs, j)
	}
}

func printRuns(runs []run) {
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
