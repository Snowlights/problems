//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1788B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 1_000_000_007
	const mx = 3000
	pow2 := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := int64(0)
	for i, ai := range a {
		l, r := i-1, i+2
		for j := i + 1; j < n; j++ {
			aj := a[j]
			for l >= 0 && a[l] >= ai*2-aj {
				l--
			}
			for r <= j || r < n && a[r] < aj*2-ai {
				r++
			}
			ans += pow2[l+1+n-r]
		}
	}
	Fprint(out, ans%mod)

}

func main() { CF1788B(os.Stdin, os.Stdout) }
