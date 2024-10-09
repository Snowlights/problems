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

	var r, g, b, n, ans int
	Fscan(in, &r, &g, &b, &n)

	for i := 0; i*r <= n; i++ {
		for j := 0; j*g+i*r <= n; j++ {
			if (n-j*g-i*r)%b == 0 {
				ans++
			}
		}
	}

	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
