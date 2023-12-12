package day09

import (
	"daniel/aoc2023/util"
	"slices"
	"strconv"
	"strings"
)

// N1 computes the results for Ex1 on the given input-file
func N1(in []string) int {
	sequences := parseGame(in)
	result := 0
	for _, s := range sequences {
		history := s.complete()
		result += history[0].numbers[len(history[0].numbers)-1]
	}

	return result
}

// N2 computes the results for Ex2 on the given input-file
func N2(in []string) int {
	sequences := parseGame(in)
	result := 0
	for _, s := range sequences {
		rev := s.reverse()
		history := rev.complete()
		result += history[0].numbers[len(history[0].numbers)-1]
	}
	return result
}

type sequence struct {
	numbers []int
}

func (s *sequence) complete() (history []sequence) {
	if s.isZeros() {
		s.numbers = append(s.numbers, 0)
		history = append(history, *s)
		return
	}
	sub := s.getSub()
	subHistory := sub.complete()
	s.numbers = append(s.numbers, s.numbers[len(s.numbers)-1]+subHistory[0].numbers[len(subHistory[0].numbers)-1])
	history = append(history, *s)
	history = append(history, subHistory...)
	return
}

func (s *sequence) reverse() (rev sequence) {
	rev.numbers = s.numbers
	slices.Reverse(rev.numbers)
	return
}

func (s *sequence) getSub() (sub sequence) {
	for i := 1; i < len(s.numbers); i++ {
		sub.numbers = append(sub.numbers, s.numbers[i]-s.numbers[i-1])
	}
	return
}

func (s *sequence) isZeros() bool {
	for _, n := range s.numbers {
		if n != 0 {
			return false
		}
	}
	return true
}

func (s *sequence) parse(l string) {
	s.numbers = util.Map(strings.Split(l, " "), func(k string) int {
		i, _ := strconv.Atoi(k)
		return i
	})
}

func parseGame(lines []string) (sequences []sequence) {
	for _, l := range lines {
		s := sequence{}
		s.parse(l)
		sequences = append(sequences, s)
	}
	return
}
