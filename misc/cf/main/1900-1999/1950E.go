//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1950E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
	o:
		for i := 1; i <= n; i++ {
			if n%i > 0 {
				continue
			}
			cnt := map[string]int{}
			for j := i; j <= n; j += i {
				cnt[s[j-i:j]]++
			}
			if len(cnt) > 2 {
				continue
			}
			if len(cnt) == 1 {
				Fprintln(out, i)
				break
			}
			var v, w string
			for s, c := range cnt {
				if c != 1 && c != n/i-1 {
					continue o
				}
				if v == "" {
					v = s
				} else {
					w = s
				}
			}
			diff := 0
			for k := range v {
				if v[k] != w[k] {
					diff++
				}
			}
			if diff == 1 {
				Fprintln(out, i)
				break
			}
		}
	}

}

func main() { CF1950E(os.Stdin, os.Stdout) }
