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

	var a, b, c int
	Fscan(in, &a, &b, &c)
	if a == b && a != c ||
		a == c && a != b || b == c && a != b {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}

}

func main() { run(os.Stdin, os.Stdout) }
