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

	var n int
	Fscan(in, &n)
	a, cnt := make([]int, n), make(map[int]int)
	vis := make([]bool, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
		cnt[a[i]] = i
	}
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if cnt[i] == i {
			continue
		}
		for cnt[i] > i {
			before := cnt[i] - 1
			if vis[before] {
				Fprintln(out, -1)
				return
			}
			cnt[i], cnt[a[before]] = before, cnt[i]
			vis[before] = true
			ans = append(ans, before)
		}
	}

	if len(ans) != n-1 {
		Fprintln(out, -1)
		return
	}

	for _, v := range ans {
		Fprintln(out, v+1)
	}

}

func main() { run(os.Stdin, os.Stdout) }
