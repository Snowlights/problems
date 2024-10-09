package _00_200

import (
	"fmt"
	"math"
)

// 100
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 102
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q, idx := []*TreeNode{root}, 0
	for len(q) > 0 {
		tmp, val := q, make([]int, 0, len(q))
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			val = append(val, v.Val)
		}
		if idx%2 == 1 {
			s, e := 0, len(val)-1
			for s < e {
				val[s], val[e] = val[e], val[s]
				s++
				e--
			}
		}
		ans = append(ans, val)
		idx++
	}
	return ans
}

// 103
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	q, idx := []*TreeNode{root}, 0
	for len(q) > 0 {
		tmp := q
		q = nil
		res := make([]int, 0, len(tmp))
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			res = append(res, v.Val)
		}

		if idx%2 == 1 {
			s, e := 0, len(res)-1
			for s < e {
				res[s], res[e] = res[e], res[s]
				s++
				e--
			}
		}
		ans = append(ans, res)
		idx++
	}

	return ans
}

// 104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 105
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	inorderIdx := 0
	for i, v := range inorder {
		if v == root.Val {
			inorderIdx = i
			break
		}
	}

	root.Left = buildTree(preorder[1:inorderIdx+1], inorder[:inorderIdx])
	root.Right = buildTree(preorder[inorderIdx+1:], inorder[inorderIdx+1:])
	return root
}

// 108
func sortedArrayToBST(nums []int) *TreeNode {

	var buildbst func(start, end int) *TreeNode
	buildbst = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		node := &TreeNode{Val: nums[mid]}
		node.Left = buildbst(start, mid-1)
		node.Right = buildbst(mid+1, end)
		return node
	}

	return buildbst(0, len(nums)-1)
}

// 110
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var h func(root *TreeNode) int
	flag := true
	h = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := h(root.Left), h(root.Right)
		if abs(l-r) > 1 {
			flag = false
		}
		return max(l, r) + 1
	}
	h(root)
	return flag
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 112

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				v.Left.Val += v.Val
				q = append(q, v.Left)
			}
			if v.Right != nil {
				v.Right.Val += v.Val
				q = append(q, v.Right)
			}
			if v.Left == nil && v.Right == nil && v.Val == targetSum {
				return true
			}
		}
	}
	return false
}

// 113
func pathSum(root *TreeNode, targetSum int) [][]int {

	ans := [][]int{}
	if root == nil {
		return ans
	}
	var dfs func(node *TreeNode, path []int, sum int)
	dfs = func(node *TreeNode, path []int, sum int) {
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				res := make([]int, len(path))
				copy(res, path)
				res = append(res, node.Val)
				ans = append(ans, res)
			}
			return
		}

		if node.Left != nil {
			dfs(node.Left, append(path, node.Val), sum)
		}
		if node.Right != nil {
			dfs(node.Right, append(path, node.Val), sum)
		}
	}
	dfs(root, []int{}, 0)
	return ans
}

// 120
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		for j := range triangle[i] {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
				continue
			}
			if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}
	fmt.Println(dp)
	mm := dp[m-1][0]
	for _, v := range dp[m-1] {
		mm = min(mm, v)
	}
	return mm
}

// 122
func maxProfit(prices []int) int {

	i, ans := 0, 0
	buy := 0
	for i < len(prices) {
		for i+1 < len(prices) && prices[i] > prices[i+1] {
			i++
		}
		if i == len(prices)-1 {
			break
		}
		buy = i
		for i+1 < len(prices) && prices[i] < prices[i+1] {
			i++
		}
		if i == len(prices) {
			i = len(prices) - 1
		}
		ans += prices[i] - prices[buy]
		i++
	}
	return ans
}

// 123
func maxProfit123(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

// 124
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
