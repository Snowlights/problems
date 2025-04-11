//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1844C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		mx := int(-1e9)
		s := [2]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			mx = max(mx, v)
			s[i%2] += max(v, 0)
		}
		if mx < 0 {
			Fprintln(out, mx)
		} else {
			Fprintln(out, max(s[0], s[1]))
		}
	}

}

func main() { CF1844C(os.Stdin, os.Stdout) }
