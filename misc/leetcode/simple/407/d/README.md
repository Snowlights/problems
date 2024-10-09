#### 题目

<p>给你两个长度相同的正整数数组 <code>nums</code> 和 <code>target</code>。</p>

<p>在一次操作中，你可以选择 <code>nums</code> 的任何<span data-keyword="subarray">子数组</span>，并将该子数组内的每个元素的值增加或减少 1。</p>

<p>返回使 <code>nums</code> 数组变为 <code>target</code> 数组所需的 <strong>最少 </strong>操作次数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [3,5,1,2], target = [4,6,2,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>执行以下操作可以使 <code>nums</code> 等于 <code>target</code>：<br />
- <code>nums[0..3]</code> 增加 1，<code>nums = [4,6,2,3]</code>。<br />
- <code>nums[3..3]</code> 增加 1，<code>nums = [4,6,2,4]</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,3,2], target = [2,1,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p>执行以下操作可以使 <code>nums</code> 等于 <code>target</code>：<br />
- <code>nums[0..0]</code> 增加 1，<code>nums = [2,3,2]</code>。<br />
- <code>nums[1..1]</code> 减少 1，<code>nums = [2,2,2]</code>。<br />
- <code>nums[1..1]</code> 减少 1，<code>nums = [2,1,2]</code>。<br />
- <code>nums[2..2]</code> 增加 1，<code>nums = [2,1,3]</code>。<br />
- <code>nums[2..2]</code> 增加 1，<code>nums = [2,1,4]</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length == target.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i], target[i] &lt;= 10<sup>8</sup></code></li>
</ul>


#### 思路

**请先阅读**：[差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看。

「子数组内的每个元素的值增加或减少 $1$」，这个操作可以转换成修改**差分数组**两个位置上的数。
设 $d_1$ 为 $\textit{nums}$ 的差分数组，$d_2$ 为 $\textit{target}$ 的差分数组。

由于差分数组和原数组是一一对应的，所以问题等价于：
- 把 $d_1$ 变成 $d_2$。
- 每次操作，可以选择两个下标 $i$ 和 $j$（或者只选一个下标 $i$，对应操作 $\textit{nums}$ 后缀的情况），把 $d_1[i]$ 加一（或减一），把 $d_1[j]$ 减一（或加一）。

从左到右遍历 $d_1$ 和 $d_2$，同时维护一个变量 $s$，表示对 $d_1[i]$ 增大/减少的累积量：
- 如果把 $d_1[i]$ 增大了 $k$，那么后面可以把 $d_1[j]$ **免费减少**，至多免费减少 $k$ 次。
- 如果把 $d_1[i]$ 减少了 $k$，那么后面可以把 $d_1[j]$ **免费增大**，至多免费增大 $k$ 次。

设 $k = d_2[i] - d_1[i]$，分类讨论：
- 如果 $k > 0$ 且 $s\ge 0$，那么必须通过操作，把 $d_1[i]$ 增大到 $d_2[i]$，操作 $k$ 次。
- 如果 $k > 0$ 且 $s < 0$，那么可以免费增大至多 $-s$ 次，如果 $k \le -s$ 则无需额外操作，否则要额外操作 $k+s$ 次。综合一下，就是额外操作 $\max(s+k,0)$ 次。
- 如果 $k \le 0$ 且 $s\le 0$，那么必须通过操作，把 $d_1[i]$ 减少到 $d_2[i]$，操作 $-k$ 次。
- 如果 $k \le 0$ 且 $s > 0$，那么可以免费减少至多 $s$ 次，如果 $-k \le s$ 则无需额外操作，否则要额外操作 $-k-s$ 次。综合一下，就是额外操作 $-\min(s+k,0)$ 次。
- 最后把 $k$ 加到 $s$ 中。

代码实现时，可以单独计算 $i=0$ 的情况，方便在计算差分数组的同时计算答案。

```
func minimumOperations(nums, target []int) int64 {
	s := target[0] - nums[0]
	ans := abs(s)
	for i := 1; i < len(nums); i++ {
		k := (target[i] - target[i-1]) - (nums[i] - nums[i-1])
		if k > 0 {
			if s >= 0 {
				ans += k
			} else {
				ans += max(k+s, 0)
			}
		} else {
			if s <= 0 {
				ans -= k
			} else {
				ans -= min(k+s, 0)
			}
		}
		s += k
	}
	return int64(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)