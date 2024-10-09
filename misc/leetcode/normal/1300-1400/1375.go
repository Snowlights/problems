package _300_1400

import "math"

// 1375
func numTimesAllBlue(flips []int) int {
	mx, ans := 0, 0
	for i, v := range flips {
		mx = max(v-1, mx)
		if mx == i {
			ans++
		}
	}
	return ans
}

// 1376
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {

	g, dp := make(map[int][]int), make([]int, n)
	for i, v := range manager {
		g[v] = append(g[v], i)
	}
	q := []int{headID}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, vv := range g[v] {
				dp[vv] = dp[v] + informTime[v]
				q = append(q, vv)
			}
		}
	}
	ans := dp[0]
	for _, v := range dp {
		ans = max(ans, v)
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 1385
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int
	for i := 0; i < len(arr1); i++ {
		isDistanceValue := true
		for j := 0; j < len(arr2); j++ {
			if int(math.Abs(float64(arr1[i]-arr2[j]))) <= d {
				isDistanceValue = false
				break
			}
		}
		if isDistanceValue {
			res++
		}
	}
	return res
}

// 1391
func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	q := [][]int{{0, 0}}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[0][0] = true
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	const (
		top    = 4
		button = 3
		right  = 2
		left   = 1
	)

	canmerge := func(from, to, loc int) bool {
		switch loc {
		case top:
			switch from {
			case 2, 5, 6:
				switch to {
				case 2, 3, 4:
					return true
				}
			}
		case button:
			switch from {
			case 2, 3, 4:
				switch to {
				case 2, 5, 6:
					return true
				}
			}
		case left:
			switch from {
			case 1, 3, 5:
				switch to {
				case 1, 4, 6:
					return true
				}
			}
		case right:
			switch from {
			case 1, 4, 6:
				switch to {
				case 1, 3, 5:
					return true
				}
			}
		}
		return false
	}

	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v[0] == m-1 && v[1] == n-1 {
				return true
			}
			for i, d := range dir {
				x, y := v[0]+d[0], v[1]+d[1]
				if 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && canmerge(grid[v[0]][v[1]], grid[x][y], i+1) {
					q = append(q, []int{x, y})
					vis[x][y] = true
				}
			}
		}
	}

	return false
}

// 1392
func longestPrefix(s string) string {
	ans := ""
	for i := 1; i < len(s); i++ {
		if s[:i] == s[len(s)-i:] {
			ans = s[:i]
		}
	}
	return ans
}

// 1397

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	mod := int64(1e9 + 7)
	m := int64(len(evil))
	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	fail := make([]int64, m+1)
	for i := range fail {
		fail[i] = -1
	}
	l, r := int64(-1), int64(0)
	for r < m {
		for r < m && (l == -1 || evil[l] == evil[r]) {
			l++
			r++
			fail[r] = l
		}
		l = fail[l]
	}

	var f func(i, e int64, equal1, equal2 bool) int64
	f = func(i, e int64, equal1, equal2 bool) int64 {
		if e == m {
			return 0
		}
		if i == int64(n) {
			return 1
		}
		if !equal1 && !equal2 && dp[i][e] != -1 {
			return dp[i][e]
		}

		down, up := byte('a'), byte('z')
		if equal1 {
			down = s1[i]
		}
		if equal2 {
			up = s2[i]
		}
		res := int64(0)
		for ; down <= up; down++ {
			next := e

			for next != -1 && down != evil[next] {
				next = fail[next]
			}
			res += f(i+1, next+1, equal1 && down == s1[i], equal2 && down == s2[i])
			res %= mod
		}

		if !equal1 && !equal2 {
			dp[i][e] = res
		}
		return res
	}
	return int(f(0, 0, true, true))
}
