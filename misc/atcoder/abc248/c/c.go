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

	var n, m, k, ans int
	const mod int = 998244353
	Fscan(in, &n, &m, &k)
	f := make([]int, k+1)
	f[0] = 1

	for ; n > 0; n-- {
		suf := make([]int, k+2)
		for i := k; i >= 0; i-- {
			suf[i] = (suf[i+1] + f[i]) % mod
		}
		for i := k; i >= 0; i-- {
			f[i] = (suf[max(i-m, 0)] - suf[i] + mod) % mod
		}
	}

	for _, v := range f {
		ans = (ans + v) % mod
	}
	Fprintln(out, ans%mod)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
