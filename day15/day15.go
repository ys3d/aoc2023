package day15

import (
	"daniel/aoc2023/util"
	"slices"
	"strconv"
	"strings"
)

// N1 computes the results for Ex1 on the given input-file
func N1(in []string) int {
	return util.Sum(util.Map(strings.Split(in[0], ","), func(s string) int {
		return getHash(s)
	}))
}

// N2 computes the results for Ex2 on the given input-file
func N2(in []string) int {
	split := strings.Split(in[0], ",")
	var boxes = make([][]Lens, 256)
	for _, o := range split {
		if o[len(o)-1] == '-' {
			label := o[:len(o)-1]
			hash := getHash(label)
			index := slices.IndexFunc(boxes[hash], func(lens Lens) bool {
				return lens.label == label
			})
			if index >= 0 {
				boxes[hash] = append(boxes[hash][:index], boxes[hash][index+1:]...)
			}
		} else {
			label := strings.Split(o, "=")[0]
			hash := getHash(label)
			index := slices.IndexFunc(boxes[hash], func(lens Lens) bool {
				return lens.label == label
			})
			focalLength, _ := strconv.Atoi(o[len(label)+1:])
			if index >= 0 {
				boxes[hash][index].focalLength = focalLength
			} else {
				boxes[hash] = append(boxes[hash], Lens{label: label, focalLength: focalLength})
			}
		}
	}
	totalFocusingPower := 0
	for boxNumber, b := range boxes {
		for slotNumber, l := range b {
			focusingPower := (1 + boxNumber) * (slotNumber + 1) * l.focalLength
			totalFocusingPower += focusingPower
		}
	}
	return totalFocusingPower
}

type Lens struct {
	label       string
	focalLength int
}

func getHash(s string) (hash int) {
	for _, c := range s {
		ascii := int(c)
		hash += ascii
		hash *= 17
		hash %= 256
	}
	return
}
