package c_106

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc106/tasks/abc106_d
//
//输入 n(≤500) m(≤2e5) q(≤1e5)。
//然后输入 m 个在 [1,n] 内的闭区间，即每行输入两个数表示闭区间 [L,R]。
//然后输入 q 个询问，每个询问也是输入两个数表示闭区间 [p,q]。
//对每个询问，输出在 [p,q] 内的完整闭区间的个数。

func AtCoder106D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, q, left, right int
	fmt.Fscan(r, &n, &m, &q)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(r, &left, &right)
		f[left][right]++
	}

	for i := n; i >= 0; i-- {
		for j := i + 1; j <= n; j++ {
			f[i][j] += f[i+1][j] + f[i][j-1] - f[i+1][j-1]
		}
	}

	for i := 0; i < q; i++ {
		fmt.Fscan(r, &left, &right)
		fmt.Fprintln(w, f[left][right])
	}

}
