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
