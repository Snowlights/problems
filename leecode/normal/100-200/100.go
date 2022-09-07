package _00_200

import "fmt"

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

// 104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
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
