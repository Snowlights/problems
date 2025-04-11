//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1730D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		cnt := [26]int{}
		cnt2 := [26][26]int{}
		for i, v := range s {
			v -= 'a'
			w := t[n-1-i] - 'a'
			cnt[v] ^= 1
			cnt[w] ^= 1
			cnt2[min(v, w)][max(v, w)] ^= 1
		}
		odd := 0
		for i, row := range cnt2 {
			for _, c := range row[i:] {
				odd += c
			}
		}
		if odd <= n%2 && cnt == [26]int{} {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}

}

func main() { CF1730D(os.Stdin, os.Stdout) }
