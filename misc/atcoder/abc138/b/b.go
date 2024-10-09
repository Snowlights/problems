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
	var v, ans float64
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		ans += 1 / v
	}
	Fprintln(out, 1/ans)
}

func main() { run(os.Stdin, os.Stdout) }
