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
	Fscan(in, &n)
	ans := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			continue
		}
		ans += i
	}

	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
