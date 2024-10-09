package _00_400

import (
	"math"
	"sort"
)

// 327
func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)

	// 计算前缀和 preSum，以及后面统计时会用到的所有数字 allNums
	allNums := make([]int, 1, 3*n+1)
	preSum := make([]int, n+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
		allNums = append(allNums, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
	}

	// 将 allNums 离散化
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i <= 3*n; i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}

	// 遍历 preSum，利用树状数组计算每个前缀和对应的合法区间数
	t := newFenwickTree(int64(k), nil)
	t.add(int64(kth[0]), 1)
	ans := 0
	for _, sum := range preSum[1:] {
		left, right := kth[sum-upper], kth[sum-lower]
		ans += int(t.query(int64(left), int64(right)))
		t.add(int64(kth[sum]), 1)
	}
	return ans

}

// 328

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// 329
func longestIncreasingPath(matrix [][]int) int {

	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	ans := 0
	m, n := len(matrix), len(matrix[0])

	var dfs func(int, int) int
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	dfs = func(i, j int) int {
		if f[i][j] != 0 {
			return f[i][j]
		}
		res := 1
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] > matrix[i][j] {
				res = max(res, dfs(x, y)+1)
			}
		}

		f[i][j] = res
		return res
	}
	for i := range matrix {
		for j := range matrix[i] {
			dfs(i, j)
		}
	}
	for _, v := range f {
		for _, vv := range v {
			ans = max(ans, vv)
		}
	}
	return ans
}

// 334
func increasingTriplet(nums []int) bool {
	s := []int{}
	for i, v := range nums {
		for len(s) > 0 && nums[s[len(s)-1]] > v {
			s = s[:len(s)-1]
		}
		s = append(s, i)
		if len(s) > 2 {
			return true
		}
	}
	return false
}

// 337
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		l, r := dfs(node.Left), dfs(node.Right)
		return []int{node.Val + l[1] + r[1], max(l[0], l[1]) + max(r[0], r[1])}
	}

	val := dfs(root)
	return max(val[0], val[1])
}

// 343
func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	q, m := n/3, n%3
	if m == 0 {
		return int(math.Pow(3, float64(q)))
	} else if m == 1 {
		return int(math.Pow(3, float64(q-1))) * 4
	} else {
		return int(math.Pow(3, float64(q))) * 2
	}
}

// 349
func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]bool)
	for _, v := range nums1 {
		hash[v] = true
	}
	res, exit := make([]int, 0), make(map[int]bool)
	for _, v := range nums2 {
		if !exit[v] && hash[v] {
			res = append(res, v)
			exit[v] = true
		}
	}

	return res
}
