package hangzhoubianchengdasai

import (
	"math"
	"sort"
)

// 1
func canReceiveAllSignals(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

// 2
func minSwaps(chess []int) int {

	pre := make([]int, len(chess)+1)
	ones := 0
	for i, v := range chess {
		if v == 1 {
			pre[i+1] = pre[i] + 1
			ones++
		} else {
			pre[i+1] = pre[i]
		}
	}
	ans := int64(0)

	for i := 0; i <= len(chess)-ones; i++ {
		ans = max(ans, int64(pre[i+ones]-pre[i]))
	}

	return int(ones) - int(ans)
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 3
func buildTransferStation(area [][]int) int {
	x, y := make([]int, 0), make([]int, 0)
	for i := range area {
		for j, b := range area[i] {
			if b == 1 {
				x = append(x, i)
				y = append(y, j)
			}
		}
	}
	sort.Ints(x)
	sort.Ints(y)
	mid, ans := len(x)/2, 0
	for _, v := range x {
		ans += int(abs(int(v) - x[mid]))
	}
	for _, v := range y {
		ans += int(abs(int(v) - y[mid]))
	}
	return ans
}

// 4
func minTransfers(distributions [][]int) int {

	n := int64(0)
	for _, v := range distributions {
		n = max(n, int64(v[0]))
		n = max(n, int64(v[1]))
	}
	n++
	cnt := make([]int, n)
	for _, v := range distributions {
		cnt[v[0]] -= v[2]
		cnt[v[1]] += v[2]
	}

	dp, sum, c := make([]int, 1<<n), make([]int, 1<<n), make([]int, 1<<n)
	for i := 1; i < (1 << n); i++ {
		for j := int64(0); j < n; j++ {
			if i&(1<<j) != 0 {
				sum[i] = sum[i-(1<<j)] + cnt[j]
				c[i] = c[i-(1<<j)] + 1
				break
			}
		}
		if sum[i] != 0 {
			dp[i] = math.MaxInt64
			continue
		}
		dp[i] = c[i] - 1
		for j := 1; j < i; j++ {
			if (i|j != i) || (sum[j] != 0) {
				continue
			}
			dp[i] = int(min(int64(dp[i]), int64(dp[j]+dp[i-j])))
		}
	}
	return dp[(1<<n)-1]
}
