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
	Fscan(in, &s)

	if strings.Contains(s, "7") {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}

}

func main() { run(os.Stdin, os.Stdout) }
