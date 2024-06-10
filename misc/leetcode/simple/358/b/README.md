#### 题目  

<p>给你一个 <strong>非空</strong> 链表的头节点 <code>head</code> ，表示一个不含前导零的非负数整数。</p>

<p>将链表 <strong>翻倍</strong> 后，返回头节点<em> </em><code>head</code><em> </em>。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/05/28/example.png" style="width: 401px; height: 81px;"/>
<pre><strong>输入：</strong>head = [1,8,9]
<strong>输出：</strong>[3,7,8]
<strong>解释：</strong>上图中给出的链表，表示数字 189 。返回的链表表示数字 189 * 2 = 378 。</pre>

<p><strong class="example">示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/05/28/example2.png" style="width: 401px; height: 81px;"/>
<pre><strong>输入：</strong>head = [9,9,9]
<strong>输出：</strong>[1,9,9,8]
<strong>解释：</strong>上图中给出的链表，表示数字 999 。返回的链表表示数字 999 * 2 = 1998 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li>链表中节点的数目在范围 <code>[1, 10<sup>4</sup>]</code> 内</li>
	<li><font face="monospace"><code>0 &lt;= Node.val &lt;= 9</code></font></li>
	<li>生成的输入满足：链表表示一个不含前导零的数字，除了数字 <code>0</code> 本身。</li>
</ul>
 
#### 思路  

看成是 $\textit{head}$ 与 $\textit{head}$ 这两个链表相加。

```go 
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// 2. 两数相加：自己和自己相加
func double(l1 *ListNode) *ListNode {
	dummy := &ListNode{} // 哨兵节点，作为新链表的头节点的前一个节点
	cur := dummy
	carry := 0 // 进位
	for l1 != nil {
		carry += l1.Val * 2                   // 节点值和进位加在一起
		cur.Next = &ListNode{Val: carry % 10} // 每个节点保存一个数位
		carry /= 10                           // 新的进位
		cur = cur.Next                        // 下一个节点
		l1 = l1.Next                          // 下一个节点
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

func doubleIt(head *ListNode) *ListNode {
	head = reverseList(head)
	res := double(head) // 反转后，就变成【2. 两数相加】了
	return reverseList(res)
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为链表的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。