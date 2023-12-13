package day13

import (
	"daniel/aoc2023/util"
)

// N1 computes the results for Ex1 on the given input-file
func N1(in []string) int {
	valleys := parse(in)
	return util.Sum(util.Map(valleys, func(v valley) int {
		r := v.reflectionIndex(0)
		return r
	}))
}

// N2 computes the results for Ex2 on the given input-file
func N2(in []string) int {
	valleys := parse(in)
	return util.Sum(util.Map(valleys, func(v valley) int {
		r := v.reflectionIndex(1)
		return r
	}))
}

type valley struct {
	c []string
}

func (v *valley) reflectionIndex(maxCorrections int) int {
	for i := 0; i < len(v.c)-1; i++ {
		if v.checkForHorizontalReflectionAt(i, maxCorrections) {
			return 100 * (i + 1)
		}
	}
	for i := 0; i < len(v.c[0])-1; i++ {
		if v.checkForVerticalReflectionAt(i, maxCorrections) {
			return i + 1
		}
	}
	return 0
}

func (v *valley) checkForHorizontalReflectionAt(i int, maxCorrections int) bool {
	for j := 0; (i-j >= 0) && (i+1+j < len(v.c)); j++ {
		for k := 0; k < len(v.c[i-j]); k++ {
			if v.c[i-j][k] != v.c[i+1+j][k] {
				maxCorrections--
				if maxCorrections < 0 {
					return false
				}
			}
		}
	}
	return maxCorrections == 0
}

func (v *valley) checkForVerticalReflectionAt(i int, maxCorrections int) bool {
	for j := 0; (i-j >= 0) && (i+1+j < len(v.c[0])); j++ {
		for _, l := range v.c {
			if l[i-j] != l[i+1+j] {
				maxCorrections--
				if maxCorrections < 0 {
					return false
				}
			}
		}
	}
	return maxCorrections == 0
}

func parse(lines []string) (valleys []valley) {
	v := valley{}
	for _, l := range lines {
		if l == "" {
			valleys = append(valleys, v)
			v = valley{}
		} else {
			v.c = append(v.c, l)
		}
	}
	valleys = append(valleys, v)
	return
}
