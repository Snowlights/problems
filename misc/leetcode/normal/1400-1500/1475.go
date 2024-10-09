package _400_1500

import (
	"math"
	"math/bits"
	"sort"
)

// 1480

func runningSum(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i] + nums[i]
	}
	return pre
}

// 1482
func minDays(bloomDay []int, m, k int) int {
	if m > len(bloomDay)/k {
		return -1
	}
	maxDay := 0
	for _, day := range bloomDay {
		if day > maxDay {
			maxDay = day
		}
	}
	return sort.Search(maxDay, func(days int) bool {
		flowers, bouquets := 0, 0
		for _, d := range bloomDay {
			if d > days {
				flowers = 0
			} else {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			}
		}
		return bouquets >= m
	})
}

// 1483
type TreeAncestor struct {
	ancestor [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	mx := bits.Len(uint(n))
	pa := make([][]int, n)
	for i := range pa {
		pa[i] = make([]int, mx)
		pa[i][0] = parent[i]
	}
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	return TreeAncestor{pa}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	pa := this.ancestor
	mx := bits.Len(uint(k))
	uptoKthPa := func(v, k int) int {
		for i := 0; i < mx && v != -1; i++ {
			if k>>i&1 > 0 {
				v = pa[v][i]
			}
		}
		return v
	}

	return uptoKthPa(node, k)
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */

// 1484
// SELECT sell_date,
//       COUNT(DISTINCT product) num_sold,
//       GROUP_CONCAT(DISTINCT product) products
// FROM Activities
// GROUP BY sell_date

// 1494
func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	// 计算每门课的先修课集合
	pre := make([]int, n)
	for _, d := range dependencies {
		pre[d[1]-1] |= 1 << (d[0] - 1)
	}
	m := 1 << n
	// 计算所有课程集合的先修课集合，不合法的标记为 -1
	totPre := make([]int, m)
	for i := range totPre {
		if bits.OnesCount(uint(i)) > k {
			totPre[i] = -1
			continue
		}
		for s := uint(i); s > 0; s &= s - 1 {
			p := pre[bits.TrailingZeros(s)]
			if p&i > 0 {
				totPre[i] = -1
				break
			}
			totPre[i] |= p
		}
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0
	for s, v := range dp {
		t := m - 1 ^ s                               // 补集
		for sub := t; sub > 0; sub = (sub - 1) & t { // 枚举下个学期要学的课
			if p := totPre[sub]; p >= 0 && s&p == p { // 这些课的先修课必须合法且在之前的学期里必须上过
				dp[s|sub] = min(dp[s|sub], v+1)
			}
		}
	}
	return dp[m-1]
}

func minNumberOfSemestersV2(n int, relations [][]int, k int) int {
	pre1 := make([]int, n)
	for _, r := range relations {
		pre1[r[1]-1] |= 1 << (r[0] - 1) // r[1] 的先修课程集合，下标改从 0 开始
	}

	u := 1<<n - 1 // 全集
	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i == 0 { // 空集
			return
		}
		p := &memo[i]
		if *p != -1 { // 之前算过了
			return *p
		}
		defer func() { *p = res }() // 记忆化
		ci := u ^ i                 // i 的补集
		i1 := 0
		for j, p := range pre1 {
			if i>>j&1 > 0 && p|ci == ci { // p 在 i 的补集中
				i1 |= 1 << j
			}
		}
		if bits.OnesCount(uint(i1)) <= k { // 如果个数小于 k，则可以全部学习，不再枚举子集
			return dfs(i^i1) + 1
		}
		res = math.MaxInt
		for j := i1; j > 0; j = (j - 1) & i1 { // 枚举 i1 的子集 j
			if bits.OnesCount(uint(j)) <= k {
				res = min(res, dfs(i^j)+1)
			}
		}
		return res
	}
	return dfs(u)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 1498
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	pow := map[int]int{0: 1}
	n, mod := len(nums), int(1e9+7)
	for i := 1; i < n; i++ {
		pow[i] = (pow[i-1] * 2) % mod
	}
	ans := 0
	for len(nums) > 0 {
		val := nums[0]
		nums = nums[1:]
		idx := sort.Search(len(nums), func(i int) bool {
			return nums[i] > target-val
		})
		if idx == 0 {
			if val*2 <= target {
				ans += pow[idx]
				ans %= mod
			}
			continue
		}
		ans += pow[idx]
		ans %= mod
	}

	return ans
}
