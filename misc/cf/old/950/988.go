package _50

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://codeforces.com/contest/988/problem/F
//
//输入正整数 dst(≤2000), n(≤(dst+1)/2) 和 m(≤2000)。
//然后输入 n 个互不相交的下雨区间，左右端点满足 0≤L<R≤dst。
//然后输入 m 行，每行两个数表示雨伞的位置 x[i](≤dst) 和重量 w[i](≤1e5)。
//
//你需要从 0 出发走到 dst。在下雨区间内移动必须打伞（区间端点处不算下雨）。
//你可以中途换伞。
//带着伞移动时，每走一单位长度，消耗的耐力等于伞的重量。
//你可以丢掉伞，空手移动时消耗的耐力为零。
//输出到达 dst 至少需要消耗多少耐力。如果无法到达 dst，输出 -1。
//
//进阶：请写一个和坐标范围无关的代码，即使坐标范围达到 1e9 也可以通过。

func CF988F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	fmt.Fscan(in, &n, &n, &m)
	seg := make([]struct{ l, r int }, n)
	for i := range seg {
		fmt.Fscan(in, &seg[i].l, &seg[i].r)
	}
	sort.Slice(seg, func(i, j int) bool { return seg[i].l < seg[j].r })
	type pair struct{ x, w int }
	a := make([]pair, m, m+1)
	for i := range a {
		fmt.Fscan(in, &a[i].x, &a[i].w)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	if a[0].x > seg[0].l {
		fmt.Fprint(out, -1)
		return
	}

	a = append(a, pair{seg[n-1].r, 0}) // 方便计算答案（也可以用 dst，但其实完全不需要这个值）
	f := make([]int, m+1)
	j := -1
	for i, p := range a {
		x := p.x
		for j+1 < n && seg[j+1].l < x { // 寻找 x 左侧最近区间左端点
			j++
		}
		if j < 0 {
			continue
		}
		x = min(x, seg[j].r) // 如果 x 在区间外（x 与左端点重合也算），那么移动到左侧区间的右端点时就应该丢伞了
		if a[i-1].x >= x {   // 如果上面没有 continue 那么这里必然 i>0
			f[i] = f[i-1]
			continue
		}
		f[i] = 1e9
		for j, q := range a[:i] {
			f[i] = min(f[i], f[j]+(x-q.x)*q.w)
		}
	}
	fmt.Fprint(out, f[m])
}

func CF988FV2(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	const (
		inf = 2e10
	)

	var dst, n, m, left, right int64
	fmt.Fscan(r, &dst, &n, &m)
	rain, umbrella := make(map[int64]bool), make([]int64, dst+1)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &left, &right)
		for j := left; j < right; j++ {
			rain[j] = true
		}
	}
	for i := range umbrella {
		umbrella[i] = inf
	}
	for i := int64(0); i < m; i++ {
		fmt.Fscan(r, &left, &right)
		umbrella[left] = min(umbrella[left], right)
	}

	dp := make([]map[int64]int64, dst+1)
	dp[0] = map[int64]int64{inf: 0}
	for i := int64(0); i < dst; i++ {
		// 打那把伞(k)的消耗是多少(v)
		dp[i+1] = make(map[int64]int64)
		for k, v := range dp[i] {
			u := umbrella[i]
			if old, ok := dp[i+1][u]; ok {
				dp[i+1][u] = min(old, v+u)
			} else {
				dp[i+1][u] = v + u
			}
			if old, ok := dp[i+1][k]; ok {
				dp[i+1][k] = min(old, v+k)
			} else {
				dp[i+1][k] = v + k
			}
			if !rain[i] {
				if old, ok := dp[i+1][inf]; ok {
					dp[i+1][inf] = min(old, v)
				} else {
					dp[i+1][inf] = v
				}
			}
		}
	}

	ans := int64(inf)
	for _, v := range dp[dst] {
		ans = min(ans, v)
	}
	if ans < inf {
		fmt.Fprintln(w, ans)
	} else {
		fmt.Fprintln(w, -1)
	}
}
