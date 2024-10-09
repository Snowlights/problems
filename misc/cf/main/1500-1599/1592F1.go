//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1592F1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	suf := make([][]byte, n+1)
	for i := range suf {
		suf[i] = make([]byte, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			suf[i][j] = suf[i][j+1] ^ suf[i+1][j] ^ suf[i+1][j+1]
			if suf[i][j] == a[i][j]&1 {
				ans++
				a[i][j] = 0
				suf[i][j] ^= 1
			}
		}
	}

	if a[n-1][m-1] == 0 {
		for _, row := range a[:n-1] {
			for j, x := range row[:m-1] {
				if x == 0 && row[m-1] == 0 && a[n-1][j] == 0 {
					Fprint(out, ans-1)
					return
				}
			}
		}
	}
	Fprint(out, ans)

}

func main() { CF1592F1(os.Stdin, os.Stdout) }
