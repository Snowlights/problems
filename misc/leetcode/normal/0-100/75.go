package __100

import (
	"math"
	"sort"
)

// 75
func sortColors(nums []int) {
	sort.Ints(nums)
}

// 76
func minWindow(s string, t string) string {
	sh, th := make(map[byte]int), make(map[byte]int)
	for _, v := range t {
		th[byte(v)]++
	}

	equals := func(a, b map[byte]int) bool {
		for k, v := range a {
			if b[k] < v {
				return false
			}
		}
		return true
	}

	st, e, ans := 0, 0, ""
	for e < len(s) {
		for e < len(s) && !equals(th, sh) {
			sh[s[e]]++
			e++
		}
		if ans == "" && equals(th, sh) {
			ans = s[st:e]
		}
		for st < e && equals(th, sh) {
			if len(ans) > e-st || (len(ans) == e-st && ans > s[st:e]) {
				ans = s[st:e]
			}
			sh[s[st]]--
			st++
		}
	}
	for st < e && equals(th, sh) {
		if len(ans) > e-st || (len(ans) == e-st && ans > s[st:e]) {
			ans = s[st:e]
		}
		sh[s[st]]--
		st++
	}
	return ans
}

// 78
func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
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

// 81
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return false
}

// 82
func deleteDuplicates82(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	pre := &ListNode{
		Val:  0,
		Next: nil,
	}
	newNode := pre
	for cur != nil && cur.Next != nil {
		found := false
		for cur.Next != nil && cur.Val == cur.Next.Val {
			found = true
			cur.Next = cur.Next.Next
		}
		if found {
			cur = cur.Next
			continue
		}
		pre.Next = cur
		pre = cur
		cur = cur.Next
	}
	pre.Next = cur
	return newNode.Next
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

// 90
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	t := []int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1]
	}
	dfs(false, 0)
	return
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

// 92
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newNode := &ListNode{
		Val:  0,
		Next: head,
	}

	cur, idx := newNode, 1
	var pre, next, reHead, reEnd *ListNode
	for idx < left {
		cur = cur.Next
		idx++
	}
	pre, reHead = cur, cur.Next
	idx++
	cur = cur.Next
	pre.Next = nil
	for idx <= right {
		cur = cur.Next
		idx++
	}
	reEnd, next = cur, cur.Next
	reEnd.Next = nil

	var p, n *ListNode
	c := reHead
	for c != nil {
		n = c.Next
		c.Next = p
		p = c
		c = n
	}

	pre.Next = reEnd
	reHead.Next = next
	return newNode.Next
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

// 95
func generateTrees(n int) []*TreeNode {

	var build func(start, end int) []*TreeNode
	build = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		allTrees := []*TreeNode{}
		// 枚举可行根节点
		for i := start; i <= end; i++ {
			// 获得所有可行的左子树集合
			leftTrees := build(start, i-1)
			// 获得所有可行的右子树集合
			rightTrees := build(i+1, end)
			// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					currTree := &TreeNode{i, nil, nil}
					currTree.Left = left
					currTree.Right = right
					allTrees = append(allTrees, currTree)
				}
			}
		}
		return allTrees
	}

	if n == 0 {
		return nil
	}
	return build(1, n)
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

// 99
func recoverTree(root *TreeNode) {

	arr := make([]*TreeNode, 0)
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr, root)
		dfs(root.Right)
	}
	dfs(root)
	s, e := 0, len(arr)-1
	for s+1 < e {
		if arr[s].Val > arr[s+1].Val {
			break
		}
		s++
	}
	for s < e-1 {
		if arr[e].Val < arr[e-1].Val {
			break
		}
		e--
	}

	arr[s].Val, arr[e].Val = arr[e].Val, arr[s].Val
}
