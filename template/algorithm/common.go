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

	// 按照从小到大的顺序，生成所有回文数
	// https://oeis.org/A002113
	// https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/
	// LC2081 https://leetcode.cn/problems/sum-of-k-mirror-numbers/
	// EXTRA: 单个数字的情况 LC564 https://leetcode.cn/problems/find-the-closest-palindrome/
	initPalindromeNumber := func() {
		const mx int = 1e9
		pal := []int{}
		// 哨兵。根据题目来定，也可以设置成 -2e9 等
		pal = append(pal, 0)

		// 严格按顺序从小到大生成所有回文数
	outer:
		for base := 1; ; base *= 10 {
			// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
			for i := base; i < base*10; i++ {
				x := i
				for t := i / 10; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				if x > mx {
					break outer
				}
				pal = append(pal, x)
			}
			// 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				if x > mx {
					break outer
				}
				pal = append(pal, x)
			}
		}

		// 哨兵。根据 mx 调整，如果 mx 是 2e9 的话要写成 mx+2
		pal = append(pal, mx+1)
	}

	_ = []interface{}{groupPrefixSum, initPalindromeNumber}
}
