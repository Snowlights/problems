package algorithm

// 推荐阅读《算法竞赛进阶指南》0x43 和 0x48 节
// https://oi-wiki.org/ds/seg/
// https://cp-algorithms.com/data_structures/segment_tree.html
// 总结得比较详细 https://www.acwing.com/blog/content/1684/
// https://codeforces.com/blog/entry/18051
// https://codeforces.com/blog/entry/89313
// https://codeforces.com/blog/entry/15890

// 目前解决的问题，CF292E, CF522D

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
	t[o].val = t.op(t[o<<1].val, t[o<<1|1].val)
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
	return t.op(t.queryRange(o<<1, l, r), t.queryRange(o<<1|1, l, r))
}

func (seg) op(a, b int) int {
	return min(a, b)
}

func sum(a, b int) int {
	return a + b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
