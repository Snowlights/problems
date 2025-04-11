//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF742B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, v, ans int
	Fscan(in, &n, &x)
	cnt := [1 << 17]int{}
	for range n {
		Fscan(in, &v)
		ans += cnt[v^x]
		cnt[v]++
	}
	Fprint(out, ans)

}

func main() { CF742B(os.Stdin, os.Stdout) }
