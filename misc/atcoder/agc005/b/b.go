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

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n, n+1)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(a, -1)
	st := []int{-1}
	for r, x := range a {
		for len(st) > 1 && a[st[len(st)-1]] >= x {
			i := st[len(st)-1]
			st = st[:len(st)-1]
			ans += a[i] * (i - st[len(st)-1]) * (r - i)
		}
		st = append(st, r)
	}
	Fprint(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
