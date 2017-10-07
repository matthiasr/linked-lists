package cons

import (
	"fmt"
	"reflect"
	"testing"

	p "github.com/tonnerre/golang-pretty"
)

func TestMake(t *testing.T) {
	if want, got := (&Cell{"x", &Cell{"x", &Cell{"x", nil}}}), Make(3, "x"); !reflect.DeepEqual(want, got) {
		t.Errorf("want: %# v\ngot: %# v", p.Formatter(want), p.Formatter(got))
	}
}

func BenchmarkMake(b *testing.B) {
	cases := []uint{1, 10, 100, 1000, 1e4, 1e5, 1e6}
	for _, c := range cases {
		b.Run(fmt.Sprintf("Make%d", c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Make(c, "x")
			}
		})
	}
}
