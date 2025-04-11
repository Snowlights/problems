//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF862A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, v, ans int
	Fscan(in, &n, &x)
	has := make([]bool, x)
	for range n {
		Fscan(in, &v)
		if v < x {
			has[v] = true
		} else if v == x {
			ans++
		}
	}
	for _, b := range has {
		if !b {
			ans++
		}
	}
	Fprint(out, ans)

}

func main() { CF862A(os.Stdin, os.Stdout) }
