//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2047B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		mx, mn := -1, -1
		for i, c := range cnt {
			if c == 0 {
				continue
			}
			if mx < 0 || c > cnt[mx] {
				mx = i
			}
			if mn < 0 || c <= cnt[mn] {
				mn = i
			}
		}
		for i, b := range s {
			if b == 'a'+byte(mn) {
				s[i] = 'a' + byte(mx)
				break
			}
		}
		Fprintln(out, string(s))
	}

}

func main() { CF2047B(os.Stdin, os.Stdout) }
