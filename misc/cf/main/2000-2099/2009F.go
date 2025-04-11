//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2009F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}
		f := func(r int) int {
			k, i := r/n, r%n
			if i <= n-k {
				return k*s[n] + s[k+i] - s[k]
			}
			return (k+1)*s[n] - s[k] + s[i-(n-k)]
		}
		for range q {
			Fscan(in, &l, &r)
			Fprintln(out, f(r)-f(l-1))
		}
	}

}

func main() { CF2009F(os.Stdin, os.Stdout) }
