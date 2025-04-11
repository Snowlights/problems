//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1955D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		cnt := map[int]int{}
		for i := 0; i < m; i++ {
			Fscan(in, &v)
			cnt[v]++
		}

		ans, c := 0, 0
		for r, v := range a {
			if cnt[v] > 0 {
				c++
			}
			cnt[v]--
			l := r + 1 - m
			if l < 0 {
				continue
			}
			if c >= k {
				ans++
			}
			cnt[a[l]]++
			if cnt[a[l]] > 0 {
				c--
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1955D(os.Stdin, os.Stdout) }
