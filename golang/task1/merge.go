package task1

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	cur := intervals[0]

	for i := 1; i <= len(intervals); i++ {
		next := intervals[i]
		if next[0] <= cur[1] {
			cur[1] = max(cur[1], next[1])
		} else {
			res = append(res, cur)
			cur = next
		}
	}
	res = append(res, cur)
	return res
}
