#### 题目

<p>如果一个整数能够被其各个数位上的数字之和整除，则称之为<strong> 哈沙德数</strong>（Harshad number）。给你一个整数 <code>x</code> 。如果 <code>x</code> 是 <strong>哈沙德数 </strong>，则返回<em> </em><code>x</code> 各个数位上的数字之和，否则，返回<em> </em><code>-1</code> 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">x = 18</span></p>

<p><strong>输出：</strong> <span class="example-io">9</span></p>

<p><strong>解释：</strong></p>

<p><code>x</code> 各个数位上的数字之和为 <code>9</code> 。<code>18</code> 能被 <code>9</code> 整除。因此 <code>18</code> 是哈沙德数，答案是 <code>9</code> 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">x = 23</span></p>

<p><strong>输出：</strong> <span class="example-io">-1</span></p>

<p><strong>解释：</strong></p>

<p><code>x</code> 各个数位上的数字之和为 <code>5</code> 。<code>23</code> 不能被 <code>5</code> 整除。因此 <code>23</code> 不是哈沙德数，答案是 <code>-1</code> 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= x &lt;= 100</code></li>
</ul>

#### 思路

模拟

```go [sol-Go]
func sumOfTheDigitsOfHarshadNumber(x int) int {
	cnt := 0
	for tmp := x; tmp > 0; tmp /= 10 {
		cnt += tmp % 10
	}
	if x%cnt == 0 {
		return cnt
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(log x)$，其中 $n$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)