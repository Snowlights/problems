package _00_900

import (
	"math"
	"sort"
)

// 801
func minSwap(nums1, nums2 []int) int {
	n := len(nums1)
	// f[i][0/1] 表示让 nums1 和 nums2 的前 i 个数严格递增所需操作的最小次数
	// 其中 f[i][0] 不交换 nums1[i] 和 nums2[i]，f[i][1] 交换 nums1[i] 和 nums2[i]

	f := make([][2]int, n)
	f[0][1] = 1
	for i := 1; i < n; i++ {
		f[i] = [2]int{n, n} // 答案不会超过 n，故初始化成 n 方便后面取 min
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			f[i][0] = f[i-1][0]
			f[i][1] = f[i-1][1] + 1
		}
		if nums2[i-1] < nums1[i] && nums1[i-1] < nums2[i] {
			f[i][0] = min(f[i][0], f[i-1][1])
			f[i][1] = min(f[i][1], f[i-1][0]+1)
		}
	}
	return min(f[n-1][0], f[n-1][1])
}

// 802
func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	rg := make([][]int, n)
	inDeg := make([]int, n)
	for x, ys := range graph {
		for _, y := range ys {
			rg[y] = append(rg[y], x)
		}
		inDeg[x] = len(ys)
	}

	q := []int{}
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		for _, x := range rg[y] {
			inDeg[x]--
			if inDeg[x] == 0 {
				q = append(q, x)
			}
		}
	}

	for i, d := range inDeg {
		if d == 0 {
			ans = append(ans, i)
		}
	}
	return
}

// 807
func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rol, col := make([]int, n), make([]int, n)

	for i := range grid {
		for j, v := range grid[i] {
			rol[i] = max(rol[i], v)
			col[j] = max(col[j], v)
		}
	}

	ans := 0
	for i := range grid {
		for j, v := range grid[i] {
			ans += min(rol[i], col[j]) - v
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

// 815
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	nodeToRoute := make(map[int][]int)
	for i, r := range routes {
		for _, v := range r {
			nodeToRoute[v] = append(nodeToRoute[v], i)
		}
	}

	q := nodeToRoute[source]
	tmp, vis := q, make(map[int]bool)
	q = nil
	for _, v := range tmp {
		for _, vv := range routes[v] {
			q = append(q, vv)
		}
		vis[v] = true
	}
	ans := 1
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == target {
				return ans
			}
			for _, route := range nodeToRoute[v] {
				if vis[route] {
					continue
				}
				vis[route] = true
				for _, r := range routes[route] {
					q = append(q, r)
				}
			}
		}
		ans++
	}

	return -1
}

// 817
func numComponents(head *ListNode, nums []int) int {
	h := make(map[int]bool)
	for _, v := range nums {
		h[v] = true
	}

	ans, flag, cur := 0, false, head
	for cur != nil {
		if h[cur.Val] {
			if !flag {
				ans++
				flag = true
			}
		} else {
			flag = false
		}
		cur = cur.Next
	}
	return ans
}

// 822
func flipgame(fronts []int, backs []int) int {
	forbidden := make(map[int]bool)
	for i, v := range fronts {
		if v == backs[i] {
			forbidden[v] = true
		}
	}
	ans := math.MaxInt32
	for i, v := range fronts {
		if !forbidden[v] {
			ans = min(ans, v)
		}
		if !forbidden[backs[i]] {
			ans = min(ans, backs[i])
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 823
func numFactoredBinaryTrees(arr []int) int {
	h, ans, mod := make(map[int]int), 0, int(1e9+7)
	sort.Ints(arr)
	for i, v := range arr {
		h[v] = 1
		for _, j := range arr[:i] {
			if v%j == 0 {
				h[v] += h[j] * h[v/j]
				h[v] %= mod
			}
		}
		ans += h[v]
	}

	return ans % mod
}
