//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF651B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	cnt := [1001]int{}
	var n, v int
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		cnt[v]++
	}
	Fprint(out, n-slices.Max(cnt[:]))

}

func main() { CF651B(os.Stdin, os.Stdout) }
