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

	const mod int = 1e9 + 7
	var n, m, v, x, y int
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		x = (x + (2*i-n+1)*v) % mod
	}
	for i := 0; i < m; i++ {
		Fscan(in, &v)
		y = (y + (2*i-m+1)*v) % mod
	}
	Fprintln(out, x%mod*y%mod)

}

func main() { run(os.Stdin, os.Stdout) }
