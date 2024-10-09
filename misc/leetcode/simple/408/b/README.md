#### 题目

<p>给你两个<strong> 正整数 </strong><code>l</code> 和 <code>r</code>。对于任何数字 <code>x</code>，<code>x</code> 的所有正因数（除了 <code>x</code> 本身）被称为 <code>x</code> 的 <strong>真因数</strong>。</p>

<p><span class="text-only" data-eleid="13" style="white-space: pre;">如果一个数字恰好仅有两个</span> <strong>真因数</strong>，则称该数字为 <strong>特殊数字</strong>。例如：</p>

<ul>
	<li>数字 4 是<strong> 特殊数字</strong>，因为它的真因数为 1 和 2。</li>
	<li>数字 6 不是 <strong>特殊数字</strong>，因为它的真因数为 1、2 和 3。</li>
</ul>

<p>返回区间 <code>[l, r]</code> 内<strong> 不是 特殊数字 </strong>的数字数量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">l = 5, r = 7</span></p>

<p><strong>输出：</strong> <span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<p>区间 <code>[5, 7]</code> 内不存在特殊数字。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">l = 4, r = 16</span></p>

<p><strong>输出：</strong> <span class="reset-io">11</span></p>

<p><strong>解释：</strong></p>

<p>区间 <code>[4, 16]</code> 内的特殊数字为 4 和 9。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= l &lt;= r &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

正难则反，统计区间 $[l,r]$ 内有多少个特殊数字。\n\n这等价于区间 $[0,r]$ 内的特殊数字个数，减去区间 $[0,l-1]$ 内的特殊数字个数。

根据题意，只有质数的平方 $p^2$ 才是特殊数字，因为 $p^2$ 恰好有两个真因数 $1$ 和 $p$。而其他的数，$1$ 没有真因数，质数只有 $1$ 个真因数，不是 $1$ 不是质数也不是质数的平方的数有至少三个真因数。
所以区间 $[0,i]$ 内的特殊数字个数等于：
- $[0,\lfloor\sqrt{i}\rfloor]$ 中有多少个质数。

预处理 $\lfloor\sqrt{10^9}\rfloor = 31622$ 内的质数，然后用前缀和计算前 $i$ 个数中的质数个数 $\pi(i)$，那么区间 $[l,r]$ 内的特殊数字个数就是

$$
\pi(\lfloor\sqrt{r}\rfloor) - \pi(\lfloor\sqrt{l-1}\rfloor)
$$

答案为区间内的整数个数，减去区间内的特殊数字个数，即

$$
r-l+1 - (\pi(\lfloor\sqrt{r}\rfloor) - \pi(\lfloor\sqrt{l-1}\rfloor))
$$

```
const mx = 31622
var pi = [mx + 1]int{}

func init() {
	for i := 2; i <= mx; i++ {
		if pi[i] == 0 { // i 是质数
			pi[i] = pi[i-1] + 1
			for j := i * i; j <= mx; j += i {
				pi[j] = -1 // 标记 i 的倍数为合数
			}
		} else {
			pi[i] = pi[i-1]
		}
	}
}

func nonSpecialCount(l, r int) int {
	cntR := pi[int(math.Sqrt(float64(r)))]
	cntL := pi[int(math.Sqrt(float64(l-1)))]
	return r - l + 1 - (cntR - cntL)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。不计入预处理的时间。
- 空间复杂度：$\mathcal{O}(1)$。不计入预处理的空间。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)