package cons

type Cell struct {
	this string
	next *Cell
}

func NewCell(this string, next *Cell) *Cell {
	return &Cell{this: this, next: next}
}

func (c *Cell) Car() string {
	return c.this
}

func (c *Cell) Cdr() *Cell {
	return c.next
}

func (c *Cell) Nil() bool {
	return c == nil
}

// Make creates a linked list of the specified
// length where all cells have the same content.
func Make(length uint, s string) *Cell {
	l := NewCell(s, nil)
	for i := uint(1); i < length; i++ {
		l = NewCell(s, l)
	}
	return l
}

// The length of the list starting at the cell. May not terminate if the list has a loop.
func (c *Cell) Length() uint {
	var l uint
	for ; !c.Nil(); c = c.Cdr() {
		l++
	}
	return l
}
