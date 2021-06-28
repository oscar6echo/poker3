package poker

import "testing"

func setup() {
	SetupKeys()
	BuildEvalFiveTables(false)
}

func TestGetRankFive(t *testing.T) {

	setup()

	cases := []struct {
		cards [5]int
		rank  int
	}{
		{[5]int{21, 33, 24, 22, 39}, 2459},
		{[5]int{51, 38, 14, 36, 17}, 3431},
		{[5]int{45, 8, 48, 34, 5}, 1171},
		{[5]int{13, 37, 33, 20, 35}, 3106},
		{[5]int{31, 26, 50, 16, 49}, 3971},
		{[5]int{28, 24, 25, 29, 2}, 4434},
		{[5]int{41, 13, 28, 25, 16}, 310},
		{[5]int{20, 36, 7, 42, 43}, 3572},
		{[5]int{38, 42, 8, 22, 44}, 761},
		{[5]int{32, 3, 18, 5, 42}, 320},
	}

	for _, c := range cases {
		got := GetRankFive(c.cards)
		if got != c.rank {
			t.Errorf("GetRankFive(%v) == %d, want %d", c.cards, got, c.rank)
		} else {
			t.Logf("GetRankFive(%v) == %d", c.cards, got)

		}
	}
}

func TestGetRankSeven(t *testing.T) {

	setup()

	cases := []struct {
		cards [7]int
		rank  int
	}{
		{[7]int{50, 6, 0, 5, 38, 7, 17}, 5124},
		{[7]int{23, 16, 34, 26, 0, 10, 8}, 1766},
		{[7]int{14, 4, 0, 7, 20, 8, 47}, 1625},
		{[7]int{10, 32, 43, 3, 25, 8, 49}, 1925},
		{[7]int{1, 16, 49, 24, 43, 42, 33}, 3676},
		{[7]int{49, 17, 1, 26, 11, 34, 20}, 887},
		{[7]int{5, 4, 18, 31, 34, 48, 22}, 1689},
		{[7]int{13, 47, 1, 25, 38, 26, 51}, 2815},
		{[7]int{44, 2, 28, 1, 3, 18, 22}, 5046},
		{[7]int{49, 27, 33, 51, 22, 1, 30}, 4000},
	}

	for _, c := range cases {
		got := GetRankSeven(c.cards)
		if got != c.rank {
			t.Errorf("GetRankSeven(%v) == %d, want %d", c.cards, got, c.rank)
		} else {
			t.Logf("GetRankSeven(%v) == %d", c.cards, got)

		}
	}
}
