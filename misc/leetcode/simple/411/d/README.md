#### 题目

<p>给你一个 <strong>二进制</strong> 字符串 <code>s</code> 和一个整数 <code>k</code>。</p>

<p>另给你一个二维整数数组 <code>queries</code> ，其中 <code>queries[i] = [l<sub>i</sub>, r<sub>i</sub>]</code> 。</p>

<p>如果一个 <strong>二进制字符串</strong> 满足以下任一条件，则认为该字符串满足 <strong>k 约束</strong>：</p>

<ul>
	<li>字符串中 <code>0</code> 的数量最多为 <code>k</code>。</li>
	<li>字符串中 <code>1</code> 的数量最多为 <code>k</code>。</li>
</ul>

<p>返回一个整数数组 <code>answer</code> ，其中 <code>answer[i]</code> 表示 <code>s[l<sub>i</sub>..r<sub>i</sub>]</code> 中满足 <strong>k 约束</strong> 的 <span data-keyword="substring-nonempty">子字符串</span> 的数量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">s = "0001111", k = 2, queries = [[0,6]]</span></p>

<p><strong>输出：</strong><span class="example-io">[26]</span></p>

<p><strong>解释：</strong></p>

<p>对于查询 <code>[0, 6]</code>， <code>s[0..6] = "0001111"</code> 的所有子字符串中，除 <code>s[0..5] = "000111"</code> 和 <code>s[0..6] = "0001111"</code> 外，其余子字符串都满足 k 约束。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">s = "010101", k = 1, queries = [[0,5],[1,4],[2,3]]</span></p>

<p><strong>输出：</strong><span class="example-io">[15,9,3]</span></p>

<p><strong>解释：</strong></p>

<p><code>s</code> 的所有子字符串中，长度大于 3 的子字符串都不满足 k 约束。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s[i]</code> 是 <code>'0'</code> 或 <code>'1'</code></li>
	<li><code>1 &lt;= k &lt;= s.length</code></li>
	<li><code>1 &lt;= queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i] == [l<sub>i</sub>, r<sub>i</sub>]</code></li>
	<li><code>0 &lt;= l<sub>i</sub> &lt;= r<sub>i</sub> &lt; s.length</code></li>
	<li>所有查询互不相同</li>
</ul>

#### 思路

**核心思路**：对于每个询问，计算以 $l$ 为右端点的合法子串个数，以 $l+1$ 为右端点的合法子串个数，……，以 $r$ 为右端点的合法子串个数。
我们需要知道以 $i$ 为右端点的合法子串，其左端点最小是多少。
由于随着 $i$ 的变大，窗口内的字符数量变多，越不能满足题目要求，所以最小左端点会随着 $i$ 的增大而增大，有**单调性**，因此可以用 [滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/) 计算。
设以 $i$ 为右端点的合法子串，其左端点**最小**是 $\textit{left}[i]$。\n\n那么以 $i$ 为右端点的合法子串，其左端点可以是 $\textit{left}[i],\textit{left}[i]+1, \cdots, i$，一共 

$$
i-\textit{left}[i]+1
$$ 

个。

回答询问时，分类讨论：
- 如果 $\textit{left}[r] \le l$，说明 $[l,r]$ 内的所有子串都是合法的，这一共有 $1+2+\cdots + (r-l+1) = \dfrac{(r-l+2)(r-l+1)}{2}$ 个。
- 否则，由于 $\textit{left}$ 是**有序数组**，我们可以在 $[l,r]$ 中 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) $\textit{left}$ 中的第一个满足 $\textit{left}[j]\ge l$ 的下标 $j$，那么：
  - 由于 $\textit{left}[j-1] < l$，所以 $[l,j-1]$ 内的所有子串都是合法的，这一共有 $1+2+\cdots + (j-l) = \dfrac{(j-l+1)(j-l)}{2}$ 个。
  - 右端点在 $[j,r]$ 内的子串，可以累加下标在 $[j,r]$ 内的所有 $i-\textit{left}[i]+1$ 的和。这可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 预处理。
  - 上述两项累加，即为答案。
  
代码实现时，两种情况可以合并为一种。

``` 
func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l
		// 计算 i-left[i]+1 的前缀和
		sum[i+1] = sum[i] + i - l + 1
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := l + sort.SearchInts(left[l:r+1], l)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。注意 $\textit{l}$ 只会增加不会减少，所以二重循环的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 优化：双指针预处理

我们可以对于每个 $i$，计算出最小的 $j$，满足 $\textit{left}[j]\ge i$。
由于 $\textit{left}$ 数组是有序的，这个过程可以用**双指针**实现。
将计算出的 $j$ 保存到 $\textit{left}_2[j]$ 中。\n\n这样回答询问时，$j$ 就可以直接通过 $\textit{left}_2[l]$ 获取到了，注意这个数不能超过 $r+1$，所以有

$$
j = \min(\textit{left}_2[l], r+1)
$$

``` 
func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l
		// 计算 i-left[i]+1 的前缀和
		sum[i+1] = sum[i] + i - l + 1
	}

	left2 := make([]int, n)
	l = 0
	for i := range left2 {
		for l < n && left[l] < i {
			l++
		}
		left2[i] = l
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := min(left2[l], r+1)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。注意 $\textit{l}$ 只会增加不会减少，所以二重循环的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

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