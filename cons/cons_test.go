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
