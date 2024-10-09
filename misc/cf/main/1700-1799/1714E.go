//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1714E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	mask, i := 0, 1
	vis := make(map[int]bool)
	for {
		if vis[i] {
			break
		}
		mask |= 1 << (i % 20)
		if i > 10 && i%2 == 0 {
			mask |= 1 << (10 + i%10/2)
		}
		vis[i] = true
		i = (i + i%10) % 20
	}
	mask |= 1 << 17
	mask |= 1 << 19

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f, m := -1, 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if v%5 == 0 {
				v += v % 10
				if f == -1 {
					f = v
				} else if v != f {
					f = -2
				}
			} else {
				m |= 1 << (v % 20)
			}
		}
		if m > 0 {
			if f == -1 && (m&mask == 0 || m&mask == m) {
				Fprintln(out, "Yes")
			} else {
				Fprintln(out, "No")
			}
		} else {
			if f != -2 {
				Fprintln(out, "Yes")
			} else {
				Fprintln(out, "No")
			}
		}

	}

}

func main() { CF1714E(os.Stdin, os.Stdout) }
