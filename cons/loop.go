package cons

// MakeLoop creates a linked list of length n, where the last element points at the first (returned) element.
func MakeLoop(n uint, s string) *Cell {
	last := NewCell(s, nil)
	first := makeFrom(n-1, s, last)
	last.next = first
	return first
}
