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

	solve := func() {
		var n, ans int
		cnt := [2]int{}
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if i > 0 {
				a[i] ^= a[i-1]
			}
		}

		for k := 0; k < 32; k++ {
			cnt[0], cnt[1] = 1, 0
			for _, v := range a {
				if v&(1<<k) != 0 {
					v = 1
				} else {
					v = 0
				}
				ans += cnt[v^1] * (1 << k)
				cnt[v]++
			}
		}

		Fprintln(out, ans)
	}

	solve()

}

func main() { run(os.Stdin, os.Stdout) }
