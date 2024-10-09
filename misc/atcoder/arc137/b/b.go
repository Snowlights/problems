package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, inc, dec, mInc, mDec int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		inc = max(inc, 0) - v*2 + 1
		mInc = max(mInc, inc)
		dec = max(dec, 0) + v*2 - 1
		mDec = max(mDec, dec)
	}
	Fprintln(out, mInc+mDec+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() { run(os.Stdin, os.Stdout) }
