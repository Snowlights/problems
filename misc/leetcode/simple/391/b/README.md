#### 题目

<p>给你两个整数 <code>numBottles</code> 和 <code>numExchange</code> 。</p>

<p><code>numBottles</code> 代表你最初拥有的满水瓶数量。在一次操作中，你可以执行以下操作之一：</p>

<ul>
	<li>喝掉任意数量的满水瓶，使它们变成空水瓶。</li>
	<li>用 <code>numExchange</code> 个空水瓶交换一个满水瓶。然后，将 <code>numExchange</code> 的值增加 1 。</li>
</ul>

<p>注意，你不能使用相同的 <code>numExchange</code> 值交换多批空水瓶。例如，如果 <code>numBottles == 3</code> 并且 <code>numExchange == 1</code> ，则不能用 <code>3</code> 个空水瓶交换成 <code>3</code> 个满水瓶。</p>

<p>返回你 <strong>最多</strong> 可以喝到多少瓶水。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/01/28/exampleone1.png" style="width: 948px; height: 482px; padding: 10px; background: #fff; border-radius: .5rem;" />
<pre>
<strong>输入：</strong>numBottles = 13, numExchange = 6
<strong>输出：</strong>15
<strong>解释：</strong>上表显示了满水瓶的数量、空水瓶的数量、numExchange 的值，以及累计喝掉的水瓶数量。
</pre>

<p><strong class="example">示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/01/28/example231.png" style="width: 990px; height: 642px; padding: 10px; background: #fff; border-radius: .5rem;" />
<pre>
<strong>输入：</strong>numBottles = 10, numExchange = 3
<strong>输出：</strong>13
<strong>解释：</strong>上表显示了满水瓶的数量、空水瓶的数量、numExchange 的值，以及累计喝掉的水瓶数量。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= numBottles &lt;= 100 </code></li>
	<li><code>1 &lt;= numExchange &lt;= 100</code></li>
</ul>

#### 思路

模拟

```go [sol-Go]
func maxBottlesDrunk(numBottles int, numExchange int) int {
	ans := numBottles

	for numBottles >= numExchange {
		numBottles++
		ans++
		numBottles -= numExchange
		numExchange++
	}

	return ans
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\sqrt{\textit{numBottles}})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)