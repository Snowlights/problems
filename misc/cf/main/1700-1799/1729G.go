//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"os"
	"sort"
)

// 先计算最小操作次数，这可以贪心。
// 找出所有 t 在 s 中的位置，然后贪心替换：
// 如果第一个 t 与第二个 t 没有重叠，那么必须替换第一个 t；
// 否则替换第二个 t ，也能让第一个 t 不在 s 中。
// 具体来说，需要在第一个 t 的【范围】内替换一个最右的 t（这个最右的 t，
// 只要第一个字母的下标在【范围】内即可）。替换后，按同样的方式替换后续的 t。
//
// 然后就可以写一个二维 DP 了，dfs(i,j) 表示当前要处理后缀 s[i:]，
// 且已经替换了 j 次。枚举下一个要替换的 t。
// 你也可以把 t 在 s 中的所有位置都处理出来，
// 设为数组 a，然后在 a 上 DP。具体见代码。

func CF1729G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 1_000_000_007

	var T int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		a := suffixarray.New(s).Lookup(t, -1)
		sort.Ints(a)
		n, m := len(a), len(t)
		minMove := 0
		for i := 0; i < n; i = sort.SearchInts(a, a[sort.SearchInts(a, a[i]+m)-1]+m) {
			minMove++
		}
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, minMove+1)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(int, int) int
		f = func(i, j int) (res int) {
			if j > minMove {
				return
			}
			if i == n {
				return 1
			}
			dv := &dp[i][j]
			if *dv != -1 {
				return *dv
			}
			defer func() { *dv = res }()
			r := sort.SearchInts(a, a[i]+m)
			for k := i; k < r; k++ {
				res = (res + f(sort.SearchInts(a, a[k]+m), j+1)) % mod
			}
			return
		}
		Fprintln(out, minMove, f(0, 0))
	}

}

func main() { CF1729G(os.Stdin, os.Stdout) }
