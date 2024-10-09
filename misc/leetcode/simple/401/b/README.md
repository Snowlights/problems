#### 题目

<p>给你两个整数 <code>n</code> 和 <code>k</code>。</p>

<p>最初，你有一个长度为 <code>n</code> 的整数数组 <code>a</code>，对所有 <code>0 &lt;= i &lt;= n - 1</code>，都有 <code>a[i] = 1</code> 。每过一秒，你会同时更新每个元素为其前面所有元素的和加上该元素本身。例如，一秒后，<code>a[0]</code> 保持不变，<code>a[1]</code> 变为 <code>a[0] + a[1]</code>，<code>a[2]</code> 变为 <code>a[0] + a[1] + a[2]</code>，以此类推。</p>

<p>返回 <code>k</code> 秒后 <code>a[n - 1]</code> 的<strong>值</strong>。</p>

<p>由于答案可能非常大，返回其对 <code>10<sup>9</sup> + 7</code> <strong>取余 </strong>后的结果。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">n = 4, k = 5</span></p>

<p><strong>输出：</strong><span class="example-io">56</span></p>

<p><strong>解释：</strong></p>

<table border="1">
	<tbody>
		<tr>
			<th>时间（秒）</th>
			<th>数组状态</th>
		</tr>
		<tr>
			<td>0</td>
			<td>[1,1,1,1]</td>
		</tr>
		<tr>
			<td>1</td>
			<td>[1,2,3,4]</td>
		</tr>
		<tr>
			<td>2</td>
			<td>[1,3,6,10]</td>
		</tr>
		<tr>
			<td>3</td>
			<td>[1,4,10,20]</td>
		</tr>
		<tr>
			<td>4</td>
			<td>[1,5,15,35]</td>
		</tr>
		<tr>
			<td>5</td>
			<td>[1,6,21,56]</td>
		</tr>
	</tbody>
</table>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">n = 5, k = 3</span></p>

<p><strong>输出：</strong><span class="example-io">35</span></p>

<p><strong>解释：</strong></p>

<table border="1">
	<tbody>
		<tr>
			<th>时间（秒）</th>
			<th>数组状态</th>
		</tr>
		<tr>
			<td>0</td>
			<td>[1,1,1,1,1]</td>
		</tr>
		<tr>
			<td>1</td>
			<td>[1,2,3,4,5]</td>
		</tr>
		<tr>
			<td>2</td>
			<td>[1,3,6,10,15]</td>
		</tr>
		<tr>
			<td>3</td>
			<td>[1,4,10,20,35]</td>
		</tr>
	</tbody>
</table>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n, k &lt;= 1000</code></li>
</ul>

#### 思路

把脑袋往左斜 45°，看示例 1，你发现了什么？

示例 1 从左下往右上看，每条斜线上是数字是

$$
\begin{align}
&[1]\\
&[1,1]\\
&[1,2,1]\\
&[1,3,3,1]\\
&[1,4,6,4]\\
&[1,5,10,10]\\
&\vdots
\end{align}
$$

这是**杨辉三角**。

我们相当于计算的是杨辉三角第 $n+k$ 排的第 $n$ 个数，即 $C(n+k-1,n-1) = C(n+k-1, k)$。

预处理阶乘及其逆元后，即可 $\mathcal{O}(1)$ 计算组合数。

``` 
const mod = 1_000_000_007
const mx = 2000
var F, invF [mx + 1]int

func init() {
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func valueAfterKSeconds(n, k int) int {
	return F[n+k-1] * invF[n-1] % mod * invF[k] % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。预处理的时间和空间不计入。
- 空间复杂度：$\mathcal{O}(1)$。

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