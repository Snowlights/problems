//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1180B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, idx int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		if a[i] >= 0 {
			a[i] = -a[i] - 1
		}
		if a[i] < a[idx] {
			idx = i
		}
	}
	if n%2 == 1 {
		a[idx] = -a[idx] - 1
	}

	for _, v := range a {
		Fprint(out, v, " ")
	}
}

func main() { CF1180B(os.Stdin, os.Stdout) }
