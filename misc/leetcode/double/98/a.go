package double_98

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1
func minMaxDifference(num int) int {
	s := strconv.Itoa(num)
	mx, mm := 0, math.MaxInt32
	for _, v := range s {
		v1 := strings.ReplaceAll(s, string(v), "9")
		v2 := strings.ReplaceAll(s, string(v), "0")

		n1, _ := strconv.Atoi(v1)
		n2, _ := strconv.Atoi(v2)

		mx = max(mx, n1)
		mm = min(mm, n2)
	}
	return mx - mm
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 2
func minimizeSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return min(min(nums[n-3]-nums[0], nums[n-2]-nums[1]), nums[n-1]-nums[2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 3
func minImpossibleOR(nums []int) int {
	h := make(map[int]bool)
	for i := range nums {
		h[nums[i]] = true
	}
	for i := 1; ; i <<= 1 {
		if !h[i] {
			return i
		}
	}
}

// 4
type seg []struct {
	l, r, cnt1 int
	rev        bool
}

func (t seg) maintain(o int) { t[o].cnt1 = t[o<<1].cnt1 + t[o<<1|1].cnt1 }

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].cnt1 = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) do(O int) {
	o := &t[O]
	o.cnt1 = o.r - o.l + 1 - o.cnt1
	o.rev = !o.rev
}

func (t seg) spread(o int) {
	if t[o].rev {
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].rev = false
	}
}

func (t seg) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t.maintain(o)
}

func handleQuery(nums1, nums2 []int, queries [][]int) (ans []int64) {
	sum := 0
	for _, x := range nums2 {
		sum += x
	}
	t := make(seg, len(nums1)*4)
	t.build(nums1, 1, 1, len(nums1))
	for _, q := range queries {
		if q[0] == 1 {
			t.update(1, q[1]+1, q[2]+1)
		} else if q[0] == 2 {
			sum += q[1] * t[1].cnt1
		} else {
			ans = append(ans, int64(sum))
		}
	}
	return
}
