package _400

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://codeforces.com/problemset/problem/1420/D
//
//输入 n, k (1≤k≤n≤3e5) 和 n 个闭区间，区间的范围在 [1,1e9]。
//你需要从 n 个区间中选择 k 个区间，且这 k 个区间的交集不为空。
//输出方案数模 998244353 的结果。

//对于每个区间的左端点 L，可能存在某些选法，是以 L 为交集的左端点的。
//如果有 m 个区间包含 L，且左端点为 L 的区间有 c 个，
//那么以 L 为交集左端点的方案数为 C(m, k) - C(m-c, k)，即减去不以 L 为交集左端点的方案数。
//
//两种实现方案：扫描线 / 差分哈希表。

func CF1420D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const (
		mod       = 998244353
		mx  int64 = 3e5 + 5
	)
	f := [mx + 1]int64{1}
	for i := int64(1); i <= mx; i++ {
		f[i] = f[i-1] * i % mod
	}
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	invF := [...]int64{mx: pow(f[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int64 {
		if k < 0 || k > n {
			return 0
		}
		return f[n] * invF[k] % mod * invF[n-k] % mod
	}

	var n, k int
	fmt.Fscan(r, &n, &k)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		fmt.Fscan(r, &a[i].l, &a[i].r)
	}

	d := map[int]int{}
	cnt := map[int]int{}
	for _, p := range a {
		d[p.l]++
		d[p.r+1]--
		cnt[p.l]++
	}
	keys := make([]int, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	sum := make(map[int]int, len(keys))
	s := 0
	for _, v := range keys {
		s += d[v]
		sum[v] = s
	}

	ans := int64(0)
	for v, c := range cnt {
		ans += C(sum[v], k) - C(sum[v]-c, k)
	}
	fmt.Fprint(w, (ans%mod+mod)%mod)

}
