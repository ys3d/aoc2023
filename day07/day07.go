package day07

import (
	"daniel/aoc2023/util"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// N1 computes the results for Ex1 on the given input-file
func N1(in []string) int {
	order := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	hands := util.Map(in, func(s string) (h hand) {
		h.fromLine(s)
		return
	})
	hands = sortByRank(hands, false, order)
	score := 0
	for i, h := range hands {
		score += (i + 1) * h.bid
	}

	return score
}

// N2 computes the results for Ex2 on the given input-file
func N2(in []string) int {
	order := []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}
	hands := util.Map(in, func(s string) (h hand) {
		h.fromLine(s)
		return
	})
	hands = sortByRank(hands, true, order)
	score := 0
	for i, h := range hands {
		score += (i + 1) * h.bid
	}

	return score
}

type chars struct {
	a     int
	k     int
	q     int
	j     int
	t     int
	nine  int
	eight int
	seven int
	six   int
	five  int
	four  int
	three int
	two   int
}

func (c *chars) zeroJ() (c2 chars) {
	c2 = *c
	c2.j = 0
	return
}

func (c *chars) numberOfGroupsOf(size int) (n int) {
	n = util.Sum(util.Map(c.toSlice(), func(i int) int {
		if i == size {
			return 1
		}
		return 0
	}))
	return
}

func (c *chars) toSlice() []int {
	return []int{
		c.a,
		c.k,
		c.q,
		c.j,
		c.t,
		c.nine,
		c.eight,
		c.seven,
		c.six,
		c.five,
		c.four,
		c.three,
		c.two,
	}
}

type hand struct {
	cards string
	bid   int
	c     chars
}

func (h *hand) fromLine(s string) {
	h.cards = strings.Split(s, " ")[0]
	h.bid, _ = strconv.Atoi(strings.Split(s, " ")[1])
	c := chars{
		a:     numberOf(h.cards, "A"),
		k:     numberOf(h.cards, "K"),
		q:     numberOf(h.cards, "Q"),
		j:     numberOf(h.cards, "J"),
		t:     numberOf(h.cards, "T"),
		nine:  numberOf(h.cards, "9"),
		eight: numberOf(h.cards, "8"),
		seven: numberOf(h.cards, "7"),
		six:   numberOf(h.cards, "6"),
		five:  numberOf(h.cards, "5"),
		four:  numberOf(h.cards, "4"),
		three: numberOf(h.cards, "3"),
		two:   numberOf(h.cards, "2"),
	}
	h.c = c
}

func sortByRank(hs []hand, enableJoker bool, order []string) []hand {
	sort.SliceStable(hs, func(i, j int) bool {
		return hs[i].comp(hs[j], enableJoker, order)
	})
	return hs
}

func (h *hand) kind(enableJoker bool) int {
	fiveOfAKind := h.c.numberOfGroupsOf(5) == 1
	fourOfAKind := h.c.numberOfGroupsOf(4) == 1
	fullHouse := h.c.numberOfGroupsOf(3) == 1 && h.c.numberOfGroupsOf(2) == 1
	threeOfAKind := h.c.numberOfGroupsOf(3) == 1
	twoPair := h.c.numberOfGroupsOf(2) == 2
	onePair := h.c.numberOfGroupsOf(2) == 1
	if enableJoker && h.c.j != 0 {
		j0 := h.c.zeroJ()
		fiveOfAKind = h.c.j == 5 || j0.numberOfGroupsOf(5-h.c.j) >= 1
		fourOfAKind = h.c.j == 4 || j0.numberOfGroupsOf(4-h.c.j) >= 1
		fullHouse = h.c.j == 3 || (h.c.j == 2 && j0.numberOfGroupsOf(2) == 1) || (h.c.j == 1 && j0.numberOfGroupsOf(2) == 2)
		threeOfAKind = h.c.j == 3 || j0.numberOfGroupsOf(3-h.c.j) >= 1
		twoPair = h.c.j == 2 || j0.numberOfGroupsOf(2) == 1
		onePair = true
	}
	if fiveOfAKind {
		return 7
	}
	if fourOfAKind {
		return 6
	}
	if fullHouse {
		return 5
	}
	if threeOfAKind {
		return 4
	}
	if twoPair {
		return 3
	}
	if onePair {
		return 2
	}

	return 1
}

func (h *hand) comp(h2 hand, enableJoker bool, order []string) bool {
	k1 := h.kind(enableJoker)
	k2 := h2.kind(enableJoker)
	if k1 < k2 {
		return true
	}
	if k1 > k2 {
		return false
	}
	for i := range h.cards {
		if h.cards[i:i+1] != h2.cards[i:i+1] {
			charComp := charRank(h.cards[i:i+1], order) > charRank(h2.cards[i:i+1], order)
			return charComp
		}
	}
	return true

}

func charRank(s string, order []string) int {
	return slices.Index(order, s)
}

func numberOf(s string, f string) int {
	s2 := strings.Replace(s, f, "", -1)
	return len(s) - len(s2)
}
