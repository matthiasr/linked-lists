package cons

import (
	"testing"
)

func TestMakeLoop(t *testing.T) {
	cases := []uint{1, 2, 3, 4, 50, 1000}
	for _, c := range cases {
		a := MakeLoop(c, "x")
		b := a
		for i := uint(0); i < c; i++ {
			b = b.Cdr()
		}

		if a != b {
			t.Errorf("expected to be back at start of the loop for depth %d, want: %# v\ngot: %# v", c, a, b)
		}
	}
}

func TestMakeWithLoop(t *testing.T) {
	cases := []struct{ n, m uint }{
		{1, 1},
		{3, 1},
		{1, 3},
		{10, 10},
		{500, 100},
		{1000, 1000},
	}

	for _, c := range cases {
		l := MakeWithLoop(c.n, c.m, "x")
		lastBeforeLoop := l
		for i := uint(0); i < c.n-1; i++ {
			lastBeforeLoop = lastBeforeLoop.Cdr()
		}
		firstInLoop := lastBeforeLoop.Cdr()
		x := firstInLoop
		for j := uint(0); j < c.m; j++ {
			x = x.Cdr()
			if x == lastBeforeLoop {
				t.Errorf("expected not to find %# v in loop, but did for case %# v", lastBeforeLoop, c)
			}
		}
		if x != firstInLoop {
			t.Errorf("expected to reach first element in loop again for case %# v, want: %# v, got: %# v", c, firstInLoop, x)
		}
	}
}
