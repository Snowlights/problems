package interview

// 面试题 04-01
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	g := make([][]int, n)
	vis := make(map[int]bool)

	for _, v := range graph {
		g[v[0]] = append(g[v[0]], v[1])
	}
	q := []int{start}
	vis[start] = true
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == target {
				return true
			}
			for _, to := range g[v] {
				if !vis[to] {
					q = append(q, to)
				}
			}
		}
	}
	return false
}

// 面试题 04-02
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := len(nums) / 2
	root := &TreeNode{Val: nums[idx]}
	root.Left = sortedArrayToBST(nums[:idx])
	root.Right = sortedArrayToBST(nums[idx+1:])
	return root
}

// 面试题 04-03
func listOfDepth(tree *TreeNode) []*ListNode {
	ans := make([]*ListNode, 0)
	q := []*TreeNode{tree}
	for len(q) > 0 {
		tmp := q
		q = nil
		root := &ListNode{}
		cur := root
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			cur.Next = &ListNode{
				Val:  v.Val,
				Next: nil,
			}
			cur = cur.Next
		}
		ans = append(ans, root.Next)
	}
	return ans
}

// 面试题 04-04
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var help func(node *TreeNode) int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	help = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := help(root.Left), help(root.Right)
		return max(left, right) + 1
	}
	return abs(help(root.Left)-help(root.Right)) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}
