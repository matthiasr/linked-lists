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
