//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func CF451E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	comb := func(n, k int) int {
		if n < k {
			return 0
		}
		n %= mod
		p, q := 1, 1
		for i := 1; i <= k; i++ {
			p = p * (n - i + 1) % mod
			q = q * i % mod
		}
		return p * pow(q, mod-2) % mod
	}

	var n, s, tot, ans int
	Fscan(in, &n, &s)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i]
	}
	if tot < s {
		Fprint(out, 0)
		return
	}
	for i := uint(0); i < 1<<n; i++ {
		s2 := s
		for j, v := range a {
			if i>>j&1 > 0 {
				s2 -= v + 1
			}
		}
		res := comb(s2+n-1, n-1)
		if bits.OnesCount(i)%2 > 0 {
			res = -res
		}
		ans += res
	}
	Fprint(out, (ans%mod+mod)%mod)

}

func main() { CF451E(os.Stdin, os.Stdout) }
