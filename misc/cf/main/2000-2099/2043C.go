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

func dp43(a []int) (minF, maxF int) {
	f, g := 0, 0
	for _, v := range a {
		f = min(f, 0) + v
		g = max(g, 0) + v
		minF = min(minF, f)
		maxF = max(maxF, g)
	}
	return
}

func CF2043C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		idx := 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] < -1 || a[i] > 1 {
				idx = i
			}
		}

		ans := []int{}
		add := func(l, r int) {
			for i := l; i <= r; i++ {
				ans = append(ans, i)
			}
		}

		add(dp43(a[:idx]))
		add(dp43(a[idx+1:]))
		var l, minL, maxL int
		for i := idx - 1; i >= 0; i-- {
			l += a[i]
			minL = min(minL, l)
			maxL = max(maxL, l)
		}
		var r, minR, maxR int
		for _, v := range a[idx+1:] {
			r += v
			minR = min(minR, r)
			maxR = max(maxR, r)
		}
		add(minL+a[idx]+minR, maxL+a[idx]+maxR)
		slices.Sort(ans)
		ans = slices.Compact(ans)

		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}

}

func main() { CF2043C(os.Stdin, os.Stdout) }
