#### 题目

<p>给你一个 <code>n</code>&nbsp;个节点的 <strong>无向带权连通</strong>&nbsp;图，节点编号为&nbsp;<code>0</code>&nbsp;到&nbsp;<code>n - 1</code>&nbsp;，再给你一个整数数组&nbsp;<code>edges</code>&nbsp;，其中&nbsp;<code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>, w<sub>i</sub>]</code>&nbsp;表示节点&nbsp;<code>a<sub>i</sub></code> 和&nbsp;<code>b<sub>i</sub></code>&nbsp;之间有一条边权为&nbsp;<code>w<sub>i</sub></code>&nbsp;的边。</p>

<p>部分边的边权为&nbsp;<code>-1</code>（<code>w<sub>i</sub> = -1</code>），其他边的边权都为 <strong>正</strong>&nbsp;数（<code>w<sub>i</sub> &gt; 0</code>）。</p>

<p>你需要将所有边权为 <code>-1</code>&nbsp;的边都修改为范围&nbsp;<code>[1, 2 * 10<sup>9</sup>]</code>&nbsp;中的 <strong>正整数</strong>&nbsp;，使得从节点&nbsp;<code>source</code>&nbsp;到节点&nbsp;<code>destination</code>&nbsp;的 <strong>最短距离</strong>&nbsp;为整数&nbsp;<code>target</code>&nbsp;。如果有 <strong>多种</strong>&nbsp;修改方案可以使&nbsp;<code>source</code> 和&nbsp;<code>destination</code>&nbsp;之间的最短距离等于&nbsp;<code>target</code>&nbsp;，你可以返回任意一种方案。</p>

<p>如果存在使 <code>source</code>&nbsp;到 <code>destination</code>&nbsp;最短距离为 <code>target</code>&nbsp;的方案，请你按任意顺序返回包含所有边的数组（包括未修改边权的边）。如果不存在这样的方案，请你返回一个 <strong>空数组</strong>&nbsp;。</p>

<p><strong>注意：</strong>你不能修改一开始边权为正数的边。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2023/04/18/graph.png" style="width: 300px; height: 300px;" /></strong></p>

<pre>
<b>输入：</b>n = 5, edges = [[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]], source = 0, destination = 1, target = 5
<b>输出：</b>[[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
<b>解释：</b>上图展示了一个满足题意的修改方案，从 0 到 1 的最短距离为 5 。
</pre>

<p><strong>示例 2：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2023/04/18/graph-2.png" style="width: 300px; height: 300px;" /></strong></p>

<pre>
<b>输入：</b>n = 3, edges = [[0,1,-1],[0,2,5]], source = 0, destination = 2, target = 6
<b>输出：</b>[]
<b>解释：</b>上图是一开始的图。没有办法通过修改边权为 -1 的边，使得 0 到 2 的最短距离等于 6 ，所以返回一个空数组。
</pre>

<p><strong>示例 3：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2023/04/19/graph-3.png" style="width: 300px; height: 300px;" /></strong></p>

<pre>
<b>输入：</b>n = 4, edges = [[1,0,4],[1,2,3],[2,3,5],[0,3,-1]], source = 0, destination = 2, target = 6
<b>输出：</b>[[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
<b>解释：</b>上图展示了一个满足题意的修改方案，从 0 到 2 的最短距离为 6 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 100</code></li>
	<li><code>1 &lt;= edges.length &lt;= n * (n - 1) / 2</code></li>
	<li><code>edges[i].length == 3</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i&nbsp;</sub>&lt;&nbsp;n</code></li>
	<li><code>w<sub>i</sub>&nbsp;= -1</code> 或者 <code>1 &lt;= w<sub>i&nbsp;</sub>&lt;= 10<sup><span style="">7</span></sup></code></li>
	<li><code>a<sub>i&nbsp;</sub>!=&nbsp;b<sub>i</sub></code></li>
	<li><code>0 &lt;= source, destination &lt; n</code></li>
	<li><code>source != destination</code></li>
	<li><code>1 &lt;= target &lt;= 10<sup>9</sup></code></li>
	<li>输入的图是连通图，且没有自环和重边。</li>
</ul>

#### 思路

## 什么时候无解？

题目要求把边权为 $-1$ 的至少修改为 $1$。如果都修改成 $1$，跑最短路（Dijkstra），发现从起点到终点的最短路长度大于 $\textit{target}$，那么由于边权变大，最短路不可能变小，所以此时无解。

另外一种无解的情况是，如果都修改成无穷大（本题限制为 $2\cdot 10^9$），发现从起点到终点的最短路长度小于 $\textit{target}$，那么由于边权变小，最短路不可能变大，所以此时也无解。

## 一个错误的思路

先把 $-1$ 都修改成 $1$，然后跑 Dijkstra。设从起点到终点的最短路长度为 $d$。

如果 $d<\textit{target}$，**看上去**把其中一个 $1$ 加上 $\textit{target}-d$，就可以使从起点到终点的最短路长度恰好为 $\textit{target}$ 了。

这是不对的，因为在增加边权后，最短路可能就走别的边了，不走刚才修改的这条边了。

只修改一条边不行，那么修改两条边呢？也是不行的，最短路仍然可以走别的边。

## 正确思路

先把 $-1$ 都修改成 $1$，然后跑第一遍 Dijkstra，设从 $\textit{source}$ 到点 $i$ 的最短路长度为 $d_{i,0}$。

如果从起点到终点的最短路长度不超过 $\textit{target}$（否则无解，返回空数组），那我们就来尝试修改某些边权，使得从起点到终点的最短路长度恰好等于 $\textit{target}$。

这会引出两个问题：

1. 要按照什么样的**顺序**修改这些边？
2. 修改成多少合适？

正所谓「牵一发而动全身」，从上面对错误思路的分析可知，仅仅修改一个边权，就可能影响很多最短路的值。

那不妨再跑一遍 Dijkstra，由于 Dijkstra 算法保证每次拿到的点的最短路就是最终的最短路，所以按照 Dijkstra 算法遍历点/边的顺序去修改，就不会对**已确定的**最短路产生影响。

对于第二遍 Dijkstra，设从 $\textit{source}$ 到点 $i$ 的最短路长度为 $d_{i,1}$。

对于一条可以修改的边 $x-y$，假设要把它的边权改为 $w$，那么 $\textit{source}-x-y-\textit{destination}$ 这条路径由三部分组成：

1. 从 $\textit{source}$ 到 $x$ 的最短路，这是第二遍 Dijkstra 算出来的，即 $d_{x,1}$。
2. 从 $x$ 到 $y$，即 $w$。
3. 从 $y$ 到 $\textit{destination}$ 的最短路，由于后面的边还没有修改，这个最短路是第一遍 Dijkstra 算出来的，即 $d_{\textit{destination},0} - d_{y,0}$。注意这个式子仅当 $y$ 在从 $\textit{source}$ 到 $\textit{destination}$ 的最短路上才成立。不过，如果 $y$ 不在最短路上，修改 $x-y$ 并不会对最短路产生影响，所以代码中并没有判断 $y$ 是否在最短路上。

这三部分之和需要等于 $\textit{target}$，所以有

$$
d_{x,1} + w + d_{\textit{destination},0} - d_{y,0} = \textit{target}
$$

解得

$$
w = \textit{target} - d_{\textit{destination},0} + d_{y,0} - d_{x,1}
$$

> 注意上式中的 $\textit{target} - d_{\textit{destination},0}$ 是一个定值，代码中用 $\textit{delta}$ 表示。

根据「什么时候无解」中的分析，如果第二遍 Dijkstra 跑完后，从起点到终点的最短路仍然小于 $\textit{target}$，那么就说明无法修改，返回空数组。

否则，答案就是我们在第二遍 Dijkstra 中作出的修改。注意第二遍 Dijkstra 跑完后可能还有些边是 $-1$（因为在 $w=1$ 的时候没有修改，或者有些边不影响最短路），把这些边都改成 $1$ 就行。

代码实现时，为了修改边权，需要在邻接表中额外记录边的编号。

此外，由于输入最坏是稠密图，可以用朴素版的 Dijkstra 算法。


```go 

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$。在稠密图（本题最坏情况）中，算法的时间复杂度与边的数量 $m=\mathcal{O}(n^2)$ 成正比。
- 空间复杂度：$\mathcal{O}(m)$，其中 $m$ 为 $\textit{edges}$ 的长度。注意输入是连通图，$m$ 至少为 $n-1$，所以 $\mathcal{O}(n+m)=\mathcal{O}(m)$
