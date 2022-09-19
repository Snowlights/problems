package didi

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

func LCDiDi001(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var np, nq, nr int
	fmt.Fscan(r, &np, &nq, &nr)
	type pair struct {
		pre, p, q, r int
	}
	h := make(map[pair]int)

	var dfs func(int, int, int, int) int
	dfs = func(pre, p, q, r int) int {
		pa := pair{pre, p, q, r}
		if val, ok := h[pa]; ok {
			return val
		}

		if p+q+r == 0 {
			return 1
		}

		res := 0
		if pre != 0 && p > 0 {
			res += dfs(0, p-1, q, r)
		}
		if pre != 1 && q > 0 {
			res += dfs(1, p, q-1, r)
		}
		if pre != 2 && r > 0 {
			res += dfs(2, p, q, r-1)
		}
		h[pa] = res
		return res
	}

	res := 0
	if np > 0 {
		res += dfs(0, np-1, nq, nr)
	}
	if nq > 0 {
		res += dfs(1, np, nq-1, nr)
	}
	if nr > 0 {
		res += dfs(2, np, nq, nr-1)
	}
	fmt.Fprintln(w, res)
}

func LCDiDi002(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	s1, s2 := "qwertasdfgzxcv", "yuiophjklbnm"
	h1, h2 := make(map[byte]bool), make(map[byte]bool)
	for i := range s1 {
		h1[s1[i]] = true
	}
	for i := range s2 {
		h2[s2[i]] = true
	}

	minDis := func(a, b string) int {
		m, n := len(a), len(b)
		if m*n == 0 {
			return m + n
		}
		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}
		for i := 0; i <= m; i++ {
			dp[i][0] = 3 * i
		}
		for i := 0; i <= n; i++ {
			dp[0][i] = 3 * i
		}
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if a[i-1] == b[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 3
					// 替换
					if (h1[a[i-1]] && h1[b[j-1]]) || (h2[a[i-1]] && h2[b[j-1]]) {
						dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
					} else if (h1[a[i-1]] && h2[b[j-1]]) || (h2[a[i-1]] && h1[b[j-1]]) {
						dp[i][j] = min(dp[i][j], dp[i-1][j-1]+2)
					}
				}
			}
		}
		return dp[m][n]
	}

	s, _, _ := r.ReadLine()
	parts := strings.Split(string(s), " ")
	target := parts[0]
	words := parts[1:]
	target = strings.ToLower(target)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	sort.Slice(words, func(i, j int) bool {
		return minDis(target, words[i]) < minDis(target, words[j])
	})

	for i := 0; i < 3 && i < len(words); i++ {
		fmt.Fprint(w, words[i], " ")
	}

}
