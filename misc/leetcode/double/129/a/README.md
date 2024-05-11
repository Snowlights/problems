### 题目

<p>给你一个二维 <code>3 x 3</code>&nbsp;的矩阵&nbsp;<code>grid</code>&nbsp;，每个格子都是一个字符，要么是&nbsp;<code>'B'</code>&nbsp;，要么是&nbsp;<code>'W'</code>&nbsp;。字符&nbsp;<code>'W'</code>&nbsp;表示白色，字符&nbsp;<code>'B'</code>&nbsp;表示黑色。</p>

<p>你的任务是改变 <strong>至多一个</strong>&nbsp;格子的颜色，使得矩阵中存在一个 <code>2 x 2</code>&nbsp;颜色完全相同的正方形。<!-- notionvc: adf957e1-fa0f-40e5-9a2e-933b95e276a7 --></p>

<p>如果可以得到一个相同颜色的 <code>2 x 2</code>&nbsp;正方形，那么返回 <code>true</code>&nbsp;，否则返回 <code>false</code>&nbsp;。</p>

<p>&nbsp;</p>
<style type="text/css">.grid-container {
  display: grid;
  grid-template-columns: 30px 30px 30px;
  padding: 10px;
}
.grid-item {
  background-color: black;
  border: 1px solid gray;
  height: 30px;
  font-size: 30px;
  text-align: center;
}
.grid-item-white {
  background-color: white;
}
</style>
<style class="darkreader darkreader--sync" media="screen" type="text/css">
</style>
<p><strong class="example">示例 1：</strong></p>

<div class="grid-container">
<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>
</div>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [["B","W","B"],["B","W","W"],["B","W","B"]]</span></p>

<p><span class="example-io"><b>输出：</b>true</span></p>

<p><strong>解释：</strong></p>

<p>修改&nbsp;<code>grid[0][2]</code> 的颜色，可以满足要求。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="grid-container">
<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>
</div>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [["B","W","B"],["W","B","W"],["B","W","B"]]</span></p>

<p><span class="example-io"><b>输出：</b>false</span></p>

<p><strong>解释：</strong></p>

<p>只改变一个格子颜色无法满足要求。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="grid-container">
<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>

<div class="grid-item grid-item-white">&nbsp;</div>
</div>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [["B","W","B"],["B","W","W"],["B","W","W"]]</span></p>

<p><span class="example-io"><b>输出：</b>true</span></p>

<p><strong>解释：</strong></p>

<p><code>grid</code>&nbsp;已经包含一个&nbsp;<code>2 x 2</code>&nbsp;颜色相同的正方形了。<!-- notionvc: 9a8b2d3d-1e73-457a-abe0-c16af51ad5c2 --></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>grid.length == 3</code></li>
	<li><code>grid[i].length == 3</code></li>
	<li><code>grid[i][j]</code>&nbsp;要么是&nbsp;<code>'W'</code>&nbsp;，要么是&nbsp;<code>'B'</code> 。</li>
</ul>

### 思路

模拟

``` go  
func canMakeSquare(grid [][]byte) bool {
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			bCnt := 0
			if grid[i][j] == 'B' {
				bCnt++
			}
			if grid[i][j+1] == 'B' {
				bCnt++
			}
			if grid[i+1][j] == 'B' {
				bCnt++
			}
			if grid[i+1][j+1] == 'B' {
				bCnt++
			}
			if bCnt <= 1 || bCnt >= 3 {
				return true
			}
		}
	}
	return false
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)