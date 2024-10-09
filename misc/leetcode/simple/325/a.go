package simple_325

import (
	"sort"
)

// 1
func closetTarget(words []string, target string, startIndex int) int {

	if words[startIndex] == target {
		return 0
	}
	n := len(words)
	l, r := (startIndex-1+n)%n, (startIndex+1)%n
	ans, step := -1, 1
	for {
		if words[l] == target || words[r] == target {
			if ans == -1 {
				ans = step
			}
		}
		l = (l - 1 + n) % n
		r = (r + 1) % n
		step++
		if step > n/2 {
			break
		}
	}

	return ans
}

// 2
func takeCharacters(s string, k int) int {
	cnt := [3]int{}
	n := len(s)
	j := n

	for cnt[0] < k || cnt[1] < k || cnt[2] < k {
		if j == 0 {
			return -1
		}
		j--
		cnt[s[j]-'a']++
	}
	ans := n - j
	for i := 0; i < n && j < n; i++ {
		cnt[s[i]-'a']++
		for j < n && cnt[s[j]-'a'] >= k {
			cnt[s[j]-'a']--
			j++
		}
		ans = min(ans, i+1+n-j)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 3
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	return sort.Search(price[len(price)-1], func(i int) bool {
		cnt := 1
		i++
		x0 := price[0]
		for _, v := range price[1:] {
			if v >= x0+i {
				cnt++
				x0 = v
			}
		}

		return cnt < k
	})
}

// 4
func countPartitions(nums []int, k int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum < k*2 {
		return 0
	}

	n := len(nums)
	mod := int(1e9 + 7)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	dp[0][0] = 1
	ans := 0

	for i, v := range nums {
		for j := 0; j < k; j++ {
			dp[i+1][j] = dp[i][j]
			if j-v >= 0 {
				dp[i+1][j] += dp[i][j-v]
			}
			dp[i+1][j] %= mod
		}
	}
	for _, v := range dp[n] {
		ans = ans + v
		ans %= mod
	}
	ans = ans * 2 % mod

	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	return (pow(2, n) - ans + mod) % mod
}
