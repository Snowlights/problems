#### 题目

<p>现有一棵 <strong>无向</strong> 树，树中包含 <code>n</code> 个节点，按从 <code>0</code> 到 <code>n - 1</code> 标记。树的根节点是节点 <code>0</code> 。给你一个长度为 <code>n - 1</code> 的二维整数数组 <code>edges</code>，其中 <code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>]</code> 表示树中节点 <code>a<sub>i</sub></code> 与节点 <code>b<sub>i</sub></code> 之间存在一条边。</p>

<p>如果一个节点的所有子节点为根的&nbsp;<span data-keyword="subtree">子树</span>&nbsp;包含的节点数相同，则认为该节点是一个 <strong>好节点</strong>。</p>

<p>返回给定树中<strong> 好节点 </strong>的数量。</p>

<p><strong>子树</strong>&nbsp;指的是一个节点以及它所有后代节点构成的一棵树。</p>

<p>&nbsp;</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]</span></p>

<p><strong>输出：</strong><span class="example-io">7</span></p>

<p><strong>说明：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/05/26/tree1.png" style="width: 360px; height: 158px;" />
<p>树的所有节点都是好节点。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">edges = [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]</span></p>

<p><strong>输出：</strong><span class="example-io">6</span></p>

<p><strong>说明：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-193552.png" style="width: 360px; height: 303px;" />
<p>树中有 6 个好节点。上图中已将这些节点着色。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>edges = [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]</span></p>

<p><span class="example-io"><b>输出：</b>12</span></p>

<p><strong>解释：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/08/08/rob.jpg" style="width: 450px; height: 277px;" />
<p>除了节点 9 以外其他所有节点都是好节点。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i</sub> &lt; n</code></li>
	<li>输入确保 <code>edges</code> 总表示一棵有效的树。</li>
</ul>

#### 思路

建树，然后从根节点 $0$ 开始 DFS 这棵树。

DFS 返回子树大小。

对于节点 $x$，如果其是叶子节点，或者其所有儿子子树大小都一样，那么答案加一。


```
func countGoodNodes(edges [][]int) (ans int) {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		size, pre, ok := 1, 0, true
		for _, y := range g[x] {
			if y == fa {
				continue // 不能递归到父节点
			}
			sz := dfs(y, x)
			if pre > 0 && sz != pre {
				ok = false
			}
			pre = sz // 记录上一个儿子子树的大小
			size += sz
		}
		if ok {
			ans++
		}
		return size
	}
	dfs(0, -1)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)