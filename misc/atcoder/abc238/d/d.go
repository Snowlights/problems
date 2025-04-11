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

	var T, and, s int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &and, &s)
		or := s - and
		if or >= 0 && or&and == and {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
