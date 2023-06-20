package exponentials

func pow(a, b int) int {
	if b == 0 {
		return 1
	}

	if a == 0 || a == 1 || b == 1 {
		return a
	}

	if b == 2 {
		return a * a
	}

	r := b % 2
	v := pow(a, b>>1)

	if r == 1 {
		return a * v * v
	}

	return v * v
}
