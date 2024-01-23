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
	"daniel/aoc2023/day13"
	"daniel/aoc2023/day14"
	"daniel/aoc2023/day15"
	"daniel/aoc2023/day16"
	"daniel/aoc2023/day18"
	"daniel/aoc2023/util"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"runtime/pprof"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

var jobs []job
var threadProfile = pprof.Lookup("threadcreate")

func main() {
	day := readDay()
	mode := readParallelMode()
	scheduleDays(day)
	fmt.Println("Scheduled", len(jobs), "jobs for execution")
	start := time.Now()
	printRuns(runScheduled(mode))
	elapsed := time.Since(start)
	fmt.Println("Execution took", elapsed)
	fmt.Println("Used threads:", threadProfile.Count())

}

func readDay() (day int) {
	day = -1
	if len(os.Args) >= 2 {
		v, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Unknown argument", os.Args[1])
		} else {
			day = v
		}
	}
	if day == -1 {
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
		day, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Unknown day", input)
			day = -1
		}
	}
	return
}

func readParallelMode() bool {
	if len(os.Args) >= 3 && os.Args[2] == "sequential" {
		return false
	}
	return true
}

func scheduleDays(day int) {
	fmt.Println("Scheduling")
	switch day {
	case 0:
		scheduleDay1()
		scheduleDay2()
		scheduleDay3()
		scheduleDay4()
		scheduleDay5()
		scheduleDay6()
		scheduleDay7()
		scheduleDay8()
		scheduleDay9()
		scheduleDay10()
		scheduleDay11()
		scheduleDay12()
		scheduleDay13()
		scheduleDay14()
		scheduleDay15()
		scheduleDay16()
		scheduleDay18()
	case 1:
		scheduleDay1()
	case 2:
		scheduleDay2()
	case 3:
		scheduleDay3()
	case 4:
		scheduleDay4()
	case 5:
		scheduleDay5()
	case 6:
		scheduleDay6()
	case 7:
		scheduleDay7()
	case 8:
		scheduleDay8()
	case 9:
		scheduleDay9()
	case 10:
		scheduleDay10()
	case 11:
		scheduleDay11()
	case 12:
		scheduleDay12()
	case 13:
		scheduleDay13()
	case 14:
		scheduleDay14()
	case 15:
		scheduleDay15()
	case 16:
		scheduleDay16()
	case 18:
		scheduleDay18()
	default:
		fmt.Println("Day does not exist")
	}
}

func scheduleDay1() {
	f := getFilesIn("day01/")
	schedule(day01.N1, f, 1, 1)
	schedule(day01.N2, f, 1, 1)
}

func scheduleDay2() {
	f := getFilesIn("day02/")
	schedule(day02.N1, f, 2, 1)
	schedule(day02.N2, f, 2, 2)
}

func scheduleDay3() {
	f := getFilesIn("day03/")
	schedule(day03.N1, f, 3, 1)
	schedule(day03.N2, f, 3, 2)
}

func scheduleDay4() {
	f := getFilesIn("day04/")
	schedule(day04.N1, f, 4, 1)
	schedule(day04.N2, f, 4, 2)
}

func scheduleDay5() {
	f := getFilesIn("day05/")
	schedule(day05.N1, f, 5, 1)
	schedule(day05.N2, f, 5, 2)
}

func scheduleDay6() {
	f := getFilesIn("day06/")
	schedule(day06.N1, f, 6, 1)
	schedule(day06.N2, f, 6, 2)
}

func scheduleDay7() {
	f := getFilesIn("day07/")
	schedule(day07.N1, f, 7, 1)
	schedule(day07.N2, f, 7, 2)
}

func scheduleDay8() {
	f1 := getFilesIn("day08/inputPart1/")
	f2 := getFilesIn("day08/inputPart2/")
	schedule(day08.N1, f1, 8, 1)
	schedule(day08.N2, f2, 8, 2)
}

func scheduleDay9() {
	f := getFilesIn("day09/")
	schedule(day09.N1, f, 9, 1)
	schedule(day09.N2, f, 9, 2)
}

func scheduleDay10() {
	f := getFilesIn("day10/")
	schedule(day10.N1, f, 10, 1)
	schedule(day10.N2, f, 10, 2)
}

func scheduleDay11() {
	f := getFilesIn("day11/")
	schedule(day11.N1, f, 11, 1)
	schedule(day11.N2, f, 11, 2)
}

func scheduleDay12() {
	f := getFilesIn("day12/")
	schedule(day12.N1, f, 12, 1)
	schedule(day12.N2, f, 12, 2)
}

func scheduleDay13() {
	f := getFilesIn("day13/")
	schedule(day13.N1, f, 13, 1)
	schedule(day13.N2, f, 13, 2)
}

func scheduleDay14() {
	f := getFilesIn("day14/")
	schedule(day14.N1, f, 14, 1)
	schedule(day14.N2, f, 14, 2)
}

func scheduleDay15() {
	f := getFilesIn("day15/")
	schedule(day15.N1, f, 15, 1)
	schedule(day15.N2, f, 15, 2)
}

func scheduleDay16() {
	f := getFilesIn("day16/")
	schedule(day16.N1, f, 16, 1)
	schedule(day16.N2, f, 16, 2)
}

func scheduleDay18() {
	f := getFilesIn("day18/")
	schedule(day18.N1, f, 18, 1)
	schedule(day18.N2, f, 18, 2)
}

type job struct {
	f    func([]string) int
	file string
	in   []string
	day  int
	ex   int
}

func runJob(j job, ch chan run, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	out := j.f(j.in)
	elapsed := time.Since(start)
	ch <- run{
		j.day,
		j.ex,
		j.file,
		out,
		elapsed,
	}
}

type run struct {
	day      int
	ex       int
	input    string
	result   int
	duration time.Duration
}

func runScheduled(parallelMode bool) (runs []run) {
	ch := make(chan run, len(jobs))
	wg := sync.WaitGroup{}
	if parallelMode {
		fmt.Println("Starting parallel execution")
	} else {
		fmt.Println("Starting sequential execution")
	}
	for _, j := range jobs {
		wg.Add(1)
		if parallelMode {
			go runJob(j, ch, &wg)
		} else {
			runJob(j, ch, &wg)
		}
	}

	wg.Wait()
	close(ch)
	for r := range ch {

		runs = append(runs, r)
	}
	slices.SortFunc(runs, func(a run, b run) int {
		if a.day < b.day {
			return -1
		} else if a.day > b.day {
			return 1
		} else if a.ex < b.ex {
			return -1
		} else if a.ex > b.ex {
			return 1
		}
		return strings.Compare(a.input, b.input)
	})
	return
}

func schedule(f func([]string) int, files []string, day int, ex int) {
	for _, file := range files {
		in, err := util.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		j := job{
			f:    f,
			file: file,
			in:   in,
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
