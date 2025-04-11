#### 题目

<p>给你两个整数数组 <code>nums1</code> 和 <code>nums2</code>，长度分别为 <code>n</code> 和 <code>m</code>。同时给你一个<strong>正整数</strong> <code>k</code>。</p>

<p>如果 <code>nums1[i]</code> 可以被 <code>nums2[j] * k</code> 整除，则称数对 <code>(i, j)</code> 为 <strong>优质数对</strong>（<code>0 &lt;= i &lt;= n - 1</code>, <code>0 &lt;= j &lt;= m - 1</code>）。</p>

<p>返回<strong> 优质数对 </strong>的总数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums1 = [1,3,4], nums2 = [1,3,4], k = 1</span></p>

<p><strong>输出：</strong><span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p>5个优质数对分别是 <code>(0, 0)</code>, <code>(1, 0)</code>, <code>(1, 1)</code>, <code>(2, 0)</code>, 和 <code>(2, 2)</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums1 = [1,2,4,12], nums2 = [2,4], k = 3</span></p>

<p><strong>输出：</strong><span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>2个优质数对分别是 <code>(3, 0)</code> 和 <code>(3, 1)</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n, m &lt;= 50</code></li>
	<li><code>1 &lt;= nums1[i], nums2[j] &lt;= 50</code></li>
	<li><code>1 &lt;= k &lt;= 50</code></li>
</ul>

#### 思路

## 方法一：统计因子

提示：

- 如果 $x$ 能被 $d$ 整除，那么 $x$ 必然有一个等于 $d$ 的因子。

思路：

为方便描述，把 $\textit{nums}_1$ 和 $\textit{nums}_2$ 记作 $a$ 和 $b$。

1. 遍历 $a$，统计所有元素的因子个数，记录到哈希表 $\textit{cnt}$ 中。
2. 遍历 $b$，那么有 $\textit{cnt}[b[i]\cdot k]$ 个数可以被 $b[i]\cdot k$ 整除，加入答案。

优化：

1. 如果 $a[i]$ 不是 $k$ 的倍数，无法被 $b[i]\cdot k$ 整除，直接跳过不统计。
2. 此外，可以改为统计 $\dfrac{a[i]}{k}$ 的因子，这样需要循环的次数会更少；相应地，遍历 $b$ 时只需要把 $\textit{cnt}[b[i]]$ 加入答案。


``` go
func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
	cnt1 := map[int]int{}
	for _, x := range nums1 {
		if x%k == 0 {
			cnt1[x/k]++
		}
	}
	cnt2 := map[int]int{}
	for _, x := range nums2 {
		cnt2[x]++
	}

	m := slices.Max(nums1) / k
	for i, c := range cnt2 {
		s := 0
		for j := i; j <= m; j += i {
			s += cnt1[j]
		}
		ans += int64(s * c)
	}
	return
}

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{U/k} + m)$，其中 $n$ 是 $\textit{nums}_1$ 的长度，$m$ 是 $\textit{nums}_2$ 的长度，$U=\max(\textit{nums}_1)$。
- 空间复杂度：$\mathcal{O}(U/k)$。不同因子个数不会超过 $U/k$。

## 方法二：枚举倍数

统计 $a[i]/k$ 和 $b[i]$ 的出现次数，分别保存到哈希表 $\textit{cnt}_1$ 和 $\textit{cnt}_2$ 中。设 $\textit{cnt}_1$ 中的最大 key 为 $u$。

枚举 $\textit{cnt}_2$ 中的元素 $i$，然后枚举 $i$ 的倍数 $i,2i,3i,\cdots$，一直到 $u$，累加这些数在 $\textit{cnt}_1$ 中的 value，乘上 $\textit{cnt}_2[i]$，加入答案。

```
func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
	cnt1 := map[int]int{}
	for _, x := range nums1 {
		if x%k == 0 {
			cnt1[x/k]++
		}
	}
	cnt2 := map[int]int{}
	for _, x := range nums2 {
		cnt2[x]++
	}

	m := slices.Max(nums1) / k
	for i, c := range cnt2 {
		s := 0
		for j := i; j <= m; j += i {
			s += cnt1[j]
		}
		ans += int64(s * c)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m + (U/k)\log m)$，其中 $n$ 是 $\textit{nums}_1$ 的长度，$m$ 是 $\textit{nums}_2$ 的长度，$U=\max(\textit{nums}_1)$。复杂度根据调和级数可得。
- 空间复杂度：$\mathcal{O}(n+m)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)