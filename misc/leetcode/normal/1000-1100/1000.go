package _000_1100

import "strconv"

// 1012
// 给定正整数 n，返回在 [1, n] 范围内具有 至少 1 位 重复数字的正整数的个数。
func numDupDigitsAtMostN(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][1 << 10]int, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(i, mask int, limit, num bool) int
	f = func(i, mask int, limit, num bool) int {
		if i == m {
			if num {
				return 1
			}
			return 0
		}
		res := 0
		if !limit && num {
			dv := &dp[i][mask]
			if *dv >= 0 {
				return *dv
			}
			defer func() {
				*dv = res
			}()
		}
		if !num {
			res += f(i+1, mask, false, false)
		}
		lower, up := 1, 9
		if limit {
			up = int(s[i] - '0')
		}
		if num {
			lower = 0
		}

		for ; lower <= up; lower++ {
			if mask>>lower&1 == 0 {
				res += f(i+1, mask|1<<lower, limit && lower == up, true)
			}
		}
		return res
	}
	return n - f(0, 0, true, false)
}

// 1014
func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func maxScoreSightseeingPair(values []int) int {

	preM, ans := int64(0), int64(0)
	for i, v := range values {
		if i == 0 {
			preM = int64(v + i)
			continue
		}
		ans = max(ans, int64(preM+int64(v-i)))
		preM = max(preM, int64(v+i))
	}

	return int(ans)
}

// 1024
func videoStitching(clips [][]int, time int) (ans int) {
	f := make([]int, time)
	for _, c := range clips {
		l, r := c[0], c[1]
		if l < time && r > f[l] {
			f[l] = r
		}
	}

	pre, last, ans := 0, 0, 0
	for i, v := range f {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}

		if i == pre {
			pre = last
			ans++
		}

	}
	return ans
}
