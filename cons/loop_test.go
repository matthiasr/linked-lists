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
