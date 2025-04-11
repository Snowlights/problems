//go:build main
// +build main

package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF2048C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, []byte{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		i0 := bytes.IndexByte(s, '0')
		if i0 < 0 {
			Fprintln(out, 1, n, 1, 1)
			continue
		}
		m := n - i0
		ans, rr := []byte{}, 0
		for r := m; r <= n; r++ {
			t := slices.Clone(s[i0:])
			for i, b := range s[r-m : r] {
				t[i] ^= b & 1
			}
			if bytes.Compare(t, ans) > 0 {
				ans, rr = t, r
			}
		}
		Fprintln(out, 1, n, rr-m+1, rr)
	}

}

func main() { CF2048C(os.Stdin, os.Stdout) }
