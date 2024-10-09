package _200

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF1207D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	const mod = 998244353

	var n, x, y int
	var ans, sa, sb, sp int64 = 1, 1, 1, 1
	fmt.Fscan(r, &n)
	type pair struct{ x, y int }
	a := make([]pair, n)
	ca := make([]int, n+1)
	cb := make([]int, n+1)
	cp := make(map[pair]int, n)
	for i := range a {
		fmt.Fscan(r, &x, &y)
		a[i] = pair{x, y}
		ca[x]++
		cb[y]++
		cp[a[i]]++
		ans = ans * int64(i+1) % mod
		sa = sa * int64(ca[x]) % mod
		sb = sb * int64(cb[y]) % mod
		sp = sp * int64(cp[a[i]]) % mod
	}
	ans -= sa + sb
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
	if sort.SliceIsSorted(a, func(i, j int) bool { return a[i].y < a[j].y }) {
		ans += sp
	}
	fmt.Fprint(w, (ans%mod+mod)%mod)
}
