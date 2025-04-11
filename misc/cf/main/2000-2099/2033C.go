//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2033C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+3)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		res := 0
		for i := 1; i+i < n; i++ {
			v1, v2 := 0, 0
			if a[i] == a[i+1] {
				v1++
			}
			if a[n-i] == a[n-i+1] {
				v1++
			}
			if a[i] == a[n-i] {
				v2++
			}
			if a[i+1] == a[n-i+1] {
				v2++
			}
			res += min(v1, v2)
		}
		if n%2 == 0 && a[n/2] == a[n/2+1] {
			res++
		}
		Fprintln(out, res)
	}

}

func main() { CF2033C(os.Stdin, os.Stdout) }
