//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2021B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		cnt := map[int]int{}
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}
		for i := 0; ; i++ {
			if cnt[i] == 0 {
				Fprintln(out, i)
				break
			}
			cnt[i+x] += cnt[i] - 1
		}
	}

}

func main() { CF2021B(os.Stdin, os.Stdout) }
