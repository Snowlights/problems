package interview

// 面试题 13

func reach(b [][]bool, i, j, m, n int) bool {
	return (i-1 >= 0 && b[i-1][j]) ||
		(j-1 >= 0 && b[i][j-1]) ||
		(i+1 < m && b[i+1][j]) ||
		(j+1 < n && b[i][j+1])
}

func dig(a, b int) int {
	ans := 0
	for a > 0 {
		ans += a % 10
		a /= 10
	}
	for b > 0 {
		ans += b % 10
		b /= 10
	}
	return ans
}

func movingCount(m int, n int, k int) int {

	b, ans := make([][]bool, m), 1
	for i := range b {
		b[i] = make([]bool, n)
	}
	b[0][0] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dig(i, j) <= k && reach(b, i, j, m, n) {
				b[i][j] = true
				ans++
			}
		}
	}
	return ans
}
