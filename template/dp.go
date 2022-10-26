package template

import (
	"math"
	"math/bits"
	"strconv"
)

// todo 背包

// 数位dp
/* 数位 DP
一般用 dp[i][j] 表示当前在第 i 位，前面维护了一个为 j 的值，且后面的位数可以随便选时的数字个数
在解释状态的含义时，网上的很多文章都漏了「后面的位数可以随便选」这个约束，只有加上这个约束，我们才能根据「是否紧贴上界」来完成相对应的代码逻辑

https://zhuanlan.zhihu.com/p/348851463
https://www.bilibili.com/video/BV1MT4y1376C
https://www.bilibili.com/video/BV1yT4y1u7jW

入门题 https://atcoder.jp/contests/abc154/tasks/abc154_e
      https://atcoder.jp/contests/dp/tasks/dp_s
      https://codeforces.com/problemset/problem/1036/C
二进制 1 的个数恰为 k 的数字个数 https://codeforces.com/problemset/problem/431/D https://www.acwing.com/problem/content/1083/
是 m 的倍数且偶数位为 d 且奇数位不为 d 的数字个数 https://codeforces.com/problemset/problem/628/D
所有数字均出现偶数次的数字个数 https://codeforces.com/problemset/problem/855/E
相邻数字约束 SC09 https://www.luogu.com.cn/problem/P2657
含有某个数字的数字个数
LC233 https://leetcode-cn.com/problems/number-of-digit-one/
      https://leetcode-cn.com/problems/number-of-2s-in-range-lcci/
      http://acm.hdu.edu.cn/showproblem.php?pid=3555
      http://acm.hdu.edu.cn/showproblem.php?pid=2089
LC600 二进制不含连续 1 的数字个数 https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/
LC902/周赛101C 最大为 N 的数字组合 https://leetcode-cn.com/contest/weekly-contest-101/problems/numbers-at-most-n-given-digit-set/
LC1012/周赛128D 有重复数字的数字个数 https://leetcode-cn.com/contest/weekly-contest-128/problems/numbers-with-repeated-digits/
LC/周赛306D 互补问题 无重复数字的数字个数 https://leetcode.cn/contest/weekly-contest-306/problems/count-special-integers/
LC1067/双周赛1D 字符 d 出现次数 https://leetcode-cn.com/contest/biweekly-contest-1/problems/digit-count-in-range/
LC1397/周赛182D 与 KMP 结合 https://leetcode-cn.com/contest/weekly-contest-182/problems/find-all-good-strings/
digsum(n)|n 的数的个数 https://www.luogu.com.cn/problem/P4127 https://www.acwing.com/problem/content/313/
http://acm.hdu.edu.cn/showproblem.php?pid=3886
http://acm.hdu.edu.cn/showproblem.php?pid=6796
注：一些第 k 小的题目需要与二分结合，或者用试填法（见后面的 kth666）
*/

func _() {
	digitDP := func(lower, upper string, sumUpper int) int64 {
		const mod int64 = 1e9 + 7

		calcGiven := func(digits []string, n int) int {
			s := strconv.Itoa(n)
			m := len(s)
			dp := make([]int, m)
			for i := range dp {
				dp[i] = -1 // dp[i] = -1 表示 i 这个状态还没被计算出来
			}
			var f func(int, bool, bool) int
			f = func(i int, isLimit, isNum bool) (res int) {
				if i == m {
					if isNum { // 如果填了数字，则为 1 种合法方案
						return 1
					}
					return
				}
				if !isLimit && isNum { // 在不受到任何约束的情况下，返回记录的结果，避免重复运算
					dv := &dp[i]
					if *dv >= 0 {
						return *dv
					}
					defer func() { *dv = res }()
				}
				if !isNum { // 前面不填数字，那么可以跳过当前数位，也不填数字
					// isLimit 改为 false，因为没有填数字，位数都比 n 要短，自然不会受到 n 的约束
					// isNum 仍然为 false，因为没有填任何数字
					res += f(i+1, false, false)
				}
				// 根据是否受到约束，决定可以填的数字的上限
				up := byte('9')
				if isLimit {
					up = s[i]
				}
				// 注意：对于一般的题目而言，如果此时 isNum 为 false，则必须从 1 开始枚举，由于本题 digits 没有 0，所以无需处理这种情况
				for _, d := range digits { // 枚举要填入的数字 d
					if d[0] > up { // d 超过上限，由于 digits 是有序的，后面的 d 都会超过上限，故退出循环
						break
					}
					// isLimit：如果当前受到 n 的约束，且填的数字等于上限，那么后面仍然会受到 n 的约束
					// isNum 为 true，因为填了数字
					res += f(i+1, isLimit && d[0] == up, true)
				}
				return
			}
			return f(0, true, false)
		}

		// 返回 <=s 的符合要求的字符串数目
		// TIPS: 某些情况下思考补集会更加容易，即求不符合要求的字符串数目
		calc := func(s string) int64 {
			const lowerC, upperC byte = '0', '9'
			dp := make([][]int64, len(s))
			for i := range dp {
				dp[i] = make([]int64, sumUpper+1)
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			var f func(p, sum int, limitUp bool) int64
			f = func(p, sum int, limitUp bool) (res int64) {
				if p == len(s) {
					return 1
				} // sum
				if !limitUp {
					dv := &dp[p][sum]
					if *dv >= 0 {
						return *dv
					} // *dv + sum*int64(math.Pow10(n-p))
					defer func() { *dv = res }()
				}
				up := upperC
				if limitUp {
					up = s[p]
				}
				for ch := lowerC; ch <= up; ch++ {
					tmp := sum

					cnt := f(p+1, tmp, limitUp && ch == up)
					res = (res + cnt) % mod
				}
				return
			}
			res := f(0, 0, true)
			return res
		}
		ansLower := calc(lower) // lower-1
		ansUpper := calc(upper)
		ans := ansUpper - ansLower
		// lower 是否算上
		//if lowerIsValid {
		//	ans++
		//}
		ans = (ans%mod + mod) % mod

		// TIPS: 对于需要判断/禁止前导零的情况，可以加一个额外的维度 fill，
		// 表示已经填入了数字（没有前导零的合法状态），最后 p>=n 的时候可以根据情况返回 1 或者 0
		// 例如 https://codeforces.com/contest/855/submission/125651587
		// 以下代码以 https://www.luogu.com.cn/problem/P2657 为例
		calc = func(s string) int64 {
			dp := make([][10]int64, len(s))
			for i := range dp {
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			var f func(p, pre int, limitUp, fill bool) int64
			f = func(p, pre int, limitUp, fill bool) (res int64) {
				if p == len(s) {
					return 1
				}
				if !limitUp && fill { // 注意这里的判断
					dv := &dp[p][pre]
					if *dv >= 0 {
						return *dv
					}
					defer func() { *dv = res }()
				}
				up := 9
				if limitUp {
					up = int(s[p] & 15)
				}
				for d := 0; d <= up; d++ {
					if !fill || abs(d-pre) > 1 {
						res += f(p+1, d, limitUp && d == up, fill || d > 0)
					}
				}
				return
			}
			return f(0, 0, true, false)
		}

		// 若需要计算的不是合法数字个数，而是合法数字之和，则需要在计算时考虑单个数位的贡献
		// 以下代码以 https://codeforces.com/problemset/problem/1073/E 为例
		calcSum := func(s string, k int) int64 {
			n := len(s)
			type pair struct{ cnt, sum int64 }
			dp := make([][1 << 10]pair, n)
			for i := range dp {
				for j := range dp[i] {
					dp[i][j] = pair{-1, -1}
				}
			}
			var f func(int, uint16, bool, bool) pair
			f = func(p int, mask uint16, limitUp, fill bool) (res pair) {
				if p == n {
					if !fill {
						return
					}
					return pair{1, 0}
				}
				if !limitUp && fill {
					dv := &dp[p][mask]
					if dv.cnt >= 0 {
						return *dv
					}
					defer func() { *dv = res }()
				}
				up := 9
				if limitUp {
					up = int(s[p] & 15)
				}
				for d := 0; d <= up; d++ {
					tmp := mask
					if fill || d > 0 {
						tmp |= 1 << d
					}
					if bits.OnesCount16(tmp) <= k {
						pr := f(p+1, tmp, limitUp && d == up, fill || d > 0)
						res.cnt = (res.cnt + pr.cnt) % mod
						res.sum = (res.sum + int64(math.Pow10(n-1-p))%mod*pr.cnt%mod*int64(d) + pr.sum) % mod
					}
				}
				return
			}
			return f(0, 0, true, false).sum
		}
		_ = []interface{}{calcGiven, calcSum}

		return ans
	}

	_ = []interface{}{digitDP}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
