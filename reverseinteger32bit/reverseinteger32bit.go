package reverseinteger32bit

import (
	"strconv"
)

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
	}

	s := []byte(strconv.Itoa(x))
	if negative {
		s = s[1:]
	}

	for i := 0; i < len(s)/2; i++ {
		aux := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = aux
	}

	if negative {
		s = append([]byte{"-"[0]}, s...)
	}

	v, err := strconv.ParseInt(string(s), 10, 32)
	if err != nil {
		print(err.Error())
		return 0
	}

	return int(v)
}
