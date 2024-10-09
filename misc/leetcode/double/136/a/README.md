### 题目

<p>给你一个整数&nbsp;<code>n</code>&nbsp;，表示在一个游戏中的玩家数目。同时给你一个二维整数数组&nbsp;<code>pick</code>&nbsp;，其中&nbsp;<code>pick[i] = [x<sub>i</sub>, y<sub>i</sub>]</code>&nbsp;表示玩家&nbsp;<code>x<sub>i</sub></code>&nbsp;获得了一个颜色为&nbsp;<code>y<sub>i</sub></code>&nbsp;的球。</p>

<p>如果玩家 <code>i</code>&nbsp;获得的球中任何一种颜色球的数目 <strong>严格大于</strong>&nbsp;<code>i</code>&nbsp;个，那么我们说玩家 <code>i</code>&nbsp;是胜利玩家。换句话说：</p>

<ul>
	<li>如果玩家 0 获得了任何的球，那么玩家 0 是胜利玩家。</li>
	<li>如果玩家 1 获得了至少 2 个相同颜色的球，那么玩家 1 是胜利玩家。</li>
	<li>...</li>
	<li>如果玩家 <code>i</code>&nbsp;获得了至少&nbsp;<code>i + 1</code>&nbsp;个相同颜色的球，那么玩家 <code>i</code>&nbsp;是胜利玩家。</li>
</ul>

<p>请你返回游戏中 <strong>胜利玩家</strong>&nbsp;的数目。</p>

<p><strong>注意</strong>，可能有多个玩家是胜利玩家。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 4, pick = [[0,0],[1,0],[1,0],[2,1],[2,1],[2,0]]</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>玩家 0 和玩家 1 是胜利玩家，玩家 2 和玩家 3 不是胜利玩家。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 5, pick = [[1,1],[1,2],[1,3],[1,4]]</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><strong>解释：</strong></p>

<p>没有胜利玩家。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 5, pick = [[1,1],[2,4],[2,4],[2,4]]</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><b>解释：</b></p>

<p>玩家 2 是胜利玩家，因为玩家 2 获得了 3 个颜色为 4 的球。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 10</code></li>
	<li><code>1 &lt;= pick.length &lt;= 100</code></li>
	<li><code>pick[i].length == 2</code></li>
	<li><code>0 &lt;= x<sub>i</sub> &lt;= n - 1 </code></li>
	<li><code>0 &lt;= y<sub>i</sub> &lt;= 10</code></li>
</ul>


### 思路

遍历 $\textit{pick}$，用一个 $n\times 11$ 大小的矩阵，统计每个玩家得到的每种颜色的球的个数。

遍历每个玩家，如果该玩家至少有一种颜色的球大于玩家编号，则把答案加一。

```
func winningPlayerCount(n int, pick [][]int) (ans int) {
	cnts := make([][11]int, n)
	for _, p := range pick {
		cnts[p[0]][p[1]]++
	}
	for i, cnt := range cnts {
		for _, c := range cnt {
			if c > i {
				ans++
				break
			}
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU+m)$，其中 $m$ 是 $\textit{pick}$ 的长度，$U$ 是 $y_i$ 的最大值。
- 空间复杂度：$\mathcal{O}(nU)$。

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