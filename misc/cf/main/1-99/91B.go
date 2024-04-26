//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF91B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := make([]int, n)
	ans[n-1] = -1
	st := []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		v := a[i]
		j := sort.Search(len(st), func(j int) bool { return a[st[j]] < v })
		if j < len(st) {
			ans[i] = st[j] - i - 1
		} else {
			ans[i] = -1
		}
		if v < a[st[len(st)-1]] {
			st = append(st, i)
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}

}

func main() { CF91B(os.Stdin, os.Stdout) }
