package cons

import (
	"reflect"
	"testing"

	p "github.com/tonnerre/golang-pretty"
)

func TestNewCons(t *testing.T) {
	if want, got := (Cell{"first", nil}), *NewCell("first", nil); want != got {
		t.Errorf("want: %+v, got: %+v", want, got)
	}

	if want, got := (Cell{"first", &Cell{"second", nil}}), *NewCell("first", NewCell("second", nil)); !reflect.DeepEqual(want, got) {
		t.Errorf("want: %# v\ngot: %# v", p.Formatter(want), p.Formatter(got))
	}
}

func TestCarCdr(t *testing.T) {
	in := NewCell("1", NewCell("2", NewCell("3", nil)))
	if want, got := "1", in.Car(); want != got {
		t.Errorf("car: want %+v, got %+v", want, got)
	}
	if want, got := "2", in.Cdr().Car(); want != got {
		t.Errorf("cdar: want %+v, got %+v", want, got)
	}
	if want, got := "3", in.Cdr().Cdr().Car(); want != got {
		t.Errorf("cddar: want %+v, got %+v", want, got)
	}
}

func TestNil(t *testing.T) {
	in := NewCell("1", NewCell("2", nil))
	if want, got := false, in.Nil(); want != got {
		t.Errorf("expected not to be nil: %# v", in)
	}
	if want, got := false, in.Cdr().Nil(); want != got {
		t.Errorf("expected not to be nil: %# v", in.Cdr())
	}
	if want, got := true, in.Cdr().Cdr().Nil(); want != got {
		t.Errorf("expected to be nil: %# v", in.Cdr().Cdr())
	}
}

func TestMake(t *testing.T) {
	if want, got := (&Cell{"x", &Cell{"x", &Cell{"x", nil}}}), Make(3, "x"); !reflect.DeepEqual(want, got) {
		t.Errorf("want: %# v\ngot: %# v", p.Formatter(want), p.Formatter(got))
	}
}

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
