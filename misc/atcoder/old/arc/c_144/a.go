package c_144

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc144/tasks/arc144_c
//
//输入 n 和 k (k<n≤3e5)。
//请构造一个字典序最小的 1~n 的排列 p（下标从 1 开始），使得 abs(p[i]-i)≥k 对于每个 i 都成立。
//如果无法构造，输出 -1。否则输出 p。

func AtCoderARC144C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k int
	fmt.Fscan(r, &n, &k)

	if k > n/2 {
		fmt.Fprint(w, -1)
		return
	}
	ans, t := make([]int, n+1), 1
	used := make(map[int]bool)
	for i := 1; i <= n; i++ {
		j := i + k
		if j > n || used[j] || j+k <= n {
			j = t
			if j > i-k {
				j = i + k
			}
		}
		ans[i] = j
		used[j] = true
		for used[t] {
			t++
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprint(w, ans[i], " ")
	}
}
