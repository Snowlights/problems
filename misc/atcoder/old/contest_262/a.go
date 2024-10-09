package contest_262

import (
	"bufio"
	"fmt"
	"io"
)

func AtCoder262A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var y int64
	fmt.Fscan(r, &y)
	for i := y; i <= y+6; i++ {
		if i%4 == 2 {
			fmt.Fprintln(w, i)
			return
		}
	}
}

func AtCoder262B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, u, v int64
	fmt.Fscan(r, &n, &m)
	h := make(map[int64]map[int64]bool)
	for i := int64(0); i < m; i++ {
		fmt.Fscan(r, &u, &v)
		ins, ok := h[u]
		if !ok {
			ins = make(map[int64]bool)
			h[u] = ins
		}
		ins[v] = true
	}
	ans := 0
	for i := int64(1); i <= n; i++ {
		for k, _ := range h[i] {
			ins, ok := h[k]
			if !ok {
				continue
			}
			for kk, _ := range ins {
				if h[i][kk] {
					ans++
				}
			}
		}
	}
	fmt.Fprintln(w, ans)
}

func AtCoder262C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	ans := int64(0)
	pre := make([]int64, n+1)
	for i, v := range arr {
		if int64(i+1) == v {
			pre[i+1] = pre[i] + 1
		} else {
			pre[i+1] = pre[i]
		}
	}

	for i, v := range arr {
		if int64(i+1) == v {
			ans += pre[n] - pre[i+1]
		} else {
			if int64(i+1) > v && arr[v-1] == int64(i+1) {
				ans++
			}
		}
	}
	fmt.Fprintln(w, ans)
}

func AtCoder262D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}

	const mod = 998244353
	ans := int64(n)
	for l := int64(2); l <= n; l++ {
		dp := make([][]int64, l+1)
		for i := range dp {
			dp[i] = make([]int64, l)
		}
		dp[0][0] = 1
		for i := range arr {
			a := arr[i] % l
			for k := l; k >= max(1, int64(l+int64(i)-n-1)); k-- {
				for j := int64(0); j < l; j++ {
					dp[k][j] += dp[k-1][(j-a+l)%l]
					dp[k][j] %= mod
				}
			}
		}
		ans += dp[l][0]
		ans %= mod
	}
	fmt.Fprintln(w, ans)
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
