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

	var n, m int
	Fscan(in, &n, &m)
	a, cnt := make([]int, n), make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	ans, exit := make([]int, 0, m), make(map[int]bool)
	for _, v := range a {
		cnt[v]--
		if exit[v] {
			continue
		}

		for len(ans) > 0 && ans[len(ans)-1] > v && cnt[ans[len(ans)-1]] > 0 {
			exit[ans[len(ans)-1]] = false
			ans = ans[:len(ans)-1]
		}

		ans = append(ans, v)
		exit[v] = true
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}

}

func main() { run(os.Stdin, os.Stdout) }
