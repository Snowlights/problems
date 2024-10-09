package _00_700

import (
	"sort"
	"strconv"
)

// 653
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	h, q := make(map[int]bool), []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if h[k-v.Val] {
				return true
			}
			h[v.Val] = true
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

// 655 https://leetcode.cn/problems/print-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	q, height := []*TreeNode{root}, 0
	for ; len(q) > 0; height++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	weight := 1<<(height) - 1
	res := make([][]string, height)
	for i := range res {
		res[i] = make([]string, weight)
	}

	var dfs func(root *TreeNode, r, c int)
	dfs = func(root *TreeNode, r, c int) {
		res[r][c] = strconv.FormatInt(int64(root.Val), 10)
		if root.Left != nil {
			dfs(root.Left, r+1, c-1<<(height-r-2))
		}
		if root.Right != nil {
			dfs(root.Right, r+1, c+1<<(height-r-2))
		}
	}
	dfs(root, 0, (weight-1)/2)

	return res
}

// 658
func findClosestElements(arr []int, k, x int) []int {
	// 稳定排序，在绝对值相同的情况下，保证更小的数排在前面
	sort.SliceStable(arr, func(i, j int) bool { return abs(arr[i]-x) < abs(arr[j]-x) })
	arr = arr[:k]
	sort.Ints(arr)
	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 660
func sumOfPower(nums []int) int {
	s, ans := 0, 0
	const mod int = 1e9 + 7
	sort.Ints(nums)
	for _, v := range nums {
		ans = (ans + v*v%mod*(v+s)) % mod
		s = ((s * 2) + v) % mod
	}
	return ans
}

// 670
func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)
	maxIdx, idx1, idx2 := n-1, -1, -1
	for i := n - 1; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			idx1, idx2 = i, maxIdx
		}
	}
	if idx1 < 0 {
		return num
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]
	v, _ := strconv.Atoi(string(s))
	return v
}

// 673
func findNumberOfLIS(nums []int) (ans int) {
	maxLen := 0
	n := len(nums)
	dp := make([]int, n)
	cnt := make([]int, n)
	for i, x := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, y := range nums[:i] {
			if x > y {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j] // 重置计数
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = cnt[i] // 重置计数
		} else if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return
}
