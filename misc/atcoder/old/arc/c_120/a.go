package c_120

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc120/tasks/arc120_c
//
//输入 n(≤2e5) 和两个长为 n 的数组 a b，元素范围在 [0,1e9]。
//每次操作你可以选择 a 中的两个相邻数字，设 x=a[i], y=a[i+1]，更新 a[i]=y+1, a[i+1]=x-1。
//输出把 a 变成 b 的最小操作次数，如果无法做到则输出 -1。

// https://atcoder.jp/contests/arc120/submissions/36718795
//
//手玩一下可以发现，a[i] 左移 i 位就 +i，右移 i 位就 -i。
//设 a[i] 最终和 b[j] 匹配，则有 a[i]+i-j=b[j]。
//移项得 a[i]+i = b[j]+j。
//设 a'[i] = a[i]+i，b'[i] = b[i]+i。
//问题变成把 a' 通过邻项交换变成数组 b'，需要的最小操作次数。
//这可以用树状数组解决，具体见代码。

func AtCoderARC120C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	h, a, b := make(map[int][]int), make([]int, n), make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
		a[i] += i
		h[a[i]] = append(h[a[i]], i)
	}

	for i := range b {
		fmt.Fscan(r, &b[i])
		b[i] += i
	}

	tree := make([]int, n+1)
	add := func(i int) {
		for i++; i <= n; i += i & (-i) {
			tree[i]++
		}
	}
	sum := func(i int) int {
		total := 0
		for i++; i > 0; i &= i - 1 {
			total += tree[i]
		}
		return total
	}

	ans := 0
	for i, v := range b {
		p := h[v]
		if len(p) == 0 {
			fmt.Fprintln(w, -1)
			return
		}
		val := p[0]
		h[v] = p[1:]

		ans += i - sum(val)
		add(val)
	}
	fmt.Fprintln(w, ans)

}
