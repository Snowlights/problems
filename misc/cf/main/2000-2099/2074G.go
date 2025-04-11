//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2074G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		f := make([][]int, n)
		for i := range a {
			Fscan(in, &a[i])
			f[i] = make([]int, n)
		}
		for i := n - 3; i >= 0; i-- {
			for j := i + 2; j < n; j++ {
				for k := i; k < j; k++ {
					f[i][j] = max(f[i][j], f[i][k]+f[k+1][j])
				}
				for k := i + 1; k < j; k++ {
					f[i][j] = max(f[i][j], f[i+1][k-1]+f[k+1][j-1]+a[i]*a[k]*a[j])
				}
			}
		}
		Fprintln(out, f[0][n-1])
	}

}

func main() { CF2074G(os.Stdin, os.Stdout) }
