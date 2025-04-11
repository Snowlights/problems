package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 998244353
	var n, m int
	var s []byte
	Fscan(in, &n, &m, &s)
	nxt := make([][26]int, 1<<n)
	for diff := range nxt {
		for c := byte('a'); c <= 'z'; c++ {
			var lcs, newLcs, newDiff int
			for j, b := range s {
				one := diff >> j & 1
				pre := newLcs
				newLcs = max(newLcs, lcs+(one|b2i(b == c)))
				newDiff |= (newLcs - pre) << j
				lcs += one
			}
			nxt[diff][c-'a'] = newDiff
		}
	}

	f := make([]int, 1<<n)
	f[0] = 1
	for i := 0; i < m; i++ {
		nf := make([]int, 1<<n)
		for diff, fv := range f {
			for _, newDiff := range nxt[diff] {
				nf[newDiff] = (nf[newDiff] + fv) % mod
			}
		}
		f = nf
	}

	ans := make([]int, n+1)
	for diff, fv := range f {
		ans[bits.OnesCount(uint(diff))] += fv
	}
	for _, v := range ans {
		Fprint(out, v%mod, " ")
	}

}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() { run(os.Stdin, os.Stdout) }
