//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF414B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	mod := int(1e9 + 7)
	var n, k, ans int
	Fscan(in, &n, &k)
	f := make([]int, n+1)
	f[1] = 1
	for i := 0; i < k; i++ {
		for j := n; j > 0; j-- {
			for l := j * 2; l <= n; l += j {
				f[l] = (f[l] + f[j]) % mod
			}
		}
	}
	for _, v := range f {
		ans = (ans + v) % mod
	}
	Fprintln(out, ans)
}

func main() { CF414B(os.Stdin, os.Stdout) }
