//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1854A1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		var pos, neg, mx, mm int
		for i := range a {
			Fscan(in, &a[i])
			if a[i] > 0 {
				pos++
			} else if a[i] < 0 {
				neg++
			}
			if a[i] > a[mx] {
				mx = i
			} else if a[i] < a[mm] {
				mm = i
			}
		}

		ans := make([][]int, 0)
		makePos := func() {
			for i, v := range a {
				if v < 0 {
					ans = append(ans, []int{i, mx})
				}
			}
			for i := 1; i < n; i++ {
				ans = append(ans, []int{i, i - 1})
			}
		}
		makeNeg := func() {
			for i, v := range a {
				if v > 0 {
					ans = append(ans, []int{i, mm})
				}
			}
			for i := n - 2; i >= 0; i-- {
				ans = append(ans, []int{i, i + 1})
			}
		}
		if a[mx] > -a[mm] {
			if neg < 13 {
				makePos()
			} else {
				for i := 0; i < 5; i++ {
					ans = append(ans, []int{mm, mm})
				}
				makeNeg()
			}
		} else {
			if pos < 13 {
				makeNeg()
			} else {
				for i := 0; i < 5; i++ {
					ans = append(ans, []int{mx, mx})
				}
				makePos()
			}
		}
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprintln(out, v[0]+1, v[1]+1)
		}
	}

}

func main() { CF1854A1(os.Stdin, os.Stdout) }
