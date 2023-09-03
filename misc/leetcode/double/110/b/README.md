### 题目  

<p>给你一个链表的头 <code>head</code> ，每个结点包含一个整数值。</p>

<p>在相邻结点之间，请你插入一个新的结点，结点值为这两个相邻结点值的 <strong>最大公约数</strong> 。</p>

<p>请你返回插入之后的链表。</p>

<p>两个数的 <strong>最大公约数</strong> 是可以被两个数字整除的最大正整数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/07/18/ex1_copy.png" style="width: 641px; height: 181px;"/></p>

<pre><b>输入：</b>head = [18,6,10,3]
<b>输出：</b>[18,6,6,2,10,1,3]
<b>解释：</b>第一幅图是一开始的链表，第二幅图是插入新结点后的图（蓝色结点为新插入结点）。
- 18 和 6 的最大公约数为 6 ，插入第一和第二个结点之间。
- 6 和 10 的最大公约数为 2 ，插入第二和第三个结点之间。
- 10 和 3 的最大公约数为 1 ，插入第三和第四个结点之间。
所有相邻结点之间都插入完毕，返回链表。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/07/18/ex2_copy1.png" style="width: 51px; height: 191px;"/></p>

<pre><b>输入：</b>head = [7]
<strong>输出：</strong>[7]
<b>解释：</b>第一幅图是一开始的链表，第二幅图是插入新结点后的图（蓝色结点为新插入结点）。
没有相邻结点，所以返回初始链表。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li>链表中结点数目在 <code>[1, 5000]</code> 之间。</li>
	<li><code>1 &lt;= Node.val &lt;= 1000</code></li>
</ul>
 
### 思路  

遍历链表，在当前节点 $\textit{cur}$ 后面插入 $\textit{gcd}$ 节点，$\textit{gcd}$ 节点指向 $\textit{cur}$ 的下一个节点。  
插入后，$\textit{cur}$ 更新为 $\textit{cur}.\textit{next}.\textit{next}$。  
循环直到 $\textit{cur}$ 没有下一个节点为止。

```go 
func insertGreatestCommonDivisors(head *ListNode) (ans *ListNode) {
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}

	cur, next := head, head.Next
	for cur != nil && next != nil {
		val := gcd(cur.Val, next.Val)
		node := &ListNode{
			Val:  val,
			Next: next,
		}
		cur.Next = node
		cur = next
		next = next.Next
	}

	return head
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为链表长度，$U$ 为节点值的最大值。
- 空间复杂度：$\mathcal{O}(1)$。返回值的空间不计入，仅用到若干额外变量。