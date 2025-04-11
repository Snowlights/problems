//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2037G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx = 1000001
	mu := [mx]int{1: 1}
	for i := 1; i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			mu[j] -= mu[i]
		}
	}
	divisors := [mx][]int32{}
	for i := int32(2); i < mx; i++ {
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	const mod = 998244353
	cnt := [mx]int{}
	var n, v, f int
	Fscan(in, &n, &v)
	for _, d := range divisors[v] {
		cnt[d] = 1
	}
	for range n - 1 {
		Fscan(in, &v)
		f = 0
		for _, d := range divisors[v] {
			f -= mu[d] * cnt[d]
		}
		for _, d := range divisors[v] {
			cnt[d] = (cnt[d] + f) % mod
		}
	}
	Fprint(out, (f%mod+mod)%mod)

}

func main() { CF2037G(os.Stdin, os.Stdout) }
