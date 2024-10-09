package _600_1700

import (
	"math"
	"math/bits"
	"sort"
)

// 1681
func minimumIncompatibility(nums []int, k int) int {
	h := make(map[int]int)
	for _, v := range nums {
		h[v]++
		if h[v] > k {
			return -1
		}
	}

	n := len(nums)
	sort.Ints(nums)
	size := n / k
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(left, pre int) int
	dfs = func(left, pre int) (res int) {
		if left == 0 {
			return 0
		}

		dv := &dp[left][pre]
		if *dv >= 0 {
			return *dv
		}
		defer func() {
			*dv = res
		}()
		res = math.MaxInt32
		if bits.OnesCount(uint(left))%size == 0 {
			lb := left & -left
			res = dfs(left^lb, bits.Len(uint(lb))-1)
			return
		}
		last := nums[pre]
		for i := pre + 1; i < n; i++ {
			if left>>i&1 == 1 && nums[i] != last {
				last = nums[i]
				res = min(res, last-nums[pre]+dfs(left^(1<<i), i))
			}
		}
		return
	}
	return dfs((1<<n)-2, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 1697
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	ids := make([]int, len(queries))
	for i := range ids {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return queries[ids[i]][2] < queries[ids[j]][2]
	})
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	ans, j := make([]bool, len(queries)), 0
	for _, id := range ids {
		val := queries[id]
		q := val[2]
		for ; j < len(edgeList) && edgeList[j][2] < q; j++ {
			merge(edgeList[j][0], edgeList[j][1])
		}
		ans[id] = find(val[0]) == find(val[1])
	}
	return ans
}
