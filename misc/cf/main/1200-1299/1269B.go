//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF1269B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	a0 := a[0]
	for i := range n - 1 {
		a[i] = a[i+1] - a[i]
	}

	pi := make([]int, n-1)
	match := 0
	for i := 1; i < n-1; i++ {
		v := a[i]
		for match > 0 && a[match] != v {
			match = pi[match-1]
		}
		if a[match] == v {
			match++
		}
		pi[i] = match
	}

	b := make([]int, n*2)
	for i := range n {
		Fscan(in, &b[i])
	}
	if n == 1 {
		Fprint(out, (b[0]-a0+m)%m)
		return
	}
	slices.Sort(b[:n])
	for i, v := range b[:n] {
		b[n+i] = v + m
	}
	oriB := slices.Clone(b)
	for i := range n*2 - 1 {
		b[i] = b[i+1] - b[i]
	}

	ans := m
	match = 0
	for i, v := range b[:n*2-1] {
		for match > 0 && a[match] != v {
			match = pi[match-1]
		}
		if a[match] == v {
			match++
		}
		if match == n-1 {
			res := oriB[i-(n-1)+1] - a0
			ans = min(ans, (res%m+m)%m)
			match = pi[match-1]
		}
	}
	Fprint(out, ans)

}

func main() { CF1269B(os.Stdin, os.Stdout) }
