#### 题目

<p>给你一个<strong>正整数</strong> <code>k</code> 。最初，你有一个数组 <code>nums = [1]</code> 。</p>

<p>你可以对数组执行以下 <strong>任意 </strong>操作 <strong>任意 </strong>次数（<strong>可能为零</strong>）：</p>

<ul>
	<li>选择数组中的任何一个元素，然后将它的值<strong> 增加</strong> <code>1</code> 。</li>
	<li>复制数组中的任何一个元素，然后将它附加到数组的末尾。</li>
</ul>

<p>返回使得最终数组元素之<strong> 和 </strong>大于或等于 <code>k</code> 所需的 <strong>最少 </strong>操作次数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">k = 11</span></p>

<p><strong>输出：</strong><span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p>可以对数组 <code>nums = [1]</code> 执行以下操作：</p>

<ul>
	<li>将元素的值增加 <code>1</code> 三次。结果数组为 <code>nums = [4]</code> 。</li>
	<li>复制元素两次。结果数组为 <code>nums = [4,4,4]</code> 。</li>
</ul>

<p>最终数组的和为 <code>4 + 4 + 4 = 12</code> ，大于等于 <code>k = 11</code> 。<br />
执行的总操作次数为 <code>3 + 2 = 5</code> 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">k = 1</span></p>

<p><strong>输出：</strong><span class="example-io">0</span></p>

<p><strong>解释：</strong></p>

<p>原始数组的和已经大于等于 <code>1</code> ，因此不需要执行操作。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

## 方法一：贪心 + 枚举

设当前数组最大值为 $m$，对它做加一操作更好（因为复制最大值最优）。

- 如果先复制 $m$，再加一，那么元素和增加了 $m+1$。
- 如果先加一，再复制 $m+1$，那么元素和增加了 $m+2$。

所以，先加一再复制更优。

所以，加一操作都应当在复制操作之前。

我们可以枚举加一操作执行了 $\textit{add}= 0,1,2,\cdots, k-1$ 次。

设 $m=1+\textit{add}$，我们还需要复制

$$
\left\lceil\dfrac{k}{m}\right\rceil-1 = \left\lfloor\dfrac{k-1}{m}\right\rfloor
$$

次，才能让元素和至少为 $k$。上式可以分类讨论 $k$ 是 $m$ 的倍数，和 $k$ 不是 $m$ 的倍数两种情况证明。

所以答案为

$$
\min\limits_{m=1}^{k} m-1 + \left\lfloor\dfrac{k-1}{m}\right\rfloor
$$

```go [sol-Go]
func minOperations(k int) int {
	ans := math.MaxInt
	for m := 1; m <= k; m++ {
		ans = min(ans, m-1+(k-1)/m)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：数学

由基本不等式，或者对勾函数性质可知，设 $\textit{rt} = \left\lfloor\sqrt{k-1}\right\rfloor$，那么当 $m$ 取 $\textit{rt}$ 或者 $\textit{rt}+1$ 时我们可以得到最小值。

为防止 $m=0$，可以和 $1$ 取最大值（或者特判）。

> 注：在本题数据范围下，开平方结果的整数部分是正确的，无需调整。

```go [sol-Go]
func minOperations(k int) int {
	rt := max(int(math.Sqrt(float64(k-1))), 1)
	return min(rt-1+(k-1)/rt, rt+(k-1)/(rt+1))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。开平方有专门的 CPU 指令，可以视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。


## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

