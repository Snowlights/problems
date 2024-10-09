//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF576A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	ans := []int{}
	notPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if notPrime[i] {
			continue
		}
		for v := i; v <= n; v *= i {
			ans = append(ans, v)
		}
		for j := i * i; j <= n; j += i {
			notPrime[j] = true
		}
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}

}

func main() { CF576A(os.Stdin, os.Stdout) }
