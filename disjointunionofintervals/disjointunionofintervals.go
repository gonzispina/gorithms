package disjointunionofintervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	var res [][]int
	interval := intervals[0]
	for len(intervals) > 0 {
		if intervals[0][0] <= interval[1] {
			if intervals[0][1] > interval[1] {
				interval[1] = intervals[0][1]
			}
		} else {
			res = append(res, interval)
			interval = intervals[0]
		}

		intervals = intervals[1:]
	}

	return append(res, interval)
}
