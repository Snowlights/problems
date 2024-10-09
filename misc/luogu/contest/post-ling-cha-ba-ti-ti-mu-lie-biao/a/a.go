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

	solve := func() {
		var n, v, ans int
		Fscan(in, &n)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			ans += (i + 1) * (n - i) * v
		}
		Fprintln(out, ans)
	}

	solve()

}

func main() { run(os.Stdin, os.Stdout) }
