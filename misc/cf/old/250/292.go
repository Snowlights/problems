package _50

import (
	"bufio"
	"fmt"
	"io"
)

// https://codeforces.com/problemset/problem/292/E
//
// 输入 n(≤1e5) 和 m (≤1e5)，两个长度都为 n 的数组 a 和 b（元素范围在 [-1e9,1e9] 内，下标从 1 开始）。
// 然后输入 m 个操作：
// 操作 1 形如 1 x y k，表示把 a 的区间 [x,x+k-1] 的元素拷贝到 b 的区间 [y,y+k-1] 上（输入保证下标不越界）。
// 操作 2 形如 2 x，输出 b[x]。
func CF292E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	a, b := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &b[i])
	}
	t := make(seg, n*4)
	t.build(1, 1, n)
	var op, x, y, k int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &op)
		switch op {
		case 1:
			fmt.Fscan(r, &x, &y, &k)
			t.updateRange(1, y, y+k-1, x-y)
		case 2:
			fmt.Fscan(r, &x)
			d := t.query(1, x)
			if d == 1e9 {
				fmt.Fprintln(w, b[x])
			} else {
				fmt.Fprintln(w, a[x+d])
			}
		}
	}

}

// 线段树模版
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

func (t seg) updateRange(o, l, r, val int) {
	if l <= t[o].l && t[o].r <= r {
		t[o].val = val
		return
	}
	if d := t[o].val; d != 1e9 {
		t[o<<1].val = d
		t[o<<1|1].val = d
		t[o].val = 1e9
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.updateRange(o<<1, l, r, val)
	}
	if m < r {
		t.updateRange(o<<1|1, l, r, val)
	}
}

func (t seg) query(o, i int) int {
	if t[o].val != 1e9 || t[o].l == t[o].r {
		return t[o].val
	}
	if i <= (t[o].l+t[o].r)>>1 {
		return t.query(o<<1, i)
	}
	return t.query(o<<1|1, i)
}

func (t seg) queryRange(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.queryRange(o<<1, l, r)
	}
	if m < l {
		return t.queryRange(o<<1|1, l, r)
	}
	return min(t.queryRange(o<<1, l, r), t.queryRange(o<<1|1, l, r))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
