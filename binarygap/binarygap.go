package binarygap

import "strconv"

// Binary gap counts the max amount of zeroes between two ones in an int's binary representation
// Solution in O(N) where N is the length of the binary representation
func BinaryGap(n int) int {
	buf := strconv.FormatInt(int64(n), 2)
	currentLength := 0
	maxLength := 0
	y := 0
	
	if n == 0 {
		return 0
	}

	for x := 0; x < len(buf); x++ {
		// x moves only over 1s, unless the number starts with zeroes
		if buf[x] == 48 {
			continue
		}

		closedGap := false
		for y = x + 1; y < len(buf); y++ {
			if buf[y] == 49 {
				closedGap = true
				break
			}

			currentLength++
		}

		if closedGap && currentLength >= maxLength {
			maxLength = currentLength
		}

		currentLength = 0
		x = y - 1// Avoid going over 0s
	}

	return maxLength
}