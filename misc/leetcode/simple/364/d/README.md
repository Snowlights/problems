#### 题目

<p>给你一棵 <code>n</code> 个节点的无向树，节点编号为 <code>1</code> 到 <code>n</code> 。给你一个整数 <code>n</code> 和一个长度为 <code>n - 1</code> 的二维整数数组 <code>edges</code> ，其中 <code>edges[i] = [u<sub>i</sub>, v<sub>i</sub>]</code> 表示节点 <code>u<sub>i</sub></code> 和 <code>v<sub>i</sub></code> 在树中有一条边。</p>

<p>请你返回树中的 <strong>合法路径数目</strong> 。</p>

<p>如果在节点 <code>a</code> 到节点 <code>b</code> 之间 <strong>恰好有一个</strong> 节点的编号是质数，那么我们称路径 <code>(a, b)</code> 是 <strong>合法的</strong> 。</p>

<p><strong>注意：</strong></p>

<ul>
	<li>路径 <code>(a, b)</code> 指的是一条从节点 <code>a</code> 开始到节点 <code>b</code> 结束的一个节点序列，序列中的节点 <strong>互不相同</strong> ，且相邻节点之间在树上有一条边。</li>
	<li>路径 <code>(a, b)</code> 和路径 <code>(b, a)</code> 视为 <strong>同一条</strong> 路径，且只计入答案 <strong>一次</strong> 。</li>
</ul>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/08/27/example1.png" style="width: 440px; height: 357px;" /></p>

<pre>
<b>输入：</b>n = 5, edges = [[1,2],[1,3],[2,4],[2,5]]
<b>输出：</b>4
<b>解释：</b>恰好有一个质数编号的节点路径有：
- (1, 2) 因为路径 1 到 2 只包含一个质数 2 。
- (1, 3) 因为路径 1 到 3 只包含一个质数 3 。
- (1, 4) 因为路径 1 到 4 只包含一个质数 2 。
- (2, 4) 因为路径 2 到 4 只包含一个质数 2 。
只有 4 条合法路径。
</pre>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/08/27/example2.png" style="width: 488px; height: 384px;" /></p>

<pre>
<b>输入：</b>n = 6, edges = [[1,2],[1,3],[2,4],[3,5],[3,6]]
<b>输出：</b>6
<b>解释：</b>恰好有一个质数编号的节点路径有：
- (1, 2) 因为路径 1 到 2 只包含一个质数 2 。
- (1, 3) 因为路径 1 到 3 只包含一个质数 3 。
- (1, 4) 因为路径 1 到 4 只包含一个质数 2 。
- (1, 6) 因为路径 1 到 6 只包含一个质数 3 。
- (2, 4) 因为路径 2 到 4 只包含一个质数 2 。
- (3, 6) 因为路径 3 到 6 只包含一个质数 3 。
只有 6 条合法路径。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n <= 10<sup>5</sup></code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>1 <= u<sub>i</sub>, v<sub>i</sub> <= n</code></li>
	<li>输入保证 <code>edges</code> 形成一棵合法的树。</li>
</ul>

#### 思路

![w364D-c.png](https://pic.leetcode.cn/1695525563-NYpzGx-w364D-c.png)

关于如何预处理质数，我之前讲过**埃氏筛**和**线性筛**两种做法，请看[【周赛 326】]()第四题。

```go  
const mx int = 1e5 + 1
var np = [mx]bool{1: true} // 质数=false 非质数=true
func init() {
	for i := 2; i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func countPaths(n int, edges [][]int) (ans int64) {
	fa := make([]int, n+1) // n+1
	size := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	sizeF := func(x int) int {
		return size[find(x)]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
		size[to] += size[from]
	}
	g := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		if np[x] && np[y] {
			merge(x, y)
		}
	}

	for x := 1; x <= n; x++ {
		if np[x] { // 跳过非质数
			continue
		}
		sum := 1
		for _, to := range g[x] { // 质数 x 把这棵树分成了若干个连通块
			if np[to] {
				ans += int64(sizeF(to) * sum)
				sum += sizeF(to)
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。