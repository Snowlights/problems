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

	var n, m, ans int
	Fscan(in, &n, &m)
	cnt := make(map[int]int)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if i < m {
			cnt[a[i]]++
		}
	}
	for cnt[ans] > 0 {
		ans++
	}
	for i, v := range a[m:] {
		cnt[v]++
		cnt[a[i]]--
		if cnt[a[i]] == 0 && a[i] < ans {
			ans = a[i]
		}
	}

	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
