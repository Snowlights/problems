//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1198B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, op, x, v int
	Fscan(in, &n)
	a, changeTime := make([]int, n), make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &n)
	sufMax := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &op)
		switch op {
		case 1:
			Fscan(in, &x, &v)
			x--
			a[x], changeTime[x] = v, i
		case 2:
			Fscan(in, &v)
			sufMax[i] = v
		}
	}

	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i], sufMax[i+1])
	}
	for i, v := range a {
		Fprint(out, max(v, sufMax[changeTime[i]]), " ")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF1198B(os.Stdin, os.Stdout) }
