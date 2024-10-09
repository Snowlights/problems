//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1237D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n, n*3)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(append(a, a...), a...)
	q := []int{}
	for i, j := 0, 0; i < n; i++ {
		for ; j < n*3 && (len(q) == 0 || a[j]*2 >= a[q[0]]); j++ {
			for len(q) > 0 && a[q[len(q)-1]] <= a[j] {
				q = q[:len(q)-1]
			}
			q = append(q, j)
		}
		if j == n*3 {
			Fprint(out, "-1 ")
		} else {
			Fprint(out, j-i, " ")
		}
		if q[0] == i {
			q = q[1:]
		}
	}

}

func main() { CF1237D(os.Stdin, os.Stdout) }
