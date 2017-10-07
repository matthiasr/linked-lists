package cons

// The length of the list starting at the cell. May not terminate if the list has a loop.
func (c *Cell) Length() uint {
	var l uint
	for ; !c.Nil(); c = c.Cdr() {
		l++
	}
	return l
}
