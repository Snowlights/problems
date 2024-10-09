//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1955C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		i, j := 0, n-1
		for i < j && k > 0 {
			if a[i] <= a[j] {
				if k < a[i]*2-1 {
					break
				}
				ans++
				k -= a[i] * 2
				a[j] -= a[i]
				i++
				if k >= 0 && a[j] == 0 {
					ans++
					j--
				}
			} else {
				if k < a[j]*2 {
					break
				}
				ans++
				k -= a[j] * 2
				a[i] -= a[j]
				j--
			}
		}
		if i == j && k >= a[i] {
			ans++
		}
		Fprintln(out, ans)
	}

}

func main() { CF1955C(os.Stdin, os.Stdout) }
