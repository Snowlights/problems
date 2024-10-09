//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1875C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		n %= m
		k := gcd(n, m)
		m /= k
		if m&(m-1) > 0 {
			Fprintln(out, -1)
			continue
		}
		n /= k
		ans := 0
		for n > 0 {
			ans += n
			n = n * 2 % m
		}
		Fprintln(out, ans*k)
	}

}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }

func main() { CF1875C(os.Stdin, os.Stdout) }
