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
	h := make(map[byte]bool)
	for i := range s {
		h[s[i]] = true
	}
	for i := '0'; i <= '9'; i++ {
		if !h[byte(i)] {
			Fprintln(out, string(i))
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
