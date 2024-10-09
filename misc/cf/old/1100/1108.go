package _100

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
)

func CF1108A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var T, l1, r1, l2, r2 int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &l1, &r1, &l2, &r2)
		if l1 != r2 {
			fmt.Fprintln(w, l1, r2)
		} else {
			fmt.Fprintln(w, r1, l2)
		}
	}

}

func CF1108B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var n int64
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}
	sort.Ints(arr)
	for i := int64(n - 1); i > 0; i-- {
		if arr[n-1]%arr[i] != 0 || arr[i-1] == arr[i] {
			fmt.Fprintln(w, arr[n-1], arr[i])
			return
		}
	}
}

func CF1108C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var n, ans int64
	var b []byte
	s, mod, ansS := []string{"RGB", "RBG", "BRG", "BGR", "GBR", "GRB"}, 3, ""
	fmt.Fscan(r, &n, &b)
	ans = n
	for _, v := range s {
		tmp := 0
		for i, vv := range b {
			if vv != v[i%mod] {
				tmp++
			}
		}
		if int64(tmp) < ans {
			ans = int64(tmp)
			ansS = v
		}
	}
	fmt.Fprintln(w, ans)
	for i := range b {
		fmt.Fprint(w, string(ansS[i%mod]))
	}

}

func CF1108D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var n, ans int64
	var b []byte
	s := "RGB"
	fmt.Fscan(r, &n, &b)
	for i := int64(1); i < n; i++ {
		if b[i] == b[i-1] {
			ans++
			for _, v := range s {
				if byte(v) != b[i-1] {
					if i+1 < n {
						if byte(v) != b[i+1] {
							b[i] = byte(v)
							break
						}
						continue
					}
					b[i] = byte(v)
					break
				}
			}

		}
	}
	fmt.Fprintln(w, ans, " ", string(b))
}

type seg []struct{ l, r, min, todo int64 }

func (t seg) build(a []int64, o, l, r int64) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].min = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg) do(o, v int64) {
	t[o].min += v
	t[o].todo += v
}

func (t seg) spread(o int64) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg) update(o, l, r, v int64) {
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
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func CF1108E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int64
	fmt.Fscan(in, &n, &m)
	a := make([]int64, n+1)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	t := make(seg, n*4)
	t.build(a, 1, 1, n)

	ps := make([]struct{ l, r int64 }, m)
	ls := make([][]int64, n+1)
	for i := range ps {
		fmt.Fscan(in, &ps[i].l, &ps[i].r)
		l, r := ps[i].l, ps[i].r
		t.update(1, l, r, -1)
		ls[l] = append(ls[l], r)
	}

	maxD, maxI := int64(0), int64(1)
	rs := make([][]int64, n+1)
	for i := int64(1); i <= n; i++ {
		for _, r := range ls[i] {
			t.update(1, i, r, 1)
			rs[r] = append(rs[r], i)
		}
		for _, l := range rs[i-1] {
			t.update(1, l, i-1, -1)
		}
		d := a[i] - t[1].min
		if d > maxD {
			maxD, maxI = d, i
		}
	}

	fmt.Fprintln(out, maxD)
	ids := []interface{}{}
	for i, p := range ps {
		if p.r < maxI || p.l > maxI {
			ids = append(ids, i+1)
		}
	}
	fmt.Fprintln(out, len(ids))
	fmt.Fprintln(out, ids...)
}

func CF1108E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int64
	mmax, mmin := int64(math.MinInt64), int64(math.MaxInt64)
	fmt.Fscan(r, &n, &m)
	arr := make([]int64, n+1)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &arr[i])
		mmax, mmin = max(mmax, arr[i]), min(mmin, arr[i])
	}
	ans, ansI := mmax-mmin, int64(0)
	left, right := make([]int64, m+1), make([]int64, m+1)
	sub, add := make(map[int64][]int64), make(map[int64][]int64)
	for i := int64(1); i <= m; i++ {
		fmt.Fscan(r, &left[i], &right[i])
		sub[left[i]] = append(sub[left[i]], i)
		add[right[i]+1] = append(add[right[i]+1], i)
	}

	for i := int64(1); i <= n; i++ {
		for j := 0; j < len(sub[i]); j++ {
			id := sub[i][j]
			for k := left[id]; k <= right[id]; k++ {
				arr[k]--
			}
		}
		for j := 0; j < len(add[i]); j++ {
			id := add[i][j]
			for k := left[id]; k <= right[id]; k++ {
				arr[k]++
			}
		}
		if len(sub[i]) > 0 || len(add[i]) > 0 {
			mi, mx := int64(math.MaxInt64), int64(math.MinInt64)
			for j := int64(1); j <= n; j++ {
				mi, mx = min(mi, arr[j]), max(mx, arr[j])
				if mx-mi > ans {
					ans = max(ans, mx-mi)
					ansI = i
				}
			}
		}
	}

	fmt.Fprintln(w, ans)
	cnt := 0
	for i := int64(1); i <= m; i++ {
		if left[i] <= ansI && right[i] >= ansI {
			cnt++
		}
	}
	fmt.Fprintln(w, cnt)
	for i := int64(1); i <= m; i++ {
		if left[i] <= ansI && right[i] >= ansI {
			fmt.Fprintln(w, i, " ")
		}
	}

}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
