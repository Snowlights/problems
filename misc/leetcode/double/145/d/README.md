#### 题目

<p>给你一个整数数组 <code>nums</code>，它表示一个循环数组。请你遵循以下规则创建一个大小&nbsp;<strong>相同&nbsp;</strong>的新数组 <code>result</code>&nbsp;：</p>
对于每个下标&nbsp;<code>i</code>（其中 <code>0 &lt;= i &lt; nums.length</code>），独立执行以下操作：

<ul>
	<li>如果 <code>nums[i] &gt; 0</code>：从下标&nbsp;<code>i</code> 开始，向&nbsp;<strong>右&nbsp;</strong>移动 <code>nums[i]</code> 步，在循环数组中落脚的下标对应的值赋给 <code>result[i]</code>。</li>
	<li>如果 <code>nums[i] &lt; 0</code>：从下标&nbsp;<code>i</code> 开始，向&nbsp;<strong>左&nbsp;</strong>移动 <code>abs(nums[i])</code> 步，在循环数组中落脚的下标对应的值赋给 <code>result[i]</code>。</li>
	<li>如果 <code>nums[i] == 0</code>：将 <code>nums[i]</code> 的值赋给 <code>result[i]</code>。</li>
</ul>

<p>返回新数组 <code>result</code>。</p>

<p><strong>注意：</strong>由于 <code>nums</code> 是循环数组，向右移动超过最后一个元素时将回到开头，向左移动超过第一个元素时将回到末尾。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [3,-2,1,1]</span></p>

<p><strong>输出：</strong> <span class="example-io">[1,1,1,3]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>对于 <code>nums[0]</code> 等于 3，向右移动 3 步到 <code>nums[3]</code>，因此 <code>result[0]</code> 为 1。</li>
	<li>对于 <code>nums[1]</code> 等于 -2，向左移动 2 步到 <code>nums[3]</code>，因此 <code>result[1]</code> 为 1。</li>
	<li>对于 <code>nums[2]</code> 等于 1，向右移动 1 步到 <code>nums[3]</code>，因此 <code>result[2]</code> 为 1。</li>
	<li>对于 <code>nums[3]</code> 等于 1，向右移动 1 步到 <code>nums[0]</code>，因此 <code>result[3]</code> 为 3。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [-1,4,-1]</span></p>

<p><strong>输出：</strong> <span class="example-io">[-1,-1,4]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>对于 <code>nums[0]</code> 等于 -1，向左移动 1 步到 <code>nums[2]</code>，因此 <code>result[0]</code> 为 -1。</li>
	<li>对于 <code>nums[1]</code> 等于 4，向右移动 4 步到 <code>nums[2]</code>，因此 <code>result[1]</code> 为 -1。</li>
	<li>对于 <code>nums[2]</code> 等于 -1，向左移动 1 步到 <code>nums[1]</code>，因此 <code>result[2]</code> 为 4。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>-100 &lt;= nums[i] &lt;= 100</code></li>
</ul>


#### 思路

下文把 $\textit{threshold}$ 简称为 $t$。

$$
\text{LCM}(x,y)\le t
$$

等价于

$$
\dfrac{x\cdot y}{\text{GCD}(x,y)} \le t
$$

设 $g = \text{GCD}(x,y)$。

**核心思路**：枚举 $g=1,2,3,\ldots,t$，以及 $g$ 的倍数 $x$ 和 $y$，对于满足 $x\cdot y\le g\cdot t$ 的 $x$ 和 $y$，把它们的下标用**并查集**连起来。

但是，$g$ 的倍数有 $\mathcal{O}(t/g)$ 个，枚举 $x$ 和 $y$ 需要 $\mathcal{O}((t/g)^2)$ 的时间，会超时。

解决办法：不需要枚举 $x$，而是找到在 $\textit{nums}$ 中的**最小的** $g$ 的倍数，作为 $x$。对于其他的 $g$ 的倍数 $y$，只要它们能和 $x$ 连起来，那么这些 $y$ 就已经在同一个连通块中了。

> 选择最小的 $x$，是因为这样做，满足 $x\cdot y\le g\cdot t$ 的 $y$ 就尽量多，我们可以把能在同一个连通块中的 $y$ 都用并查集合并。

注意题目保证所有元素互不相同。

## 答疑

**问**：如果枚举的过程中，出现 $g=2,x=4,y=8$ 的情况怎么办？此时 $g\ne \text{GCD}(x,y)$。

**答**：虽然 $\text{GCD}$ 对不上，但不影响正确性。如果此时 $\dfrac{x\cdot y}{g} \le t$ 成立，那么 $\dfrac{x\cdot y}{\text{GCD}(x,y)} \le t$ 也必然成立。

```
func countComponents(nums []int, threshold int) int {
	n := len(nums)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	// 记录每个数的下标
	idx := make([]int, threshold+1)
	for i, x := range nums {
		if x <= threshold {
			idx[x] = i + 1 // 这里 +1 了，下面减掉
		}
	}

	for g := 1; g <= threshold; g++ {
		minX := -1
		for x := g; x <= threshold; x += g {
			if idx[x] > 0 { // idx[x] == 0 表示不存在
				minX = x
				break
			}
		}
		if minX < 0 {
			continue
		}
		fi := find(idx[minX] - 1)
		for y := minX + g; y <= threshold && y <= g*threshold/minX; y += g {
			if idx[y] > 0 {
				fj := find(idx[y] - 1)
				if fj != fi {
					fa[fj] = fi // 合并 idx[x] 和 idx[y]
					n-- // 连通块个数减一
				}
			}
		}
	}
	return n
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + \alpha t\log t)$，其中 $n$ 是 $\textit{nums}$ 的长度，$t$ 是 $\textit{threshold}$，$\alpha$ 是并查集单次合并的均摊复杂度。根据调和级数，二重循环（不考虑并查集）的循环次数为 $\mathcal{O}(t\log t)$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + t)$，取决于实现。

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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)