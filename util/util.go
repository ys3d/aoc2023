package util

import (
	"os"
	"strings"
)

func ReadFile(path string) (s []string, err error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return
	}
	s = strings.Split(string(dat), "\n")
	return
}

func Sum[K int | float64](arr []K) K {
	var sum K
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

func ToText(i int) string {
	switch i {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return ""
	}
}
