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
	Fscan(in, &n)

	Fprintln(out, (n+99)*100)
	for j := 0; j < 100; j++ {
		for i := 0; i < n; i++ {
			Fprintln(out, i)
		}
		for i := 0; i < 99; i++ {
			Fprintln(out, 0)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
