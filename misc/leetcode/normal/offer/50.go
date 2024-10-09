package offer

import (
	"sort"
	"strings"
)

// offer 50
func pathSum50(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	q, ans := []*TreeNode{root}, 0

	var dfs func(tree *TreeNode, val int)
	dfs = func(tree *TreeNode, val int) {
		if tree == nil {
			return
		}
		val += tree.Val
		if val == targetSum {
			ans++
		}
		dfs(tree.Left, val)
		dfs(tree.Right, val)
	}

	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			dfs(v, 0)
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	return ans
}

// offer 51
func reversePairs(nums []int) int {
	ans := 0
	arr := make([]int, 0)
	arr = append(arr, nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		idx := sort.Search(len(arr), func(j int) bool {
			return arr[j] < nums[i]
		})
		if idx == len(arr) {
			ans += len(arr)
			arr = append(arr, nums[i])
		} else {
			ans += idx + 1
			arr = append(arr[:idx], append([]int{nums[i]}, arr[idx+1:]...)...)
		}
	}

	return ans
}

// offer 55-1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l, r := dfs(node.Left), dfs(node.Right)
		return max(l, r) + 1
	}

	return max(dfs(root.Left), dfs(root.Right)) + 1
}

// offer 55-2
func isBalanced(root *TreeNode) bool {
	flag := true
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l, r := dfs(node.Left), dfs(node.Right)
		if abs(l-r) > 1 {
			flag = false
		}
		return max(l, r) + 1
	}
	dfs(root)
	return flag
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// offer 57
func twoSum(nums []int, target int) []int {

	h := make(map[int]int)
	for i, v := range nums {
		if _, ok := h[target-v]; ok {
			return []int{target - v, v}
		}
		h[v] = i
	}
	return nil
}

// offer 57-2
func findContinuousSequence(target int) [][]int {
	ans := make([][]int, 0)
	for i := 1; i <= target>>1; i++ {
		base, sum, tmp := i, 0, []int{}
		for sum < target {
			sum += base
			tmp = append(tmp, base)
			base++
		}
		if sum == target {
			ans = append(ans, tmp)
		}
	}

	return ans
}

// offer 58
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// offer 58-1
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	parts := strings.Split(s, " ")
	realParts := make([]string, 0, len(parts))
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}
		realParts = append(realParts, p)
	}
	st, e := 0, len(realParts)-1
	for st < e {
		realParts[st], realParts[e] = realParts[e], realParts[st]
		st++
		e--
	}
	return strings.Join(realParts, " ")
}

// 59 -1
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	q := []int{}
	for i := 0; i < k; i++ {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	ans := []int{nums[q[0]]}
	for i := k; i < len(nums); i++ {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		for len(q) > 0 && q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}

	return ans
}

// offer 60
func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	for i := range dp {
		dp[i] = 1.0 / 6.0
	}
	for i := 2; i <= n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j] / 6.0
			}
		}
		dp = tmp
	}
	return dp
}

// offer 62
func lastRemaining(n int, m int) int {
	ans := 0
	// 最后一轮剩下2个人，所以从2开始反推
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

// offer 63
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans, preMin := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-preMin > ans {
			ans = prices[i] - preMin
		}
		if preMin > prices[i] {
			preMin = prices[i]
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// offer 66
func constructArr(a []int) []int {
	zeroCount := 0
	total := 1
	for _, v := range a {
		if v == 0 {
			zeroCount++
			continue
		}
		total *= v
	}
	ans := make([]int, len(a))
	for i, v := range a {
		if v == 0 {
			if zeroCount > 1 {
				ans[i] = 0
			} else {
				ans[i] = total
			}
		} else {
			if zeroCount > 0 {
				ans[i] = 0
			} else {
				ans[i] = total / v
			}
		}
	}
	return ans
}

// offer 68-2
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath, qPath := []*TreeNode{}, []*TreeNode{}
	var bfs func(*TreeNode, []*TreeNode)
	bfs = func(root *TreeNode, path []*TreeNode) {
		if root == nil {
			return
		}
		if root.Val == p.Val {
			arr := make([]*TreeNode, len(path))
			copy(arr, path)
			arr = append(arr, root)
			pPath = arr
		}
		if root.Val == q.Val {
			arr := make([]*TreeNode, len(path))
			copy(arr, path)
			arr = append(arr, root)
			qPath = arr
		}

		bfs(root.Left, append(path, root))
		bfs(root.Right, append(path, root))
	}
	bfs(root, []*TreeNode{})
	pMap := make(map[int]*TreeNode)
	for _, v := range pPath {
		pMap[v.Val] = v
	}
	for i := len(qPath) - 1; i >= 0; i-- {
		if node, ok := pMap[qPath[i].Val]; ok {
			return node
		}
	}
	return nil
}
