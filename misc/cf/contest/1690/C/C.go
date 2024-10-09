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

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s, e := make([]int, n), make([]int, n)
		for i := range s {
			Fscan(in, &s[i])
		}
		for i := range e {
			Fscan(in, &e[i])
		}
		cur := s[0]
		for i, v := range s {
			Fprint(out, e[i]-max(v, cur), " ")
			cur = e[i]
		}
		Fprintln(out)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
