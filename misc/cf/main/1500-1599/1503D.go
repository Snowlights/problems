//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF1503D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	to := make([]int, n*2+1)
	pos := make([]int, n*2+1)
	ps := make([]struct{ x, y int }, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
		pos[ps[i].y] = 1
		to[ps[i].x] = ps[i].y
		to[ps[i].y] = ps[i].x
	}

	a := make([]int, n)
	al, ar := 0, n-1
	b := make([]int, n)
	bl, br := 0, n-1
	vis := make([]bool, n*2+2)
	mn, mx := 1, n*2
	for {
		for vis[mn] {
			mn++
		}
		if mn > n*2 {
			break
		}

		cnt := [2]int{}
		t := mn
		for {
			last := 0
			for ; mn <= t; mn++ {
				if vis[mn] {
					continue
				}
				cnt[pos[mn]]++

				vis[mn] = true
				a[al] = mn
				al++

				last = to[mn]
				vis[last] = true
				b[bl] = last
				bl++
			}

			if last == 0 {
				break
			}

			t = last
			last = 0
			for ; mx >= t; mx-- {
				if vis[mx] {
					continue
				}
				cnt[pos[mx]]++

				vis[mx] = true
				a[ar] = mx
				ar--

				last = to[mx]
				vis[last] = true
				b[br] = last
				br--
			}

			if last == 0 {
				break
			}
			t = last
		}
		ans += min(cnt[0], cnt[1])
	}
	slices.Reverse(b)
	if !slices.IsSorted(a) || !slices.IsSorted(b) {
		Fprint(out, -1)
		return
	}
	Fprint(out, ans)

}

func main() { CF1503D(os.Stdin, os.Stdout) }
