//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1336C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 998244353
	var s, t string
	Fscan(in, &s, &t)
	n, m := len(s), len(t)
	f := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		if i >= m || s[0] == t[i] {
			f[i][i] = 2
		}
		for j := i + 1; j < n; j++ {
			if i >= m || s[j-i] == t[i] {
				f[i][j] = f[i+1][j]
			}
			if j >= m || s[j-i] == t[j] {
				f[i][j] = (f[i][j] + f[i][j-1]) % mod
			}
		}
	}
	ans := 0
	for _, v := range f[0][m-1:] {
		ans = (ans + v) % mod
	}
	Fprint(out, ans)

}

func main() { CF1336C(os.Stdin, os.Stdout) }
