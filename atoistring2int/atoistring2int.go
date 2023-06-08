package atoistring2int

import (
	"math"
)

func myAtoi(s string) int {
	negative := false
	start := -1
	i := 0

	for i = 0; i < len(s) && start == -1; i++ {
		switch rune(s[i]) {
		case '-':
			start = i + 1
			negative = true
		case '+':
			start = i + 1
		case '0':
			start = i
		case '1':
			start = i
		case '2':
			start = i
		case '3':
			start = i
		case '4':
			start = i
		case '5':
			start = i
		case '6':
			start = i
		case '7':
			start = i
		case '8':
			start = i
		case '9':
			start = i
			break
		case ' ':
			continue
		default:
			return 0
		}
	}

	if start == -1 {
		return 0
	}

	res := 0
	value := 0
	for i = start; i < len(s); i++ {
		switch rune(s[i]) {
		case '0':
			value = 0
		case '1':
			value = 1
		case '2':
			value = 2
		case '3':
			value = 3
		case '4':
			value = 4
		case '5':
			value = 5
		case '6':
			value = 6
		case '7':
			value = 7
		case '8':
			value = 8
		case '9':
			value = 9
		default:
			if negative {
				res = -res
			}

			return res
		}

		if negative && (-1*res) < (math.MinInt32/10) {
			return math.MinInt32
		} else if !negative && res > (math.MaxInt32/10) {
			return math.MaxInt32
		} else {
			if negative && (-1*res) == (math.MinInt32/10) && value > 8 {
				return math.MinInt32
			} else if !negative && res == (math.MaxInt32/10) && value > 7 {
				return math.MaxInt32
			}

			res = res*10 + value
		}
	}

	if negative {
		res = -res
	}

	return res
}
