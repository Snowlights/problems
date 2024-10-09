package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s, &s)
	a, b := strings.Count(s, "T"), strings.Count(s, "A")
	if a > b || a == b && s[len(s)-1] == 'A' {
		Fprintln(out, "T")
	} else {
		Fprintln(out, "A")
	}
}

func main() { run(os.Stdin, os.Stdout) }
