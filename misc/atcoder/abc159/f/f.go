package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

//下标从 0 开始。
//定义 f[i][j] 表示子数组右端点为 i（左端点任意），子序列和为 j 的方案数。
//本题要求的答案就是 f[0][s]+f[1][s]+...+f[n-1][s]。
//考虑 a[i] 选或不选，有
//f[i][j] = f[i-1][j] + f[i-1][j-a[i]]
//初始值：f[i][0] = i+1。例如 f[1][0] 表示子数组 {a[1]} 中有 1 个子序列和为 0，子数组 {a[0],a[1]} 中有 1 个子序列和为 0，所以 f[1][0]=2。

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int = 998244353
	var n, s int
	Fscan(in, &n, &s)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f, ans := make([]int, s+1), 0
	for _, v := range a {
		f[0]++
		for j := s; j >= v; j-- {
			f[j] = (f[j] + f[j-v]) % mod
		}
		ans = (ans + f[s]) % mod
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
