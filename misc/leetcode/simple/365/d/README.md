#### 题目

<p>现有一个有向图，其中包含 <code>n</code> 个节点，节点编号从 <code>0</code> 到 <code>n - 1</code> 。此外，该图还包含了 <code>n</code> 条有向边。</p>

<p>给你一个下标从 <strong>0</strong> 开始的数组 <code>edges</code> ，其中 <code>edges[i]</code> 表示存在一条从节点 <code>i</code> 到节点 <code>edges[i]</code> 的边。</p>

<p>想象在图上发生以下过程：</p>

<ul>
	<li>你从节点 <code>x</code> 开始，通过边访问其他节点，直到你在<strong> 此过程 </strong>中再次访问到之前已经访问过的节点。</li>
</ul>

<p>返回数组 <code>answer</code> 作为答案，其中 <code>answer[i]</code> 表示如果从节点 <code>i</code> 开始执行该过程，你可以访问到的不同节点数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/08/31/graaphdrawio-1.png" />
<pre>
<strong>输入：</strong>edges = [1,2,0,0]
<strong>输出：</strong>[3,3,3,4]
<strong>解释：</strong>从每个节点开始执行该过程，记录如下：
- 从节点 0 开始，访问节点 0 -&gt; 1 -&gt; 2 -&gt; 0 。访问的不同节点数是 3 。
- 从节点 1 开始，访问节点 1 -&gt; 2 -&gt; 0 -&gt; 1 。访问的不同节点数是 3 。
- 从节点 2 开始，访问节点 2 -&gt; 0 -&gt; 1 -&gt; 2 。访问的不同节点数是 3 。
- 从节点 3 开始，访问节点 3 -&gt; 0 -&gt; 1 -&gt; 2 -&gt; 0 。访问的不同节点数是 4 。
</pre>

<p><strong class="example">示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/08/31/graaph2drawio.png" style="width: 191px; height: 251px;" />
<pre>
<strong>输入：</strong>edges = [1,2,3,4,0]
<strong>输出：</strong>[5,5,5,5,5]
<strong>解释：</strong>无论从哪个节点开始，在这个过程中，都可以访问到图中的每一个节点。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == edges.length</code></li>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= edges[i] &lt;= n - 1</code></li>
	<li><code>edges[i] != i</code></li>
</ul>

#### 思路

本题给出的图叫做**内向基环树**。
之前写过一篇题解，介绍了处理基环树问题的一些通用技巧，请看 [内向基环树：拓扑排序+分类讨论](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/)  
对于本题来说：
- 对于在基环上的点，其可以访问到的节点数，就是基环的大小。
- 对于不在基环上的点 $x$，其可以访问到的节点数，是基环的大小，再加上点 $x$ 的深度。这里的深度是指以基环上的点 $\\textit{root}$ 为根的树枝作为一棵树，点 $x$ 在这棵树中的深度。这可以从 $\\textit{root}$ 出发，在反图上 DFS 得到。
  
注意题目给出的图可能不是连通的，可能有多棵内向基环树。

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