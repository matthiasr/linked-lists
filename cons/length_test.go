package cons

import (
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
