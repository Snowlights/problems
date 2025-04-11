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

	var n, k, d, v int
	Fscan(in, &n, &k, &d)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, d)
		for j := range f[i] {
			f[i][j] = -1e18
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		for j := min(i, k); j > 0; j-- {
			for s := 0; s < d; s++ {
				f[j][s] = max(f[j][s], f[j-1][((s-v)%d+d)%d]+v)
			}
		}
	}
	ans := f[k][0]
	if ans < 0 {
		ans = -1
	}
	Fprint(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
