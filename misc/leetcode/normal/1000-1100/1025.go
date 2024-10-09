package _000_1100

// 1027
func longestArithSeqLength(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = map[int]int{}
	}
	ans := 0

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for i, v := range nums {
		for j := 0; j < i; j++ {
			diff := nums[j] - v

			cnt, ok := dp[j][diff]
			if ok {
				dp[i][diff] = cnt + 1
			} else {
				dp[i][diff] = 2
			}

			ans = max(ans, dp[i][diff])
		}
	}

	return ans
}

// 1039
func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if r-l < 2 {
			return 0
		}

		if dp[l][r] >= 0 {
			return dp[l][r]
		}

		if r-l == 2 {
			return values[l] * values[l+1] * values[r]
		}

		res := int(1e9)
		cur := values[l] * values[r]
		for j := l + 1; j < r; j++ {
			res = min(res, cur*values[j]+dfs(l, j)+dfs(j, r))
		}

		dp[l][r] = res
		return res
	}

	return dfs(0, n-1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 1042

func gardenNoAdj(n int, paths [][]int) []int {
	var res = make([]int, n)
	// 构建邻接矩阵
	var c = make([][]int, n)
	for _, line := range paths {
		c[line[0]-1] = append(c[line[0]-1], line[1]-1)
		c[line[1]-1] = append(c[line[1]-1], line[0]-1)
	}
	for i := 0; i < n; i++ {
		used := make(map[int]struct{})
		for _, point := range c[i] {
			if res[point] != 0 {
				used[res[point]] = struct{}{}
			}
		}
		for colour := 1; colour <= 4; colour++ {
			if _, ok := used[colour]; !ok {
				res[i] = colour
				break
			}
		}
	}
	return res
}

// 1048
func longestStrChain(words []string) int {
	m := make(map[string]int)
	for i := range words {
		m[words[i]] = 0
	}
	var dfs func(s string) int
	dfs = func(s string) int {
		res := m[s]
		if res > 0 {
			return res
		}
		for i := range s {
			t := s[:i] + s[i+1:]
			if _, ok := m[t]; ok {
				res = int(max(int64(res), int64(dfs(t))))
			}
		}
		m[s] = res + 1
		return res + 1
	}
	ans := 0
	for i := range words {
		ans = int(max(int64(ans), int64(dfs(words[i]))))
	}
	return ans
}
