//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func CF1677B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		Fscan(in, &s)
		col := make(map[int]bool, m)
		f := make([]int, m)
		window := strings.Count(s[:m], "1")
		for i, v := range s {
			if i < m {
				if i == 0 {
					f[i] = int(v & 1)
				} else {
					f[i] = f[i-1] | int(v&1)
				}
			} else {
				window += int(v&1) - int(s[i-m]&1)
				if window > 0 {
					f[i%m]++
				}
			}
			if v == '1' && !col[i%m] {
				col[i%m] = true
			}
			Fprint(out, f[i%m]+len(col), " ")
		}
		Fprintln(out)
	}

}

func main() { CF1677B(os.Stdin, os.Stdout) }
