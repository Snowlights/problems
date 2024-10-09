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

	var n, v int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v%2 == 0 && v%3 != 0 && v%5 != 0 {
			Fprintln(out, "DENIED")
			return
		}
	}
	Fprintln(out, "APPROVED")
}

func main() { run(os.Stdin, os.Stdout) }
