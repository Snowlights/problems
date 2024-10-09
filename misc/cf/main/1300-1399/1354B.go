//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func CF1354B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var (
		T int
		s string
	)
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		l, cnt, ans := 0, make(map[byte]int), math.MaxInt
		for r := range s {
			cnt[s[r]]++
			for len(cnt) == 3 {
				ans = min(ans, r-l+1)
				cnt[s[l]]--
				if cnt[s[l]] == 0 {
					delete(cnt, s[l])
				}
				l++
			}
		}
		if ans == math.MaxInt {
			ans = 0
		}
		Fprintln(out, ans)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1354B(os.Stdin, os.Stdout) }
