package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, l, r int
	Fscan(in, &n)
	pos := make([][]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		pos[v] = append(pos[v], i)
	}

	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &l, &r, &v)
		Fprintln(out, sort.SearchInts(pos[v], r)-sort.SearchInts(pos[v], l-1))
	}

}

func main() { run(os.Stdin, os.Stdout) }
