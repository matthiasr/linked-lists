package cons

// Make creates a linked list of the specified
// length where all cells have the same content.
func Make(n uint, s string) *Cell {
	return makeFrom(n, s, nil)
}

func makeFrom(n uint, s string, c *Cell) *Cell {
	for i := uint(0); i < n; i++ {
		c = NewCell(s, c)
	}
	return c
}
