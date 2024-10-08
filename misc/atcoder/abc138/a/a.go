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

	var n int
	var s string
	Fscan(in, &n, &s)
	if n >= 3200 {
		Fprintln(out, s)
	} else {
		Fprintln(out, "red")
	}

}

func main() { run(os.Stdin, os.Stdout) }
