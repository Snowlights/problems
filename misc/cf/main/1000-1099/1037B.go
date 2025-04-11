//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1037B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	m := n / 2
	if a[m] > k {
		for i := m; i >= 0 && a[i] > k; i-- {
			ans += a[i] - k
		}
	} else {
		for i := m; i < n && a[i] < k; i++ {
			ans += k - a[i]
		}
	}
	Fprint(out, ans)

}

func main() { CF1037B(os.Stdin, os.Stdout) }
