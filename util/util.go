package util

import (
	"github.com/fxtlabs/primes"
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

func Map[K any, T any](arr []K, f func(K) T) []T {
	var out []T
	for _, v := range arr {
		out = append(out, f(v))
	}
	return out
}

func Filter[K any](arr []K, f func(K) bool) (k []K) {
	for _, v := range arr {
		if f(v) {
			k = append(k, v)
		}
	}
	return
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

func PrimeFactors(i int) (out []int) {
	primesSelection := primes.Sieve(i)
	for _, p := range primesSelection {
		li := i
		for li%p == 0 {
			out = append(out, p)
			li /= p
		}
	}
	return
}
func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
