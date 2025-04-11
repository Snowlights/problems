#### 题目

<p>给你一个由字符 <code>'N'</code>、<code>'S'</code>、<code>'E'</code> 和 <code>'W'</code> 组成的字符串 <code>s</code>，其中 <code>s[i]</code> 表示在无限网格中的移动操作：</p>

<ul>
	<li><code>'N'</code>：向北移动 1 个单位。</li>
	<li><code>'S'</code>：向南移动 1 个单位。</li>
	<li><code>'E'</code>：向东移动 1 个单位。</li>
	<li><code>'W'</code>：向西移动 1 个单位。</li>
</ul>

<p>初始时，你位于原点 <code>(0, 0)</code>。你 <strong>最多</strong> 可以修改 <code>k</code> 个字符为任意四个方向之一。</p>

<p>请找出在 <strong>按顺序</strong> 执行所有移动操作过程中的 <strong>任意时刻</strong> ，所能达到的离原点的&nbsp;<strong>最大曼哈顿距离&nbsp;</strong>。</p>

<p><strong>曼哈顿距离&nbsp;</strong>定义为两个坐标点 <code>(x<sub>i</sub>, y<sub>i</sub>)</code> 和 <code>(x<sub>j</sub>, y<sub>j</sub>)</code> 的横向距离绝对值与纵向距离绝对值之和，即 <code>|x<sub>i</sub> - x<sub>j</sub>| + |y<sub>i</sub> - y<sub>j</sub>|</code>。</p>

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "NWSE", k = 1</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><b>解释：</b></p>

<p>将&nbsp;<code>s[2]</code>&nbsp;从&nbsp;<code>'S'</code>&nbsp;改为&nbsp;<code>'N'</code> ，字符串&nbsp;<code>s</code>&nbsp;变为&nbsp;<code>"NWNE"</code> 。</p>

<table style="border: 1px solid black;">
	<thead>
		<tr>
			<th style="border: 1px solid black;">移动操作</th>
			<th style="border: 1px solid black;">位置 (x, y)</th>
			<th style="border: 1px solid black;">曼哈顿距离</th>
			<th style="border: 1px solid black;">最大值</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td style="border: 1px solid black;">s[0] == 'N'</td>
			<td style="border: 1px solid black;">(0, 1)</td>
			<td style="border: 1px solid black;">0 + 1 = 1</td>
			<td style="border: 1px solid black;">1</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">s[1] == 'W'</td>
			<td style="border: 1px solid black;">(-1, 1)</td>
			<td style="border: 1px solid black;">1 + 1 = 2</td>
			<td style="border: 1px solid black;">2</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">s[2] == 'N'</td>
			<td style="border: 1px solid black;">(-1, 2)</td>
			<td style="border: 1px solid black;">1 + 2 = 3</td>
			<td style="border: 1px solid black;">3</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">s[3] == 'E'</td>
			<td style="border: 1px solid black;">(0, 2)</td>
			<td style="border: 1px solid black;">0 + 2 = 2</td>
			<td style="border: 1px solid black;">3</td>
		</tr>
	</tbody>
</table>

<p>执行移动操作过程中，距离原点的最大曼哈顿距离是 3 。</p>
</div>

<p><b>示例 2：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "NSWWEW", k = 3</span></p>

<p><span class="example-io"><b>输出：</b>6</span></p>

<p><b>解释：</b></p>

<p>将&nbsp;<code>s[1]</code>&nbsp;从&nbsp;<code>'S'</code>&nbsp;改为&nbsp;<code>'N'</code> ，将&nbsp;<code>s[4]</code>&nbsp;从&nbsp;<code>'E'</code>&nbsp;改为&nbsp;<code>'W'</code> 。字符串&nbsp;<code>s</code>&nbsp;变为&nbsp;<code>"NNWWWW"</code>&nbsp;。</p>

<p>执行移动操作过程中，距离原点的最大曼哈顿距离是 6&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= k &lt;= s.length</code></li>
	<li><code>s</code>&nbsp;仅由&nbsp;<code>'N'</code>、<code>'S'</code>、<code>'E'</code>&nbsp;和&nbsp;<code>'W'</code> 。</li>
</ul>

#### 思路

## 方法一

对于曼哈顿距离，由于水平方向的移动和垂直方向的移动互不影响，我们可以把横纵坐标分别计算。

设当前向西走了 $a$ 步，向东走了 $b$ 步。比如 $a=2,\ b=5$。

贪心地，把其中更小的 $a$ 改成向东走。

如果把 $a$ 减少 $d$，那么 $b$ 就能增大 $d$，所以修改后的当前位置的横坐标为

$$
(b+d) - (a-d) = b-a+2d
$$

如果 $b$ 更小，那么横坐标的绝对值为

$$
a-b+2d
$$

综合两种情况，修改后的当前位置的横坐标的绝对值为

$$
|a-b|+2d
$$

其中

$$
d = \min(a,b,k)
$$

然后把 $k$ 减少 $d$，按照同样的方法继续计算纵坐标。

用修改后的横纵坐标绝对值之和，更新答案的最大值。

```
func maxDistance(s string, k int) (ans int) {
	cnt := ['X']int{} // 'W' + 1 = 'X'
	for _, ch := range s {
		cnt[ch]++
		left := k
		f := func(a, b int) int {
			d := min(a, b, left)
			left -= d
			return abs(a-b) + d*2
		}
		ans = max(ans, f(cnt['N'], cnt['S'])+f(cnt['E'], cnt['W']))
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 方法二

设当前位置为 $(x,y)$，那么到原点的曼哈顿距离为 $|x|+|y|$。

通过方法一可知，每操作一次，曼哈顿距离都会增大 $2$，但这不会超过移动的次数 $i+1$。

所以执行完 $s[i]$ 后的答案为

$$
\min(|x|+|y|+2k,i+1)
$$

```
func maxDistance(s string, k int) int {
	ans, x, y := 0, 0, 0
	for i, c := range s {
		switch c {
		case 'N': y++
		case 'S': y--
		case 'E': x++
		default:  x--
		}
		ans = max(ans, min(abs(x)+abs(y)+k*2, i+1))
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)