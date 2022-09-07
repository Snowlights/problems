package _300_1400

// 1351
func countNegatives(grid [][]int) int {
	ans := 0
	for _, row := range grid {
		for _, value := range row {
			if value < 0 {
				ans++
			}
		}
	}
	return ans
}

// 1367
type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(n *ListNode, root *TreeNode) bool
	dfs = func(n *ListNode, root *TreeNode) bool {
		if n == nil && root == nil {
			return true
		}
		if n == nil {
			return true
		}
		if root == nil {
			return false
		}
		if n.Val != root.Val {
			return false
		}
		return dfs(n.Next, root.Left) || dfs(n.Next, root.Right)
	}

	if (root == nil && head != nil) || (root != nil && head == nil) {
		return false
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			cur := head
			if dfs(cur, v) {
				return true
			}
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	return false
}
