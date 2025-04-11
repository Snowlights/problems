//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1993B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := [2][]int{}
		for range n {
			Fscan(in, &v)
			a[v&1] = append(a[v&1], v)
		}
		if len(a[1]) == 0 {
			Fprintln(out, 0)
			continue
		}
		mx := slices.Max(a[1])
		slices.Sort(a[0])
		ans := len(a[0])
		for _, v := range a[0] {
			if v > mx {
				ans++
				break
			}
			mx += v
		}
		Fprintln(out, ans)
	}

}

func main() { CF1993B(os.Stdin, os.Stdout) }
