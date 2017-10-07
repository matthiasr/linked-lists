package cons

// MakeLoop creates a linked list of length n, where the last element points at the first (returned) element.
func MakeLoop(n uint, s string) *Cell {
	last := NewCell(s, nil)
	first := makeFrom(n-1, s, last)
	last.next = first
	return first
}

// MakeWithLoop creates a straight linked list
// of length n with a loop of length m attached.
// In other words, the Cdr of element n+m is the
// element n+1.
func MakeWithLoop(n, m uint, s string) *Cell {
	return makeFrom(n, s, MakeLoop(m, s))
}
