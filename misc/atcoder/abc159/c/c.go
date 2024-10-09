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

	var n float64
	Fscan(in, &n)
	Fprintln(out, Sprintf("%.12f", (n/3)*(n/3)*(n/3)))
}

func main() { run(os.Stdin, os.Stdout) }
