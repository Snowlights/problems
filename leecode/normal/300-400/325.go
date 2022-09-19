package _00_400

import "math"

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
