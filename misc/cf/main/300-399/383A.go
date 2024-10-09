//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF383A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, cnt1, ans int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if v > 0 {
			cnt1++
		} else {
			ans += cnt1
		}
	}
	Fprint(out, ans)
}

func main() { CF383A(os.Stdin, os.Stdout) }
