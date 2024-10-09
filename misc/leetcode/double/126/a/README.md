### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;，数组中的元素都是&nbsp;<strong>正</strong>&nbsp;整数。定义一个加密函数&nbsp;<code>encrypt</code>&nbsp;，<code>encrypt(x)</code>&nbsp;将一个整数 <code>x</code>&nbsp;中 <strong>每一个</strong>&nbsp;数位都用 <code>x</code>&nbsp;中的&nbsp;<strong>最大</strong>&nbsp;数位替换。比方说&nbsp;<code>encrypt(523) = 555</code> 且&nbsp;<code>encrypt(213) = 333</code>&nbsp;。</p>

<p>请你返回数组中所有元素加密后的 <strong>和</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">nums = [1,2,3]</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">6</span></p>

<p><b>解释：</b>加密后的元素位&nbsp;<code>[1,2,3]</code>&nbsp;。加密元素的和为&nbsp;<code>1 + 2 + 3 == 6</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">nums = [10,21,31]</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">66</span></p>

<p><b>解释：</b>加密后的元素为&nbsp;<code>[11,22,33]</code>&nbsp;。加密元素的和为&nbsp;<code>11 + 22 + 33 == 66</code> 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 50</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 1000</code></li>
</ul>

### 思路

对于 $x=\textit{nums}[i]$，遍历 $x$ 的每个数位，计算其最大数位 $\textit{mx}$。

遍历的同时，计算出一个等于 $111\cdots$ 的值 $\textit{base}$。

那么加密后的结果就是 $\textit{mx}\cdot \textit{base}$，加入答案。

```go [sol-Go]
func sumOfEncryptedInt(nums []int) (ans int) {
	for _, x := range nums {
		mx, base := 0, 0
		for ; x > 0; x /= 10 {
			mx = max(mx, x%10)
			base = base*10 + 1
		}
		ans += mx * base
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
