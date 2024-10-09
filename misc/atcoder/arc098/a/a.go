package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n, &s)
	pre, suf := make([]int, n+1), make([]int, n+1)
	for i, v := range s {
		pre[i+1] = pre[i]
		if v == 'W' {
			pre[i+1]++
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		suf[i] = suf[i+1]
		if s[i] == 'E' {
			suf[i]++
		}
	}
	ans := math.MaxInt32
	for i, v := range pre {
		if v+suf[i] < ans {
			ans = v + suf[i]
		}
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
