package double_91

import (
	"fmt"
	"math"
	"sort"
)

// 1
func distinctAverages(nums []int) int {
	h := make(map[int]bool)
	sort.Ints(nums)
	for i := 0; i < len(nums)/2; i++ {
		h[nums[i]+nums[len(nums)-i-1]] = true
	}
	return len(h)
}

// 2
func countGoodStrings(low int, high int, zero int, one int) int {
	f := make([]int, high+1)
	const mod = 1e9 + 7

	f[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			f[i] = (f[i] + f[i-zero]) % mod
		}
		if i >= one {
			f[i] = (f[i] + f[i-one]) % mod
		}
	}

	ans := 0
	for _, v := range f[low:] {
		ans = (ans + v) % mod
	}
	return ans
}

// 3
func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make(map[int][]int)
	for _, v := range edges {
		from, to := v[0], v[1]
		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}

	bobTime := make([]int, n)
	for i := range bobTime {
		bobTime[i] = math.MaxInt32
	}
	var dfsBob func(fa, i, cnt int) bool
	dfsBob = func(fa, j, cnt int) bool {
		if j == 0 {
			bobTime[j] = cnt
			return true
		}
		found := false
		for _, to := range g[j] {
			if to != fa {
				if dfsBob(j, to, cnt+1) {
					found = true
				}
			}
		}
		if found {
			bobTime[j] = cnt
		}
		return found
	}
	dfsBob(-1, bob, 0)

	ans := math.MinInt32

	g[0] = append(g[0], -1)
	var dfsAlice func(fa, i, t, total int)
	dfsAlice = func(fa, i, t, total int) {
		if t < bobTime[i] {
			total += amount[i]
		} else if t == bobTime[i] {
			total += amount[i] / 2
		}

		if len(g[i]) == 1 {
			ans = max(ans, total)
			return
		}

		for _, to := range g[i] {
			if to != fa {
				dfsAlice(i, to, t+1, total)
			}
		}
	}
	dfsAlice(-1, 0, 0, 0)

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 4
func splitMessage(message string, limit int) []string {

	n, suf := len(message), 5
	cap, num := 0, 0
	for {
		num++
		if num == 10 {
			cap -= 9
			suf = 7
		} else if num == 100 {
			cap -= 99
			suf = 9
		} else if num == 1000 {
			cap -= 999
			suf = 11
		}
		if limit-suf <= 0 {
			return nil
		}
		cap += limit - suf
		if cap < n {
			continue
		}

		ans := make([]string, 0)
		idx := 0
		for i := 1; i <= num && idx <= n; i++ {
			s := fmt.Sprintf("<%d/%d>", i, num)
			m := idx + limit - len(s)
			ans = append(ans, message[idx:min(m, n)]+s)
			idx = m
		}
		return ans
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
