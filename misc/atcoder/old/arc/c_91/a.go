package c_91

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc091/tasks/arc091_c
//
//输入 n a b (≤3e5)。
//构造一个 1~n 的排列，其 LIS 的长度为 a，LDS 的长度为 b。
//如果不存在这样的排列，输出 -1，否则输出任意一个满足要求的排列。
//注：LIS 指最长上升子序列，LDS 指最长下降子序列。

func AtCoderARC91C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, a, b, i int
	fmt.Fscan(r, &n, &a, &b)
	if a+b-1 > n || a*b < n {
		fmt.Fprintln(w, -1)
		return
	}

	for a--; a >= 0; a-- {
		sz := min(b, n-a) // 后面至少要留下 a 个，这样才能满足 |LIS|=a 的条件
		for j := i + sz; j > i; j-- {
			fmt.Fprint(w, j, " ")
		}
		i += sz
		n -= sz
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
