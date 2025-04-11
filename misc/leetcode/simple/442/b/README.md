#### 题目

<p>给你一个二维整数数组 <code>properties</code>，其维度为 <code>n x m</code>，以及一个整数 <code>k</code>。</p>

<p>定义一个函数 <code>intersect(a, b)</code>，它返回数组 <code>a</code> 和 <code>b</code> 中<strong> 共有的不同整数的数量 </strong>。</p>

<p>构造一个 <strong>无向图</strong>，其中每个索引 <code>i</code> 对应 <code>properties[i]</code>。如果且仅当 <code>intersect(properties[i], properties[j]) &gt;= k</code>（其中 <code>i</code> 和 <code>j</code> 的范围为 <code>[0, n - 1]</code> 且 <code>i != j</code>），节点 <code>i</code> 和节点 <code>j</code> 之间有一条边。</p>

<p>返回结果图中<strong> 连通分量 </strong>的数量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">properties = [[1,2],[1,1],[3,4],[4,5],[5,6],[7,7]], k = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<p>生成的图有 3 个连通分量：</p>

<p><img src="https://pic.leetcode.cn/1742665594-CDVPWz-image.png" style="width: 279px; height: 171px;" /></p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">properties = [[1,2,3],[2,3,4],[4,3,5]], k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">1</span></p>

<p><strong>解释：</strong></p>

<p>生成的图有 1 个连通分量：</p>

<p><img alt="" src="https://pic.leetcode.cn/1742665565-NzYlYH-screenshot-from-2025-02-27-23-58-34.png" style="width: 219px; height: 171px;" /></p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">properties = [[1,1],[1,1]], k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p><code>intersect(properties[0], properties[1]) = 1</code>，小于 <code>k</code>。因此在图中 <code>properties[0]</code> 和 <code>properties[1]</code> 之间没有边。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == properties.length &lt;= 100</code></li>
	<li><code>1 &lt;= m == properties[i].length &lt;= 100</code></li>
	<li><code>1 &lt;= properties[i][j] &lt;= 100</code></li>
	<li><code>1 &lt;= k &lt;= m</code></li>
</ul>

#### 思路

暴力枚举所有 $\textit{properties}[i]$ 和 $\textit{properties}[j]$，如果交集大小 $\ge k$，
那么用并查集合并 $i$ 和 $j$。
初始连通块个数 $\textit{cc}=n$，每成功合并一次，就把 $\textit{cc}$ 减一

```
type uf struct {
	fa []int
	cc int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa, n}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return
	}
	u.fa[x] = y
	u.cc--
}

func numberOfComponents(properties [][]int, k int) int {
	sets := make([]map[int]bool, len(properties))
	for i, a := range properties {
		sets[i] = map[int]bool{}
		for _, x := range a {
			sets[i][x] = true
		}
	}

	u := newUnionFind(len(properties))
	for i, a := range sets {
		for j, b := range sets[:i] {
			cnt := 0
			for x := range b {
				if a[x] {
					cnt++
				}
			}
			if cnt >= k {
				u.merge(i, j)
			}
		}
	}
	return u.cc
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2m)$ 或者 $\mathcal{O}(n^2m/w)$，其中 $n$ 是 $\textit{properties}$ 的长度，$m$ 是 $\textit{properties}[i]$ 的长度，$w$ 等于 $32$ 或 $64$。
- 空间复杂度：$\mathcal{O}(nm)$ 或者 $\mathcal{O}(nm/w)$。

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