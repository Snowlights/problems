//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF940B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, a, b, ans int
	Fscan(in, &n, &k, &a, &b)
	for n != 1 {
		ans += (n % k) * a
		n -= n % k
		if (n-n/k)*a < b {
			ans += (n - 1) * a
			break
		}
		ans += b
		n /= k
	}
	Fprintln(out, ans)
}

func main() { CF940B(os.Stdin, os.Stdout) }
