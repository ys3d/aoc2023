package day08

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
	inst, nodes := parseGame(in)
	ni := slices.IndexFunc(nodes, func(n node) bool {
		return n.name == "AAA"
	})
	steps := 0
	for nodes[ni].name != "ZZZ" {
		instruct := inst.next()
		steps++
		if instruct == "L" {
			ni = slices.IndexFunc(nodes, func(n node) bool {
				return n.name == nodes[ni].left
			})
		} else {
			ni = slices.IndexFunc(nodes, func(n node) bool {
				return n.name == nodes[ni].right
			})
		}
	}

	return strconv.Itoa(steps)
}

// N2 computes the results for Ex2 on the given input-file
func N2(file string) (out string) {
	// Assumption the path length from the start to Z is equal to the length from Z to Z
	in, err := util.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	inst, nodes := parseGame(in)

	starts := util.Filter(nodes, func(n node) bool {
		return n.name[2:3] == "A"
	})

	indices := util.Map(starts, func(k node) int {
		return slices.IndexFunc(nodes, func(n node) bool {
			return n.name == k.name
		})
	})
	pathLength := util.Map(indices, func(k int) int {
		return 0
	})

	for i := range indices {
		for nodes[indices[i]].name[2:3] != "Z" {
			instruct := inst.next()
			pathLength[i]++
			if instruct == "L" {
				indices[i] = slices.IndexFunc(nodes, func(n node) bool {
					return n.name == nodes[indices[i]].left
				})
			} else {
				indices[i] = slices.IndexFunc(nodes, func(n node) bool {
					return n.name == nodes[indices[i]].right
				})
			}
		}
	}
	var factors = make(map[int]int)
	for _, p := range pathLength {
		localFactors := util.PrimeFactors(p)
		slices.Sort(localFactors)
		dict := make(map[int]int)
		for _, num := range localFactors {
			dict[num] = dict[num] + 1
		}
		for k, v := range dict {
			number, ok := factors[k]
			if ok && number < v {
				number = 0
			} else if !ok {
				factors[k] = v
			}
		}
	}
	res := 1
	for k, v := range factors {
		for i := 0; i < v; i++ {
			res *= k
		}
	}
	return strconv.Itoa(res)
}

type instruction struct {
	inst string
	pos  int
}

func (i *instruction) next() (s string) {
	s = i.inst[i.pos : i.pos+1]
	i.pos++
	if i.pos == len(i.inst) {
		i.pos = 0
	}
	return
}

type node struct {
	name  string
	left  string
	right string
}

func (n *node) parseNode(l string) {
	split := strings.Split(l, " ")
	n.name = split[0]
	n.left = split[2][1:4]
	n.right = split[3][:3]
}

func parseGame(lines []string) (instructions instruction, nodes []node) {
	instructions = instruction{lines[0], 0}
	for _, l := range lines[2:] {
		n := node{}
		n.parseNode(l)
		nodes = append(nodes, n)
	}
	return
}
