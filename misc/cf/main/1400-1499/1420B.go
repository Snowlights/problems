//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1420B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		cnt := [30]int{}
		for range n {
			Fscan(in, &v)
			b := bits.Len(v) - 1
			ans += cnt[b]
			cnt[b]++
		}
		Fprintln(out, ans)
	}

}

func main() { CF1420B(os.Stdin, os.Stdout) }
