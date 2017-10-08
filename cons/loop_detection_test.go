package cons

import (
	"testing"
)

var cases = []struct {
	desc string
	want bool
	l    *Cell
}{
	{"nil", false, nil},
	{"Make_1", false, Make(1, "x")},
	{"Make_10", false, Make(10, "x")},
	{"Make_100", false, Make(100, "x")},
	{"Make_1000", false, Make(1000, "x")},
	{"Make_10000", false, Make(10000, "x")},
	{"MakeLoop_1", true, MakeLoop(1, "x")},
	{"MakeLoop_10", true, MakeLoop(10, "x")},
	{"MakeLoop_100", true, MakeLoop(100, "x")},
	{"MakeLoop_1000", true, MakeLoop(1000, "x")},
	{"MakeLoop_10000", true, MakeLoop(10000, "x")},
	{"MakeWithLoop_1_1", true, MakeWithLoop(1, 1, "x")},
	{"MakeWithLoop_10_10", true, MakeWithLoop(10, 10, "x")},
	{"MakeWithLoop_100_100", true, MakeWithLoop(100, 100, "x")},
	{"MakeWithLoop_1000_1000", true, MakeWithLoop(1000, 1000, "x")},
	{"MakeWithLoop_10000_10000", true, MakeWithLoop(10000, 10000, "x")},
	{"MakeWithLoop_1_10", true, MakeWithLoop(1, 10, "x")},
	{"MakeWithLoop_1_100", true, MakeWithLoop(1, 100, "x")},
	{"MakeWithLoop_1_1000", true, MakeWithLoop(1, 1000, "x")},
	{"MakeWithLoop_1_10000", true, MakeWithLoop(1, 10000, "x")},
	{"MakeWithLoop_10_1", true, MakeWithLoop(10, 1, "x")},
	{"MakeWithLoop_100_1", true, MakeWithLoop(100, 1, "x")},
	{"MakeWithLoop_1000_1", true, MakeWithLoop(1000, 1, "x")},
	{"MakeWithLoop_10000_1", true, MakeWithLoop(10000, 1, "x")},
}

func TestDetectLoopSimple(t *testing.T) {
	for _, c := range cases {
		if got := c.l.DetectLoopSimple(); got != c.want {
			t.Errorf("unexpected result for test case %v, want: %v, got: %v", c.desc, c.want, got)
		}
	}
}

func BenchmarkDetectLoopSimple(b *testing.B) {
	for _, c := range cases {
		b.Run(c.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.l.DetectLoopSimple()
			}
		})
	}
}
