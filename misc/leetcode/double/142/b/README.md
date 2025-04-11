### 题目

<p>给你一棵 <code>n</code>&nbsp;个节点且根节点为编号 0 的树，节点编号为&nbsp;<code>0</code>&nbsp;到&nbsp;<code>n - 1</code>&nbsp;。这棵树用一个长度为&nbsp;<code>n</code>&nbsp;的数组&nbsp;<code>parent</code>&nbsp;表示，其中&nbsp;<code>parent[i]</code>&nbsp;是第 <code>i</code>&nbsp;个节点的父亲节点的编号。由于节点 0 是根，<code>parent[0] == -1</code>&nbsp;。</p>

<p>给你一个长度为 <code>n</code>&nbsp;的字符串&nbsp;<code>s</code>&nbsp;，其中&nbsp;<code>s[i]</code>&nbsp;是节点 <code>i</code>&nbsp;对应的字符。</p>

<p>对于节点编号从 <code>1</code>&nbsp;到 <code>n - 1</code>&nbsp;的每个节点 <code>x</code>&nbsp;，我们 <strong>同时</strong> 执行以下操作 <strong>一次</strong>&nbsp;：</p>

<ul>
	<li>找到距离节点 <code>x</code>&nbsp;<strong>最近</strong>&nbsp;的祖先节点 <code>y</code>&nbsp;，且&nbsp;<code>s[x] == s[y]</code>&nbsp;。</li>
	<li>如果节点 <code>y</code>&nbsp;不存在，那么不做任何修改。</li>
	<li>否则，将节点 <code>x</code>&nbsp;与它父亲节点之间的边 <strong>删除</strong>&nbsp;，在 <code>x</code>&nbsp;与 <code>y</code>&nbsp;之间连接一条边，使&nbsp;<code>y</code>&nbsp;变为 <code>x</code>&nbsp;新的父节点。</li>
</ul>

<p>请你返回一个长度为 <code>n</code>&nbsp;的数组&nbsp;<code>answer</code>&nbsp;，其中&nbsp;<code>answer[i]</code>&nbsp;是 <strong>最终</strong>&nbsp;树中，节点 <code>i</code>&nbsp;为根的子树的 <strong>大小</strong>&nbsp;。</p>

<p>一个 <strong>子树</strong>&nbsp;<code>subtree</code>&nbsp;指的是节点 <code>subtree</code>&nbsp;和它所有的后代节点。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>parent = [-1,0,0,1,1,1], s = "abaabc"</span></p>

<p><span class="example-io"><b>输出：</b>[6,3,1,1,1,1]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/08/15/graphex1drawio.png" style="width: 230px; height: 277px;" /></p>

<p>节点 3 的父节点从节点 1 变为节点 0 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>parent = [-1,0,4,0,1], s = "abbba"</span></p>

<p><span class="example-io"><b>输出：</b>[5,2,1,1,1]</span></p>

<p><b>解释：</b></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/08/20/exgraph2drawio.png" style="width: 160px; height: 308px;" /></p>

<p>以下变化会同时发生：</p>

<ul>
	<li>节点 4 的父节点从节点 1 变为节点 0 。</li>
	<li>节点 2 的父节点从节点 4 变为节点 1 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == parent.length == s.length</code></li>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li>对于所有的&nbsp;<code>i &gt;= 1</code>&nbsp;，都有&nbsp;<code>0 &lt;= parent[i] &lt;= n - 1</code>&nbsp;。</li>
	<li><code>parent[0] == -1</code></li>
	<li><code>parent</code>&nbsp;表示一棵合法的树。</li>
	<li><code>s</code>&nbsp;只包含小写英文字母。</li>
</ul>

### 思路

## 方法一：两次 DFS

不了解递归的同学请先看[【基础算法精讲 09】](https://www.bilibili.com/video/BV1UD4y1Y769/)。

### 第一次 DFS

写一个自顶向下的 DFS，同时维护路径上的每个字母对应的最深节点 $\textit{ancestor}$，初始化成一个全为 $-1$ 的数组（或者哈希表）。递归到节点 $x$ 时，更新 $\textit{ancestor}[s[x]]=x$。

例如示例 1，递归到节点 $1$ 时，有 $\textit{ancestor}[\texttt{a}]=0,\ \textit{ancestor}[\texttt{b}]=1$。

对于节点 $x$ 及其儿子 $y$，如果 $\textit{ancestor}[s[y]]\ne -1$，则把 $y$ 加到 $\textit{ancestor}[s[y]]$ 的子节点列表中，同时把 $x$ 的儿子 $y$ 改成 $-1$，在第二次 DFS 中，不去递归等于 $-1$ 的儿子。

注意要写 [回溯](https://www.bilibili.com/video/BV1mG4y1A7Gu/)，递归返回前，要恢复 $\textit{ancestor}[s[x]]$ 的原始值（递归到 $x$ 之前的值）。

### 第二次 DFS

写一个自底向上的 DFS，返回子树大小。

当前子树大小为 $1$（自己）加上所有儿子子树的大小之和。

⚠**注意**：在遍历列表的同时，往列表中加入数据，会导致我们遍历到新加入列表的数据。可以考虑固定遍历次数为 $x$ 的子节点列表的长度。

> Go 语言不用考虑这个问题。

```
func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	ancestor := [26]int{}
	for i := range ancestor {
		ancestor[i] = -1
	}
	var rebuild func(int)
	rebuild = func(x int) {
		sx := s[x] - 'a'
		old := ancestor[sx]
		ancestor[sx] = x
		for i, y := range g[x] {
			if anc := ancestor[s[y]-'a']; anc != -1 {
				g[anc] = append(g[anc], y)
				g[x][i] = -1 // -1 表示删除 y
			}
			rebuild(y)
		}
		ancestor[sx] = old // 恢复现场
	}
	rebuild(0)

	size := make([]int, n)
	var dfs func(int)
	dfs = func(x int) {
		size[x] = 1
		for _, y := range g[x] {
			if y != -1 { // y 没被删除
				dfs(y)
				size[x] += size[y]
			}
		}
	}
	dfs(0)
	return size
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：一次 DFS

把两个 DFS 结合起来。在 $\textit{dfs}(y)$ 结束后，分类讨论：

- 如果 $y$ 没有对应的祖先，把 $\textit{size}[x]$ 增加 $\textit{size}[y]$。
- 如果 $y$ 有对应的祖先 $\textit{anc}$，把 $\textit{size}[\textit{anc}]$ 增加 $\textit{size}[y]$。

```
func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	size := make([]int, n)
	ancestor := [26]int{}
	for i := range ancestor {
		ancestor[i] = -1
	}
	var dfs func(int)
	dfs = func(x int) {
		size[x] = 1
		sx := s[x] - 'a'
		old := ancestor[sx]
		ancestor[sx] = x
		for _, y := range g[x] {
			dfs(y)
			anc := ancestor[s[y]-'a']
			if anc < 0 {
				anc = x
			}
			size[anc] += size[y]
		}
		ancestor[sx] = old // 恢复现场
	}
	dfs(0)
	return size
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见 [链表、二叉树与一般树](https://leetcode.cn/circle/discuss/K0n2gO/) 中的「**§2.2 自顶向下 DFS**」和「**§2.3 自底向上 DFS**」。

## 题单


- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)