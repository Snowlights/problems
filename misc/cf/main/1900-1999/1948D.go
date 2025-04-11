//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1948D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, ""
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		for m := n / 2; m > 0; m-- {
			cnt := 0
			for i := m; i < n; i++ {
				if s[i] != '?' && s[i-m] != '?' && s[i] != s[i-m] {
					cnt = 0
					continue
				}
				cnt++
				if cnt == m {
					Fprintln(out, m*2)
					continue o
				}
			}
		}
		Fprintln(out, 0)
	}

}

func main() { CF1948D(os.Stdin, os.Stdout) }
