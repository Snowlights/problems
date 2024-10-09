package _200

import (
	"bufio"
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"io"
	"sort"
)

func CF1249A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n, val int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)
		h := make(map[int64]bool)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &val)
			h[val] = true
		}
		found := false
		for k, _ := range h {
			if h[k-1] || h[k+1] {
				found = true
				fmt.Fprintln(w, 2)
				break
			}
		}
		if !found {
			fmt.Fprintln(w, 1)
		}
	}

}

func CF1249B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)
		arr := make([]int64, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &arr[i])
		}

		exit := make(map[int64]map[int64]bool)
		for i := int64(0); i < n; i++ {
			if _, ok := exit[arr[i]]; ok {
				continue
			}
			h, x := make(map[int64]bool), arr[i]
			for ; h[x] == false; x = arr[x-1] {
				h[x] = true
				exit[x] = h
			}
		}
		for i := int64(0); i < n; i++ {
			v, _ := exit[i+1]
			fmt.Fprint(w, len(v), " ")
		}
		fmt.Fprintln(w)
	}

}

// C1
var ans []int64

func init() {
	h := make(map[int64]bool, 2000)
	base, num := int64(3), int64(3)
	h[1] = true
	for i := 0; i < 10; i++ {
		newH := make(map[int64]bool, 2000)
		newH[num] = true
		for k, _ := range h {
			newH[k] = true
			newH[k+num] = true
		}
		h = newH
		num *= base
	}
	ans = make([]int64, len(h))
	for k, _ := range h {
		ans = append(ans, k)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
}

func CF1249C1(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)
		idx := sort.Search(len(ans), func(i int) bool {
			return ans[i] >= n
		})
		fmt.Fprintln(w, ans[idx])
	}
}

func CF1249C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)
		base, num, ans := int64(3), int64(3), int64(1)
		for ans < n {
			ans += num
			num *= base
		}
		for num > 0 {
			if ans-num >= n {
				ans -= num
			}
			num /= base
		}
		fmt.Fprintln(w, ans)
	}
}

type pair struct {
	idx, right int64
}

func CF1249D1(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	type pair struct {
		idx, right int64
	}

	var n, k int64
	var left, right int64
	h := make(map[int64][]*pair)
	fmt.Fscan(r, &n, &k)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &left, &right)
		h[left] = append(h[left], &pair{
			idx:   i,
			right: right,
		})
	}

	ans, set := make([]int64, 0), make([]*pair, 0)
	for i := int64(1); i < 2e5+10; i++ {
		for _, v := range h[i] {
			set = append(set, v)
		}
		sort.Slice(set, func(i, j int) bool {
			return set[i].right < set[j].right ||
				(set[i].right == set[j].right && set[i].idx < set[j].idx)
		})
		for int64(len(set)) > k {
			back := set[len(set)-1]
			set = set[:len(set)-1]
			ans = append(ans, back.idx)
		}
		for len(set) > 0 && set[0].right == i {
			set = set[1:]
		}
	}
	fmt.Fprintln(w, len(ans))
	for _, a := range ans {
		fmt.Fprint(w, a, " ")
	}
}

func CF1249D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k int64
	var left, right int64
	h := make(map[int64][]*pair)
	fmt.Fscan(r, &n, &k)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &left, &right)
		h[left] = append(h[left], &pair{
			idx:   i,
			right: right,
		})
	}

	ans := make([]int64, 0)
	c := func(a, b interface{}) int {
		av, bv := a.(*pair), b.(*pair)
		if av.right > bv.right {
			return 1
		}
		if av.right == bv.right {
			if av.idx > bv.idx {
				return 1
			}
			if av.idx == bv.idx {
				return 0
			}
		}
		return -1
	}

	rbt := redblacktree.NewWith(c)
	for i := int64(1); i < 2e5+10; i++ {
		for _, v := range h[i] {
			rbt.Put(v, v)
		}
		for int64(rbt.Size()) > k {
			back := rbt.Right()
			ans = append(ans, back.Value.(*pair).idx)
			rbt.Remove(back.Key)
		}
		for rbt.Size() > 0 && rbt.Left().Value.(*pair).right == i {
			rbt.Remove(rbt.Left().Key)
		}
	}

	fmt.Fprintln(w, len(ans))
	for _, a := range ans {
		fmt.Fprint(w, a, " ")
	}
}

func CF1249E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, c int64
	fmt.Fscan(r, &n, &c)
	a, b := make([]int64, n), make([]int64, n)
	for i := int64(1); i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	for i := int64(1); i < n; i++ {
		fmt.Fscan(r, &b[i])
	}
	ans := make([][2]int64, n)
	ans[0][0], ans[0][1] = 0, c
	fmt.Fprint(w, 0, " ")
	for i := int64(1); i < n; i++ {
		ans[i][0] = min(ans[i-1][0], ans[i-1][1]) + a[i]
		ans[i][1] = min(ans[i-1][1]+b[i], ans[i-1][0]+b[i]+c)
		fmt.Fprint(w, min(ans[i][0], ans[i][1]), " ")
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
