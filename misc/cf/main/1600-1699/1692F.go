//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1692F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt, a := make([]int, 10), make([]int, 0, 30)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			v %= 10
			if cnt[v] < 3 {
				cnt[v]++
				a = append(a, v)
			}
		}
		for i, x := range a {
			for j, y := range a[:i] {
				for _, z := range a[:j] {
					if (x+y+z)%10 == 3 {
						Fprintln(out, "YES")
						continue o
					}
				}
			}
		}
		Fprintln(out, "NO")
	}

}

func main() { CF1692F(os.Stdin, os.Stdout) }
