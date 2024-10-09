#### 题目

<p>给你一个整数数组 <code>colors</code> 和一个二维整数数组 <code>queries</code> 。<code>colors</code>表示一个由红色和蓝色瓷砖组成的环，第 <code>i</code>&nbsp;块瓷砖的颜色为&nbsp;<code>colors[i]</code>&nbsp;：</p>

<ul>
	<li><code>colors[i] == 0</code>&nbsp;表示第&nbsp;<code>i</code>&nbsp;块瓷砖的颜色是 <strong>红色</strong>&nbsp;。</li>
	<li><code>colors[i] == 1</code>&nbsp;表示第 <code>i</code>&nbsp;块瓷砖的颜色是 <strong>蓝色</strong>&nbsp;。</li>
</ul>

<p>环中连续若干块瓷砖的颜色如果是 <strong>交替</strong>&nbsp;颜色（也就是说这组瓷砖中除了第一块和最后一块瓷砖以外，中间瓷砖的颜色与它<strong>&nbsp;左边</strong>&nbsp;和 <strong>右边</strong>&nbsp;的颜色都不同），那么它被称为一个 <strong>交替组</strong>。</p>

<p>你需要处理两种类型的查询：</p>

<ul>
	<li><code>queries[i] = [1, size<sub>i</sub>]</code>，确定大小为<code>size<sub>i</sub></code>的<strong> </strong><strong>交替组</strong> 的数量。</li>
	<li><code>queries[i] = [2, index<sub>i</sub>, color<sub>i</sub>]</code>，将<code>colors[index<sub>i</sub>]</code>更改为<code>color<sub>i</sub></code>。</li>
</ul>

<p>返回数组 <code>answer</code>，数组中按顺序包含第一种类型查询的结果。</p>

<p><b>注意</b>&nbsp;，由于&nbsp;<code>colors</code>&nbsp;表示一个 <strong>环</strong>&nbsp;，<strong>第一块</strong>&nbsp;瓷砖和 <strong>最后一块</strong>&nbsp;瓷砖是相邻的。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">colors = [0,1,1,0,1], queries = [[2,1,0],[1,4]]</span></p>

<p><strong>输出：</strong><span class="example-io">[2]</span></p>

<p><strong>解释：</strong></p>

<p>第一次查询：</p>

<p>将 <code>colors[1]</code> 改为 0。</p>

<p><img alt="" data-darkreader-inline-bgcolor="" data-darkreader-inline-bgimage="" src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-20-25.png" style="width: 150px; height: 150px; padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; --darkreader-inline-bgimage: initial; --darkreader-inline-bgcolor: #181a1b;" /></p>

<p>第二次查询：</p>

<p>统计大小为 4 的交替组的数量：</p>

<p><img alt="" data-darkreader-inline-bgcolor="" data-darkreader-inline-bgimage="" src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-25-02-2.png" style="width: 150px; height: 150px; padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; --darkreader-inline-bgimage: initial; --darkreader-inline-bgcolor: #181a1b;" /><img alt="" data-darkreader-inline-bgcolor="" data-darkreader-inline-bgimage="" src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-24-12.png" style="width: 150px; height: 150px; padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; --darkreader-inline-bgimage: initial; --darkreader-inline-bgcolor: #181a1b;" /></p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">colors = [0,0,1,0,1,1], queries = [[1,3],[2,3,0],[1,5]]</span></p>

<p><strong>输出：</strong><span class="example-io">[2,0]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" data-darkreader-inline-bgcolor="" data-darkreader-inline-bgimage="" src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-35-50.png" style="width: 150px; height: 150px; padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; --darkreader-inline-bgimage: initial; --darkreader-inline-bgcolor: #181a1b;" /></p>

<p>第一次查询：</p>

<p>统计大小为 3 的交替组的数量。</p>

<p><img alt="" data-darkreader-inline-bgcolor="" data-darkreader-inline-bgimage="" src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-37-13.png" style="width: 150px; height: 150px; padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; --darkreader-inline-bgimage: initial; --darkreader-inline-bgcolor: #181a1b;" /><img alt="" data-darkreader-inline-bgcolor="" data-darkreader-inline-bgimage="" src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-36-40.png" style="width: 150px; height: 150px; padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; --darkreader-inline-bgimage: initial; --darkreader-inline-bgcolor: #181a1b;" /></p>

<p>第二次查询：<code>colors</code>不变。</p>

<p>第三次查询：不存在大小为 5 的交替组。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>4 &lt;= colors.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>0 &lt;= colors[i] &lt;= 1</code></li>
	<li><code>1 &lt;= queries.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>queries[i][0] == 1</code> 或 <code>queries[i][0] == 2</code></li>
	<li>对于所有的<code>i</code>：
	<ul>
		<li><code>queries[i][0] == 1</code>： <code>queries[i].length == 2</code>, <code>3 &lt;= queries[i][1] &lt;= colors.length - 1</code></li>
		<li><code>queries[i][0] == 2</code>： <code>queries[i].length == 3</code>, <code>0 &lt;= queries[i][1] &lt;= colors.length - 1</code>, <code>0 &lt;= queries[i][2] &lt;= 1</code></li>
	</ul>
	</li>
</ul>


#### 思路

为方便讨论，下文把 $\textit{colors}$ 数组简记为 $a$。

## 一、非环形数组 + 没有修改操作

这是最简单的情况。

以 $a=[0,1,1,0,1,0,0,1,0]$ 为例，它可以分为三个交替段：$[0,1],[1,0,1,0],[0,1,0]$。

有多少个长度恰好为 $\textit{size}=3$ 的交替子数组？有 $3$ 个：$[1,0,1,0]$ 中有 $2$ 个，$[0,1,0]$ 中有 $1$ 个。

如果遍历所有交替段，时间复杂度就太高了。

一般地，对于长为 $k$ 的交替段，其中有

$$
k - (\textit{size} - 1)
$$

个长度恰好为 $\textit{size}$ 的交替子数组。

假设有 $3$ 个长度均 $\ge \textit{size}$ 的交替段，长度分别为 $k_1,k_2,k_3$，那么其中有

$$
\begin{aligned}
& k_1 - (\textit{size} - 1) + k_2 - (\textit{size} - 1) + k_3 - (\textit{size} - 1)      \\
={} & (k_1+k_2+k_3) - 3 \cdot (\textit{size} - 1)        \\
\end{aligned}
$$

个长度恰好为 $\textit{size}$ 的交替子数组。

这启发我们维护长度 $\ge \textit{size}$ 的交替段的**个数**以及**元素和**。考虑到后续要执行修改操作，用**树状数组**维护。

## 二、非环形数组 + 有修改操作

考虑交替段的结束位置，即 $i=n-1$ 或者 $a[i]=a[i+1]$ 的位置。

如果修改 $a[i]$ 的值，**会影响哪些结束位置**？**会对交替段的长度产生什么影响**？

结束位置 $i-1$ 和 $i$ 会受到影响。

为避免复杂的分类讨论，在修改之前，先移除掉结束位置 $i-1$ 和 $i$（如果是结束位置的话），然后再根据 $a[i-1]=a[i]$ 以及 $a[i]=a[i+1]$ 是否成立，添加结束位置 $i-1$ 和 $i$。
- 当我们添加结束位置 $i$ 时，设 $\textit{pre}$ 和 $\textit{nxt}$ 是 $i$ 前后两个相邻的结束位置，那么一个长为 $\textit{nxt} - \textit{pre}$ 的交替段会被拆分为两个长度分别为 $i-\textit{pre}$ 和 $\textit{nxt}-i$ 的交替段。
- 当我们移除结束位置 $i$ 时，设 $\textit{pre}$ 和 $\textit{nxt}$ 是 $i$ 前后两个相邻的结束位置，那么两个长度分别为 $i-\textit{pre}$ 和 $\textit{nxt}-i$ 的交替段会合并成一个长为 $\textit{nxt} - \textit{pre}$ 的交替段。

如何快速找到 $i$ 的前后两个相邻的结束位置？

用有序集合维护所有结束位置。

## 三、环形数组 + 有修改操作

$a$ 变成环形数组后，哪些计算会发生变化？

1. $i=n-1$ 不一定是结束位置了，必须判断 $a[n-1]=a[0]$ 是否成立才行。
2. 如果小于 $i$ 的最大结束位置不存在，那么取所有结束位置的最大值。
3. 如果大于 $i$ 的最小结束位置不存在，那么取所有结束位置的最小值。
4. 对于交替段长度，例如 $\textit{nxt} - \textit{pre}$，可能会算出负数或 $0$。可以通过 $(\textit{nxt} - \textit{pre} + n - 1)\bmod n + 1$ 调整成 $[1,n]$ 中的数。


### 细节

1. 如果没有结束位置，那么无论询问多长的交替子数组，答案都是 $n$。
2. 如果修改操作不改变 $a[i]$ 的值，则直接 `continue`。

``` 
type fenwickTree [][2]int

// op=1，添加一个 size
// op=-1，移除一个 size
func (t fenwickTree) update(size, op int) {
	for i := len(t) - size; i < len(t); i += i & -i {
		t[i][0] += op
		t[i][1] += op * size
	}
}

// 返回 >= size 的元素个数，元素和
func (t fenwickTree) query(size int) (cnt, sum int) {
	for i := len(t) - size; i > 0; i &= i - 1 {
		cnt += t[i][0]
		sum += t[i][1]
	}
	return
}

func numberOfAlternatingGroups(a []int, queries [][]int) (ans []int) {
	n := len(a)
	set := redblacktree.New[int, struct{}]()
	t := make(fenwickTree, n+1)

	// op=1，添加一个结束位置 i
	// op=-1，移除一个结束位置 i
	update := func(i, op int) {
		prev, ok := set.Floor(i)
		if !ok {
			prev = set.Right()
		}
		pre := prev.Key

		next, ok := set.Ceiling(i)
		if !ok {
			next = set.Left()
		}
		nxt := next.Key

		t.update((nxt-pre+n-1)%n+1, -op) // 移除/添加旧长度
		t.update((i-pre+n)%n, op)
		t.update((nxt-i+n)%n, op) // 添加/移除新长度
	}

	// 添加一个结束位置 i
	add := func(i int) {
		if set.Empty() {
			t.update(n, 1)
		} else {
			update(i, 1)
		}
		set.Put(i, struct{}{})
	}

	// 移除一个结束位置 i
	del := func(i int) {
		set.Remove(i)
		if set.Empty() {
			t.update(n, -1)
		} else {
			update(i, -1)
		}
	}

	for i, c := range a {
		if c == a[(i+1)%n] {
			add(i) // i 是一个结束位置
		}
	}
	for _, q := range queries {
		if q[0] == 1 {
			if set.Empty() {
				ans = append(ans, n) // 每个长为 size 的子数组都符合要求
			} else {
				cnt, sum := t.query(q[1])
				ans = append(ans, sum-cnt*(q[1]-1))
			}
		} else {
			i := q[1]
			if a[i] == q[2] { // 无影响
				continue
			}
			pre, nxt := (i-1+n)%n, (i+1)%n
			// 修改前，先去掉结束位置
			if a[pre] == a[i] {
				del(pre)
			}
			if a[i] == a[nxt] {
				del(i)
			}
			a[i] ^= 1
			// 修改后，添加新的结束位置
			if a[pre] == a[i] {
				add(pre)
			}
			if a[i] == a[nxt] {
				add(i)
			}
		}
	}
	return
}
```


#### 复杂度

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $n$ 是 $\textit{colors}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计


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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)