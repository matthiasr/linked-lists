package cons

import (
	"fmt"
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

func BenchmarkLoopCdr(b *testing.B) {
	cases := []uint{1, 10, 100, 1000, 1e4, 1e5, 1e6, 1e7}
	for _, c := range cases {
		l := MakeLoop(c, "x")
		b.Run(fmt.Sprintf("LoopCdr%d", c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l = l.Cdr()
			}
		})
	}
}
