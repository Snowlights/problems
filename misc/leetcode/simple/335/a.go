package simple_335

import "sort"

// 1
func passThePillow(n, time int) int {
	t := time % (n - 1)
	if time/(n-1)%2 > 0 {
		return n - t
	}
	return 1 + t
}

// 2
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	sum := []int{}
	for len(q) > 0 {
		tmp, s := q, 0
		q = nil
		for _, node := range tmp {
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		sum = append(sum, s)
	}
	n := len(sum)
	if n < k {
		return -1
	}
	sort.Ints(sum) // 也可以用快速选择
	return int64(sum[n-k])
}

// 3
func findValidSplit(nums []int) int {
	left := map[int]int{}           // left[p] 表示质数 p 首次出现的下标
	right := make([]int, len(nums)) // right[i] 表示左端点为 i 的区间的右端点的最大值
	f := func(p, i int) {
		if l, ok := left[p]; ok {
			right[l] = i // 记录左端点 l 对应的右端点的最大值
		} else {
			left[p] = i // 第一次遇到质数 p
		}
	}

	for i, x := range nums {
		for d := 2; d*d <= x; d++ { // 分解质因数
			if x%d == 0 {
				f(d, i)
				for x /= d; x%d == 0; x /= d {
				}
			}
		}
		if x > 1 {
			f(x, i)
		}
	}

	maxR := 0
	for l, r := range right {
		if l > maxR { // 最远可以遇到 maxR
			return maxR // 也可以写 l-1
		}
		maxR = max(maxR, r)
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 4
func waysToReachTarget(target int, types [][]int) int {
	const mod int = 1e9 + 7
	f := make([]int, target+1)
	f[0] = 1
	for _, p := range types {
		count, marks := p[0], p[1]
		for j := target; j > 0; j-- {
			for k := 1; k <= count && k <= j/marks; k++ {
				f[j] += f[j-k*marks]
			}
			f[j] %= mod
		}
	}
	return f[target]
}
