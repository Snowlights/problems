#### 题目

<p>给你一个整数数组 <code>groups</code>，其中 <code>groups[i]</code> 表示第 <code>i</code> 组的大小。另给你一个整数数组 <code>elements</code>。</p>

<p>请你根据以下规则为每个组分配&nbsp;<strong>一个&nbsp;</strong>元素：</p>

<ul>
	<li>如果 <code>groups[i]</code> 能被 <code>elements[j]</code> 整除，则元素 <code>j</code> 可以分配给组 <code>i</code>。</li>
	<li>如果有多个元素满足条件，则分配下标最小的元素 &nbsp;<code>j</code> 。</li>
	<li>如果没有元素满足条件，则分配 -1 。</li>
</ul>

<p>返回一个整数数组 <code>assigned</code>，其中 <code>assigned[i]</code> 是分配给组 <code>i</code> 的元素的索引，若无合适的元素，则为 -1。</p>

<p><strong>注意：</strong>一个元素可以分配给多个组。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">groups = [8,4,3,2,4], elements = [4,2]</span></p>

<p><strong>输出：</strong> <span class="example-io">[0,0,-1,1,0]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><code>elements[0] = 4</code> 被分配给组 0、1 和 4。</li>
	<li><code>elements[1] = 2</code> 被分配给组 3。</li>
	<li>无法为组 2 分配任何元素，分配 -1 。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">groups = [2,3,5,7], elements = [5,3,3]</span></p>

<p><strong>输出：</strong> <span class="example-io">[-1,1,0,-1]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><code>elements[1] = 3</code> 被分配给组 1。</li>
	<li><code>elements[0] = 5</code> 被分配给组 2。</li>
	<li>无法为组 0 和组 3 分配任何元素，分配 -1 。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">groups = [10,21,30,41], elements = [2,1]</span></p>

<p><strong>输出：</strong> <span class="example-io">[0,1,0,1]</span></p>

<p><strong>解释：</strong></p>

<p><code>elements[0] = 2</code> 被分配给所有偶数值的组，而 <code>elements[1] = 1</code> 被分配给所有奇数值的组。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= groups.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= elements.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= groups[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= elements[i] &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

设 $\textit{groups}$ 中的最大值为 $\textit{mx}$。我们直接预处理 $1,2,3,\ldots,\textit{mx}$ 中的每个数能被哪个 $\textit{elements}[i]$ 整除。
如果有多个相同的 $\textit{elements}[i]$，只考虑最左边的那个。

从左到右遍历 $\textit{elements}$，设 $x=\textit{elements}[i]$。枚举 $x$ 的倍数 $y$，标记 $y$ 可以被下标为 $i$ 的元素整除， 记作 $\textit{target}[y]=i$。标记过的数字不再重复标记。

⚠**注意**：如果我们之前遍历过 $x$ 的因子 $d$，那么不用枚举 $x$ 的倍数，因为这些数必然已被 $d$ 标记。例如 $\textit{elements}=[2,4]$，由于 $4$ 的倍数一定都是偶数（$2$ 的倍数），所以 $4$ 的倍数一定都被 $2$ 标记，所以无需枚举 $4$ 的倍数。

> 这个做法也保证了每个数 $x$ 我们只会循环枚举其倍数一次，从而保证时间复杂度不会退化成暴力。

最后，回答询问，对于 $\textit{groups}[i]$ 来说，答案为 $\textit{target}[\textit{groups}[i]]$。
**小技巧**：为了方便计算 $-1$ 的情况，可以初始化 $\textit{target}[y]=-1$。

```
func assignElements(groups []int, elements []int) []int {
	mx := slices.Max(groups)
	target := make([]int, mx+1)
	for i := range target {
		target[i] = -1
	}

	for i, x := range elements {
		if x > mx || target[x] >= 0 { // x 及其倍数一定已被标记，跳过
			continue
		}
		for y := x; y <= mx; y += x { // 枚举 x 的倍数 y
			if target[y] < 0 { // 没有标记过
				target[y] = i // 标记 y 可以被 x 整除
			}
		}
	}

	// 回答询问
	for i, x := range groups {
		groups[i] = target[x] // 原地修改
	}
	return groups
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(U\log n + m)$，其中 $n$ 是 $\textit{elements}$ 的长度，$m$ 是 $\textit{groups}$ 的长度，$U=\max(\textit{groups})$。代码中的二重循环，根据**调和级数**可得，循环次数为 $\mathcal{O}(U\log n)$。
- 空间复杂度：$\mathcal{O}(U)$。返回值不计入。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)