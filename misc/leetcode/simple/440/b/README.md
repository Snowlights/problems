#### 题目

<p>给你两个整数数组，<code>nums1</code> 和 <code>nums2</code>，长度均为 <code>n</code>，以及一个正整数 <code>k</code> 。</p>

<p>对从 <code>0</code> 到 <code>n - 1</code> 每个下标 <code>i</code> ，执行下述操作：</p>

<ul>
	<li>找出所有满足 <code>nums1[j]</code> 小于 <code>nums1[i]</code> 的下标 <code>j</code> 。</li>
	<li>从这些下标对应的 <code>nums2[j]</code> 中选出 <strong>至多</strong> <code>k</code> 个，并 <strong>最大化</strong> 这些值的总和作为结果。</li>
</ul>

<p>返回一个长度为 <code>n</code> 的数组 <code>answer</code> ，其中 <code>answer[i]</code> 表示对应下标 <code>i</code> 的结果。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums1 = [4,2,1,5,3], nums2 = [10,20,30,40,50], k = 2</span></p>

<p><strong>输出：</strong><span class="example-io">[80,30,0,80,50]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>对于 <code>i = 0</code> ：满足 <code>nums1[j] &lt; nums1[0]</code> 的下标为 <code>[1, 2, 4]</code> ，选出其中值最大的两个，结果为 <code>50 + 30 = 80</code> 。</li>
	<li>对于 <code>i = 1</code> ：满足 <code>nums1[j] &lt; nums1[1]</code> 的下标为 <code>[2]</code> ，只能选择这个值，结果为 <code>30</code> 。</li>
	<li>对于 <code>i = 2</code> ：不存在满足 <code>nums1[j] &lt; nums1[2]</code> 的下标，结果为 <code>0</code> 。</li>
	<li>对于 <code>i = 3</code> ：满足 <code>nums1[j] &lt; nums1[3]</code> 的下标为 <code>[0, 1, 2, 4]</code> ，选出其中值最大的两个，结果为 <code>50 + 30 = 80</code> 。</li>
	<li>对于 <code>i = 4</code> ：满足 <code>nums1[j] &lt; nums1[4]</code> 的下标为 <code>[1, 2]</code> ，选出其中值最大的两个，结果为 <code>30 + 20 = 50</code> 。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums1 = [2,2,2,2], nums2 = [3,1,2,3], k = 1</span></p>

<p><strong>输出：</strong><span class="example-io">[0,0,0,0]</span></p>

<p><strong>解释：</strong>由于 <code>nums1</code> 中的所有元素相等，不存在满足条件 <code>nums1[j] &lt; nums1[i]</code>，所有位置的结果都是 0 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == nums1.length == nums2.length</code></li>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums1[i], nums2[i] &lt;= 10<sup>6</sup></code></li>
	<li><code>1 &lt;= k &lt;= n</code></li>
</ul>

#### 思路

**前置题目**：[703. 数据流中的第 K 大元素](https://leetcode.cn/problems/kth-largest-element-in-a-stream/)

根据输入，创建 $n$ 个三元组 $(\textit{nums}_1[i], \textit{nums}_2[i], i)$，然后按照 $\textit{nums}_1[i]$ 从小到大排序。

排序后，小于 $\textit{nums}_1[i]$ 的数，都在 $\textit{nums}_1[i]$ 左边，这样方便我们增量地处理。

遍历三元组列表，同时用一个**最小堆**维护 $\textit{nums}_2[i]$ 的前 $k$ 大元素：

- 把 $\textit{nums}_2[i]$ 入堆。
- 如果堆的大小超过 $k$，弹出堆顶（这 $k+1$ 个数的最小值）。
- 入堆出堆的过程中，用一个变量 $s$ 维护堆中元素之和。

## 写法一（更通用）

可能存在多个 $\textit{nums}_1[i]$ 相同的情况，要把这些相同的数一起处理，原理见 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)。

```
func findMaxSum(nums1, nums2 []int, k int) []int64 {
	n := len(nums1)
	type tuple struct{ x, y, i int }
	a := make([]tuple, n)
	for i, x := range nums1 {
		a[i] = tuple{x, nums2[i], i}
	}
	slices.SortFunc(a, func(p, q tuple) int { return p.x - q.x })

	ans := make([]int64, n)
	h := &hp{}
	s := 0
	// 分组循环模板
	for i := 0; i < n; {
		start := i
		// 找到所有相同的 nums1[i]，这些数的答案都是一样的
		x := a[start].x
		for ; i < n && a[i].x == x; i++ {
			ans[a[i].i] = int64(s)
		}
		// 把这些相同的 nums1[i] 对应的 nums2[i] 入堆
		for ; start < i; start++ {
			y := a[start].y
			s += y
			heap.Push(h, y)
			if h.Len() > k {
				s -= heap.Pop(h).(int)
			}
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

## 写法二（针对本题）

上面是通用写法。针对本题，可以判断 $\textit{nums}_1[i]$ 和 $\textit{nums}_1[i-1]$ 是否相等，如果相等那么 $\textit{nums}_1[i]$ 的答案就是 $\textit{nums}_1[i-1]$ 的答案。

```
func findMaxSum(nums1, nums2 []int, k int) []int64 {
	n := len(nums1)
	type tuple struct{ x, y, i int }
	a := make([]tuple, n)
	for i, x := range nums1 {
		a[i] = tuple{x, nums2[i], i}
	}
	slices.SortFunc(a, func(p, q tuple) int { return p.x - q.x })

	ans := make([]int64, n)
	h := &hp{}
	s := 0
	for i, t := range a {
		if i > 0 && t.x == a[i-1].x {
			ans[t.i] = ans[a[i-1].i]
		} else {
			ans[t.i] = int64(s)
		}
		s += t.y
		heap.Push(h, t.y)
		if h.Len() > k {
			s -= heap.Pop(h).(int)
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。
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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)