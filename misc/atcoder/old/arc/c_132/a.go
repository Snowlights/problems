package c_132

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc132/tasks/arc132_c
//
//输入 n(≤500) d(≤5) 和长为 n 的数组 a。
//a 原本是一个 1~n 的排列 p，不过有些数字被替换成了 -1。
//你需要还原 p，使得 abs(p[i]-i) ≤ d 对每个 i 都成立。
//输出有多少个这样的 p，模 998244353。

// https://atcoder.jp/contests/arc132/submissions/37103023
//定义 f[i][mask] 表示前 i 个数字构成的集合为 mask 时的方案数。
//枚举第 i 个数字要填的数字 j，转移给 f[i+1][mask|(1<<j)]。
//初始 f[0][0]=1，答案为 sum(f[n])。
//你以为 mask 有 2^n-1 那么大？
//注意到第 i 个数字能填的范围是有限的，所以 mask 至多是 2^(2d)-1。

func AtCoderARC132C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const mod = 998244353
	var n, d, v, ans int
	fmt.Fscan(r, &n, &d)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 1<<(d*2))
	}
	f[0][0] = 1
	for i, fs := range f[:n] {
		fmt.Fscan(r, &v)
		for mask, res := range fs {
			for j := max(i-d, 0); j <= min(i+d, n-1); j++ {
				if (v == -1 || j == v-1) && mask>>(j-i+d)&1 == 0 {
					m := (mask | 1<<(j-i+d)) >> 1
					f[i+1][m] = (f[i+1][m] + res) % mod
				}
			}
		}
	}
	for _, res := range f[n] {
		ans += res
	}
	fmt.Fprintln(w, ans%mod)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
