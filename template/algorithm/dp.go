package algorithm

import (
	"math"
	"math/bits"
	"strconv"
)

// todo 背包

func _() {
	/* 区间 DP
	一般来说转移是合并区间或者分解区间
	套路 https://www.luogu.com.cn/blog/BreakPlus/ou-jian-dp-zong-jie-ti-xie

	① 将序列分成 K 个连续区间，求解这些区间的某个最优性质
	一般定义 dp[i][k] 表示将 a[:i] 分成 k 个连续区间得到的最优解
	此时可以枚举最后一个区间的左端点 j，从 dp[j-1][k-1] 转移到 dp[i][k]，转移时考虑 a[j:i] 对最优解的影响
	LC410 https://leetcode.cn/problems/split-array-largest-sum/
	LC813 https://leetcode.cn/problems/largest-sum-of-averages/
	LC1278 https://leetcode.cn/problems/palindrome-partitioning-iii/
	       至多 k 个回文串 https://codeforces.com/problemset/problem/137/D
	LC1335 https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/

	② 求解关于某个序列的最优性质，要求大区间的最优解可以依赖于小区间的最优解
	一般定义 dp[i][j] 表示 a[i:j] 的最优解
	此时可以枚举区间大小和区间左端点，从小区间转移到大区间
	LC375 https://leetcode.cn/problems/guess-number-higher-or-lower-ii/
	戳气球（好题）LC312 https://leetcode.cn/problems/burst-balloons/
	移除盒子（状态定义和转移的好题）LC546 https://leetcode.cn/problems/remove-boxes/
	打印机（好题）LC664 https://leetcode.cn/problems/strange-printer/
	最优三角剖分 LC1039 https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/
	插入形成回文 LC1312 https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/ https://www.luogu.com.cn/problem/P1435
	另见 LPS

	[1800·hot10] https://codeforces.com/problemset/problem/1509/C
	容斥 https://atcoder.jp/contests/abc106/tasks/abc106_d
	染色【套路】https://codeforces.com/problemset/problem/1114/D
	同色消除【套路】https://www.luogu.com.cn/problem/P4170 https://codeforces.com/problemset/problem/1132/F
	回文消除 https://codeforces.com/problemset/problem/607/B
	二维区间 DP https://codeforces.com/problemset/problem/1198/D
	状态设计的好题 https://codeforces.com/problemset/problem/1025/D
	https://codeforces.com/problemset/problem/149/D
	https://blog.csdn.net/weixin_43914593/article/details/106163859 算法竞赛专题解析（14）：DP应用--区间DP
	todo https://atcoder.jp/contests/abc159/tasks/abc159_f
	     https://codeforces.com/problemset/problem/245/H
	*/
	// 石子合并
	// https://atcoder.jp/contests/dp/tasks/dp_n
	// https://ac.nowcoder.com/acm/contest/1043/A https://ac.nowcoder.com/acm/problem/51170
	// 环形的情况 https://www.luogu.com.cn/problem/P1880
	// 相邻 k 堆的情况（综合①②）LC1000 https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
	mergeStones := func(a []int) int {
		n := len(a)
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, n)
			for j := range dp[i] {
				dp[i][j] = 1e9
			}
			dp[i][i] = 0
		}
		for sz := 2; sz <= n; sz++ {
			for l := 0; l+sz <= n; l++ {
				r := l + sz - 1
				for i := l; i < r; i++ {
					dp[l][r] = min(dp[l][r], dp[l][i]+dp[i+1][r])
				}
				dp[l][r] += sum[r+1] - sum[l]
			}
		}
		return dp[0][n-1]
	}

	// 统计区间内回文串个数
	// 返回一个二维数组 dp, dp[i][j] 表示 [i,j] 内的回文串的个数
	// https://codeforces.com/problemset/problem/245/H
	countPalindromes := func(s string) [][]int {
		n := len(s)
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, n)
			dp[i][i] = 1
			if i+1 < n && s[i] == s[i+1] {
				dp[i][i+1] = 1
			}
		}
		for i := n - 3; i >= 0; i-- {
			for j := i + 2; j < n; j++ {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
		// 到这里为止，dp[i][j] = 1 表示 s[i:j+1] 是回文串
		for i := n - 2; i >= 0; i-- {
			for j := i + 1; j < n; j++ {
				dp[i][j] += dp[i][j-1] + dp[i+1][j] - dp[i+1][j-1] // 容斥
			}
		}
		return dp
	}

	_ = []interface{}{mergeStones, countPalindromes}

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
	{
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
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
