package offer

import (
	"sort"
	"strconv"
)

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

// offer 28
type TNode struct {
	Val   int
	Prev  *TNode
	Next  *TNode
	Child *TNode
}

func flatten(root *TNode) *TNode {
	if root == nil {
		return root
	}
	q := &TNode{}
	cur := q
	var dfs func(node *TNode)
	dfs = func(node *TNode) {
		for node != nil {
			cur.Next = node
			node.Prev = cur
			cur = cur.Next

			tmp := node.Next
			if node.Child != nil {
				dfs(node.Child)
			}

			node.Child = nil
			node = tmp
		}
	}
	dfs(root)
	q.Next.Prev = nil
	return q.Next
}

// offer 29
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)

	xS, xE, yS, yE := 0, m-1, 0, n-1
	for xS <= xE && yS <= yE {
		s, e := yS, yE
		for i := s; i <= e; i++ {
			ans = append(ans, matrix[xS][i])
		}
		s, e = xS+1, xE
		for i := s; i <= e; i++ {
			ans = append(ans, matrix[i][yE])
		}
		if xS < xE && yS < yE {
			s, e = yE-1, yS
			for i := s; i >= e; i-- {
				ans = append(ans, matrix[xE][i])
			}
			s, e = xE-1, xS+1
			for i := s; i >= e; i-- {
				ans = append(ans, matrix[i][xS])
			}
		}
		xS++
		xE--
		yS++
		yE--
	}

	return ans
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

// offer 31
func validateStackSequences(pushed, popped []int) bool {
	st := []int{}
	j := 0
	for _, x := range pushed {
		st = append(st, x)
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
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

// offer 33
func verifyPostorder(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}

	index := len(postorder) - 1   // 区分左右子树：左子树上的值全都比根节点小，右子树上的值全都比根节点大
	rootValue := postorder[index] // 用来记录根节点的值

	for k, v := range postorder {
		// 当出现第一个大于根节点的值时，这个值往后全是右子树
		if index == len(postorder)-1 && v > rootValue {
			index = k
		}
		// 在右子树中出现小于根节点的值时，则该树不是二叉搜索树
		if index != len(postorder)-1 && rootValue > v {
			return false
		}
	}
	return verifyPostorder(postorder[:index]) && verifyPostorder(postorder[index:len(postorder)-1])
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

// offer 38
func permutation(s string) []string {

	var bfs func(int, []byte)
	vis := make(map[int]bool)
	exit := make(map[string]bool)
	ans := make([]string, 0)
	bfs = func(idx int, path []byte) {
		if idx == len(s) {
			if exit[string(path)] {
				return
			}
			exit[string(path)] = true
			ans = append(ans, string(path))
			return
		}
		for i, v := range s {
			if vis[i] {
				continue
			}
			vis[i] = true
			bfs(idx+1, append(path, byte(v)))
			vis[i] = false
		}
	}

	for i := range s {
		vis[i] = true
		bfs(1, append([]byte{}, byte(s[i])))
		vis[i] = false
	}

	return ans
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

// offer 43
func countDigitOne(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int, bool) int
	f = func(i, cnt1 int, isLimit bool) (res int) {
		if i == m {
			return cnt1
		}
		if !isLimit {
			dv := &dp[i][cnt1]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for d := 0; d <= up; d++ { // 枚举要填入的数字 d
			c := cnt1
			if d == 1 {
				c++
			}
			res += f(i+1, c, isLimit && d == up)
		}
		return
	}
	return f(0, 0, true)
}

// offer 44
func findNthDigit(n int) int {
	digit, digitNum, count := 1, 1, 9
	for n > count {
		n -= count
		digit++
		digitNum *= 10
		count = 9 * digit * digitNum
	}
	num := digitNum + (n-1)/digit

	index := (n - 1) % digit

	numStr := strconv.Itoa(num)
	return int(numStr[index] - '0')
}

// offer 45
func minNumber(nums []int) string {
	//将整数数组按字符串形式排序
	sort.Slice(nums, func(i, j int) bool {
		a, b := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return a+b < b+a
	})

	res := ""
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}
	return res
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

// offer 49
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}
