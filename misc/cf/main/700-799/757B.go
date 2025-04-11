//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF757B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx int = 1e5
	cnt := [mx + 1]int{}
	var n, v int
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		cnt[v]++
	}
	ans := 1
	for i := 2; i <= mx; i++ {
		s := 0
		for j := i; j <= mx; j += i {
			s += cnt[j]
		}
		ans = max(ans, s)
	}
	Fprint(out, ans)

}

func main() { CF757B(os.Stdin, os.Stdout) }
