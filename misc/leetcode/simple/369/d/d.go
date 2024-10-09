package main

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][14]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, fa int) int {
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res1 := coins[i]>>j - k
		res2 := coins[i] >> (j + 1)
		for _, ch := range g[i] {
			if ch != fa {
				res1 += dfs(ch, j, i) // 不右移
				if j < 13 {           // j >= 14 相当于 res2 += 0 无需递归
					res2 += dfs(ch, j+1, i) // 右移
				}
			}
		}
		*p = max(res1, res2)
		return *p
	}
	return dfs(0, 0, -1)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
