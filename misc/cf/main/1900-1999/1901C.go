//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1901C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		mn, mx := int(1e9), 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			mn = min(mn, v)
			mx = max(mx, v)
		}
		if mn == mx {
			Fprintln(out, 0)
			continue
		}
		u := 1<<bits.Len(uint(mx)) - 1
		d := u - mx
		ans := 1 + bits.Len(uint((mn+d)/2^u/2))
		Fprintln(out, ans)
		if ans <= n {
			Fprintln(out, d, strings.Repeat("0 ", ans-1))
		}
	}

}

func main() { CF1901C(os.Stdin, os.Stdout) }
