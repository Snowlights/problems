package c_185

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc185/tasks/abc185_e
//
//输入 n(≤1000) 和 m(≤1000)，长度分别为 n 和 m 的数组 a 和 b，元素范围 [1,1e9]。
//从 a 中移除若干元素，得到一个子序列 a'；从 b 中移除若干元素，得到一个子序列 b'。
//要求 a' 和 b' 的长度相同。
//输出 (a和b总共移除的元素个数) + (a'[i]≠b'[i]的i的个数) 的最小值。

// https://atcoder.jp/contests/abc185/submissions/36352936
//
//类似 LCS，定义 f[i][j] 表示 a 的前 i 个元素和 b 的前 j 的元素算出的答案。
//
//- 不选 a[i] 选 b[j]：f[i-1][j]+1
//- 选 a[i] 不选 b[j]：f[i][j-1]+1
//- 选 a[i] 选 b[j]：f[i-1][j-1] + (a[i]==b[j] ? 0 : 1)
//取最小值。
//注：都不选是不用考虑的，这已经包含在 f[i-1][j] 或者 f[i][j-1] 中了。
//也可以这么想：都不选是不如都选的。
//边界：f[0][i]=f[i][0]=i。
//答案：f[n][m]。

func AtCoderABC185E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	a, b := make([]int, n+1), make([]int, m+1)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &a[i])
		f[i][0] = i
	}
	for i := 1; i <= m; i++ {
		fmt.Fscan(r, &b[i])
		f[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i] == b[j] {
				f[i][j] = min(min(f[i-1][j], f[i][j-1])+1, f[i-1][j-1])
			} else {
				f[i][j] = min(min(f[i-1][j], f[i][j-1]), f[i-1][j-1]) + 1
			}
		}
	}
	fmt.Fprintln(w, f[n][m])

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
