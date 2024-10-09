package double_92

// 1
func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	}
	return n
}

// 2
func onesMinusZeros(grid [][]int) [][]int {

	m, n := len(grid), len(grid[0])
	onesRow, onesCol := make([]int, m), make([]int, n)
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i, v := range grid {
		for j, vv := range v {
			onesRow[i] += vv
			onesCol[j] += vv
		}
	}

	// diff[i][j] = onesRowi + onesColj - zerosRowi - zerosColj
	for i := range ans {
		for j := range ans[i] {
			ans[i][j] = onesRow[i] + onesCol[j] - (m - onesRow[i]) - (n - onesCol[j])
		}
	}

	return ans
}

// 3
func bestClosingTime(customers string) int {
	res, ans := 0, 0
	for i := range customers {
		if customers[i] == 'Y' {
			res++
		} else {
			res--
		}
		ans = max(ans, res)
	}
	if ans == 0 {
		return 0
	}

	now := 0
	for i := range customers {
		if customers[i] == 'Y' {
			now++
		} else {
			now--
		}
		if now == ans {
			return i + 1
		}
	}
	return 0
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 4
func countPalindromes(s string) (ans int) {
	const mod int = 1e9 + 7
	n := len(s)
	suf := [10]int{}
	suf2 := [10][10]int{}
	for i := n - 1; i >= 0; i-- {
		d := s[i] - '0'
		for j, c := range suf {
			suf2[d][j] += c
		}
		suf[d]++
	}

	pre := [10]int{}
	pre2 := [10][10]int{}
	for _, d := range s[:n-1] {
		d -= '0'
		suf[d]--
		for j, c := range suf {
			suf2[d][j] -= c // 撤销
		}
		for j, sf := range suf2 {
			for k, c := range sf {
				ans += pre2[j][k] * c // 枚举所有字符组合
			}
		}
		for j, c := range pre {
			pre2[d][j] += c
		}
		pre[d]++
	}
	return ans % mod
}
