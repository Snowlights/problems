//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1584F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := [123][10][]int{}
		for i := 0; i < n; i++ {
			pos[0][i] = []int{-1}
			Fscan(in, &s)
			for j, b := range s {
				pos[b][i] = append(pos[b][i], j)
			}
		}

		memo := make([][123]int, 1<<n)
		for i := range memo {
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		type pair struct {
			mask int
			c    byte
		}
		from := make([][123]pair, 1<<n)
		var dfs func(int, byte) int
		dfs = func(mask int, c byte) (res int) {
			p := &memo[mask][c]
			if *p != -1 {
				return *p
			}
			ps := pos[c][:n]
			mxMask := 0
			var mxC byte
			for ch := byte('A'); ch <= 'z'; {
				mask2 := 0
				var r int
				for i, ps2 := range pos[ch][:n] {
					if ps2 == nil {
						goto nxt
					}
					if ps2[0] > ps[i][mask>>i&1] {
						// 0
					} else if len(ps2) > 1 && ps2[1] > ps[i][mask>>i&1] {
						mask2 |= 1 << i
					} else {
						goto nxt
					}
				}
				r = dfs(mask2, ch)
				if r > res {
					res = r
					mxC = ch
					mxMask = mask2
				}
			nxt:
				if ch == 'Z' {
					ch = 'a'
				} else {
					ch++
				}
			}
			from[mask][c] = pair{mxMask, mxC}
			res++
			*p = res
			return
		}
		Fprintln(out, dfs(0, 0)-1)

		lcs := []byte{}
		for p := from[0][0]; p.c > 0; p = from[p.mask][p.c] {
			lcs = append(lcs, p.c)
		}
		Fprintf(out, "%s\n", lcs)
	}
}

func main() { CF1584F(os.Stdin, os.Stdout) }
