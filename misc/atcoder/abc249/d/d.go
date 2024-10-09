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

	const mx = 1e5
	var n, v, ans int
	cnt := [mx + 1]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
	}
	for i := 1; i <= mx; i++ {
		for j := 1; i*j <= mx; j++ {
			ans += cnt[i] * cnt[j] * cnt[i*j]
		}
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
