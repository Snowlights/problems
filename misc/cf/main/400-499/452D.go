//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF452D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var k, n1, n2, n3, t1, t2, t3, finish int
	Fscan(in, &k, &n1, &n2, &n3, &t1, &t2, &t3)
	f1 := make([]int, n1)
	f2 := make([]int, n2)
	f3 := make([]int, n3)
	for i := 0; i < k; i++ {
		finish = max(max(f1[i%n1]+t1+t2+t3, f2[i%n2]+t2+t3), f3[i%n3]+t3)
		f1[i%n1] = finish - t2 - t3
		f2[i%n2] = finish - t3
		f3[i%n3] = finish
	}
	Fprint(out, finish)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() { CF452D(os.Stdin, os.Stdout) }
