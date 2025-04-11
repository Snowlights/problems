//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1225B2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &d, &d)
		a := make([]int, n)
		cnt := map[int]int{}
		ans := n
		for i := range a {
			Fscan(in, &a[i])
			cnt[a[i]]++
			l := i - d + 1
			if l < 0 {
				continue
			}
			ans = min(ans, len(cnt))
			v := a[l]
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1225B2(os.Stdin, os.Stdout) }
