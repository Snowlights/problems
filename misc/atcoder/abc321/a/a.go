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

	var s string
	Fscan(in, &s)
	for i := 1; i < len(s); i++ {
		if s[i] >= s[i-1] {
			Fprintln(out, "No")
			return
		}
	}
	Fprintln(out, "Yes")
}

func main() { run(os.Stdin, os.Stdout) }
