//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF650A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	type pair struct{ x, y int }
	xx, yy, ph := make(map[int]int), make(map[int]int), make(map[pair]int)

	var n, x, y, ans int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x, &y)
		p := pair{x, y}
		ans += xx[x] + yy[y] - ph[p]
		xx[x]++
		yy[y]++
		ph[p]++
	}
	Fprint(out, ans)
	
}

func main() { CF650A(os.Stdin, os.Stdout) }
