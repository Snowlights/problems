#### 题目

<p>给你两个整数 <code>red</code> 和 <code>blue</code>，分别表示红色球和蓝色球的数量。你需要使用这些球来组成一个三角形，满足第 1 行有 1 个球，第 2 行有 2 个球，第 3 行有 3 个球，依此类推。</p>

<p>每一行的球必须是 <strong>相同 </strong>颜色，且相邻行的颜色必须<strong> 不同</strong>。</p>

<p>返回可以实现的三角形的 <strong>最大 </strong>高度。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">red = 2, blue = 4</span></p>

<p><strong>输出：</strong> 3</p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/16/brb.png" style="width: 300px; height: 240px; padding: 10px;" /></p>

<p>上图显示了唯一可能的排列方式。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">red = 2, blue = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/16/br.png" style="width: 150px; height: 135px; padding: 10px;" /><br />
上图显示了唯一可能的排列方式。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">red = 1, blue = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">1</span></p>
</div>

<p><strong class="example">示例 4：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">red = 10, blue = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/16/br.png" style="width: 150px; height: 135px; padding: 10px;" /><br />
上图显示了唯一可能的排列方式。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= red, blue &lt;= 100</code></li>
</ul>

#### 思路

奇数行放红球，偶数行放蓝球；或者奇数行放蓝球，偶数行放红球。

计算最多能放多少排。两种情况取最大值。

## 奇数行

设奇数行有 $k$ 行，那么需要

$$
1+3+5+\cdots + (2k-1) = k^2
$$

个球。（等差数列求和公式）

假设我们有 $n$ 个球，那么有

$$
n\ge k^2
$$

解得

$$
k \le \left\lfloor\sqrt n\right\rfloor
$$

## 偶数行

设偶数行有 $k$ 行，那么需要

$$
2+4+6+\cdots + 2k = k^2 + k
$$

个球。（等差数列求和公式）

假设我们有 $n$ 个球，那么有

$$
n\ge k^2 + k
$$

解得

$$
k \le \left\lfloor\dfrac{\sqrt{4n+1}-1}{2}\right\rfloor
$$

## 答案

设有 $\textit{odd}$ 个奇数行，$\textit{even}$ 个偶数行，那么总行数为

$$
\begin{cases}
2\cdot \textit{even} + 1, & odd > even      \\
2\cdot \textit{odd}, & \text{otherwise}     \\
\end{cases}
$$

```
func f(n, m int) int {
	odd := int(math.Sqrt(float64(n)))
	even := int((math.Sqrt(float64(m*4+1)) - 1) / 2)
	if odd > even {
		return even*2 + 1
	}
	return odd * 2
}

func maxHeightOfTriangle(red, blue int) int {
	return max(f(red, blue), f(blue, red))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。CPU 有专门的计算平方根的指令，可以视作 $\mathcal{O}(1)$。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)