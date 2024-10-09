#### 题目  

<p>给你一个 <code>n</code> 个节点的无向无根树，节点编号从 <code>0</code> 到 <code>n - 1</code> 。给你整数 <code>n</code> 和一个长度为 <code>n - 1</code> 的二维整数数组 <code>edges</code> ，其中 <code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>]</code> 表示树中节点 <code>a<sub>i</sub></code> 和 <code>b<sub>i</sub></code> 之间有一条边。再给你一个长度为 <code>n</code> 的数组 <code>coins</code> ，其中 <code>coins[i]</code> 可能为 <code>0</code> 也可能为 <code>1</code> ，<code>1</code> 表示节点 <code>i</code> 处有一个金币。</p>

<p>一开始，你需要选择树中任意一个节点出发。你可以执行下述操作任意次：</p>

<ul>
	<li>收集距离当前节点距离为 <code>2</code> 以内的所有金币，或者</li>
	<li>移动到树中一个相邻节点。</li>
</ul>

<p>你需要收集树中所有的金币，并且回到出发节点，请你返回最少经过的边数。</p>

<p>如果你多次经过一条边，每一次经过都会给答案加一。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/03/01/graph-2.png" style="width: 522px; height: 522px;"/></p>

<pre><b>输入：</b>coins = [1,0,0,0,0,1], edges = [[0,1],[1,2],[2,3],[3,4],[4,5]]
<b>输出：</b>2
<b>解释：</b>从节点 2 出发，收集节点 0 处的金币，移动到节点 3 ，收集节点 5 处的金币，然后移动回节点 2 。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/03/02/graph-4.png" style="width: 522px; height: 522px;"/></p>

<pre><b>输入：</b>coins = [0,0,0,1,1,0,0,1], edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[5,6],[5,7]]
<b>输出：</b>2
<b>解释：</b>从节点 0 出发，收集节点 4 和 3 处的金币，移动到节点 2 处，收集节点 7 处的金币，移动回节点 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == coins.length</code></li>
	<li><code>1 &lt;= n &lt;= 3 * 10<sup>4</sup></code></li>
	<li><code>0 &lt;= coins[i] &lt;= 1</code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i</sub> &lt; n</code></li>
	<li><code>a<sub>i</sub> != b<sub>i</sub></code></li>
	<li><code>edges</code> 表示一棵合法的树。</li>
</ul>
 
#### 思路  

#### 提示 1

去掉不包含金币的子树，访问其中任何一个点都毫无意义。  
做法：从没有金币的叶子出发，跑拓扑排序。  
注意，去掉这些子树后，某些原来不是叶子的节点会变成叶子。

#### 提示 2

只需要考虑有金币的的叶子，因为不在叶子上的金币**顺路**就能收集到。

#### 提示 3

从有金币的的叶子出发，再次跑拓扑排序。在拓扑排序的同时，标记每个点入队的时间 $\textit{time}$。  
**注意是入队的时间，不是访问到这个节点的时间。**
- 叶子入队的时间为 $0$；
- 去掉这些叶子后，又产生了**新的叶子**，这些叶子入队的时间为 $1$；
- 去掉这些叶子后，又产生了**新的叶子**，这些叶子入队的时间为 $2$；
- ……

示例 2 如下图，数字表示节点入队的时间：
![t4.png](https://pic.leetcode.cn/1679802238-QZehnH-t4.png)
那么只要走到 $\textit{time}[x]=2$ 的节点 $x$，就能收集到在叶子上的金币。  
遍历所有边 $x-y$，如果满足 $\textit{time}[x]\ge 2$ 且 $\textit{time}[y]\ge 2$（上图蓝色边），那么这条边需要恰好经过 $2$ 次（因为需要回到出发点），答案加 $2$；如果不满足，则无需经过。
> 注：从任意被蓝色边连接的点出发，算出来的答案都是一样的。

```go 
func collectTheCoins(coins []int, edges [][]int) (ans int) {
	n := len(coins)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
		deg[x]++
		deg[y]++
	}

	// 用拓扑排序「剪枝」：去掉没有金币的子树
	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 1 && coins[i] == 0 { // 无金币叶子
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 && coins[y] == 0 {
				q = append(q, y)
			}
		}
	}

	// 再次拓扑排序
	for i, d := range deg {
		if d == 1 && coins[i] == 1 { // 有金币叶子
			q = append(q, i)
		}
	}
	if len(q) <= 1 { // 至多一个有金币的叶子，直接收集
		return
	}
	time := make([]int, n)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 {
				time[y] = time[x] + 1 // 记录入队时间
				q = append(q, y)
			}
		}
	}

	// 统计答案
	for _, e := range edges {
		if time[e[0]] >= 2 && time[e[1]] >= 2 {
			ans += 2
		}
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{coins}$ 的长度。
- 空间复杂度：$O(n)$。