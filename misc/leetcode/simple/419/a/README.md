#### 题目

<p>给你一个由 <code>n</code> 个整数组成的数组 <code>nums</code>，以及两个整数 <code>k</code> 和 <code>x</code>。</p>

<p>数组的 <strong>x-sum</strong> 计算按照以下步骤进行：</p>

<ul>
	<li>统计数组中所有元素的出现次数。</li>
	<li>仅保留出现次数最多的前 <code>x</code> 个元素的每次出现。如果两个元素的出现次数相同，则数值<strong> 较大 </strong>的元素被认为出现次数更多。</li>
	<li>计算结果数组的和。</li>
</ul>

<p><strong>注意</strong>，如果数组中的不同元素少于 <code>x</code> 个，则其 <strong>x-sum</strong> 是数组的元素总和。</p>

<p>返回一个长度为 <code>n - k + 1</code> 的整数数组 <code>answer</code>，其中 <code>answer[i]</code> 是 <span data-keyword="subarray-nonempty">子数组</span> <code>nums[i..i + k - 1]</code> 的 <strong>x-sum</strong>。</p>

<p><strong>子数组</strong> 是数组内的一个连续<b> 非空</b> 的元素序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [1,1,2,2,3,4,2,3], k = 6, x = 2</span></p>

<p><strong>输出：</strong><span class="example-io">[6,10,12]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>对于子数组 <code>[1, 1, 2, 2, 3, 4]</code>，只保留元素 1 和 2。因此，<code>answer[0] = 1 + 1 + 2 + 2</code>。</li>
	<li>对于子数组 <code>[1, 2, 2, 3, 4, 2]</code>，只保留元素 2 和 4。因此，<code>answer[1] = 2 + 2 + 2 + 4</code>。注意 4 被保留是因为其数值大于出现其他出现次数相同的元素（3 和 1）。</li>
	<li>对于子数组 <code>[2, 2, 3, 4, 2, 3]</code>，只保留元素 2 和 3。因此，<code>answer[2] = 2 + 2 + 2 + 3 + 3</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [3,8,7,8,7,5], k = 2, x = 2</span></p>

<p><strong>输出：</strong><span class="example-io">[11,15,15,15,12]</span></p>

<p><strong>解释：</strong></p>

<p>由于 <code>k == x</code>，<code>answer[i]</code> 等于子数组 <code>nums[i..i + k - 1]</code> 的总和。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == nums.length &lt;= 50</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 50</code></li>
	<li><code>1 &lt;= x &lt;= k &lt;= nums.length</code></li>
</ul>

#### 思路

请先完成本题的简单版本 [3013. 将数组分成最小总代价的子数组 II](https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/)。

本题需要在 3013 的基础上变形，额外用一个哈希表 $\textit{cnt}$ 维护窗口中的元素及其出现次数。
两个**有序集合**维护的是二元组 $(\textit{cnt}[x], x)$。\n\n当元素进入窗口时：
- 先把 $(\textit{cnt}[x], x)$ 从有序集合中移除。
- 然后把 $\textit{cnt}[x]$ 加一。
- 最后把 $(\textit{cnt}[x], x)$ 加入有序集合。

当元素离开窗口时：
- 先把 $(\textit{cnt}[x], x)$ 从有序集合中移除。
- 然后把 $\textit{cnt}[x]$ 减一。
- 最后把 $(\textit{cnt}[x], x)$ 加入有序集合。

```
type pair struct{ c, x int } // 出现次数，元素值

func less(p, q pair) int {
	return cmp.Or(p.c-q.c, p.x-q.x)
}

func findXSum(nums []int, k, x int) []int64 {
	L := redblacktree.NewWith[pair, struct{}](less)
	R := redblacktree.NewWith[pair, struct{}](less)

	sumL := 0 // L 的元素和
	cnt := map[int]int{}
	add := func(x int) {
		p := pair{cnt[x], x}
		if p.c == 0 {
			return
		}
		if !L.Empty() && less(p, L.Left().Key) > 0 { // p 比 L 中最小的还大
			sumL += p.c * p.x
			L.Put(p, struct{}{})
		} else {
			R.Put(p, struct{}{})
		}
	}
	del := func(x int) {
		p := pair{cnt[x], x}
		if p.c == 0 {
			return
		}
		if _, ok := L.Get(p); ok {
			sumL -= p.c * p.x
			L.Remove(p)
		} else {
			R.Remove(p)
		}
	}
	l2r := func() {
		p := L.Left().Key
		sumL -= p.c * p.x
		L.Remove(p)
		R.Put(p, struct{}{})
	}
	r2l := func() {
		p := R.Right().Key
		sumL += p.c * p.x
		R.Remove(p)
		L.Put(p, struct{}{})
	}

	ans := make([]int64, len(nums)-k+1)
	for r, in := range nums {
		// 添加 in
		del(in)
		cnt[in]++
		add(in)

		l := r + 1 - k
		if l < 0 {
			continue
		}

		// 维护大小
		for !R.Empty() && L.Size() < x {
			r2l()
		}
		for L.Size() > x {
			l2r()
		}
		ans[l] = int64(sumL)

		// 移除 out
		out := nums[l]
		del(out)
		cnt[out]--
		add(out)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)