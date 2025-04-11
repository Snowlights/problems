#### 题目

<p>给你两个长度为 <code>n</code>&nbsp;的整数数组，<code>fruits</code> 和 <code>baskets</code>，其中 <code>fruits[i]</code> 表示第 <code>i</code>&nbsp;种水果的 <strong>数量</strong>，<code>baskets[j]</code> 表示第 <code>j</code>&nbsp;个篮子的 <strong>容量</strong>。</p>

<p>你需要对 <code>fruits</code> 数组从左到右按照以下规则放置水果：</p>

<ul>
	<li>每种水果必须放入第一个 <strong>容量大于等于</strong> 该水果数量的 <strong>最左侧可用篮子</strong> 中。</li>
	<li>每个篮子只能装 <b>一种</b> 水果。</li>
	<li>如果一种水果 <b>无法放入</b> 任何篮子，它将保持 <b>未放置</b>。</li>
</ul>

<p>返回所有可能分配完成后，剩余未放置的水果种类的数量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">fruits = [4,2,5], baskets = [3,5,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">1</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><code>fruits[0] = 4</code> 放入 <code>baskets[1] = 5</code>。</li>
	<li><code>fruits[1] = 2</code> 放入 <code>baskets[0] = 3</code>。</li>
	<li><code>fruits[2] = 5</code> 无法放入 <code>baskets[2] = 4</code>。</li>
</ul>

<p>由于有一种水果未放置，我们返回 1。</p>
</div>

<p><strong class="example">示例 2</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">fruits = [3,6,1], baskets = [6,4,7]</span></p>

<p><strong>输出：</strong> <span class="example-io">0</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><code>fruits[0] = 3</code> 放入 <code>baskets[0] = 6</code>。</li>
	<li><code>fruits[1] = 6</code> 无法放入 <code>baskets[1] = 4</code>（容量不足），但可以放入下一个可用的篮子 <code>baskets[2] = 7</code>。</li>
	<li><code>fruits[2] = 1</code> 放入 <code>baskets[1] = 4</code>。</li>
</ul>

<p>由于所有水果都已成功放置，我们返回 0。</p>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>n == fruits.length == baskets.length</code></li>
	<li><code>1 &lt;= n &lt;= 100</code></li>
	<li><code>1 &lt;= fruits[i], baskets[i] &lt;= 1000</code></li>
</ul>

#### 思路

**前置知识**：线段树。

用线段树维护 $\textit{baskets}$ 的区间最大值。

对于 $x=\textit{fruits}[i]$，在线段树上二分找第一个 $\ge x$ 的数。

- 如果整个区间的最大值都小于 $x$，那么没有这样的数，返回 $-1$。
- 如果能递归到叶子，返回叶子对应的区间端点。
- 先递归左子树。
- 如果左子树没找到，再递归右子树。

如果没有找到这样的数，把答案加一。

否则，把对应的位置改成 $-1$，表示不能放水果。

```
type seg []int

func (t seg) maintain(o int) {
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 初始化线段树
func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 找区间内的第一个 >= x 的数，并更新为 -1，返回这个数的下标（没有则返回 -1）
func (t seg) findFirstAndUpdate(o, l, r, x int) int {
	if t[o] < x { // 区间没有 >= x 的数
		return -1
	}
	if l == r {
		t[o] = -1 // 更新为 -1，表示不能放水果
		return l
	}
	m := (l + r) >> 1
	i := t.findFirstAndUpdate(o<<1, l, m, x) // 先递归左子树
	if i < 0 { // 左子树没找到
		i = t.findFirstAndUpdate(o<<1|1, m+1, r, x) // 再递归右子树
	}
	t.maintain(o)
	return i
}

func newSegmentTree(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func numOfUnplacedFruits(fruits, baskets []int) (ans int) {
	t := newSegmentTree(baskets)
	for _, x := range fruits {
		if t.findFirstAndUpdate(1, 0, len(baskets)-1, x) < 0 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{fruits}$ 的长度，也是 $\textit{baskets}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)