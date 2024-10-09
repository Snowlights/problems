#### 题目

<p>如果数组的每一对相邻元素都是两个奇偶性不同的数字，则该数组被认为是一个 <strong>特殊数组</strong> 。</p>

<p>周洋哥有一个整数数组 <code>nums</code> 和一个二维整数矩阵 <code>queries</code>，对于 <code>queries[i] = [from<sub>i</sub>, to<sub>i</sub>]</code>，请你帮助周洋哥检查子数组 <code>nums[from<sub>i</sub>..to<sub>i</sub>]</code> 是不是一个 <strong>特殊数组 </strong>。</p>

<p>返回布尔数组 <code>answer</code>，如果 <code>nums[from<sub>i</sub>..to<sub>i</sub>]</code> 是特殊数组，则 <code>answer[i]</code> 为 <code>true</code> ，否则，<code>answer[i]</code> 为 <code>false</code> 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [3,4,1,2,6], queries = [[0,4]]</span></p>

<p><strong>输出：</strong><span class="example-io">[false]</span></p>

<p><strong>解释：</strong></p>

<p>子数组是 <code>[3,4,1,2,6]</code>。2 和 6 都是偶数。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [4,3,1,6], queries = [[0,2],[2,3]]</span></p>

<p><strong>输出：</strong><span class="example-io">[false,true]</span></p>

<p><strong>解释：</strong></p>

<ol>
	<li>子数组是 <code>[4,3,1]</code>。3 和 1 都是奇数。因此这个查询的答案是 <code>false</code>。</li>
	<li>子数组是 <code>[1,6]</code>。只有一对：<code>(1,6)</code>，且包含了奇偶性不同的数字。因此这个查询的答案是 <code>true</code>。</li>
</ol>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>0 &lt;= queries[i][0] &lt;= queries[i][1] &lt;= nums.length - 1</code></li>
</ul>

#### 思路

定义一个长为 $n-1$ 的数组 $a$，其中

$$
a[i] =
\begin{cases}
0,\ &\textit{nums}[i]\bmod 2 \ne \textit{nums}[i+1]\bmod 2\\
1,\ &\textit{nums}[i]\bmod 2 = \textit{nums}[i+1]\bmod 2
\end{cases}
$$

那么 $a$ 的从 $\textit{from}$ 到 $\textit{to}-1$ 的子数组和等于 $0$，就说明询问的子数组是特殊数组。

计算 $a$ 的前缀和 $s$，可以快速判断子数组和是否为 $0$。

代码实现时，相邻两数的异或和的最低位取反，即为 $a[i]$。

``` 
func isArraySpecial(nums []int, queries [][]int) []bool {
	s := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		s[i] = s[i-1] + (nums[i]^nums[i-1]^1)&1
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = s[q[0]] == s[q[1]]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，其中 $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)