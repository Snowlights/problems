package offer

import "strconv"

// 25
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p := &ListNode{}
	cur := p
	l1, l2 := list1, list2
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		} else {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
	}
	for l1 != nil {
		cur.Next = l1
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		cur.Next = l2
		cur = cur.Next
		l2 = l2.Next
	}

	return p.Next
}

// 26
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	q := []*TreeNode{A}
	for len(q) > 0 {
		t := q
		q = nil
		for _, v := range t {
			if equal(v, B) {
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

func equal(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return true
	}
	return p.Val == q.Val && equal(p.Left, q.Left) && equal(p.Right, q.Right)
}

// 27
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			v.Left, v.Right = v.Right, v.Left
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	return root
}

// 28
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// offer 30
type MinStack struct {
	list, stack []int
}

/** initialize your data structure here. */
func Constructor30() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) > 0 {
		val := this.stack[len(this.list)-1]
		if val < x {
			this.stack = append(this.stack, val)
		} else {
			this.stack = append(this.stack, x)
		}
	} else {
		this.stack = append(this.stack, x)
	}

	this.list = append(this.list, x)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) Min() int {
	return this.stack[len(this.list)-1]
}

// offer 32
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp, val := q, make([]int, len(q))
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
		ans = append(ans, val...)
	}
	return ans
}

// offer 34
func pathSum(root *TreeNode, target int) [][]int {

	var dfs func(node *TreeNode, val int, path []int)
	ans := make([][]int, 0)
	dfs = func(node *TreeNode, val int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		val += node.Val
		if val == target && node.Left == nil && node.Right == nil {
			arr := make([]int, len(path))
			copy(arr, path)
			ans = append(ans, arr)
		}
		dfs(node.Left, val, path)
		dfs(node.Right, val, path)
	}
	dfs(root, 0, []int{})
	return ans
}

// offer 35
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		nodeMap[cur] = &Node{
			Val: cur.Val,
		}
	}

	for cur := head; cur != nil; cur = cur.Next {
		n, ok := nodeMap[cur]
		if ok {
			n.Random = nodeMap[cur.Random]
			n.Next = nodeMap[cur.Next]
		}
	}
	return nodeMap[head]
}

// offer 42
func maxSubArray(nums []int) int {
	ans, tmp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if tmp < 0 && tmp < nums[i] {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}
		ans = max(tmp, ans)
	}
	return ans
}

// offer 46
func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}

// offer 47
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// offer 48
func lengthOfLongestSubstring(s string) int {
	h := make(map[byte]bool)
	l, r, ans := 0, 0, 0
	for l <= r && r < len(s) {
		for r < len(s) && !h[s[r]] {
			h[s[r]] = true
			r++
		}
		ans = max(ans, r-l)
		h[s[l]] = false
		l++
	}
	ans = max(ans, r-l)
	return ans
}
