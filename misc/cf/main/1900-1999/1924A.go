//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"strings"
)

func CF1924A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, m int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &m, &s)
		t := s[:0]
		mask := 0
		for _, c := range s {
			mask |= 1 << (c - 'a')
			if mask == 1<<k-1 {
				t = append(t, c)
				mask = 0
			}
		}
		if len(t) >= n {
			Fprintln(out, "YES")
		} else {
			Fprintf(out, "NO\n%s%c%s\n", t, 'a'+bits.TrailingZeros(uint(^mask)), strings.Repeat("a", n-1-len(t)))
		}
	}

}

func main() { CF1924A(os.Stdin, os.Stdout) }
