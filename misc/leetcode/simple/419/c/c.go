package main

func countWinningSequences(s string) int {
	const mod = 1_000_000_007
	n := len(s)
	pow2 := make([]int, (n+1)/2)
	pow2[0] = 1
	for i := 1; i < len(pow2); i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	mp := [...]int{'F': 0, 'W': 1, 'E': 2}
	memo := make([][][3]int, n)
	for i := range memo {
		memo[i] = make([][3]int, n*2+1)
		for j := range memo[i] {
			memo[i][j] = [3]int{-1, -1, -1} // -1 表示没有计算过
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, ban int) (res int) {
		if -j > i { // 必败
			return
		}
		if j > i+1 { // 必胜
			return pow2[i+1]
		}
		p := &memo[i][j+n][ban]
		if *p != -1 { // 之前计算过
			return *p
		}
		for k := 0; k < 3; k++ { // 枚举当前召唤的生物
			// 判断 i == n-1，避免讨论 k == ban 的情况
			if i == n-1 || k != ban { // 不能连续两个回合召唤相同的生物
				score := (k - mp[s[i]] + 3) % 3
				if score == 2 {
					score = -1
				}
				res += dfs(i-1, j+score, k)
			}
		}
		res %= mod
		*p = res // 记忆化
		return
	}
	return dfs(n-1, 0, 0) // ban=0,1,2 都可以
}
