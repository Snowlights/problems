//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF482A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	if k == 1 {
		for i := 1; i <= n; i++ {
			Fprint(out, i, " ")
		}
		return
	}
	k--
	Fprint(out, "1 ", n)
	i := 2
	for k > 2 {
		Fprint(out, " ", i, " ", n+1-i)
		k -= 2
		i++
	}
	if k == 1 {
		for j := n + 1 - i; j >= i; j-- {
			Fprint(out, " ", j)
		}
	} else {
		for j := i; j <= n+1-i; j++ {
			Fprint(out, " ", j)
		}
	}

}

func main() { CF482A(os.Stdin, os.Stdout) }
