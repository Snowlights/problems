package main

func minimumCost(n int, edges, query [][]int) []int {
	fa := make([]int, n)
	and := make([]int, n)
	for i := range fa {
		fa[i] = i
		and[i] = -1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for _, e := range edges {
		x, y := find(e[0]), find(e[1])
		and[y] &= e[2]
		if x != y {
			and[y] &= and[x]
			fa[x] = y
		}
	}

	ans := make([]int, len(query))
	for i, q := range query {
		s, t := q[0], q[1]
		if s == t {
			continue
		}
		if find(s) != find(t) {
			ans[i] = -1
		} else {
			ans[i] = and[find(s)]
		}
	}
	return ans
}
