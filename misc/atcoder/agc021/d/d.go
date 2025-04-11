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

	s, K := "", 0
	Fscan(in, &s, &K)
	n := len(s)
	pre := make([][]int, n)
	f := make([][]int, n)
	for i := range f {
		pre[i] = make([]int, n)
		pre[i][i] = 1
		f[i] = make([]int, n)
		f[i][i] = 1
	}
	for k := 0; k <= K; k++ {
		for i := n - 2; i >= 0; i-- {
			for j := i + 1; j < n; j++ {
				if s[i] == s[j] {
					f[i][j] = f[i+1][j-1] + 2
				} else {
					f[i][j] = max(f[i+1][j], f[i][j-1])
					if k > 0 {
						f[i][j] = max(f[i][j], pre[i+1][j-1]+2)
					}
				}
			}
		}
		pre, f = f, pre
	}
	Fprint(out, pre[0][n-1])

}

func main() { run(os.Stdin, os.Stdout) }
