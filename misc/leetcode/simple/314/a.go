package simple_314

// 1
func hardestWorker(n int, logs [][]int) (ans int) {
	ans = logs[0][0]
	mx := logs[0][1]
	for i := 1; i < len(logs); i++ {
		t := logs[i][1] - logs[i-1][1]
		if t > mx || t == mx && logs[i][0] < ans {
			ans = logs[i][0]
			mx = t
		}
	}
	return
}

// 2
func findArray(a []int) (ans []int) {
	ans = append(ans, a[0])
	for i := 1; i < len(a); i++ {
		ans = append(ans, a[i]^a[i-1])
	}
	return
}

// 3
func robotWithString(s string) string {
	ans := make([]byte, 0, len(s))
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	min := byte(0) // 剩余最小字母
	st := []byte{}
	for _, c := range s {
		cnt[c-'a']--
		for min < 25 && cnt[min] == 0 { // 找到第一个正数
			min++
		}
		st = append(st, byte(c))
		for len(st) > 0 && st[len(st)-1]-'a' <= min {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return string(ans)
}

// 4
func numberOfPaths(grid [][]int, k int) int {

	m, n, mod := len(grid), len(grid[0]), int(1e9+7)
	dp := make([][][51]int, m+1)
	for i := range dp {
		dp[i] = make([][51]int, n+1)
	}
	// dp[1][0][0] = 1
	dp[1][0][0] = 1
	for i, row := range grid {
		for j, x := range row {
			for v := 0; v < k; v++ {
				dp[i+1][j+1][(v+x)%k] = (dp[i+1][j+1][(v+x)%k] + dp[i+1][j][v] + dp[i][j+1][v]) % mod
			}
		}
	}

	return dp[m][n][0]
}

// func numberOfPaths(grid [][]int, k int) int {
//	mod := int(1e9 + 7)
//	var dfs func(i, j, k int) int
//
//	type pair struct {
//		i, j, v int
//	}
//	h := make(map[pair]int)
//
//	dfs = func(i, j, v int) int {
//		p := pair{i: i, j: j, v: v}
//		if i < 0 || j < 0 {
//			h[p] = 0
//			return 0
//		}
//		v = (v + grid[i][j]) % k
//		p.v = v
//		if vv, ok := h[p]; ok {
//			return vv
//		}
//		if i == 0 && j == 0 {
//			if v == 0 {
//				h[p] = 1
//				return 1
//			}
//			h[p] = 0
//			return 0
//		}
//		ans := (dfs(i-1, j, v) + dfs(i, j-1, v)) % mod
//		h[p] = ans
//		return ans
//	}
//	return dfs(len(grid)-1, len(grid[0])-1, 0)
//}
