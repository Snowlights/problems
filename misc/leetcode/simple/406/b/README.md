#### 题目

<p>给你一个整数数组 <code>nums</code> 和一个链表的头节点 <code>head</code>。从链表中<strong>移除</strong>所有存在于 <code>nums</code> 中的节点后，返回修改后的链表的头节点。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,3], head = [1,2,3,4,5]</span></p>

<p><strong>输出：</strong> <span class="example-io">[4,5]</span></p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/06/11/linkedlistexample0.png" style="width: 400px; height: 66px;" /></strong></p>

<p>移除数值为 1, 2 和 3 的节点。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1], head = [1,2,1,2,1,2]</span></p>

<p><strong>输出：</strong> <span class="example-io">[2,2,2]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/11/linkedlistexample1.png" style="height: 62px; width: 450px;" /></p>

<p>移除数值为 1 的节点。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [5], head = [1,2,3,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">[1,2,3,4]</span></p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/06/11/linkedlistexample2.png" style="width: 400px; height: 83px;" /></strong></p>

<p>链表中不存在值为 5 的节点。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>nums</code> 中的所有元素都是唯一的。</li>
	<li>链表中的节点数在 <code>[1, 10<sup>5</sup>]</code> 的范围内。</li>
	<li><code>1 &lt;= Node.val &lt;= 10<sup>5</sup></code></li>
	<li>输入保证链表中至少有一个值没有在&nbsp;<code>nums</code> 中出现过。</li>
</ul>

#### 思路

如何在遍历链表的同时，删除链表节点？请看[【基础算法精讲 08】](https://www.bilibili.com/video/BV1VP4y1Q71e/)。

对于本题，由于直接判断节点值是否在 $\textit{nums}$ 中，需要遍历 $\textit{nums}$，时间复杂度为 $\mathcal{O}(n)$。通过把 $\textit{nums}$ 中的元素加到一个哈希集合中，然后判断节点值是否在哈希集合中，这样可以做到每次判断时间复杂度为 $\mathcal{O}(1)$。

具体做法：

1. 把 $\textit{nums}$ 中的元素加到一个哈希集合中。
2. 由于头节点可能会被删除，在头节点前面插入一个哨兵节点 $\textit{dummy}$，以简化代码逻辑。
3. 初始化 $\textit{cur} = \textit{dummy}$。
4. 不断循环，直到 $\textit{cur}$ 没有下一个节点。
5. 如果 $\textit{cur}$ 的下一个节点的值在哈希集合中，则需要删除，更新 $\textit{cur}.\textit{next}$ 为 $\textit{cur}.\textit{next}.\textit{next}$；否则不删除，更新 $\textit{cur}$ 为 $\textit{cur}.\textit{next}$。
6. 循环结束后，返回 $\textit{dummy}.\textit{next}$。

⚠注意：$\textit{dummy}$ 和 $\textit{cur}$ 是同一个节点的引用，修改 $\textit{cur}.\textit{next}$ 也会修改 $\textit{dummy}.\textit{next}$。

```
func modifiedList(nums []int, head *ListNode) *ListNode {
	has := make(map[int]bool, len(nums)) // 预分配空间
	for _, x := range nums {
		has[x] = true
	}
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if has[cur.Next.Val] {
			cur.Next = cur.Next.Next // 删除
		} else {
			cur = cur.Next // 向后移动
		}
	}
	return dummy.Next
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是链表的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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