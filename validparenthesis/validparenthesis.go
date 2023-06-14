package validparenthesis

func counterPart(s byte) byte {
	if s == ')' {
		return '('
	} else if s == ']' {
		return '['
	} else { // if s == '{'
		return '{'
	}
}

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	open := make([]byte, len(s)/2)
	ptr := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			ptr++
			if ptr >= len(s)/2 {
				return false
			}
			open[ptr] = s[i]
			continue
		}

		if ptr > -1 && open[ptr] == counterPart(s[i]) {
			ptr--
			continue
		}

		return false
	}

	return ptr == -1
}
