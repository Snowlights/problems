package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 枚举子序列的长度。
//
//考虑子序列长度固定为 m 时，有多少个平均值为整数的子序列。
//相当于子序列的元素和模 m 为 0。
//
//用选或不选来思考。
//定义 f[i][j][k] 表示从前 i 个数中选 j 个数，元素和模 m 为 k 的方案数。
//
//为方便计算取模，用刷表法（用查表法的话，需要算 (k-a[i])%m，可能会算出负数）：
//f[i][j][(k+a[i])%m] = f[i-1][j][(k+a[i])%m] + f[i-1][j-1][k]
//
//答案为 f[n][m][0]。
//
//代码实现时，第一个维度可以去掉，然后像 0-1 背包那样倒序循环 j。初始值 f[0][0] = 1。

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	mod := 998244353
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	f, ans := make([][]int, n+1), 0
	for m := 1; m <= n; m++ {
		for i := range f {
			f[i] = make([]int, n+1)
		}
		f[0][0] = 1

		for _, v := range a {
			for j := m; j > 0; j-- {
				for k := 0; k < m; k++ {
					f[j][(k+v)%m] = (f[j][(k+v)%m] + f[j-1][k]) % mod
				}
			}
		}
		ans = (ans + f[m][0]) % mod
	}

	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
