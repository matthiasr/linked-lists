package cons

import (
	"testing"
)

func TestNewCons(t *testing.T) {
	if want, got := (Cell{"first", nil}), *NewCell("first", nil); want != got {
		t.Errorf("want: %+v, got: %+v", want, got)
	}
}
