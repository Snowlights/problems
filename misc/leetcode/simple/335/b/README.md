#### 题目  

<p>给你一棵二叉树的根节点 <code>root</code> 和一个正整数 <code>k</code> 。</p>

<p>树中的 <strong>层和</strong> 是指 <strong>同一层</strong> 上节点值的总和。</p>

<p>返回树中第 <code>k</code> 大的层和（不一定不同）。如果树少于 <code>k</code> 层，则返回 <code>-1</code> 。</p>

<p><strong>注意</strong>，如果两个节点与根节点的距离相同，则认为它们在同一层。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2022/12/14/binaryytreeedrawio-2.png" style="width: 301px; height: 284px;"/></p>

<pre><strong>输入：</strong>root = [5,8,9,2,1,3,7,4,6], k = 2
<strong>输出：</strong>13
<strong>解释：</strong>树中每一层的层和分别是：
- Level 1: 5
- Level 2: 8 + 9 = 17
- Level 3: 2 + 1 + 3 + 7 = 13
- Level 4: 4 + 6 = 10
第 2 大的层和等于 13 。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2022/12/14/treedrawio-3.png" style="width: 181px; height: 181px;"/></p>

<pre><strong>输入：</strong>root = [1,2,null,3], k = 1
<strong>输出：</strong>3
<strong>解释：</strong>最大的层和是 3 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中的节点数为 <code>n</code></li>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= Node.val &lt;= 10<sup>6</sup></code></li>
	<li><code>1 &lt;= k &lt;= n</code></li>
</ul>
 
#### 思路  

BFS 二叉树，记录每一层的节点值之和，排序后取第 $k$ 大（也可以用快速选择）。

```go 
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	sum := []int{}
	for len(q) > 0 {
		tmp, s := q, 0
		q = nil
		for _, node := range tmp {
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		sum = append(sum, s)
	}
	n := len(sum)
	if n < k {
		return -1
	}
	sort.Ints(sum) // 也可以用快速选择
	return int64(sum[n-k])
}

```

#### 复杂度分析  

- 时间复杂度：$O(n\log n)$，其中 $n$ 为二叉树的节点个数。
- 空间复杂度：$O(n)$。