//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2074C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		if x&1 == 0 {
			y = x&-x | 1
		} else {
			y = (x+1)&^x | 1
		}
		if y >= x {
			y = -1
		}
		Fprintln(out, y)
	}

}

func main() { CF2074C(os.Stdin, os.Stdout) }
