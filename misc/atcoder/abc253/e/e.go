package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 前缀和优化 DP。
// 定义 f[i][j] 表示考虑前 i 个数，其中 a[i]=j 的方案数。
// 根据要求，从所有 abs(j-j')  >= k 的 f[i-1][j'] 转移过来，这个和式可以用前缀和优化成 O(1)。
// 初始值 f[0][j] = 1。
// 答案为 sum(f[n-1])。

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k int
	Fscan(in, &n, &m, &k)
	mod := int(998244353)
	f := make([]int, m+1)
	for i := range f {
		f[i] = i
	}
	for i := 1; i < n; i++ {
		newF := make([]int, m+1)
		for j := 0; j < m; j++ {
			if k > 0 {
				newF[j+1] = ((newF[j]+f[m])%mod - (f[min(j+k, m)]-f[max(j-k+1, 0)])%mod) % mod
			} else {
				newF[j+1] = ((newF[j] + f[m]) % mod) % mod
			}
		}
		f = newF
	}

	Fprintln(out, (f[m]+mod)%mod)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
