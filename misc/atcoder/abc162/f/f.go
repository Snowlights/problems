package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	dp := make([][3]int, n)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = 1e18
		}
	}
	var f func(int, int) int
	f = func(i, skip int) (res int) {
		if i >= n {
			if n&1 > 0 && skip == 2 || n&1 == 0 && skip == 1 {
				return
			}
			return -1e18
		}
		dv := &dp[i][skip]
		if *dv != 1e18 {
			return *dv
		}
		defer func() { *dv = res }()
		if n&1 > 0 {
			if skip == 2 {
				return a[i] + f(i+2, skip)
			}
			return max(a[i]+f(i+2, skip+b2i(i == n-2)), f(i+1, skip+1))
		} else {
			if skip == 1 {
				return a[i] + f(i+2, skip)
			}
			return max(a[i]+f(i+2, skip+b2i(i == n-2)), f(i+1, skip+1))
		}
	}
	ans := f(0, 0)
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = 1e18
		}
	}
	ans = max(ans, f(0, 0))
	Fprint(out, ans)

}

func runV2(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, s, v int
	Fscan(r, &n, &s)
	f := make([]int, n+1)
	for i := 2; i <= n; i++ {
		Fscan(r, &v)
		if i&1 == 0 {
			f[i] = max(s, f[i-2]+v)
		} else {
			s += v
			f[i] = max(f[i-1], f[i-2]+v)
		}
	}

	Fprintln(w, f[n])
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() { runV2(os.Stdin, os.Stdout) }
