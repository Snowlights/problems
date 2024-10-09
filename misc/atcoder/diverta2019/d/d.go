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

	var n, ans int
	Fscan(in, &n)
	for m := 1; (m+1)*m < n; m++ {
		if n%m == 0 {
			ans += n/m - 1
		}
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
