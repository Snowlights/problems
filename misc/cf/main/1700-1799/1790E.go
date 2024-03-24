//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1790E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := n >> 1
		if n&1 > 0 || n&ans > 1 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, n|ans, ans)
		}
	}

}

func main() { CF1790E(os.Stdin, os.Stdout) }
