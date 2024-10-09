//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF437C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		ans += min(a[v-1], a[w-1])
	}
	Fprint(out, ans)

}

func main() { CF437C(os.Stdin, os.Stdout) }
