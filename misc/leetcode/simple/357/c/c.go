package main

func maximumSafenessFactor(g [][]int) (ans int) {
	n, q := len(g), [][]int{}
	dis := make([][]int, n)
	for i, v := range g {
		dis[i] = make([]int, n)
		for j, vv := range v {
			dis[i][j] = -1
			if vv == 1 {
				q = append(q, []int{i, j})
				dis[i][j] = 0
			}
		}
	}

	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	group := [][][]int{q}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, d := range dir {
				x, y := v[0]+d[0], v[1]+d[1]
				if 0 <= x && x < n && 0 <= y && y < n && dis[x][y] < 0 {
					q = append(q, []int{x, y})
					dis[x][y] = len(group)
				}
			}
		}
		group = append(group, q)
	}

	fa := make([]int, n*n)
	for i := range fa {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	for i := len(group) - 2; i >= 0; i-- {
		for _, v := range group[i] {
			for _, d := range dir {
				x, y := v[0]+d[0], v[1]+d[1]
				if 0 <= x && x < n && 0 <= y && y < n && dis[x][y] >= dis[v[0]][v[1]] {
					merge(x*n+y, v[0]*n+v[1])
				}
			}
		}
		if find(0) == find(n*n-1) {
			return i
		}
	}

	return
}
