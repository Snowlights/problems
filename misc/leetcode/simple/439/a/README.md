#### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;和一个整数&nbsp;<code>k</code> 。</p>

<p>如果整数 <code>x</code>&nbsp;恰好仅出现在&nbsp;<code>nums</code>&nbsp;中的一个大小为 <code>k</code>&nbsp;的子数组中，则认为&nbsp;<code>x</code>&nbsp;是 <code>nums</code>&nbsp;中的几近缺失（<strong>almost missing</strong>）整数。</p>

<p>返回 <code>nums</code> 中 <strong>最大的几近缺失</strong> 整数，如果不存在这样的整数，返回&nbsp;<code>-1</code>&nbsp;。</p>
<strong>子数组</strong> 是数组中的一个连续元素序列。

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [3,9,2,1,7], k = 3</span></p>

<p><span class="example-io"><b>输出：</b>7</span></p>

<p><b>解释：</b></p>

<ul>
	<li>1 出现在两个大小为 3 的子数组中：<code>[9, 2, 1]</code>、<code>[2, 1, 7]</code></li>
	<li>2 出现在三个大小为&nbsp;3 的子数组中：<code>[3, 9, 2]</code>、<code>[9, 2, 1]</code>、<code>[2, 1, 7]</code></li>
	<li index="2">3 出现在一个大小为 3 的子数组中：<code>[3, 9, 2]</code></li>
	<li index="3">7 出现在一个大小为 3 的子数组中：<code>[2, 1, 7]</code></li>
	<li index="4">9 出现在两个大小为 3 的子数组中：<code>[3, 9, 2]</code>、<code>[9, 2, 1]</code></li>
</ul>

<p>返回 7 ，因为它满足题意的所有整数中最大的那个。</p>
</div>

<p><b>示例 2：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [3,9,7,2,1,7], k = 4</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><b>解释：</b></p>

<ul>
	<li>1 出现在两个大小为 3 的子数组中：<code>[9, 7, 2, 1]</code>、<code>[7, 2, 1, 7]</code></li>
	<li>2 出现在三个大小为 3 的子数组中：<code>[3, 9, 7, 2]</code>、<code>[9, 7, 2, 1]</code>、<code>[7, 2, 1, 7]</code></li>
	<li>3 出现在一个大小为 3 的子数组中：<code>[3, 9, 7, 2]</code></li>
	<li>7 出现在三个大小为 3 的子数组中：<code>[3, 9, 7, 2]</code>、<code>[9, 7, 2, 1]</code>、<code>[7, 2, 1, 7]</code></li>
	<li>9 出现在两个大小为 3 的子数组中：<code>[3, 9, 7, 2]</code>、<code>[9, 7, 2, 1]</code></li>
</ul>

<p>返回 3&nbsp;，因为它满足题意的所有整数中最大的那个。</p>
</div>

<p><b>示例 3：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [0,0], k = 1</span></p>

<p><span class="example-io"><b>输出：</b>-1</span></p>

<p><b>解释：</b></p>

<p>不存在满足题意的整数。</p>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 50</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 50</code></li>
	<li><code>1 &lt;= k &lt;= nums.length</code></li>
</ul>

#### 思路

分类讨论：

- 如果 $k=n$，只有一个子数组，所以每个数都满足要求，返回 $\textit{nums}$ 中的最大值。
- 如果 $k=1$，有 $n$ 个长为 $1$ 的子数组，只有出现次数等于 $1$ 的元素满足要求，返回出现次数等于 $1$ 的最大元素。
- 如果 $1 < k < n$，只有 $\textit{nums}[0]$ 和 $\textit{nums}[n-1]$ 是可能满足要求的数，因为**其他元素至少出现在两个子数组中**。返回这两个数中的出现次数等于 $1$ 的最大元素。

```
func largestInteger(nums []int, k int) int {
    n := len(nums)
    if k == n {
        return slices.Max(nums)
    }
    if k == 1 {
        cnt := map[int]int{}
        for _, x := range nums {
            cnt[x]++
        }
        ans := -1
        for x, c := range cnt {
            if c == 1 {
                ans = max(ans, x)
            }
        }
        return ans
    }
    // nums[0] 不能出现在其他地方，nums[n-1] 同理
    return max(f(nums[1:], nums[0]), f(nums[:n-1], nums[n-1]))
}

func f(a []int, x int) int {
    if slices.Contains(a, x) {
        return -1
    }
    return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)