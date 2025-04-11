//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1926D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		cnt := map[int]int{}
		for range n {
			Fscan(in, &v)
			w := 1<<31 - 1 ^ v
			if cnt[w] > 0 {
				cnt[w]--
			} else {
				ans++
				cnt[v]++
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1926D(os.Stdin, os.Stdout) }
