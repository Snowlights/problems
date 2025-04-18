//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1965B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		hb := 1 << (bits.Len(uint(k)) - 1)
		a := []any{k - hb, k + 1, hb<<1 | k}
		for i := 1; i <= n; i <<= 1 {
			if i != hb {
				a = append(a, i)
			}
		}
		Fprintln(out, len(a))
		Fprintln(out, a...)
	}

}

func main() { CF1965B(os.Stdin, os.Stdout) }
