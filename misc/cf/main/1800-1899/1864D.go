//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1864D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)

	defer out.Flush()
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		f := make([]byte, n)
		diag1 := make([]byte, n*2)
		diag2 := make([]byte, n*2)
		for i := 0; i < n; i++ {
			Fscan(in, &s)
			for j, c := range s {
				f[j] ^= diag1[i-j+n] ^ diag2[i+j]
				if f[j] != c&1 {
					ans++
					f[j] ^= 1
					diag1[i-j+n] ^= 1
					diag2[i+j] ^= 1
				}
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1864D(os.Stdin, os.Stdout) }
