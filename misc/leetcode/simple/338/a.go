package simple_338

import "sort"

// 1
func kItemsWithMaximumSum(numOnes, numZeros, _, k int) int {
	if k <= numOnes {
		return k
	}
	if k <= numOnes+numZeros {
		return numOnes
	}
	return numOnes*2 + numZeros - k
}

// 2
var p = []int{0} // 哨兵，避免二分越界

func init() {
	const mx = 1000
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			p = append(p, i) // 预处理质数列表
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func primeSubOperation(nums []int) bool {
	pre := 0
	for _, x := range nums {
		if x <= pre {
			return false
		}
		pre = x - p[sort.SearchInts(p, x-pre)-1] // 减去 < x-pre 的最大质数
	}
	return true
}

// 3
func minOperations(nums, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)
	sum := make([]int, n+1) // 前缀和
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	ans := make([]int64, len(queries))
	for i, q := range queries {
		j := sort.SearchInts(nums, q)
		left := q*j - sum[j]               // 蓝色面积
		right := sum[n] - sum[j] - q*(n-j) // 绿色面积
		ans[i] = int64(left + right)
	}
	return ans
}

// 4
func collectTheCoins(coins []int, edges [][]int) (ans int) {
	n := len(coins)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
		deg[x]++
		deg[y]++
	}

	// 用拓扑排序「剪枝」：去掉没有金币的子树
	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 1 && coins[i] == 0 { // 无金币叶子
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 && coins[y] == 0 {
				q = append(q, y)
			}
		}
	}

	// 再次拓扑排序
	for i, d := range deg {
		if d == 1 && coins[i] == 1 { // 有金币叶子
			q = append(q, i)
		}
	}
	if len(q) <= 1 { // 至多一个有金币的叶子，直接收集
		return
	}
	time := make([]int, n)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 {
				time[y] = time[x] + 1 // 记录入队时间
				q = append(q, y)
			}
		}
	}

	// 统计答案
	for _, e := range edges {
		if time[e[0]] >= 2 && time[e[1]] >= 2 {
			ans += 2
		}
	}
	return
}
