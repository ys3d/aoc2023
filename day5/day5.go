package day5

import (
	"daniel/aoc2023/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// N1 computes the results for Ex1 on the given input-file
func N1(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	seeds, soil, fertilizer, water, light, temp, humidity, location := parseGame(in)
	pos := apply(seeds, soil)
	pos = apply(pos, fertilizer)
	pos = apply(pos, water)
	pos = apply(pos, light)
	pos = apply(pos, temp)
	pos = apply(pos, humidity)
	pos = apply(pos, location)

	return strconv.Itoa(slices.Min(pos))
}

// N2 computes the results for Ex2 on the given input-file
func N2(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	seeds, soil, fertilizer, water, light, temp, humidity, location := parseGameInterval(in)
	pos := applyInterval(seeds, soil)
	pos = applyInterval(pos, fertilizer)
	pos = applyInterval(pos, water)
	pos = applyInterval(pos, light)
	pos = applyInterval(pos, temp)
	pos = applyInterval(pos, humidity)
	pos = applyInterval(pos, location)

	return strconv.Itoa(slices.Min(util.Map(pos, func(i Interval) int {
		return i.start
	})))
}

// Interval represents an interval.
// It  includes values in [start, start+length)
type Interval struct {
	start  int
	length int
}

// cut computes an Interval that is contained in this Interval and gets modified by the given map
// If a cut is found the boolean is set to true and the Interval is returned.
// In case no cut exists, false and an undefined Interval get returned
func (i *Interval) cut(m Map) (bool, Interval) {
	if i.start < m.sourceRangeStart+m.rangeLength && i.start+i.length > m.sourceRangeStart {
		start := max(i.start, m.sourceRangeStart)
		length := min(i.start+i.length, m.sourceRangeStart+m.rangeLength) - start
		j := Interval{
			start:  start,
			length: length,
		}
		return true, j
	}
	return false, Interval{}
}

// Remove computes remaining intervals by removing the Interval j.
// If both intervals do not cut each other, no intervals get returned
func (i *Interval) remove(j Interval) (k []Interval) {
	if j.start > i.start {
		k = append(k, Interval{
			start:  i.start,
			length: j.start - i.start,
		})
	}
	if j.start+j.length < i.start+i.length {
		k = append(k, Interval{
			start:  j.start + j.length,
			length: i.start + i.length - (j.start + j.length),
		})
	}
	return
}

// Map represents a mapping from one space to another
type Map struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

// modify modifies an Interval.
// It returns three values:
//   - [bool] if a modification took place
//   - [Interval] The modified intervals
//   - [[]Interval] Remaining intervals
func (m *Map) modify(i Interval) (bool, Interval, []Interval) {
	found, c := i.cut(*m)
	if found {
		return true, Interval{
			start:  c.start + m.destinationRangeStart - m.sourceRangeStart,
			length: c.length,
		}, i.remove(c)
	}
	return false, Interval{}, []Interval{}
}

// fromLine parses a Map
func (m *Map) fromLine(line string) {
	v := util.Map(strings.Split(line, " "), func(s string) int {
		out, _ := strconv.Atoi(s)
		return out
	})
	m.destinationRangeStart = v[0]
	m.sourceRangeStart = v[1]
	m.rangeLength = v[2]
}

// contains checks if a given index is contained in the Map
func (m *Map) contains(seed int) bool {
	return seed > m.sourceRangeStart && seed < m.sourceRangeStart+m.rangeLength
}

// indexOf returns the new index of the given old index after the modification by the Map took place
// In case the map does not contain this index, the value is arbitrary
func (m *Map) indexOf(seed int) int {
	return m.destinationRangeStart + (seed - m.sourceRangeStart)
}

// apply applies a sequence of maps on a sequence of positions.
func apply(pos []int, maps []Map) (newPos []int) {
	for _, p := range pos {
		found := false
		for _, m := range maps {
			if !found && m.contains(p) {
				found = true
				newPos = append(newPos, m.indexOf(p))
			}
		}
		if !found {
			newPos = append(newPos, p)
		}
	}
	return
}

// apply applies a sequence of maps on a sequence of intervals.
func applyInterval(pos []Interval, maps []Map) (newPos []Interval) {
	for _, m := range maps {
		j := 0
		for j < len(pos) {
			i := pos[j]
			found, newI, remainingI := m.modify(i)
			if found {
				newPos = append(newPos, newI)
				pos = append(pos, remainingI...)
				pos = slices.Delete(pos, j, j+1)
			} else {
				j++
			}

		}
	}
	for _, i := range pos {
		newPos = append(newPos, i)
	}
	return
}

func parseGame(in []string) (seeds []int, soil []Map, fertilizer []Map, water []Map, light []Map, temp []Map, humidity []Map, location []Map) {
	seeds = util.Map(strings.Split(in[0], " ")[1:], func(s string) int {
		out, _ := strconv.Atoi(s)
		return out
	})
	mapsPart := in[2:]
	end := slices.Index(mapsPart, "")
	soil = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	fertilizer = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	water = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	light = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	temp = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	humidity = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	location = util.Map(mapsPart[1:], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	return
}

func parseGameInterval(in []string) (seeds []Interval, soil []Map, fertilizer []Map, water []Map, light []Map, temp []Map, humidity []Map, location []Map) {
	splits := strings.Split(in[0], " ")[1:]
	for i := range splits {
		if i%2 == 1 {
			start, _ := strconv.Atoi(splits[i-1])
			length, _ := strconv.Atoi(splits[i])
			seeds = append(seeds, Interval{start: start, length: length})
		}
	}
	mapsPart := in[2:]
	end := slices.Index(mapsPart, "")
	soil = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	fertilizer = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	water = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	light = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	temp = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	end = slices.Index(mapsPart, "")
	humidity = util.Map(mapsPart[1:end], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	mapsPart = mapsPart[end+1:]
	location = util.Map(mapsPart[1:], func(s string) (m Map) {
		m.fromLine(s)
		return
	})
	return
}
