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

	var n, k, q, l, r int
	Fscan(in, &n, &k)

	a := make([]int, n)
	sum := make([]int, n+k) // int64
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i, v := range a {
		sum[i+k] = sum[i] + v
	}
	pre := func(x, t int) int {
		if x%k <= t {
			return sum[x/k*k+t]
		}
		return sum[(x+k-1)/k*k+t]
	}
	// 求下标在 [l,r) 范围内且下标模 k 同余于 t 的所有元素之和
	query := func(l, r, t int) int {
		t %= k
		return pre(r, t) - pre(l, t)
	}

loop:
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		l--
		s := query(l, r, 0)
		for i := 1; i < k; i++ {
			if query(l, r, i) != s {
				Fprintln(out, "No")
				continue loop
			}
		}
		Fprintln(out, "Yes")
	}

}

func main() { run(os.Stdin, os.Stdout) }
