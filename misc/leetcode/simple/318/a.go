package simple_318

import (
	"container/heap"
	"sort"
)

// 1
func applyOperations(a []int) []int {
	n := len(a)
	b := a[:0]
	for i := 0; i < n-1; i++ {
		if a[i] > 0 {
			if a[i] == a[i+1] {
				a[i] *= 2
				a[i+1] = 0
			}
			b = append(b, a[i]) // 非零数字排在前面
		}
	}
	if a[n-1] > 0 {
		b = append(b, a[n-1])
	}
	for i := len(b); i < n; i++ {
		a[i] = 0
	}
	return a
}

// 2
func maximumSubarraySum(nums []int, k int) int64 {
	ans, sum := 0, 0
	cnt := map[int]int{}
	for _, x := range nums[:k-1] {
		cnt[x]++
		sum += x
	}
	for i := k - 1; i < len(nums); i++ {
		cnt[nums[i]]++ // 移入元素
		sum += nums[i]
		if len(cnt) == k && sum > ans {
			ans = sum
		}
		x := nums[i+1-k]
		cnt[x]-- // 移出元素
		if cnt[x] == 0 {
			delete(cnt, x) // 重要：及时移除个数为 0 的数据
		}
		sum -= x
	}
	return int64(ans)
}

// 3
func totalCost(costs []int, k, candidates int) int64 {
	ans := 0
	if n := len(costs); candidates*2 < n {
		pre := hp{costs[:candidates]}
		heap.Init(&pre) // 原地建堆
		suf := hp{costs[n-candidates:]}
		heap.Init(&suf)
		for i, j := candidates, n-1-candidates; k > 0 && i <= j; k-- {
			if pre.IntSlice[0] <= suf.IntSlice[0] {
				ans += pre.IntSlice[0]
				pre.IntSlice[0] = costs[i]
				heap.Fix(&pre, 0)
				i++
			} else {
				ans += suf.IntSlice[0]
				suf.IntSlice[0] = costs[j]
				heap.Fix(&suf, 0)
				j--
			}
		}
		costs = append(pre.IntSlice, suf.IntSlice...)
	}
	sort.Ints(costs)
	for _, c := range costs[:k] { // 也可以用快速选择算法求前 k 小
		ans += c
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (hp) Push(interface{})     {} // 没有用到，留空即可
func (hp) Pop() (_ interface{}) { return }

// 4

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	sort.Ints(robot)
	n, m := len(factory), len(robot)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, j int) (res int) {
		if j >= m {
			return
		}
		dv := &dp[i][j]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		if i == n-1 {
			if m-j > factory[i][1] {
				return 1e18
			}
			for _, x := range robot[j:] {
				res += abs(x - factory[i][0])
			}
			return
		}
		res = f(i+1, j)
		for s, k := 0, 1; k <= factory[i][1] && j+k-1 < m; k++ {
			s += abs(robot[j+k-1] - factory[i][0])
			res = min(res, s+f(i+1, j+k))
		}
		return
	}
	return int64(f(0, 0))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// func minimumTotalDistance(robot []int, factory [][]int) int64 {
//	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
//	sort.Ints(robot)
//	m := len(robot)
//	f := make([]int, m+1)
//	for i := range f {
//		f[i] = 1e18
//	}
//	f[0] = 0
//	for _, fa := range factory {
//		for j := m; j > 0; j-- {
//			for k, cost := 1, 0; k <= min(j, fa[1]); k++ {
//				cost += abs(robot[j-k] - fa[0])
//				f[j] = min(f[j], f[j-k]+cost)
//			}
//		}
//	}
//	return int64(f[m])
//}
//
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
