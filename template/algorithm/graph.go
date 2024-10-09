package algorithm

import (
	"container/heap"
	. "fmt"
	"io"
	"math/bits"
)

func _() {
	const (
		n         = 1000
		inf int64 = 1e18
	)
	type edge struct{ to, wt int }
	dijkstra := func(g [][]edge, start int) []int64 {
		dist := make([]int64, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[start] = 0
		// 虽然可以用 dist 来判断是否需要 relax，但是对于一些变形题，用 vis 是最稳的
		vis := make([]bool, n)
		fa := make([]int, n)
		for i := range fa {
			fa[i] = -1
		}
		h := vdHeap{{start, 0}}
		for len(h) > 0 {
			p := h.pop()
			v := p.v
			if vis[v] { // p.dis > dist[v]
				continue
			}
			vis[v] = true
			for _, e := range g[v] {
				w := e.to
				if newD := dist[v] + int64(e.wt); newD < dist[w] {
					dist[w] = newD
					fa[w] = v
					h.push(vdPair{w, dist[w]})
				}
			}
		}
		return dist
	}

	// 最近公共祖先 · 其一 · 基于树上倍增和二分搜索
	lcaBinarySearch := func(n, root int, g [][]int, max func(int, int) int) {
		// O(nlogn) 预处理，O(logn) 查询
		// 适用于查询量和节点数等同的情形
		// NOTE: 多个点的 LCA 等于 dfn_min 和 dfn_max 的 LCA
		// https://oi-wiki.org/graph/lca/#_5
		//
		// 模板题 https://www.luogu.com.cn/problem/P3379
		// 到两点距离相同的点的数量 https://codeforces.com/problemset/problem/519/E
		// https://atcoder.jp/contests/arc060/tasks/arc060_c
		// https://codeforces.com/problemset/problem/33/D
		// 路径点权乘积 https://ac.nowcoder.com/acm/contest/6913/C
		// 树上倍增应用（静态路径查询）：代码见下面的 EXTRA 部分
		//    维护最大值（与 MST 结合）https://codeforces.com/problemset/problem/609/E
		//       变体 https://codeforces.com/problemset/problem/733/F
		//    维护最大值（与 MST 结合）LC1697 https://leetcode-cn.com/problems/checking-existence-of-edge-length-limited-paths/
		//    维护最大值（与 MST 结合）LC1724（上面这题的在线版）https://leetcode-cn.com/problems/checking-existence-of-edge-length-limited-paths-ii/
		//    维护最大值和严格次大值（严格次小 MST）：见 graph.go 中的 strictlySecondMST
		//    维护前十大（点权）https://codeforces.com/problemset/problem/587/C
		// 树上倍增-查询深度最小的未被标记的点 https://codeforces.com/problemset/problem/980/E
		// 题目推荐 https://cp-algorithms.com/graph/lca.html#toc-tgt-2

		const mx = 17 // bits.Len(最大节点数)
		pa := make([][mx]int, n)
		dep := make([]int, n)
		var buildPa func(v, p, d int)
		buildPa = func(v, p, d int) {
			pa[v][0] = p
			dep[v] = d
			for _, w := range g[v] {
				if w != p {
					buildPa(w, v, d+1)
				}
			}
		}
		buildPa(root, -1, 0) // d 从 0 开始
		// 倍增
		for i := 0; i+1 < mx; i++ {
			for v := range pa {
				if p := pa[v][i]; p != -1 {
					pa[v][i+1] = pa[p][i]
				} else {
					pa[v][i+1] = -1
				}
			}
		}

		// 从 v 开始，向上跳到指定深度 d
		// https://en.wikipedia.org/wiki/Level_ancestor_problem
		// https://codeforces.com/problemset/problem/1535/E
		uptoDep := func(v, d int) int {
			if d > dep[v] {
				panic(-1)
			}
			for i := 0; i < mx; i++ {
				if (dep[v]-d)>>i&1 > 0 {
					v = pa[v][i]
				}
			}
			return v
		}
		_lca := func(v, w int) int {
			if dep[v] > dep[w] {
				v, w = w, v
			}
			w = uptoDep(w, dep[v])
			if w == v {
				return v
			}
			for i := mx - 1; i >= 0; i-- {
				if pv, pw := pa[v][i], pa[w][i]; pv != pw {
					v, w = pv, pw
				}
			}
			return pa[v][0]
		}
		disVW := func(v, w int) int { return dep[v] + dep[w] - dep[_lca(v, w)]*2 }

		// EXTRA: 输入 v 和 to，to 可能是 v 的子孙，返回从 v 到 to 路径上的第二个节点（v 的一个儿子）
		// 如果 v 不是 to 的子孙，返回 -1
		down1 := func(v, to int) int {
			if dep[v] >= dep[to] {
				return -1
			}
			to = uptoDep(to, dep[v]+1)
			if pa[to][0] == v {
				return to
			}
			return -1
		}

		// EXTRA: 从 v 出发，向 to 方向走一步
		// 输入需要保证 v != to
		move1 := func(v, to int) int {
			if v == to {
				panic(-1)
			}
			if _lca(v, to) == v { // to 在 v 下面
				return uptoDep(to, dep[v]+1)
			}
			// lca 在 v 上面
			return pa[v][0]
		}

		// EXTRA: 从 v 开始，向上跳 k 步
		// 不存在则返回 -1
		// O(1) 求法见长链剖分
		uptoKthPa := func(v, k int) int {
			for i := 0; i < mx && v != -1; i++ {
				if k>>i&1 > 0 {
					v = pa[v][i]
				}
			}
			return v
		}

		// EXTRA: 输入 v 和 w，返回 v 到 w 路径上的中点
		// 返回值是一个数组，因为可能有两个中点
		// 在有两个中点的情况下，保证返回值的第一个中点离 v 更近
		midPath := func(v, w int) []int {
			lca := _lca(v, w)
			dv := dep[v] - dep[lca]
			dw := dep[w] - dep[lca]
			if dv == dw {
				return []int{lca}
			}
			if dv > dw {
				mid := uptoKthPa(v, (dv+dw)/2)
				if (dv+dw)%2 == 0 {
					return []int{mid}
				}
				return []int{mid, pa[mid][0]}
			} else {
				mid := uptoKthPa(w, (dv+dw)/2)
				if (dv+dw)%2 == 0 {
					return []int{mid}
				}
				return []int{pa[mid][0], mid} // pa[mid][0] 离 v 更近
			}
		}

		{
			// 加权树上二分
			var dep []int64 // 加权深度，dfs 预处理略
			// 从 v 开始向根移动至多 d 距离，返回最大移动次数，以及能移动到的离根最近的点
			// 变形 https://codeforces.com/problemset/problem/932/D
			uptoDep := func(v int, d int64) (int, int) {
				step := 0
				dv := dep[v]
				for i := mx - 1; i >= 0; i-- {
					if p := pa[v][i]; p != -1 && dv-dep[p] <= d {
						step |= 1 << i
						v = p
					}
				}
				return step, v
			}
			_ = uptoDep
		}

		{
			// EXTRA: 倍增的时候维护其他属性，如边权最值等
			// 下面的代码来自 https://codeforces.com/problemset/problem/609/E
			// EXTRA: 额外维护最值边的下标，见 https://codeforces.com/contest/733/submission/120955685
			type nb struct{ to, wt int }
			g := make([][]nb, n)
			// read g ...
			const mx = 18
			type pair struct{ p, maxWt int }
			pa := make([][mx]pair, n)
			dep := make([]int, n)
			var build func(v, p, d int)
			build = func(v, p, d int) {
				pa[v][0].p = p
				dep[v] = d
				for _, e := range g[v] {
					if w := e.to; w != p {
						pa[w][0].maxWt = e.wt
						build(w, v, d+1)
					}
				}
			}
			build(0, -1, 0)

			for i := 0; i+1 < mx; i++ {
				for v := range pa {
					if p := pa[v][i]; p.p != -1 {
						pp := pa[p.p][i]
						pa[v][i+1] = pair{pp.p, max(p.maxWt, pp.maxWt)}
					} else {
						pa[v][i+1].p = -1
					}
				}
			}

			// 求 LCA(v,w) 的同时，顺带求出 v-w 上的边权最值
			_lca := func(v, w int) (lca, maxWt int) {
				if dep[v] > dep[w] {
					v, w = w, v
				}
				for i := 0; i < mx; i++ {
					if (dep[w]-dep[v])>>i&1 > 0 {
						p := pa[w][i]
						maxWt = max(maxWt, p.maxWt)
						w = p.p
					}
				}
				if w != v {
					for i := mx - 1; i >= 0; i-- {
						if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
							maxWt = max(maxWt, max(pv.maxWt, pw.maxWt))
							v, w = pv.p, pw.p
						}
					}
					maxWt = max(maxWt, max(pa[v][0].maxWt, pa[w][0].maxWt))
					v = pa[v][0].p
				}
				// 如果是点权的话这里加上 maxWt = max(maxWt, pa[v][0].maxWt)
				lca = v
				return
			}

			_ = _lca
		}

		_ = []interface{}{disVW, uptoKthPa, down1, move1, midPath}
	}

	// 最近公共祖先 · 其二 · 基于 RMQ
	lcaRMQ := func(root int, g [][]int) {
		// O(nlogn) 预处理，O(1) 查询
		// 由于预处理 ST 表是基于一个长度为 2n 的序列，所以常数上是比倍增算法要大的。内存占用也比倍增要大一倍左右（这点可忽略）
		// 优点是查询的复杂度低，适用于查询量大的情形
		// https://oi-wiki.org/graph/lca/#rmq
		// https://codeforces.com/problemset/problem/342/E

		vs := make([]int, 0, 2*len(g)-1)  // 欧拉序列
		pos := make([]int, len(g))        // pos[v] 表示 v 在 vs 中第一次出现的位置编号
		dep := make([]int, 0, 2*len(g)-1) // 深度序列，和欧拉序列一一对应
		disRoot := make([]int, len(g))    // disRoot[v] 表示 v 到 root 的距离
		var build func(v, p, d int)       // 若有边权需额外传参 dis
		build = func(v, p, d int) {
			pos[v] = len(vs)
			vs = append(vs, v)
			dep = append(dep, d)
			disRoot[v] = d
			for _, w := range g[v] {
				if w != p {
					build(w, v, d+1) // d+e.wt
					vs = append(vs, v)
					dep = append(dep, d)
				}
			}
		}
		build(root, -1, 0)
		type stPair struct{ v, i int }
		const mx = 17 + 1 // bits.Len(len(dep))
		var st [][mx]stPair
		stInit := func(a []int) {
			n := len(a)
			st = make([][mx]stPair, n)
			for i, v := range a {
				st[i][0] = stPair{v, i}
			}
			for j := 1; 1<<j <= n; j++ {
				for i := 0; i+1<<j <= n; i++ {
					if a, b := st[i][j-1], st[i+1<<(j-1)][j-1]; a.v < b.v {
						st[i][j] = a
					} else {
						st[i][j] = b
					}
				}
			}
		}
		stInit(dep)
		stQuery := func(l, r int) int { // [l,r) 注意 l r 是从 0 开始算的
			k := bits.Len(uint(r-l)) - 1
			a, b := st[l][k], st[r-1<<k][k]
			if a.v < b.v {
				return a.i
			}
			return b.i
		}
		// 注意下标的换算，打印 LCA 的话要 +1
		_lca := func(v, w int) int {
			pv, pw := pos[v], pos[w]
			if pv > pw {
				pv, pw = pw, pv
			}
			return vs[stQuery(pv, pw+1)]
		}
		_d := func(v, w int) int { return disRoot[v] + disRoot[w] - disRoot[_lca(v, w)]*2 }

		_ = _d
	}

	// 最近公共祖先 · 其三 · Tarjan 离线算法
	lcaTarjan := func(in io.Reader, n, q, root int) []int {
		// 时间复杂度 O(n+qα)
		// 原论文 https://dl.acm.org/doi/pdf/10.1145/800061.808753
		// https://core.ac.uk/download/pdf/82125836.pdf
		// https://oi-wiki.org/graph/lca/#tarjan
		// https://cp-algorithms.com/graph/lca_tarjan.html
		// 扩展：Tarjan RMQ https://codeforces.com/blog/entry/48994
		// LC2646 https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			v, w := 0, 0
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		lca := make([]int, q)
		dis := make([]int, q) // dis(q.v,q.w)
		type query struct{ w, i int }
		qs := make([][]query, n)
		for i := 0; i < q; i++ {
			v, w := 0, 0
			Fscan(in, &v, &w)
			v--
			w--

			// 第一种写法：保证在 v=w 时恰好只更新一个（结合下面的 if w := q.w; w == v || ... 理解）
			qs[v] = append(qs[v], query{w, i})
			if v != w {
				qs[w] = append(qs[w], query{v, i})
			}

			// 第二种写法：单独处理 v==w 的情况
			//if v != w {
			//	qs[v] = append(qs[v], query{w, i})
			//	qs[w] = append(qs[w], query{v, i})
			//} else {
			//	// do v==w...
			//	lca[i] = v
			//	dis[i] = 0
			//}
		}

		_fa := make([]int, n)
		for i := range _fa {
			_fa[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			if _fa[x] != x {
				_fa[x] = find(_fa[x])
			}
			return _fa[x]
		}

		dep := make([]int, n)
		// 为什么不用 bool 数组？
		// 对于下面代码中的 do(v, w, lcaVW)
		// 如果 v 是 w 的祖先节点，那么 w 递归结束后会触发一次，v 递归结束后又会触发一次
		// 如果 do 中有增量更新，这样就错了
		// 而三色标记法可以保证只会触发一次
		color := make([]int8, n)
		var tarjan func(v, d int)
		tarjan = func(v, d int) {
			dep[v] = d
			color[v] = 1
			for _, w := range g[v] {
				if color[w] == 0 {
					tarjan(w, d+1)
					_fa[w] = v // 相当于把 w 的子树节点全部 merge 到 v
				}
			}
			for _, q := range qs[v] {
				w := q.w
				// color[w] == 2 意味着 y 所在子树已经遍历完
				// 也就意味着 w 已经 merge 到它和 v 的 LCA 上了
				if w == v || color[w] == 2 {
					// do(v, w, lcaVW)...
					lcaVW := find(w)
					lca[q.i] = lcaVW
					dis[q.i] = dep[v] + dep[w] - dep[lcaVW]*2
				}
			}
			color[v] = 2
		}
		tarjan(root, 0)
		return lca
	}

	// 割点（割顶） cut vertices / articulation points
	// https://codeforces.com/blog/entry/68138
	// https://oi-wiki.org/graph/cut/#_1
	// low(v): 在不经过 v 父亲的前提下能到达的最小的时间戳
	// 模板题 https://www.luogu.com.cn/problem/P3388
	// LC928 https://leetcode-cn.com/problems/minimize-malware-spread-ii/
	findCutVertices := func(n int, g [][]int, min func(int, int) int) (isCut []bool) {
		isCut = make([]bool, n)
		dfn := make([]int, n) // DFS 到结点 v 的时间（从 1 开始）
		// low[v]: v 的儿子及其邻居的 dfn 的最小值
		dfsClock := 0
		var tarjan func(v, fa int) int
		tarjan = func(v, fa int) int { // 无需考虑重边
			dfsClock++
			dfn[v] = dfsClock
			lowV := dfsClock
			childCnt := 0
			for _, w := range g[v] {
				if dfn[w] == 0 {
					childCnt++
					lowW := tarjan(w, v)
					if lowW >= dfn[v] { // 以 w 为根的子树中没有反向边能连回 v 的祖先（可以连到 v 上，这也算割顶）
						isCut[v] = true
					}
					lowV = min(lowV, lowW)
				} else if w != fa { // （w!=fa 可以省略，但为了保证某些题目没有重复统计所以保留）   找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
					lowV = min(lowV, dfn[w])
				}
			}
			if fa == -1 && childCnt == 1 { // 特判：只有一个儿子的树根，删除后并没有增加连通分量的个数，这种情况下不是割顶
				isCut[v] = false
			}
			return lowV
		}
		for v, timestamp := range dfn {
			if timestamp == 0 {
				tarjan(v, -1)
			}
		}

		cuts := []int{}
		for v, is := range isCut {
			if is {
				cuts = append(cuts, v) // v+1
			}
		}

		return
	}

	// 桥（割边）
	// https://oi-wiki.org/graph/cut/#_4
	// https://algs4.cs.princeton.edu/41graph/Bridge.java.html
	// 模板题 LC1192 https://leetcode-cn.com/problems/critical-connections-in-a-network/
	//       https://codeforces.com/problemset/problem/1000/E
	// 题目推荐 https://cp-algorithms.com/graph/bridge-searching.html#toc-tgt-2
	// 与 MST 结合 https://codeforces.com/problemset/problem/160/D
	// 与最短路结合 https://codeforces.com/problemset/problem/567/E
	// https://codeforces.com/problemset/problem/118/E
	// todo 构造 https://codeforces.com/problemset/problem/550/D
	findBridges := func(in io.Reader, n, m int, min func(int, int) int) (isBridge []bool) {
		type neighbor struct{ to, eid int }
		g := make([][]neighbor, n)
		type edge struct{ v, w int }
		edges := make([]edge, m)
		for i := 0; i < m; i++ {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], neighbor{w, i})
			g[w] = append(g[w], neighbor{v, i})
			edges[i] = edge{v, w}
		}
		isBridge = make([]bool, len(edges))
		dfn := make([]int, len(g)) // 值从 1 开始
		dfsClock := 0
		var tarjan func(int, int) int
		tarjan = func(v, fid int) int { // 使用 fid 而不是 fa，可以兼容重边的情况
			dfsClock++
			dfn[v] = dfsClock
			lowV := dfsClock
			for _, e := range g[v] {
				if w := e.to; dfn[w] == 0 {
					lowW := tarjan(w, e.eid)
					if lowW > dfn[v] { // 以 w 为根的子树中没有反向边能连回 v 或 v 的祖先，所以 v-w 必定是桥
						isBridge[e.eid] = true
					}
					lowV = min(lowV, lowW)
				} else if e.eid != fid { // 找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
					lowV = min(lowV, dfn[w])
				}
			}
			return lowV
		}
		for v, timestamp := range dfn {
			if timestamp == 0 {
				tarjan(v, -1)
			}
		}

		// EXTRA: 所有桥边的下标
		bridgeEIDs := []int{}
		for eid, b := range isBridge {
			if b {
				bridgeEIDs = append(bridgeEIDs, eid)
			}
		}

		return
	}

	// 任意两点最短路 Floyd-Warshall  O(n^3)  本质是求 Min-plus matrix multiplication
	// 传入邻接矩阵 dis
	// dis[v][w] == inf 表示没有 v-w 边
	// 带你发明 Floyd 算法！https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/
	// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
	// https://en.wikipedia.org/wiki/Min-plus_matrix_multiplication
	// https://cp-algorithms.com/graph/all-pair-shortest-path-floyd-warshall.html#toc-tgt-5
	//
	// 模板题 https://www.luogu.com.cn/problem/B3647
	// 传递闭包 https://www.luogu.com.cn/problem/B3611
	// https://codeforces.com/problemset/problem/33/B
	// https://codeforces.com/problemset/problem/1204/C
	// LC1334 https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
	// LC1462 https://leetcode.cn/problems/course-schedule-iv/
	// 动态加点 https://codeforces.com/problemset/problem/295/B
	// 动态加边 LC2642 https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/
	// - https://codeforces.com/problemset/problem/25/C LC2646 https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/
	// todo https://atcoder.jp/contests/abc243/tasks/abc243_e
	// 传递闭包 UVa247 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=4&page=show_problem&problem=183
	// 注：求传递闭包时，若 i-k 不连通，则最内层循环无需运行
	// 任意两点最大边权最小路径 UVa10048 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=12&page=show_problem&problem=989
	shortestPathFloydWarshall := func(dis [][]int) [][]int {
		// dis[k][i][j] 表示「经过若干个编号不超过 k 的中间节点」时，从 i 到 j 的最短路长度，其中第一维可以压缩掉
		// 为什么可以把第一维度去掉？dis[i][k] 和 dis[k][j] 不会被覆盖掉吗？
		// 见算法导论第三版练习 25.2-4（网络上有习题解答）

		// 初始化，注意 dis[i][i] = 0
		//dis := make([][]int, n)
		//for i := range dis {
		//	dis[i] = make([]int, n)
		//	for j := range dis[i] {
		//		if j != i {
		//			dis[i][j] = math.MaxInt / 2
		//		}
		//	}
		//}
		for k := range dis {
			for i := range dis {
				for j := range dis {
					// 不选 k，选 k
					dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
				}
			}
		}

		// 如果出现 dis[i][i] < 0 则说明有负环

		// 动态加边
		// https://codeforces.com/problemset/problem/25/C
		// LC2642 https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/
		for i := range dis {
			// 注意 from=i 或者 to=j 时，下面的 dis[i][from] 和 dis[to][j] 都需要 dis[i][i] 这样的值
			// 所以初始化成 0 方便计算
			dis[i][i] = 0
		}
		addEdge := func(from, to, wt int) {
			// 无法让任何最短路变短
			if wt >= dis[from][to] {
				return
			}
			for i := range dis {
				for j := range dis {
					dis[i][j] = min(dis[i][j], dis[i][from]+wt+dis[to][j])
				}
			}
		}
		_ = addEdge

		return dis
	}

	_ = []interface{}{dijkstra, lcaBinarySearch, lcaRMQ, lcaTarjan, findCutVertices, findBridges, shortestPathFloydWarshall}
}

type vdPair struct {
	v   int
	dis int64
}
type vdHeap []vdPair

func (h vdHeap) Len() int              { return len(h) }
func (h vdHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h vdHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *vdHeap) Push(v interface{})   { *h = append(*h, v.(vdPair)) }
func (h *vdHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *vdHeap) push(v vdPair)        { heap.Push(h, v) }
func (h *vdHeap) pop() vdPair          { return heap.Pop(h).(vdPair) }
