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

	var a, b int
	Fscan(in, &a, &b)
	if a+b >= 15 && b >= 8 {
		Fprintln(out, 1)
	} else if a+b >= 10 && b >= 3 {
		Fprintln(out, 2)
	} else if a+b >= 3 {
		Fprintln(out, 3)
	} else {
		Fprintln(out, 4)
	}

}

func main() { run(os.Stdin, os.Stdout) }
