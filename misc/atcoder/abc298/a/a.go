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
	o, x := false, false
	for _, v := range s {
		if v == 'o' {
			o = true
		}
		if v == 'x' {
			x = true
		}
	}

	if o && !x {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}

}

func main() { run(os.Stdin, os.Stdout) }
