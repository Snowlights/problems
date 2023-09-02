//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1215B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, cur, pos, neg int
	Fscan(in, &n)
	cnt := [2]int{1}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v < 0 {
			cur ^= 1
		}
		pos += cnt[cur^1]
		neg += cnt[cur]
		cnt[cur]++
	}
	Fprintln(out, pos, neg)
}

func main() { CF1215B(os.Stdin, os.Stdout) }
