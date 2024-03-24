//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF750C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, score, sum, div int
	mm, mx := int(-1e9), int(1e9)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &score, &div)
		switch div {
		case 1:
			mm = max(mm, 1900-sum)
		case 2:
			mx = min(mx, 1899-sum)
		}
		sum += score
	}

	if mm > mx {
		Fprintln(out, "Impossible")
	} else if mx == int(1e9) {
		Fprintln(out, "Infinity")
	} else {
		Fprintln(out, mx+sum)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF750C(os.Stdin, os.Stdout) }
