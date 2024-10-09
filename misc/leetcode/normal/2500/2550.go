package _500

// 2569
// lazy线段树
type lazySeg []struct {
	l, r, val int
	flip      bool
}

func (s lazySeg) maintain(o int) {
	s[o].val = s[o<<1].val + s[o<<1|1].val
}

func (s lazySeg) build(a []int, o, l, r int) {
	s[o].l, s[o].r = l, r
	if l == r {
		s[o].val = a[l-1]
		return
	}
	m := (l + r) >> 1
	s.build(a, o<<1, l, m)
	s.build(a, o<<1|1, m+1, r)
	s.maintain(o)
}

func (s lazySeg) do(o int) {
	p := &s[o]
	p.val = p.r - p.l + 1 - p.val
	p.flip = !p.flip
}

func (s lazySeg) spread(o int) {
	if s[o].flip {
		s.do(o << 1)
		s.do(o<<1 | 1)
		s[o].flip = false
	}
}

func (s lazySeg) update(o, l, r int) {
	if l <= s[o].l && s[o].r <= r {
		s.do(o)
		return
	}
	s.spread(o)

	m := (s[o].l + s[o].r) >> 1
	if m >= l {
		s.update(o<<1, l, r)
	}
	if m < r {
		s.update(o<<1|1, l, r)
	}
	s.maintain(o)
}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n := len(nums1)
	sum := 0
	for _, v := range nums2 {
		sum += v
	}
	seg := make(lazySeg, n*4)
	seg.build(nums1, 1, 1, n)
	ans := make([]int64, 0)
	for _, v := range queries {
		switch v[0] {
		case 1:
			seg.update(1, v[1]+1, v[2]+1)
		case 2:
			sum += seg[1].val * v[1]
		case 3:
			ans = append(ans, int64(sum))
		}
	}
	return ans
}
