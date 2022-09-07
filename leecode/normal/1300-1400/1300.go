package _300_1400

// 1314
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pre[i][j] = mat[i-1][j-1] + pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1]
		}
	}

	ans := make([][]int, m)
	for i := range mat {
		ans[i] = make([]int, n)
		for j := range mat[i] {
			left, right := max(0, j-k), min(n, j+k+1)
			top, button := max(0, i-k), min(m, i+k+1)
			ans[i][j] = pre[button][right] - pre[button][left] - pre[top][right] + pre[top][left]
		}
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 1319
func makeConnected(n int, connections [][]int) (ans int) {
	if len(connections) < n-1 {
		return -1
	}

	graph := make([][]int, n)
	for _, c := range connections {
		v, w := c[0], c[1]
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}

	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for _, to := range graph[from] {
			if !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return ans - 1
}
