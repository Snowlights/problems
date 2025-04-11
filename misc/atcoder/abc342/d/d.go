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

	const mx int = 2e5 + 1
	core := [mx]int{}
	for i := 1; i < mx; i++ {
		if core[i] == 0 {
			for j := 1; i*j*j < mx; j++ {
				core[i*j*j] = i
			}
		}
	}

	cnt := [mx]int{}
	var n, v, ans int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v == 0 {
			ans += i
			cnt[0]++
		} else {
			v = core[v]
			ans += cnt[v] + cnt[0]
			cnt[v]++
		}
	}
	Fprint(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
