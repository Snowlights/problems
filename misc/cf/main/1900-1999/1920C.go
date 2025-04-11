//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1920C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		f := func(d int) {
			g := 0
			for i := 0; i < d; i++ {
				for j := i + d; j < n; j += d {
					g = gcd(g, a[j]-a[j-d])
				}
			}
			if g != 1 && g != -1 {
				ans++
			}
		}
		for d := 1; d*d <= n; d++ {
			if n%d == 0 {
				f(d)
				if d*d < n {
					f(n / d)
				}
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1920C(os.Stdin, os.Stdout) }
