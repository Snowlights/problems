package __100

import (
	"math"
	"sort"
)

// 75
func sortColors(nums []int) {
	sort.Ints(nums)
}

// 79
func exist(board [][]byte, word string) bool {

	var dfs func(int, int, int) bool
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	m, n := len(board), len(board[0])
	h := make([][]bool, len(board))
	for i := range h {
		h[i] = make([]bool, len(board[0]))
	}
	dfs = func(x, y, i int) bool {
		if board[x][y] != word[i] {
			return false
		}
		if i == len(word)-1 {
			return true
		}
		for _, d := range dir {
			xx, yy := x+d[0], y+d[1]
			if 0 <= xx && xx < m && 0 <= yy && yy < n && !h[xx][yy] {
				h[xx][yy] = true
				if dfs(xx, yy, i+1) {
					return true
				}
				h[xx][yy] = false
			}
		}
		return false
	}

	for i := range board {
		for j := range board[i] {
			h[i][j] = true
			if dfs(i, j, 0) {
				return true
			}
			h[i][j] = false
		}
	}

	return false
}

// 83
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

// 88
func merge(nums1 []int, m int, nums2 []int, n int) {
	ms, ns := 0, 0
	ans := make([]int, 0, m+n)
	for ms < m && ns < n {
		if nums1[ms] > nums2[ns] {
			ans = append(ans, nums2[ns])
			ns++
		} else {
			ans = append(ans, nums1[ms])
			ms++
		}
	}

	for ms < m {
		ans = append(ans, nums1[ms])
		ms++
	}
	for ns < n {
		ans = append(ans, nums2[ns])
		ns++
	}
	for i, v := range ans {
		nums1[i] = v
	}
}

// 91
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

// 94
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

// 97
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[n][m]
}

// 98
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
