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
