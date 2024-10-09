package main

// 预处理每个数的不包括自己的因子，时间复杂度 O(mx*log(mx))
const mx = 200

var divisors [mx + 1][]int

func init() {
	for i := 1; i <= mx; i++ {
		for j := i * 2; j <= mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func calc(s string) int {
	n := len(s)
	res := n
	for _, d := range divisors[n] {
		cnt := 0
		for i0 := 0; i0 < d; i0++ {
			t := []byte{}
			for i := i0; i < n; i += d {
				t = append(t, s[i])
			}
			for i, m := 0, len(t); i < m/2; i++ {
				v, w := t[i], t[m-1-i]
				if v != w {
					cnt++
				}
			}
		}
		res = min(res, cnt)
	}
	return res
}

func minimumChanges(s string, k int) (ans int) {
	n := len(s)
	modify := make([][]int, n-1)
	for l := range modify {
		modify[l] = make([]int, n)
		for r := l + 1; r < n; r++ { // 半回文串长度至少为 2
			modify[l][r] = calc(s[l : r+1])
		}
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return modify[0][j]
		}
		p := &memo[j][i]
		if *p != -1 {
			return *p
		}
		res := n
		for L := i * 2; L < j; L++ {
			res = min(res, dfs(i-1, L-1)+modify[L][j])
		}
		*p = res
		return res
	}
	return dfs(k-1, n-1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
