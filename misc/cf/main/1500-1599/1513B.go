//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1513B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int = 1e9 + 7
	multi := func(x int) int {
		res := 1
		for i := 2; i < x; i++ {
			res *= i
			res %= mod
		}
		return res
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		h, mask := make(map[int]int), -1
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			h[v]++
			mask &= v
		}
		cnt := h[mask]
		Fprintln(out, cnt*(cnt-1)%mod*multi(n-1)%mod)
	}
}

func main() { CF1513B(os.Stdin, os.Stdout) }
