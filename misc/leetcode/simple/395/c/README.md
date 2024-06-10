#### 题目

<p>给你两个整数 <code>n</code> 和 <code>x</code> 。你需要构造一个长度为 <code>n</code> 的 <strong>正整数 </strong>数组 <code>nums</code> ，对于所有 <code>0 &lt;= i &lt; n - 1</code> ，满足 <code>nums[i + 1]</code><strong> 大于 </strong><code>nums[i]</code> ，并且数组 <code>nums</code> 中所有元素的按位 <code>AND</code> 运算结果为 <code>x</code> 。</p>

<p>返回 <code>nums[n - 1]</code> 可能的<strong> 最小 </strong>值。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">n = 3, x = 4</span></p>

<p><strong>输出：</strong><span class="example-io">6</span></p>

<p><strong>解释：</strong></p>

<p>数组 <code>nums</code> 可以是 <code>[4,5,6]</code> ，最后一个元素为 <code>6</code> 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">n = 2, x = 7</span></p>

<p><strong>输出：</strong><span class="example-io">15</span></p>

<p><strong>解释：</strong></p>

<p>数组 <code>nums</code> 可以是 <code>[7,15]</code> ，最后一个元素为 <code>15</code> 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n, x &lt;= 10<sup>8</sup></code></li>
</ul>

#### 思路

从集合的视角看，$x$ 是每个 $\textit{nums}[i]$ 的**子集**。换句话说，$\textit{nums}[i]$ 一定是 $x$ 的**超集**。
例如 $x = 100100$，那么 $\textit{nums}[i]$ 一定在如下序列中：

$$
1\underline{00}1\underline{00},1\underline{00}1\underline{01},1\underline{00}1\underline{10},1\underline{00}1\underline{11},1\underline{01}1\underline{00},1\underline{01}1\underline{01}\cdots
$$

只看下划线上的数，是一个自然数序列

$$
0000,0001,0010,0011,0100,0101,\cdots
$$

为了让 $\textit{nums}[n-1]$ 尽量小，我们应当选择 $x$ 的超集中最小的 $n$ 个数。
所以把 $x$ 的二进制中的 $0$ 视作「空位」，往空位上填入 $n-1$，即为最小的 $\textit{nums}[n-1]$。
如果空位不足，往 $x$ 的前面添加前导零即可。

``` go
func minEnd(n, x int) int64 {
	n-- // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
	i, j := 0, 0
	for n>>j > 0 {
		// x 的第 i 个比特值是 0，即「空位」
		if x>>i&1 == 0 {
			// 空位填入 n 的第 j 个比特值
			x |= n >> j & 1 << i
			j++
		}
		i++
	}
	return int64(x)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
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
