package c_203

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://atcoder.jp/contests/abc203/tasks/abc203_d
//
//输入 n k (1≤k≤n≤800) 和一个 n*n 的矩阵，元素范围 [0,1e9]。
//定义 k*k 子矩阵的中位数为子矩阵的第 floor(k*k/2)+1 大的数。
//输出中位数的最小值。
//注：「第 x 大」中的 x 从 1 开始。

func AtCoderABC203E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k int
	fmt.Fscan(r, &n, &k)
	a, pre := make([][]int, n), make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			fmt.Fscan(r, &a[i][j])
		}
	}
	for i := range pre {
		pre[i] = make([]int, n+1)
	}

	fmt.Fprintln(w, sort.Search(1e9, func(ans int) bool {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1]
				if a[i-1][j-1] <= ans {
					pre[i][j]++
				}
				if i >= k && j >= k && pre[i][j]-pre[i-k][j]-
					pre[i][j-k]+pre[i-k][j-k] >= (k*k+1)/2 {
					return true
				}
			}
		}
		return false
	}))

}
