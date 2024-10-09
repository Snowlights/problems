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

	var n, m, k int
	var s string
	Fscan(in, &n, &m, &k, &s)
	pos := make([]int, 0)
	for i, v := range s {
		if v == 'x' {
			pos = append(pos, i)
		}
	}

	cntX := len(pos)
	ans := min(k/cntX*n+pos[k%cntX], n*m)
	for _, p := range pos {
		k++
		res := min(k/cntX*n+pos[k%cntX], n*m) - p - 1
		ans = max(ans, res)
	}
	Fprint(out, ans)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() { run(os.Stdin, os.Stdout) }
