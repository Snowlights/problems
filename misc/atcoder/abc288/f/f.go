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

	const mod int = 998244353

	var n, f, sum int
	var s string
	Fscan(in, &n, &s)
	sum = 1
	for _, c := range s {
		f = (10*f + sum*int(c-'0')) % mod
		sum = (sum + f) % mod
	}

	Fprintln(out, f)
}

func main() { run(os.Stdin, os.Stdout) }
