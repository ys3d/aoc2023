package day12

import (
	"daniel/aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

// N1 computes the results for Ex1 on the given input-file
func N1(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	instances := util.Map(in, func(s string) (i instance) {
		i.fromString(s)
		return
	})
	result := util.Sum(util.Map(instances, func(i instance) int {
		return i.alternatives()
	}))
	return strconv.Itoa(result)
}

// N2 computes the results for Ex2 on the given input-file
func N2(file string) (out string) {
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	instances := util.Map(in, func(l string) (i instance) {
		i.fromString(l)
		// Enlarge input
		s := i.s
		g := i.groups
		for j := 0; j < 4; j++ {
			i.s += "?" + s
			i.groups = append(i.groups, g...)
		}
		return
	})
	result := util.Sum(util.Map(instances, func(i instance) int {
		r := i.alternatives()
		//fmt.Println(i.s, " ", i.groups, " : ", i.alternatives())
		return r
	}))
	return strconv.Itoa(result)
}

type instance struct {
	s      string
	groups []int
}

func (i *instance) fromString(s string) {
	split := strings.Split(s, " ")
	i.s = split[0]
	i.groups = util.Map(strings.Split(split[1], ","), func(s string) (j int) {
		j, _ = strconv.Atoi(s)
		return
	})
}

func (i *instance) alternatives() int {
	if val, ok := cache[i.s+fmt.Sprint(i.groups)]; ok {
		return val
	}
	if len(i.groups) == 0 {
		if !strings.Contains(i.s, "#") {
			cache[i.s+fmt.Sprint(i.groups)] = 1
			return 1
		} else {
			cache[i.s+fmt.Sprint(i.groups)] = 0
			return 0
		}
	}
	for strings.HasPrefix(i.s, ".") {
		i.s = strings.TrimPrefix(i.s, ".")
	}
	if strings.Count(i.s, "#")+strings.Count(i.s, "?") < i.groups[0] {
		cache[i.s+fmt.Sprint(i.groups)] = 0
		return 0
	}
	switch i.s[0] {
	case '?':
		newI1 := instance{
			s:      i.s[1:],
			groups: i.groups,
		}
		newI2 := instance{
			s:      "#" + i.s[1:],
			groups: i.groups,
		}
		result := newI1.alternatives() + newI2.alternatives()
		cache[i.s+fmt.Sprint(i.groups)] = result
		return result
	case '#':
		if hasPrefixOfSize(i.s, i.groups[0]) {
			var newI instance
			if len(i.s) == i.groups[0] {
				newI = instance{
					s:      "",
					groups: i.groups[1:],
				}
			} else {
				newI = instance{
					s:      i.s[i.groups[0]+1:],
					groups: i.groups[1:],
				}
			}
			result := newI.alternatives()
			cache[i.s+fmt.Sprint(i.groups)] = result
			return result
		}
	}

	return 0
}

func hasPrefixOfSize(s string, size int) bool {
	for i := 0; i < size; i++ {
		if i == len(s) || s[i] == '.' {
			return false
		}
	}
	if len(s) > size {
		if s[size] == '#' {
			return false
		}
	}

	return true
}
