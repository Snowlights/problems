package _600_1700

import "sort"

// 1697
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	ids := make([]int, len(queries))
	for i := range ids {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return queries[ids[i]][2] < queries[ids[j]][2]
	})
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	ans, j := make([]bool, len(queries)), 0
	for _, id := range ids {
		val := queries[id]
		q := val[2]
		for ; j < len(edgeList) && edgeList[j][2] < q; j++ {
			merge(edgeList[j][0], edgeList[j][1])
		}
		ans[id] = find(val[0]) == find(val[1])
	}
	return ans
}
