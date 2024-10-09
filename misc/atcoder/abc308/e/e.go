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
	var s string
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &s)

	h := map[int]int{
		1:               1,
		1 | 1<<2:        1,
		1 | 1<<1:        2,
		1 | 1<<1 | 1<<2: 3,
	}

	pre, suf := make([]int, 3), make([]int, 3)
	for i, v := range s {
		switch v {
		case 'X':
			suf[a[i]]++
		}
	}
	for i, v := range s {
		switch v {
		case 'M':
			pre[a[i]]++
		case 'E':
			for j, vj := range pre {
				for k, vk := range suf {
					ans += vj * vk * (h[1<<k|1<<j|1<<a[i]])
				}
			}
		case 'X':
			suf[a[i]]--
		}
	}
	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
