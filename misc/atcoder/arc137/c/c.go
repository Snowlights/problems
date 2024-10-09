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

	var n, x, y int
	Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &x)
	}
	Fscan(in, &y)

	if x+1 < y || (y-n)%2 == 0 {
		Fprintln(out, "Alice")
	} else {
		Fprintln(out, "Bob")
	}

}

func main() { run(os.Stdin, os.Stdout) }
