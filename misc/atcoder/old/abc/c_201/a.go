package c_201

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc201/tasks/abc201_e
//
//输入 n(≤2e5) 和一棵有边权的树的 n-1 条边，节点编号从 1 开始，边权范围 [0,2^60)。
//定义 xor(i,j) 表示从 i 到 j 的简单路径上的边权的异或和。
//累加所有满足 1≤i<j≤n 的 xor(i,j)，对结果模 1e9+7 后输出。

// https://atcoder.jp/contests/abc201/submissions/36166044
//
//提示 1：逐位考虑。也就是假设边权只有 0 和 1。
//
//提示 2：xor(i,j) = xor(1,i) ^ xor(1,j)。
//
//提示 3：DFS，统计 xor(1,i) 中 1 的个数，记作 c。
//由于只有 1 和 0 异或才能是 1，这个比特位上的答案为 c * (n-c)。

// e
func AtCoderABC201E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)

	const mod int = 1e9 + 7
	var n, v, w, wt, ans int
	fmt.Fscan(r, &n)
	type pair struct{ to, wt int }
	g := make([][]pair, n+1)
	for i := 1; i < n; i++ {
		fmt.Fscan(r, &v, &w, &wt)
		g[v] = append(g[v], pair{w, wt})
		g[w] = append(g[w], pair{v, wt})
	}

	cnt := [64]int{}
	var dfs func(int, int, int)
	dfs = func(v, fa, xor int) {
		for i := range cnt {
			cnt[i] += xor >> i & 1
		}
		for _, e := range g[v] {
			if e.to != fa {
				dfs(e.to, v, xor^e.wt)
			}
		}
	}
	dfs(1, 0, 0)

	for i, c := range cnt {
		ans += 1 << i % mod * c % mod * (n - c)
	}
	fmt.Fprint(out, ans%mod)
}
