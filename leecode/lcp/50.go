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
