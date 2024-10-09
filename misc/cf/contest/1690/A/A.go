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

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%3 == 0 {
			Fprintln(out, n/3, n/3+1, n/3-1)
		} else if n%3 == 1 {
			Fprintln(out, n/3, n/3+2, n/3-1)
		} else if n%3 == 2 {
			Fprintln(out, n/3+1, n/3+2, n/3-1)
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
