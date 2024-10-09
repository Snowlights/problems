//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1879D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var n, v, xor, ans int
	Fscan(in, &n)
	cnt := [30][2]int{}
	for i := range cnt {
		cnt[i][0] = 1
	}
	sum := [30][2]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		xor ^= v
		for j := 0; j < 30; j++ {
			b := xor >> j & 1
			ans = (ans + (i*cnt[j][b^1]-sum[j][b^1])%mod<<j) % mod
			cnt[j][b]++
			sum[j][b] += i
		}
	}
	Fprint(out, ans)

}

func main() { CF1879D(os.Stdin, os.Stdout) }
