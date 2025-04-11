//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2025D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, r, c0 int
	Fscan(in, &n, &m)
	f := make([]int, m+1)
	cnt := make([][2]int, m+1)
	for i := range n {
		Fscan(in, &r)
		if r < 0 {
			cnt[-r][0]++
		} else if r > 0 {
			cnt[r][1]++
		}
		if r != 0 && i < n-1 {
			continue
		}
		for j := 2; j <= c0; j++ {
			cnt[j][0] += cnt[j-1][0]
			cnt[j][1] += cnt[j-1][1]
		}
		for j := c0; j > 0; j-- {
			f[j] = max(f[j], f[j-1]) + cnt[j][0] + cnt[c0-j][1]
		}
		f[0] += cnt[c0][1]
		c0++
		clear(cnt)
	}
	Fprint(out, slices.Max(f))

}

func main() { CF2025D(os.Stdin, os.Stdout) }
