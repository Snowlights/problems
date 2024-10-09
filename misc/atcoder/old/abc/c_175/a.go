package c_175

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc175/tasks/abc175_e
//
//输入 n m (1≤n,m≤3000) k(≤min(2e5,r*c))，表示一个 n*m 的网格，和网格中的 k 个物品。
//接下来 k 行，每行三个数 x y v(≤1e9) 表示物品的行号、列号和价值（行列号从 1 开始）。
//每个网格至多有一个物品。
//
//你从 (1,1) 出发走到 (n,m)，每步只能向下或向右。
//经过物品时，你可以选或不选，且每行至多可以选三个物品。
//输出你选到的物品的价值和的最大值。

// https://atcoder.jp/contests/abc175/submissions/36563526
//
//定义 f[i][j][0/1/2/3] 表示从 (1,1) 走到 (i,j)，且当前行选了 0/1/2/3 个物品时的最大价值和。
//转移方程如下：
//f[i][j][0] = max(f[i-1][j])
//f[i][j][1] = max(f[i][j-1][1], f[i][j][0]+a[i][j])
//f[i][j][2] = max(f[i][j-1][2], f[i][j-1][1]+a[i][j])
//f[i][j][3] = max(f[i][j-1][3], f[i][j-1][2]+a[i][j])
//答案为 max(f[n][m])。
//实际计算 max(f[][]) 时，f[][][0] 可以不计入。

func AtCoderABC185E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, k, x, y, v int
	fmt.Fscan(r, &n, &m, &k)
	f, a := make([][][4]int, n+1), make([][]int, n+1)
	for i := range f {
		f[i] = make([][4]int, m+1)
		a[i] = make([]int, m+1)
	}

	for i := 0; i < k; i++ {
		fmt.Fscan(r, &x, &y, &v)
		a[x][y] = v
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j][0] = max(f[i-1][j][1], max(f[i-1][j][2], f[i-1][j][3]))
			f[i][j][1] = max(f[i][j-1][1], f[i][j][0]+a[i][j])
			f[i][j][2] = max(f[i][j-1][2], f[i][j-1][1]+a[i][j])
			f[i][j][3] = max(f[i][j-1][3], f[i][j-1][2]+a[i][j])
		}
	}

	fmt.Fprintln(w, max(max(f[n][m][1], f[n][m][2]), f[n][m][3]))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
