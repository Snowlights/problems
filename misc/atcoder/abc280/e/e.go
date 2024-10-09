package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 期望 DP 入门题。
//
// 用 f[i] 表示血量为 i 时的攻击次数的期望。
// 那么 f[i] = p/100 * (f[i-2]+1) + (1-p/100) * (f[i-1]+1)
// 初始值 f[0]=0, f[1]=1。
// 答案为 f[n]。

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	mod := 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var n, p int
	Fscan(in, &n, &p)
	f, inv := make([]int, n+1), pow(100, mod-2)
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = (p*f[i-2] + (100-p)*f[i-1] + 100) % mod * inv % mod
	}

	Fprintln(out, f[n])
}

func main() { run(os.Stdin, os.Stdout) }
