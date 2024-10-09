package _00

import (
	"bufio"
	"fmt"
	"io"
)

// https://codeforces.com/problemset/problem/522/D
// 输入 n(2≤n≤2e5) 和 m(1≤m≤2e5)；然后输入一个长为 n 的数组 a(-1e9≤a[i]≤1e9)，
// 数组下标从 1 开始；然后输入 m 个询问，每个询问表示一个数组 a 内的闭区间 [L,R] (1≤L≤R≤n)。
// 对每个询问，输出区间内的相同元素下标之间的最小差值。如果区间内不存在相同元素，输出 -1。

// https://codeforces.com/contest/522/submission/173138916
//本题用到的技巧十分经典（指提示 4），如果你对此题没有思路，请务必掌握这一技巧。
//提示 1：在相同元素中，只需要考虑相邻的元素。换句话说，这些元素对可以看成是若干线段。
//提示 2：问题转换成查询区间内的最短线段。
//提示 3：把询问离线，按照右端点排序。
//提示 4：把提示 1 中的线段长度记录在线段的左端点上。
//提示 5：遍历排序后的询问 [L,R]，如果位置 R 有线段的右端点，
// 则更新线段左端点处的线段长度，我们最后要做的就是查询区间 [L,R] 内的最小值。
//提示 6：用线段树实现。
//线段长度可以一边遍历数组 a，一边用一个 last map 记录。

type seg []struct{ l, r, val int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].val = l, r, 1e9
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].val = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t[o].val = min(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return min(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func CF522D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, r int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	type pair struct{ l, i int }
	q := make([][]pair, n+1)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &l, &r)
		q[r] = append(q[r], pair{l, i})
	}

	ans := make([]int, m)
	t := make(seg, n*4)
	t.build(1, 1, n)
	last := map[int]int{}
	for i := 1; i <= n; i++ {
		if p := last[a[i]]; p > 0 {
			t.update(1, p, i-p)
		}
		for _, p := range q[i] {
			res := t.query(1, p.l, i)
			if res == 1e9 {
				res = -1
			}
			ans[p.i] = res
		}
		last[a[i]] = i
	}
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}
