package c_173

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc173/tasks/abc173_f
// 图的连通块
//输入 n (2≤n≤2e5) 和一棵树的 n-1 条边（节点编号从 1 开始）。
//定义 f(L,R) 表示用节点编号在 [L,R] 内的点组成的连通块的个数（边的两个端点必须都在 [L,R] 内）。
//输出满足 1≤L≤R≤n 的所有 f(L,R) 的和。

func AtCoderABC173F(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, from, to, sum int64
	fmt.Fscan(r, &n)
	// 长为i的线段的数量是(n-i+1)个
	for i := int64(1); i <= n; i++ {
		sum += i * (n - i + 1)
	}
	// l，r 区间内的连通块数量是
	// 比如5， 9， 至少包含5-9的子段都会-1
	// 从左边数，可以是1，2，3，4，5
	// 从右边数，可以是9，10，11，12...n
	for i := int64(0); i < n-1; i++ {
		fmt.Fscan(r, &from, &to)
		if from > to {
			from, to = to, from
		}
		sum -= from * (n - to + 1)
	}
	fmt.Fprintln(w, sum)
}
