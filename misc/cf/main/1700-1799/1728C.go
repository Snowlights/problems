//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1728C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		small := [10]int{}
		big := map[int]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if v < 10 {
				small[v]++
			} else {
				big[v]++
			}
		}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if v < 10 {
				small[v]--
			} else {
				big[v]--
			}
		}

		ans := 0
		for v, c := range big {
			ans += abs(c)
			k := 1
			for ; v > 9; v /= 10 {
				k++
			}
			small[k] += c
		}
		for _, c := range small[2:] {
			ans += abs(c)
		}
		Fprintln(out, ans)
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() { CF1728C(os.Stdin, os.Stdout) }
