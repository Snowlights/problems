//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1920E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 998244353
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, k)
		}
		for j := range f[0] {
			f[0][j] = 1
		}
		for i := 1; i <= n; i++ {
			for j := 0; j < k; j++ {
				for x := 0; x+j < k && (x+1)*(j+1) <= i; x++ {
					f[i][j] = (f[i][j] + f[i-(x+1)*(j+1)][x]) % mod
				}
			}
		}
		ans := 0
		for _, v := range f[n] {
			ans += v
		}
		Fprintln(out, ans%mod)
	}

}

func main() { CF1920E(os.Stdin, os.Stdout) }
