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
		var n, v, ans int
		Fscan(in, &n)
		f := make([]int, 1<<17)

		f[0] = 1
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			for j := len(f) - 1; j >= v; j-- {
				f[j] = f[j] ^ f[j-v]
			}
		}
		for i, v := range f {
			if v == 1 {
				ans ^= i
			}
		}
		Fprintln(out, ans)
	}

	solve()

}

func main() { run(os.Stdin, os.Stdout) }
