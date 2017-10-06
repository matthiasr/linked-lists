package cons

func rec(n uint64) error {
	if n == 0 {
		return nil
	}
	return rec(n - 1)
}
