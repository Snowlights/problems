package c_130

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc130/tasks/abc130_e
//
//输入 n(≤2000) 和 m(≤2000)，长度分别为 n 和 m 的数组 a 和 b，元素范围 [1,1e5]。
//从 a 和 b 中分别选出一个子序列（允许为空），要求这两个子序列相同。
//输出有多少种不同的选法，模 1e9+7。
//注意：选出的子序列不同，当且仅当下标不同（即使子序列的元素是相同的，也算不同）。

// https://atcoder.jp/contests/abc130/submissions/36480640
//定义 f[i][j] 表示在 a 的前 i 个数和 b 的前 j 个数中选择子序列，能得到的答案。
//考虑转移来源，可以是 f[i-1][j] 和 f[i][j-1]，但这两个都包含 f[i-1][j-1]，所以要减掉重复的 f[i-1][j-1]（类比二维前缀和）。
//如果 a[i]=b[j]，那么我们可以把这两加到所有 f[i-1][j-1] 的末尾，再加上 1，即 a[i] 和 b[j] 单独组成子序列。
//因此状态转移方程为：
//f[i][j] = f[i-1][j] + f[i][j-1] - f[i-1][j-1] + (a[i]==b[j] ? f[i-1][j-1]+1 : 0)
//最后答案为 f[n][m]+1（需要把空子序列加上）。

func AtCoderABC130E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	const (
		mod int = 1e9 + 7
	)
	fmt.Fscan(r, &n, &m)
	a, b := make([]int, n), make([]int, m)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	for i := range b {
		fmt.Fscan(r, &b[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i, u := range a {
		for j, v := range b {
			if u == v {
				dp[i+1][j+1] = (dp[i][j+1] + dp[i+1][j] + 1) % mod
			} else {
				dp[i+1][j+1] = (dp[i][j+1] + dp[i+1][j] - dp[i][j] + mod) % mod
			}
		}
	}

	fmt.Fprintln(w, dp[n][m]+1)
}
