package cons

import (
	"reflect"
	"testing"

	p "github.com/tonnerre/golang-pretty"
)

func TestMake(t *testing.T) {
	if want, got := (&Cell{"x", &Cell{"x", &Cell{"x", nil}}}), Make(3, "x"); !reflect.DeepEqual(want, got) {
		t.Errorf("want: %# v\ngot: %# v", p.Formatter(want), p.Formatter(got))
	}
}
