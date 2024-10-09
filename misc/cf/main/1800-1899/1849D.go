//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1849D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := 0; i < n; {
		if a[i] == 0 {
			ans++
			i++
			continue
		}
		st := i
		has2 := false
		for ; i < n && a[i] > 0; i++ {
			if a[i] == 2 {
				has2 = true
			}
		}
		ans++
		if st > 0 && a[st-1] == 0 {
			ans--
			if !has2 {
				continue
			}
		}
		if i < n {
			a[i] = 1
			i++
		}
	}
	Fprint(out, ans)

}

func main() { CF1849D(os.Stdin, os.Stdout) }
