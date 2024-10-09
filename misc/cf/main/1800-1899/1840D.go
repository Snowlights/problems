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

func CF1840D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		Fprintln(out, sort.Search(1e9, func(m int) bool {
			i := sort.SearchInts(a, a[0]+m*2+1)
			b := a[i:]
			if len(b) == 0 {
				return true
			}
			i = sort.SearchInts(b, b[0]+m*2+1)
			return i == len(b) || b[len(b)-1]-b[i] <= m*2
		}))
	}

}

func main() { CF1840D(os.Stdin, os.Stdout) }
