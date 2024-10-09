#### 题目

<p>给你一个整数数组 <code>rewardValues</code>，长度为 <code>n</code>，代表奖励的值。</p>

<p>最初，你的总奖励 <code>x</code> 为 0，所有下标都是<strong> 未标记 </strong>的。你可以执行以下操作 <strong>任意次 </strong>：</p>

<ul>
	<li>从区间 <code>[0, n - 1]</code> 中选择一个 <strong>未标记 </strong>的下标 <code>i</code>。</li>
	<li>如果 <code>rewardValues[i]</code> <strong>大于</strong> 你当前的总奖励 <code>x</code>，则将 <code>rewardValues[i]</code> 加到 <code>x</code> 上（即 <code>x = x + rewardValues[i]</code>），并<strong> 标记</strong> 下标 <code>i</code>。</li>
</ul>

<p>以整数形式返回执行最优操作能够获得的<strong> 最大</strong><em> </em>总奖励。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">rewardValues = [1,1,3,3]</span></p>

<p><strong>输出：</strong><span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p>依次标记下标 0 和 2，总奖励为 4，这是可获得的最大值。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">rewardValues = [1,6,4,3,2]</span></p>

<p><strong>输出：</strong><span class="example-io">11</span></p>

<p><strong>解释：</strong></p>

<p>依次标记下标 0、2 和 1。总奖励为 11，这是可获得的最大值。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= rewardValues.length &lt;= 2000</code></li>
	<li><code>1 &lt;= rewardValues[i] &lt;= 2000</code></li>
</ul>

#### 思路

对于 $\textit{rewardValues}$ 中的数，如果先选大的，就没法再选小的，所以**按照从小到大的顺序选**是最好的。

把 $\textit{rewardValues}$ 从小到大排序。

排序后，问题变成一个标准的 0-1 背包问题，请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

对于本题，定义 $f[i][j]$ 表示能否从前 $i$ 个数中得到总奖励 $j$。

考虑 $v=\textit{rewardValues}[i]$ 选或不选：

- 不选 $v$：$f[i][j] = f[i-1][j]$。
- 选 $v$：$f[i][j] = f[i-1][j-v]$，前提是 $j\ge v$ 且 $j-v < v$，即 $v\le j<2v$。

满足其一即可，得

$$
f[i][j] = f[i-1][j] \vee f[i-1][j-v]
$$

其中 $\vee$ 即编程语言中的 `||`。

初始值 $f[0][0] = \texttt{true}$。

答案为最大的满足 $f[n][j]=\texttt{true}$ 的 $j$。

这可以解决周赛第三题，但第四题范围很大，需要去掉第一个维度，并用 **bitset** 优化（也可以用 **bigint**）。

用一个二进制数 $f$ 保存状态，二进制从低到高第 $j$ 位为 $1$ 表示 $f[j]=\texttt{true}$，为 $0$ 表示 $f[j]=\texttt{false}$。

对于上面的转移方程 $f[i][j] = f[i-1][j] \vee f[i-1][j-v]$，其中 $0\le j-v < v$，相当于取 $f$ 的低 $v$ 位，再左移 $v$ 位，然后与 $f$ 取按位或。

初始值 $f=1$。

答案为 $f$ 的最高位，即 $f$ 的二进制长度减一。

代码实现时，可以把 $\textit{rewardValues}$ 中的重复数字去掉，毕竟无法选两个一样的数。

``` 
const w = bits.UintSize

type bitset []uint

// b <<= k
func (b bitset) lsh(k int) bitset {
	shift, offset := k/w, k%w
	if offset == 0 {
		// Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(w-offset)
		}
		b[shift] = b[0] << offset
	}
	clear(b[:shift])
	return b
}

// 把 >= start 的清零
func (b bitset) resetRange(start int) bitset {
	i := start / w
	b[i] &= ^(^uint(0) << (start % w))
	clear(b[i+1:])
	return b
}

// b |= c
func (b bitset) unionFrom(c bitset) {
	for i, v := range c {
		b[i] |= v
	}
}

func (b bitset) lastIndex1() int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i*w | (bits.Len(b[i]) - 1)
		}
	}
	return -1
}

func maxTotalReward2(rewardValues []int) int {
	m := slices.Max(rewardValues)
	has := map[int]bool{}
	for _, v := range rewardValues {
		if v == m-1 {
			return m*2 - 1
		}
		if has[v] {
			continue
		}
		if has[m-1-v] {
			return m*2 - 1
		}
		has[v] = true
	}

	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重
	f := make(bitset, m*2/w+1)
	f[0] = 1
	for _, v := range rewardValues {
		f.unionFrom(slices.Clone(f).lsh(v).resetRange(v * 2))
	}
	return f.lastIndex1()
}

func maxTotalReward(rewardValues []int) int {
	m := slices.Max(rewardValues)
	has := map[int]bool{}
	for _, v := range rewardValues {
		if v == m-1 {
			return m*2 - 1
		}
		if has[v] {
			continue
		}
		if has[m-1-v] {
			return m*2 - 1
		}
		has[v] = true
	}

	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重

	one := big.NewInt(1)
	f := big.NewInt(1)
	p := new(big.Int)
	for _, v := range rewardValues {
		mask := p.Sub(p.Lsh(one, uint(v)), one)
		f.Or(f, p.Lsh(p.And(f, mask), uint(v)))
	}
	return f.BitLen() - 1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm/w)$，其中 $n$ 是 $\textit{rewardValues}$ 的长度，$m=\max(\textit{rewardValues})$，$w=64$ 或 $32$。
- 空间复杂度：$\mathcal{O}(n + m/w)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)