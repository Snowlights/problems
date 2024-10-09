//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1365C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, ans int
	Fscan(in, &n)
	m, cnt := make(map[int]int, n), make(map[int]int)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		m[v] = i
	}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		e, _ := m[v]
		if e < i {
			e += n
		}
		cnt[e-i]++
	}
	for _, v := range cnt {
		if v > ans {
			ans = v
		}
	}
	Fprintln(out, ans)
}

func main() { CF1365C(os.Stdin, os.Stdout) }
