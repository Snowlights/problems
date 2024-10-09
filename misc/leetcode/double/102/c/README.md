### 题目  

<p>给你一棵二叉树的根 <code>root</code> ，请你将每个节点的值替换成该节点的所有 <strong>堂兄弟节点值的和 </strong>。</p>

<p>如果两个节点在树中有相同的深度且它们的父节点不同，那么它们互为 <strong>堂兄弟</strong> 。</p>

<p>请你返回修改值之后，树的根<em> </em><code>root</code><em> </em>。</p>

<p><strong>注意</strong>，一个节点的深度指的是从树根节点到这个节点经过的边数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/01/11/example11.png" style="width: 571px; height: 151px;"/></p>

<pre><b>输入：</b>root = [5,4,9,1,10,null,7]
<b>输出：</b>[0,0,0,7,7,null,11]
<b>解释：</b>上图展示了初始的二叉树和修改每个节点的值之后的二叉树。
- 值为 5 的节点没有堂兄弟，所以值修改为 0 。
- 值为 4 的节点没有堂兄弟，所以值修改为 0 。
- 值为 9 的节点没有堂兄弟，所以值修改为 0 。
- 值为 1 的节点有一个堂兄弟，值为 7 ，所以值修改为 7 。
- 值为 10 的节点有一个堂兄弟，值为 7 ，所以值修改为 7 。
- 值为 7 的节点有两个堂兄弟，值分别为 1 和 10 ，所以值修改为 11 。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/01/11/diagram33.png" style="width: 481px; height: 91px;"/></p>

<pre><b>输入：</b>root = [3,1,2]
<b>输出：</b>[0,0,0]
<b>解释：</b>上图展示了初始的二叉树和修改每个节点的值之后的二叉树。
- 值为 3 的节点没有堂兄弟，所以值修改为 0 。
- 值为 1 的节点没有堂兄弟，所以值修改为 0 。
- 值为 2 的节点没有堂兄弟，所以值修改为 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中节点数目的范围是 <code>[1, 10<sup>5</sup>]</code> 。</li>
	<li><code>1 &lt;= Node.val &lt;= 10<sup>4</sup></code></li>
</ul>
 
### 思路 

### 提示 1

下文将具有相同父节点的节点互称为**兄弟节点**。  
对于一个节点 $x$ 来说，它的所有**堂兄弟节点**值的和，等价于 $x$ 这一层的所有节点值之和， 减去 $x$ 及其兄弟节点的值之和。
例如样例 1：
- $4$ 的左右儿子的节点值，都被更新成了 $7$，也就是左右儿子这一层的节点值之和 $1+10+7=18$，减去 $4$ 的左右儿子的节点值之和 $1+10=11$，得到 $7$。
- $9$ 的右儿子的节点值，被更新成了 $11$，也就是右儿子这一层的节点值之和 $1+10+7=18$，减去 $9$ 的右儿子的节点值 $7$，得到 $11$。

### 提示 2

怎么实现呢？  
用 BFS 遍历二叉树，对于每一层：
- 首先，遍历当前层的每个节点，通过节点的左右儿子，计算下一层的节点值之和 $\textit{nextLevelSum}$；
- 然后，再次遍历当前层的每个节点 $x$，计算 $x$ 的左右儿子的节点值之和 $\textit{childrenSum}$，更新 $x$ 的左右儿子的节点值为 $\textit{nextLevelSum}-\textit{childrenSum}$。

```go 
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		nextLevelSum := 0 // 下一层的节点值之和
		for _, node := range tmp {
			if node.Left != nil {
				q = append(q, node.Left)
				nextLevelSum += node.Left.Val
			}
			if node.Right != nil {
				q = append(q, node.Right)
				nextLevelSum += node.Right.Val
			}
		}

		// 再次遍历，更新下一层的节点值
		for _, node := range tmp {
			childrenSum := 0 // node 左右儿子的节点值之和
			if node.Left != nil {
				childrenSum += node.Left.Val
			}
			if node.Right != nil {
				childrenSum += node.Right.Val
			}
			// 更新 node 左右儿子的节点值
			if node.Left != nil {
				node.Left.Val = nextLevelSum - childrenSum
			}
			if node.Right != nil {
				node.Right.Val = nextLevelSum - childrenSum
			}
		}
	}
	return root
}
```

### 复杂度分析  

- 时间复杂度：$O(n)$，其中 $n$ 为二叉树的节点个数。
- 空间复杂度：$O(n)$。