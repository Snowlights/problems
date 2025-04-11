//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2037C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n < 5 {
			Fprintln(out, -1)
			continue
		}
		for i := 1; i <= n; i += 2 {
			if i != 5 {
				Fprint(out, i, " ")
			}
		}
		Fprint(out, "5 4")
		for i := 2; i <= n; i += 2 {
			if i != 4 {
				Fprint(out, " ", i)
			}
		}
		Fprintln(out)
	}

}

func main() { CF2037C(os.Stdin, os.Stdout) }
