package _00

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"sort"
)

// 提示 1：转换成每对相邻区间，能放下的桥的长度的范围，记到一个数组 b 中。
// 提示 2：从小到大排序 a 和 b，其中 b 按照左端点排序。
// 提示 3：遍历 a，对于 b 中范围左端点在 a[i] 左边且没有使用过的范围，从中选出右端点最小且 ≥a[i] 的范围。
// 提示 4：用堆模拟。
func CF555B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	var left, right, ll, rr int64
	fmt.Fscan(r, &n, &m, &ll, &rr)
	type tuple struct {
		l, r int64
		i    int
	}
	b := make([]tuple, n-1)
	for i := range b {
		fmt.Fscan(r, &left, &right)
		b[i] = tuple{
			l: left - rr,
			r: right - ll,
			i: i,
		}
		ll, rr = left, right
	}
	sort.Slice(b, func(i, j int) bool { return b[i].l < b[j].l })

	a := make([]pair, m)
	for i := range a {
		fmt.Fscan(r, &a[i].x)
		a[i].i = i + 1
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

	ans := make([]int, n-1)
	h, j := hp{}, 0
	for _, p := range a {
		for j < n-1 && b[j].l <= p.x {
			heap.Push(&h, pair{b[j].r, b[j].i})
			j++
		}
		if len(h) > 0 {
			if h[0].x < p.x {
				break
			}
			ans[heap.Pop(&h).(pair).i] = p.i
		}
	}

	if len(h) > 0 || j < n-1 {
		fmt.Fprintln(w, "No")
	} else {
		fmt.Fprintln(w, "Yes")
		for _, v := range ans {
			fmt.Fprint(w, v, " ")
		}
	}

}

type pair struct {
	x int64
	i int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
