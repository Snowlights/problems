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

	var n, m int
	Fscan(in, &n, &m)
	Fprintln(out, n*(n-1)/2+m*(m-1)/2)

}

func main() { run(os.Stdin, os.Stdout) }
