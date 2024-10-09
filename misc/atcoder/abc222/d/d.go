package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	// f[i][j] = 前 i 个数，其中第 i 个数填 <=j 的方案数
	// f[i][j] = f[i][j-1] + (恰好填 j） f[i-1][min(j, b[i])]
	f := make([][3001]int, n)
	for j := a[0]; j <= b[0]; j++ {
		f[0][j] = j - a[0] + 1
	}
	for i := 1; i < n; i++ {
		low, up := a[i], b[i]
		for j := low; j <= up; j++ {
			f[i][j] = (f[i][max(0, j-1)] + f[i-1][min(j, b[i-1])]) % mod
		}
	}
	Fprint(out, f[n-1][b[n-1]])

}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() { run(os.Stdin, os.Stdout) }
