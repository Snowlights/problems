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

	var a, b, k, ans int
	Fscan(in, &a, &b, &k)
	for a < b {
		a = a * k
		ans++
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
