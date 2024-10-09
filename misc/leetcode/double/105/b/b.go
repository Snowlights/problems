package main

func minExtraChar(s string, dictionary []string) (ans int) {
	d := make(map[string]bool)
	for _, v := range dictionary {
		d[v] = true
	}

	n := len(s)
	mm := make([]int, n)
	for i := range mm {
		mm[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		dv := &mm[i]
		if *dv >= 0 {
			return *dv
		}

		res := dfs(i-1) + 1

		for j := 0; j <= i; j++ {
			if d[s[j:i+1]] {
				res = min(res, dfs(j-1))
			}
		}
		*dv = res
		return res
	}

	return dfs(n - 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
