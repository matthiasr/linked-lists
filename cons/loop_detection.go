package cons

// DetectLoopSimple checks for loops by walking
// the whole list again at each item. It is
// O(n²) but easy to see that it is correct.
func (c *Cell) DetectLoopSimple() bool {
	first := c
	pos := uint(0)
	for !c.Nil() {
		for i, cc := uint(0), first; i < pos; i++ {
			if cc == c {
				return true
			}
			cc = cc.Cdr()
		}
		c = c.Cdr()
		pos++
	}
	return false
}

// DetectLoop2 checks for loops by comparing the
// current element with a marker that is moved
// along the list at half speed. If there is a
// loop, the marker will eventually reach it,
// and some time later the current item will
// overtake it. This is O(n) but with some
// constant factor that deopends on the shape of
// the list and loop.
func (c *Cell) DetectLoop2() bool {
	marker := c
	pos := uint(0)
	for !c.Nil() {
		pos++
		// Move to the next item before checking,
		// otherwise it will always be the same as
		// the marker for the first element.
		c = c.Cdr()
		if c == marker {
			return true
		}
		if pos&1 == 0 {
			marker = marker.Cdr()
		}
	}
	return false
}
