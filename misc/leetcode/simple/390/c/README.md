#### 题目

<p>你需要在一个集合里动态记录 ID 的出现频率。给你两个长度都为 <code>n</code>&nbsp;的整数数组&nbsp;<code>nums</code> 和&nbsp;<code>freq</code>&nbsp;，<code>nums</code>&nbsp;中每一个元素表示一个 ID ，对应的 <code>freq</code>&nbsp;中的元素表示这个 ID 在集合中此次操作后需要增加或者减少的数目。</p>

<ul>
	<li><strong>增加 ID 的数目：</strong>如果&nbsp;<code>freq[i]</code>&nbsp;是正数，那么&nbsp;<code>freq[i]</code>&nbsp;个 ID 为&nbsp;<code>nums[i]</code>&nbsp;的元素在第 <code>i</code>&nbsp;步操作后会添加到集合中。</li>
	<li><strong>减少 ID 的数目：</strong>如果&nbsp;<code>freq[i]</code>&nbsp;是负数，那么&nbsp;<code>-freq[i]</code>&nbsp;个 ID 为&nbsp;<code>nums[i]</code>&nbsp;的元素在第 <code>i</code>&nbsp;步操作后会从集合中删除。</li>
</ul>

<p>请你返回一个长度为 <code>n</code>&nbsp;的数组 <code>ans</code>&nbsp;，其中&nbsp;<code>ans[i]</code>&nbsp;表示第 <code>i</code>&nbsp;步操作后出现频率最高的 ID <strong>数目</strong>&nbsp;，如果在某次操作后集合为空，那么 <code>ans[i]</code>&nbsp;为 0 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,3,2,1], freq = [3,2,-3,1]</span></p>

<p><span class="example-io"><b>输出：</b>[3,3,2,2]</span></p>

<p><strong>解释：</strong></p>

<p>第 0 步操作后，有 3 个 ID 为 2 的元素，所以&nbsp;<code>ans[0] = 3</code>&nbsp;。<br />
第 1 步操作后，有 3 个 ID 为 2 的元素和 2 个 ID 为 3 的元素，所以&nbsp;<code>ans[1] = 3</code>&nbsp;。<br />
第 2 步操作后，有 2 个 ID 为 3 的元素，所以&nbsp;<code>ans[2] = 2</code>&nbsp;。<br />
第 3 步操作后，有 2 个 ID 为 3 的元素和 1 个 ID 为 1 的元素，所以&nbsp;<code>ans[3] = 2</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [5,5,3], freq = [2,-2,1]</span></p>

<p><span class="example-io"><b>输出：</b>[2,0,1]</span></p>

<p><strong>解释：</strong></p>

<p>第 0 步操作后，有 2 个 ID 为 5 的元素，所以&nbsp;<code>ans[0] = 2</code>&nbsp;。<br />
第 1 步操作后，集合中没有任何元素，所以&nbsp;<code>ans[1] = 0</code>&nbsp;。<br />
第 2 步操作后，有 1 个 ID 为 3 的元素，所以&nbsp;<code>ans[2] = 1</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length == freq.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>5</sup> &lt;= freq[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>freq[i] != 0</code></li>
	<li>输入保证任何操作后，集合中的元素出现次数不会为负数。</li>
</ul>

#### 思路

## 方法一：哈希表 + 有序集合

做法类似前几天的每日一题 [2671. 频率跟踪器](https://leetcode.cn/problems/frequency-tracker/)：

- 用哈希表 $\textit{cnt}$ 记录 $x=\textit{nums}[i]$ 的出现次数 $\textit{cnt}[x]$（用 $\textit{freq}$ 更新出现次数）。
- 用有序集合记录 $\textit{cnt}[x]$ 的出现次数，从而可以 $\mathcal{O}(\log n)$ 知道最大的 $\textit{cnt}[x]$ 是多少。


```go [sol-Go]
// https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha
func mostFrequentIDs(nums, freq []int) []int64 {
	cnt := map[int]int{}
	t := redblacktree.New[int, int]()
	ans := make([]int64, len(nums))
	for i, x := range nums {
		// 减少一次 cnt[x] 的出现次数
		node := t.GetNode(cnt[x])
		if node != nil {
			node.Value--
			if node.Value == 0 {
				t.Remove(node.Key)
			}
		}

		cnt[x] += freq[i]

		// 增加一次 cnt[x] 的出现次数
		node = t.GetNode(cnt[x])
		if node == nil {
			t.Put(cnt[x], 1)
		} else {
			node.Value++
		}
		ans[i] = int64(t.Right().Key)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：哈希表 + 懒删除堆

也可以不用有序集合，而是用一个**最大堆**保存数对 $(\textit{cnt}[x], x)$。

在堆中查询 $\textit{cnt}[x]$ 的最大值时，如果堆顶保存的数据并不是目前实际的 $\textit{cnt}[x]$，那么就弹出堆顶。

```go [sol-Go]
func mostFrequentIDs(nums, freq []int) []int64 {
	ans := make([]int64, len(nums))
	cnt := make(map[int]int)
	h := hp{}
	heap.Init(&h)
	for i, x := range nums {
		cnt[x] += freq[i]
		heap.Push(&h, pair{cnt[x], x})
		for h[0].c != cnt[h[0].x] { // 堆顶保存的数据已经发生变化
			heap.Pop(&h) // 删除
		}
		ans[i] = int64(h[0].c)
	}
	return ans
}

type pair struct{ c, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].c > h[j].c } // 最大堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。由于至多入堆 $\mathcal{O}(n)$ 次，所以出堆也至多 $\mathcal{O}(n)$ 次，二重循环的时间复杂度为 $\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

