package simple_298

import (
	"math/bits"
	"strconv"
	"strings"
)

// 1
func greatestLetter(s string) string {
	h := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		h[s[i]] = true
	}
	for tmp := 'Z'; tmp >= 'A'; tmp-- {
		last := strings.ToLower(string(tmp))
		if h[byte(tmp)] && h[[]byte(last)[0]] {
			return string(tmp)
		}
	}

	return ""
}

// 2

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	if num < k {
		return -1
	}
	if k == 0 {
		if num%10 == 0 {
			return 1
		}
		return -1
	}
	count := 1
	for {
		r := num - count*k
		if r < 0 {
			return -1
		}
		if r%10 == 0 {
			return count
		}
		count++
	}
}

// 3
func longestSubsequence(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if n < m {
		return n
	}
	ans := m
	if v, _ := strconv.ParseInt(s[n-m:], 2, 0); int(v) > k {
		ans--
	}
	return ans + strings.Count(s[:n-m], "0")
}

// 4
func sellingWood(m, n int, prices [][]int) int64 {
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for _, price := range prices {
		pr[price[0]][price[1]] = price[2]
	}

	f := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			f[i][j] = pr[i][j]
			for k := 1; k < j; k++ { // 垂直切割
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k < i; k++ { // 水平切割
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return int64(f[m][n])
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
