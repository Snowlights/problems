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

	s := ""
	Fscan(in, &s)
	n := len(s)
	ans := n
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans = min(ans, max(i, n-i))
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
