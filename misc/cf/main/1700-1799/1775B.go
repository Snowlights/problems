//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1775B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
p:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, cnt := make([][]int, n), make(map[int]int)
		for i := 0; i < n; i++ {
			Fscan(in, &k)
			a[i] = make([]int, k)
			for j := range a[i] {
				Fscan(in, &a[i][j])
				cnt[a[i][j]]++
			}
		}
	o:
		for _, v := range a {
			for _, vv := range v {
				if cnt[vv] == 1 {
					continue o
				}
			}
			Fprintln(out, "Yes")
			continue p
		}
		Fprintln(out, "No")
	}

}

func main() { CF1775B(os.Stdin, os.Stdout) }
