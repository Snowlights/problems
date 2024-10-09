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

	var n, a, b int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &a, &b)
		Fprintln(out, a+b)
	}

}

func main() { run(os.Stdin, os.Stdout) }
