package _500_1600

import "sort"

// 1584
func minCostConnectPoints(points [][]int) int {
	type pair struct {
		from, to, val int
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	val := func(i, j int) int {
		a, b := points[i], points[j]
		return abs(a[0]-b[0]) + abs(a[1]-b[1])
	}
	pairList := make([]pair, 0, len(points))
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			pairList = append(pairList, pair{
				from: i,
				to:   j,
				val:  val(i, j),
			})
		}
	}
	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].val < pairList[j].val
	})

	fa := make([]int, len(points))
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			fa[from] = find(fa[from])
		}
		return fa[from]
	}
	same := func(from, to int) bool {
		return find(from) == find(to)
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	ans := 0
	for _, pair := range pairList {
		if !same(pair.from, pair.to) {
			merge(pair.from, pair.to)
			ans += pair.val
		}
	}

	return ans
}
