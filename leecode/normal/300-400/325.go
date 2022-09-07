package _00_400

// 329
func longestIncreasingPath(matrix [][]int) int {

	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	ans := 0
	m, n := len(matrix), len(matrix[0])

	var dfs func(int, int) int
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	dfs = func(i, j int) int {
		if f[i][j] != 0 {
			return f[i][j]
		}
		res := 1
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] > matrix[i][j] {
				res = max(res, dfs(x, y)+1)
			}
		}

		f[i][j] = res
		return res
	}
	for i := range matrix {
		for j := range matrix[i] {
			dfs(i, j)
		}
	}
	for _, v := range f {
		for _, vv := range v {
			ans = max(ans, vv)
		}
	}
	return ans
}
