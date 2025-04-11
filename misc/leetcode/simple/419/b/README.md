#### 题目

<p>给你一棵 <strong>二叉树 </strong>的根节点 <code>root</code> 和一个整数<code>k</code>。</p>

<p>返回第 <code>k</code> 大的 <strong>完美二叉</strong><span data-keyword="subtree"><strong>子树</strong> </span>的大小，如果不存在则返回 <code>-1</code>。</p>

<p><strong>完美二叉树 </strong>是指所有叶子节点都在同一层级的树，且每个父节点恰有两个子节点。</p>

<p><strong>子树 </strong>是指树中的某一个节点及其所有后代形成的树。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">root = [5,3,6,5,2,5,7,1,8,null,null,6,8], k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/21/image.jpg" style="width: 300px; height: 175px;" /></p>

<p>完美二叉子树的根节点在图中以黑色突出显示。它们的大小按降序排列为 <code>[3, 3, 1, 1, 1, 1, 1, 1]</code>。<br />
第 <code>2</code> 大的完美二叉子树的大小是 3。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">root = [1,2,3,4,5,6,7], k = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">7</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/21/image1.jpg" style="width: 300px; height: 149px;" /></p>

<p>完美二叉子树的大小按降序排列为 <code>[7, 3, 3, 1, 1, 1, 1]</code>。最大的完美二叉子树的大小是 7。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">root = [1,2,3,null,4], k = 3</span></p>

<p><strong>输出：</strong> <span class="example-io">-1</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/21/image4.jpg" style="width: 150px; height: 130px;" /></p>

<p>完美二叉子树的大小按降序排列为 <code>[1, 1]</code>。完美二叉子树的数量少于 3。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中的节点数目在 <code>[1, 2000]</code> 范围内。</li>
	<li><code>1 &lt;= Node.val &lt;= 2000</code></li>
	<li><code>1 &lt;= k &lt;= 1024</code></li>
</ul>

#### 思路

根据完美二叉树的定义，一棵高为 $h$ 的完美二叉树，从上往下，每一层恰好有 $1,2,4,8,\cdots,2^{h-1}$ 个节点，所以子树大小为

$$
1+2+4+8+\cdots+2^{h-1} = 2^h-1
$$

由于知道高度就知道了子树大小，DFS 只需返回子树高度。

DFS 的同时，用一个数组 $\textit{hs}$ 维护合法子树的高度。
分类讨论：
- 如果当前节点是空节点，返回 $0$。
- 否则，判断左右子树的高度是否都不为 $-1$ 且相同，如果不是，返回 $-1$；如果是，把子树高度加一加入 $\textit{hs}$，然后返回子树高度加一。


DFS 结束后：
- 设 $m$ 为 $\textit{hs}$ 的大小。
- 如果 $k>m$，返回 $-1$。
- 否则设 $\textit{hs}$ 从大到小排序后的下标为 $k-1$ 的元素为 $h$（也可以是从小到大排序后的下标为 $m-k$ 的元素），返回 $2^h - 1$。

```
func kthLargestPerfectSubtree(root *leetcode.TreeNode, k int) int {
	hs := []int{}
	var dfs func(*leetcode.TreeNode) int
	dfs = func(node *leetcode.TreeNode) int {
		if node == nil {
			return 0
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH < 0 || leftH != rightH {
			return -1
		}
		hs = append(hs, leftH+1)
		return leftH + 1
	}
	dfs(root)

	if k > len(hs) {
		return -1
	}
	slices.Sort(hs)
	return 1<<hs[len(hs)-k] - 1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\log m)$ 或者 $\mathcal{O}(n)$，其中 $n$ 是二叉树的节点个数，$m$ 是 $\textit{hs}$ 的长度。如果使用快速选择，则时间复杂度为 $\mathcal{O}(n)$，
- 空间复杂度：$\mathcal{O}(n)$。递归需要 $\mathcal{O}(n)$ 的栈空间。
- 
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)