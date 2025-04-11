//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1881D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := map[int]int{}
		for range n {
			Fscan(in, &x)
			for i := 2; i*i <= x; i++ {
				if x%i > 0 {
					continue
				}
				e := 1
				for x /= i; x%i == 0; x /= i {
					e++
				}
				cnt[i] += e
			}
			if x > 1 {
				cnt[x]++
			}
		}
		for _, c := range cnt {
			if c%n > 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}

}

func main() { CF1881D(os.Stdin, os.Stdout) }
