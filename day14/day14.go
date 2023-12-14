package day14

// N1 computes the results for Ex1 on the given input-file
func N1(in []string) int {
	f := parse(in)
	f = f.tiltNorth()
	return eval(f)
}

// N2 computes the results for Ex2 on the given input-file
func N2(in []string) int {
	rounds := 1000000000
	f := parse(in)
	cache := make(map[string]int)
	for i := 0; i < rounds; i++ {
		f = f.tiltNorth()
		f = f.tiltWest()
		f = f.tiltSouth()
		f = f.tiltEast()
		if val, ok := cache[f.fieldToString()]; ok {
			// Cycle detected, skip to the end
			for j := 0; j < (((rounds - 1) - i) % (i - val)); j++ {
				f = f.tiltNorth()
				f = f.tiltWest()
				f = f.tiltSouth()
				f = f.tiltEast()
			}
			return eval(f)
		}
		cache[f.fieldToString()] = i
	}
	return eval(f)
}

type field struct {
	f [][]rune
}

func (f *field) fieldToString() (s string) {
	var runes []rune
	for _, l := range f.f {
		runes = append(runes, l...)
	}
	s = string(runes)
	return
}

func (f *field) tiltNorth() field {
	space := make([]int, len(f.f[0]))
	for i := range f.f {
		for j := range f.f[i] {
			switch f.f[i][j] {
			case '#':
				space[j] = i + 1
			case 'O':
				tmp := f.f[space[j]][j]
				f.f[space[j]][j] = f.f[i][j]
				f.f[i][j] = tmp
				space[j]++
			}
		}
	}
	out := field{f: f.f}
	return out
}

func (f *field) tiltWest() field {
	for i := range f.f {
		space := 0
		for j := range f.f[i] {
			switch f.f[i][j] {
			case '#':
				space = j + 1
			case 'O':
				tmp := f.f[i][space]
				f.f[i][space] = f.f[i][j]
				f.f[i][j] = tmp
				space++
			}
		}
	}
	out := field{f: f.f}
	return out
}

func (f *field) tiltSouth() field {
	space := make([]int, len(f.f[0]))
	for i := range space {
		space[i] = len(f.f) - 1
	}
	for i := len(f.f) - 1; i >= 0; i-- {
		for j := range f.f[i] {
			switch f.f[i][j] {
			case '#':
				space[j] = i - 1
			case 'O':
				tmp := f.f[space[j]][j]
				f.f[space[j]][j] = f.f[i][j]
				f.f[i][j] = tmp
				space[j]--
			}
		}
	}
	out := field{f: f.f}
	return out
}

func (f *field) tiltEast() field {
	for i := range f.f {
		space := len(f.f[i]) - 1
		for j := len(f.f[i]) - 1; j >= 0; j-- {
			switch f.f[i][j] {
			case '#':
				space = j - 1
			case 'O':
				tmp := f.f[i][space]
				f.f[i][space] = f.f[i][j]
				f.f[i][j] = tmp
				space--
			}
		}
	}
	out := field{f: f.f}
	return out
}

func eval(in field) (out int) {
	for x, l := range in.f {
		for _, c := range l {
			if c == 'O' {
				out += len(in.f) - x
			}
		}
	}
	return
}

func parse(lines []string) (out field) {
	for _, l := range lines {
		out.f = append(out.f, []rune(l))
	}
	return
}
