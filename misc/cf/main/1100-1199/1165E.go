//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF1165E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 998244353
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] *= (n - i) * (i + 1)
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	sort.Ints(b)

	for i, v := range a {
		ans = (ans + v%mod*b[i]) % mod
	}
	Fprint(out, ans)

}

func main() { CF1165E(os.Stdin, os.Stdout) }
