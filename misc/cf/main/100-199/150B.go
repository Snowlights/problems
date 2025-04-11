//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF150B(_r io.Reader, _w io.Writer) {
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
	var n, m, k int
	Fscan(in, &n, &m, &k)
	if k == 1 || k > n {
		Fprint(out, pow(m, n))
	} else if k == n {
		Fprint(out, pow(m, (n+1)/2))
	} else if k%2 == 0 {
		Fprint(out, m)
	} else {
		Fprint(out, m*m)
	}

}

func main() { CF150B(os.Stdin, os.Stdout) }
