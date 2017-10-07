package cons

import (
	"fmt"
	"testing"

	p "github.com/tonnerre/golang-pretty"
)

func TestLength(t *testing.T) {
	cases := []struct {
		l uint
		c *Cell
	}{
		{0, nil},
		{1, &Cell{"x", nil}},
		{2, &Cell{"x", &Cell{"x", nil}}},
		{1, NewCell("x", nil)},
		{2, NewCell("x", NewCell("x", nil))},
		{100, Make(100, "x")},
	}

	for _, c := range cases {
		if want, got := c.l, c.c.Length(); want != got {
			t.Errorf("unexpected length: want %d, got: %d for: %# v", want, got, p.Formatter(c.c))
		}
	}
}

func BenchmarkLength(b *testing.B) {
	cases := []uint{1, 10, 100, 1000, 1e4, 1e5, 1e6, 1e7}
	for _, c := range cases {
		l := Make(c, "x")
		b.Run(fmt.Sprintf("Length%d", c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l.Length()
			}
		})
	}
}
