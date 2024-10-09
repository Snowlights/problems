package lcp

// lcp 52
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNumber(root *TreeNode, ops [][]int) int {
	ans, n := 0, len(ops)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if root == nil {
			return
		}
		for i := n - 1; i >= 0; i-- {
			if root.Val >= ops[i][1] && root.Val <= ops[i][2] {
				ans += ops[i][0]
				break
			}
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	return ans
}

// lcp 56

func conveyorBelt(matrix []string, start []int, end []int) int {
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = m * n
		}
	}
	dp[start[0]][start[1]] = 0
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	row := []byte{'<', '>', 'v', '^'}
	q := [][]int{start}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for i, d := range dir {
				x, y := v[0]+d[0], v[1]+d[1]
				if 0 <= x && x < n && 0 <= y && y < m {
					if matrix[v[0]][v[1]] == row[i] {
						if dp[x][y] > dp[v[0]][v[1]] {
							dp[x][y] = dp[v[0]][v[1]]
							q = append(q, []int{x, y})
						}
					} else {
						if dp[x][y] > dp[v[0]][v[1]]+1 {
							dp[x][y] = dp[v[0]][v[1]] + 1
							q = append(q, []int{x, y})
						}
					}
				}
			}
		}
	}

	return dp[end[0]][end[1]]
}
