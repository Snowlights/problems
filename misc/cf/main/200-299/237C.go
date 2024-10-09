//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF237C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b, k int
	Fscan(in, &a, &b, &k)

	primes := []int{a - 1}
	pid := [1e6 + 1]int{-1, -1}
	for i := 2; i <= b; i++ {
		if pid[i] == 0 {
			if i >= a {
				primes = append(primes, i)
			}
			pid[i] = len(primes)
			for j := i * i; j <= b; j += i {
				pid[j] = -1
			}
		}
	}
	primes = append(primes, b+1)

	if len(primes)-2 < k {
		Fprintln(out, -1)
		return
	}

	ans := 0
	for i := k; i < len(primes); i++ {
		ans = max(ans, primes[i]-primes[i-k])
	}
	Fprintln(out, ans)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF237C(os.Stdin, os.Stdout) }
