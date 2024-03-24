//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF1701C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		cnt := make([]int, n)
		for i := 0; i < m; i++ {
			Fscan(in, &v)
			cnt[v-1]++
		}
		Fprintln(out, sort.Search(2*m, func(t int) bool {
			left := 0
			for _, c := range cnt {
				if c > t {
					left += c - t
				} else {
					left -= (t - c) / 2
				}
			}
			return left <= 0
		}))
	}

}

func main() { CF1701C(os.Stdin, os.Stdout) }
