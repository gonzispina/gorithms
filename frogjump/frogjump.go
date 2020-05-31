package frogjump

import "math"

// FrogJump determines how many jumps a frog has to do to get from X to Y if the frog cover d distance in one jump
// Solution on O(K)
func FrogJump(x int, y int, d int) int {
	if d == 0 {
		return math.MaxInt64
	}

	dist := y - x
	if dist < 0 {
		dist = int(math.Abs(float64(dist)))
	}

	jumps := dist/d
	rem := dist % d

	if rem != 0 {
		return jumps + 1
	}

	return jumps
}


