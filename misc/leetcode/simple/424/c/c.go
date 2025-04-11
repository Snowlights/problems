package main

import "math/bits"

type seg []struct {
	l, r, mx, todo int
}

func (t seg) do(o, v int) {
	t[o].mx -= v
	t[o].todo += v
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg) maintain(o int) {
	t[o].mx = max(t[o<<1].mx, t[o<<1|1].mx)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].mx = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	if t[1].mx <= 0 {
		return 0
	}
	for i, q := range queries {
		t.update(1, q[0], q[1], q[2])
		if t[1].mx <= 0 {
			return i + 1
		}
	}
	return -1
}
