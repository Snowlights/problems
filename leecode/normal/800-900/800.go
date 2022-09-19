package _00_900

// 802
func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	rg := make([][]int, n)
	inDeg := make([]int, n)
	for x, ys := range graph {
		for _, y := range ys {
			rg[y] = append(rg[y], x)
		}
		inDeg[x] = len(ys)
	}

	q := []int{}
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		for _, x := range rg[y] {
			inDeg[x]--
			if inDeg[x] == 0 {
				q = append(q, x)
			}
		}
	}

	for i, d := range inDeg {
		if d == 0 {
			ans = append(ans, i)
		}
	}
	return
}

// 807
func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rol, col := make([]int, n), make([]int, n)

	for i := range grid {
		for j, v := range grid[i] {
			rol[i] = max(rol[i], v)
			col[j] = max(col[j], v)
		}
	}

	ans := 0
	for i := range grid {
		for j, v := range grid[i] {
			ans += min(rol[i], col[j]) - v
		}
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 815
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	nodeToRoute := make(map[int][]int)
	for i, r := range routes {
		for _, v := range r {
			nodeToRoute[v] = append(nodeToRoute[v], i)
		}
	}

	q := nodeToRoute[source]
	tmp, vis := q, make(map[int]bool)
	q = nil
	for _, v := range tmp {
		for _, vv := range routes[v] {
			q = append(q, vv)
		}
		vis[v] = true
	}
	ans := 1
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == target {
				return ans
			}
			for _, route := range nodeToRoute[v] {
				if vis[route] {
					continue
				}
				vis[route] = true
				for _, r := range routes[route] {
					q = append(q, r)
				}
			}
		}
		ans++
	}

	return -1
}
