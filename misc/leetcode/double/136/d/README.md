#### 题目

<p>给你一棵 <strong>无向</strong>&nbsp;树，树中节点从 <code>0</code>&nbsp;到 <code>n - 1</code>&nbsp;编号。同时给你一个长度为 <code>n - 1</code>&nbsp;的二维整数数组&nbsp;<code>edges</code>&nbsp;，其中&nbsp;<code>edges[i] = [u<sub>i</sub>, v<sub>i</sub>]</code>&nbsp;表示节点&nbsp;<code>u<sub>i</sub></code> 和&nbsp;<code>v<sub>i</sub></code>&nbsp;在树中有一条边。</p>

<p>一开始，<strong>所有</strong>&nbsp;节点都 <strong>未标记</strong>&nbsp;。对于节点 <code>i</code>&nbsp;：</p>

<ul>
	<li>当&nbsp;<code>i</code>&nbsp;是奇数时，如果时刻 <code>x - 1</code>&nbsp;该节点有 <strong>至少</strong>&nbsp;一个相邻节点已经被标记了，那么节点 <code>i</code>&nbsp;会在时刻 <code>x</code>&nbsp;被标记。</li>
	<li>当&nbsp;<code>i</code>&nbsp;是偶数时，如果时刻 <code>x - 2</code>&nbsp;该节点有 <strong>至少</strong>&nbsp;一个相邻节点已经被标记了，那么节点 <code>i</code>&nbsp;会在时刻 <code>x</code>&nbsp;被标记。</li>
</ul>

<p>请你返回一个数组&nbsp;<code>times</code>&nbsp;，表示如果你在时刻 <code>t = 0</code>&nbsp;标记节点 <code>i</code>&nbsp;，那么时刻 <code>times[i]</code>&nbsp;时，树中所有节点都会被标记。</p>

<p>请注意，每个 <code>times[i]</code> 的答案都是独立的，即当你标记节点 <code>i</code> 时，所有其他节点都未标记。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>edges = [[0,1],[0,2]]</span></p>

<p><b>输出：</b>[2,4,3]</p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122236.png" style="width: 500px; height: 241px;" /></p>

<ul>
	<li>对于&nbsp;<code>i = 0</code>&nbsp;：
	<ul>
		<li>节点 1 在时刻&nbsp;<code>t = 1</code>&nbsp;被标记，节点 2 在时刻&nbsp;<code>t = 2</code>&nbsp;被标记。</li>
	</ul>
	</li>
	<li>对于&nbsp;<code>i = 1</code>&nbsp;：
	<ul>
		<li>节点 0 在时刻&nbsp;<code>t = 2</code>&nbsp;被标记，节点 2 在时刻&nbsp;<code>t = 4</code>&nbsp;被标记。</li>
	</ul>
	</li>
	<li>对于&nbsp;<code>i = 2</code>&nbsp;：
	<ul>
		<li>节点 0 在时刻&nbsp;<code>t = 2</code>&nbsp;被标记，节点 1 在时刻&nbsp;<code>t = 3</code>&nbsp;被标记。</li>
	</ul>
	</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>edges = [[0,1]]</span></p>

<p><b>输出：</b>[1,2]</p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122249.png" style="width: 500px; height: 257px;" /></p>

<ul>
	<li>对于&nbsp;<code>i = 0</code>&nbsp;：
	<ul>
		<li>节点 1 在时刻&nbsp;<code>t = 1</code>&nbsp;被标记。</li>
	</ul>
	</li>
	<li>对于&nbsp;<code>i = 1</code>&nbsp;：
	<ul>
		<li>节点 0 在时刻&nbsp;<code>t = 2</code>&nbsp;被标记。</li>
	</ul>
	</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>edges = </span>[[2,4],[0,1],[2,3],[0,2]]</p>

<p><b>输出：</b>[4,6,3,5,5]</p>

<p><b>解释：</b></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-210550.png" style="height: 266px; width: 500px;" /></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>0 &lt;= edges[i][0], edges[i][1] &lt;= n - 1</code></li>
	<li>输入保证&nbsp;<code>edges</code>&nbsp;表示一棵合法的树。</li>
</ul>

#### 思路

**前置知识**：[【图解】一张图秒懂换根 DP！](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/)

本题相当于对每个节点，计算以该节点为根时，树的最大深度。

其中从 $x\rightarrow y$ 的有向边的边权为 $(y+1)\bmod 2 + 1$，即当 $y$ 是奇数时，边权为 $1$；当 $y$ 是偶数时，边权为 $2$。

⚠**注意**：如果 $x$ 和 $y$ 的奇偶性不同，那么从 $x\rightarrow y$ 的有向边和从 $y\rightarrow x$ 的有向边的边权是不一样的。

考虑换根 DP。

首先，通过一次 DFS，计算以 $0$ 为根节点时，树的最大深度。

在 DFS 的过程中，额外保存：

- 子树 $x$ 的**最大**深度 $\textit{maxD}$。
- 子树 $x$ 的**次大**深度 $\textit{maxD}_2$。
- 子树 $x$ 通过其儿子 $\textit{my}$ 取到的最大深度。

然后，再通过一次 DFS，计算出本题的答案。

对于节点 $x$，其答案是以下两种情况的最大值：

- 子树 $x$ 的最大深度。
- $x$ 往上走到某个节点再往下拐弯的路径长度。

为了算出第二种情况，我们可以在 DFS 额外传入一个参数 $\textit{fromUp}$。

如果 $x$ 的儿子 $y = \textit{my}$，那么往下传入的参数更新为

$$
\max(\textit{fromUp}, \textit{maxD}_2 + (x+1)\bmod 2 + 1)
$$

如果 $x$ 的儿子 $y\ne \textit{my}$，那么往下传入的参数更新为

$$
\max(\textit{fromUp}, \textit{maxD} + (x+1)\bmod 2 + 1)
$$

注：我把[【图解】一张图秒懂换根 DP](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/) 这题叫做**第一类换根 DP**，本题需要额外维护次大信息，我称其为**第二类换根 DP**。

```
func timeTaken(edges [][]int) []int {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// nodes[x] 保存子树 x 的最大深度 maxD，次大深度 maxD2，以及最大深度要往儿子 y 走
	nodes := make([]struct{ maxD, maxD2, y int }, len(g))
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		p := &nodes[x]
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			maxD := dfs(y, x) + (y+1)%2 + 1 // 从 x 出发，往 y 方向的最大深度
			if maxD > p.maxD {
				p.maxD2 = p.maxD
				p.maxD = maxD
				p.y = y
			} else if maxD > p.maxD2 {
				p.maxD2 = maxD
			}
		}
		return p.maxD
	}
	dfs(0, -1)

	ans := make([]int, len(g))
	var reroot func(int, int, int)
	reroot = func(x, fa, fromUp int) {
		p := nodes[x]
		ans[x] = max(fromUp, p.maxD)
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			w := (x+1)%2 + 1 // 从 y 到 x 的边权
			if y == p.y { // 对于 y 来说，上面要选次大的
				reroot(y, x, max(fromUp, p.maxD2)+w)
			} else { // 对于 y 来说，上面要选最大的
				reroot(y, x, max(fromUp, p.maxD)+w)
			}
		}
	}
	reroot(0, -1, 0)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)