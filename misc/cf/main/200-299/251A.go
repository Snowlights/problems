//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF251A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, d, l, ans int
	Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		for a[l] < a[i]-d {
			l++
		}
		ans += (i - l) * (i - l - 1) / 2
	}
	Fprint(out, ans)

}

func main() { CF251A(os.Stdin, os.Stdout) }
