//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1748B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		h := make(map[int]bool)
		for i := range s {
			h[int(s[i]-'0')] = true
		}
		ans := 0
		for i := range s {
			cnt, c, mxc := make([]int, 10), 0, 0
			for _, v := range s[i:] {
				v -= '0'
				if cnt[v] == 10 {
					break
				}
				if cnt[v] == 0 {
					c++
				}
				cnt[v]++
				if cnt[v] > mxc {
					mxc = cnt[v]
				}
				if mxc <= c {
					ans++
				}
			}
		}
		Fprintln(out, ans)

	}

}

func main() { CF1748B(os.Stdin, os.Stdout) }
