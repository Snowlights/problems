#### 题目

<p>你正在维护一个项目，该项目有 <code>n</code> 个方法，编号从 <code>0</code> 到 <code>n - 1</code>。</p>

<p>给你两个整数 <code>n</code> 和 <code>k</code>，以及一个二维整数数组 <code>invocations</code>，其中 <code>invocations[i] = [a<sub>i</sub>, b<sub>i</sub>]</code> 表示方法 <code>a<sub>i</sub></code> 调用了方法 <code>b<sub>i</sub></code>。</p>

<p>已知如果方法 <code>k</code> 存在一个已知的 bug。那么方法 <code>k</code> 以及它直接或间接调用的任何方法都被视为<strong> </strong><strong>可疑方法</strong> ，我们需要从项目中移除这些方法。</p>

<p><span class="text-only" data-eleid="13" style="white-space: pre;">只有当一组方法没有被这组之外的任何方法调用时，这组方法才能被移除。</span></p>

<p>返回一个数组，包含移除所有<strong> </strong><strong>可疑方法</strong> 后剩下的所有方法。你可以以任意顺序返回答案。如果无法移除<strong> 所有 </strong>可疑方法，则<strong> 不 </strong>移除任何方法。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1:</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">n = 4, k = 1, invocations = [[1,2],[0,1],[3,2]]</span></p>

<p><strong>输出:</strong> <span class="example-io">[0,1,2,3]</span></p>

<p><strong>解释:</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/18/graph-2.png" style="width: 200px; height: 200px;" /></p>

<p>方法 2 和方法 1 是可疑方法，但它们分别直接被方法 3 和方法 0 调用。由于方法 3 和方法 0 不是可疑方法，我们无法移除任何方法，故返回所有方法。</p>
</div>

<p><strong class="example">示例 2:</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">n = 5, k = 0, invocations = [[1,2],[0,2],[0,1],[3,4]]</span></p>

<p><strong>输出:</strong> <span class="example-io">[3,4]</span></p>

<p><strong>解释:</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/18/graph-3.png" style="width: 200px; height: 200px;" /></p>

<p>方法 0、方法 1 和方法 2 是可疑方法，且没有被任何其他方法直接调用。我们可以移除它们。</p>
</div>

<p><strong class="example">示例 3:</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">n = 3, k = 2, invocations = [[1,2],[0,1],[2,0]]</span></p>

<p><strong>输出:</strong> <span class="example-io">[]</span></p>

<p><strong>解释:</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/20/graph.png" style="width: 200px; height: 200px;" /></p>

<p>所有方法都是可疑方法。我们可以移除它们。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= k &lt;= n - 1</code></li>
	<li><code>0 &lt;= invocations.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>invocations[i] == [a<sub>i</sub>, b<sub>i</sub>]</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i</sub> &lt;= n - 1</code></li>
	<li><code>a<sub>i</sub> != b<sub>i</sub></code></li>
	<li><code>invocations[i] != invocations[j]</code></li>
</ul>

#### 思路

## 题意

一个项目有 $n$ 个方法，其中第 $k$ 个方法有 bug。可能是第 $k$ 个方法自己的 bug，也可能是第 $k$ 个方法调用的更底层的方法有 bug。

你需要删除所有可能有 bug 的方法。如果删除后无法编译（剩余的方法调用了删除的方法），那么返回数组 $[0,1,2,\cdots,n-1]$。

如果可以正常删除，返回剩余的方法编号。

## 思路

1. 从 $k$ 开始 DFS 图，标记所有可能有 bug 的方法（节点）。题目把这些方法叫做**可疑方法**。
2. 遍历 $\textit{invocations}$，如果存在从「非可疑方法」到「可疑方法」的边，则删除后无法编译，返回数组 $[0,1,2,\cdots,n-1]$。
3. 否则可以正常删除，把非可疑方法加入答案。

⚠**注意**：图中可能有环，为避免 DFS 无限递归下去，只需 DFS 没有访问过（没有被标记）的节点。

```
func remainingMethods(n, k int, invocations [][]int) (ans []int) {
	g := make([][]int, n)
	for _, e := range invocations {
		g[e[0]] = append(g[e[0]], e[1])
	}

	// 标记所有可疑方法
	isSuspicious := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		isSuspicious[x] = true
		for _, y := range g[x] {
			if !isSuspicious[y] { // 避免无限递归
				dfs(y)
			}
		}
	}
	dfs(k)

	// 检查是否有【非可疑方法】->【可疑方法】的边
	for _, e := range invocations {
		if !isSuspicious[e[0]] && isSuspicious[e[1]] {
			// 无法移除可疑方法
			for i := range n {
				ans = append(ans, i)
			}
			return
		}
	}

	// 移除所有可疑方法
	for i, b := range isSuspicious {
		if !b {
			ans = append(ans, i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 是 $\textit{invocations}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。
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