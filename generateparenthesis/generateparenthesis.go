package generateparenthesis

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	if n == 1 {
		return []string{"()"}
	}

	type node struct {
		value  string
		open   int
		closed int
	}

	nodes := []node{{
		value:  "(",
		open:   1,
		closed: 0,
	}}

	opened := 1
	for opened > 0 {
		current := nodes[0]
		nodes = nodes[1:]

		if current.closed == n {
			nodes = append(nodes, current)
			continue
		}

		opened -= 1
		if current.open > 0 {
			nodes = append(nodes, node{
				open:   current.open - 1,
				value:  fmt.Sprintf("%s)", current.value),
				closed: current.closed + 1,
			})

			if current.closed+1 < n {
				opened += 1
			}
		}

		if current.open+current.closed < n {
			nodes = append(nodes, node{
				open:   current.open + 1,
				value:  fmt.Sprintf("%s(", current.value),
				closed: current.closed,
			})

			opened += 1
		}
	}

	res := make([]string, len(nodes))

	for i, no := range nodes {
		res[i] = no.value
	}

	return res
}
