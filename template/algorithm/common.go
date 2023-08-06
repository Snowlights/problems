package algorithm

func _() {
	// 同余前缀和，a 的下标从 0 开始，k 为模数
	// 具体用法见 query 上的注释
	// LC1664 https://leetcode-cn.com/problems/ways-to-make-a-fair-array/
	// https://atcoder.jp/contests/abc288/tasks/abc288_d
	groupPrefixSum := func(a []int, k int) {
		sum := make([]int, len(a)+k) // int64
		for i, v := range a {
			sum[i+k] = sum[i] + v
		}
		pre := func(x, t int) int {
			if x%k <= t {
				return sum[x/k*k+t]
			}
			return sum[(x+k-1)/k*k+t]
		}
		// 求下标在 [l,r) 范围内且下标模 k 同余于 t 的所有元素之和
		query := func(l, r, t int) int {
			t %= k
			return pre(r, t) - pre(l, t)
		}
		_ = query
	}
	_ = []interface{}{groupPrefixSum}
}
