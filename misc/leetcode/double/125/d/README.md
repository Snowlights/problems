### 题目

<p>给你一棵 <code>n</code> 个节点的 <strong>无向</strong> 树，节点从 <code>0</code> 到 <code>n - 1</code> 编号。树以长度为 <code>n - 1</code> 下标从 <strong>0</strong> 开始的二维整数数组 <code>edges</code> 的形式给你，其中 <code>edges[i] = [u<sub>i</sub>, v<sub>i</sub>]</code> 表示树中节点 <code>u<sub>i</sub></code> 和 <code>v<sub>i</sub></code> 之间有一条边。同时给你一个 <strong>正</strong> 整数 <code>k</code> 和一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的 <strong>非负</strong> 整数数组 <code>nums</code> ，其中 <code>nums[i]</code> 表示节点 <code>i</code> 的 <strong>价值</strong> 。</p>

<p>日增哥哥想 <strong>最大化</strong> 树中所有节点价值之和。为了实现这一目标，日增哥哥可以执行以下操作 <strong>任意</strong> 次（<strong>包括</strong><strong> 0 次</strong>）：</p>

<ul>
	<li>选择连接节点 <code>u</code> 和 <code>v</code> 的边 <code>[u, v]</code> ，并将它们的值更新为：
	<ul>
		<li><code>nums[u] = nums[u] XOR k</code></li>
		<li><code>nums[v] = nums[v] XOR k</code></li>
	</ul>
	</li>
</ul>

<p>请你返回日增哥哥通过执行以上操作 <strong>任意次</strong> 后，可以得到所有节点 <strong>价值之和</strong> 的 <strong>最大值</strong> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012513.png" style="width: 300px; height: 277px;padding: 10px; background: #fff; border-radius: .5rem;" /></p>

<pre>
<b>输入：</b>nums = [1,2,1], k = 3, edges = [[0,1],[0,2]]
<b>输出：</b>6
<b>解释：</b>日增哥哥可以通过一次操作得到最大价值和 6 ：
- 选择边 [0,2] 。nums[0] 和 nums[2] 都变为：1 XOR 3 = 2 ，数组 nums 变为：[1,2,1] -> [2,2,2] 。
所有节点价值之和为 2 + 2 + 2 = 6 。
6 是可以得到最大的价值之和。
</pre>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/01/09/screenshot-2024-01-09-220017.png" style="padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; width: 300px; height: 239px;" /></p>

<pre>
<b>输入：</b>nums = [2,3], k = 7, edges = [[0,1]]
<b>输出：</b>9
<b>解释：</b>日增哥哥可以通过一次操作得到最大和 9 ：
- 选择边 [0,1] 。nums[0] 变为：2 XOR 7 = 5 ，nums[1] 变为：3 XOR 7 = 4 ，数组 nums 变为：[2,3] -> [5,4] 。
所有节点价值之和为 5 + 4 = 9 。
9 是可以得到最大的价值之和。
</pre>

<p><strong class="example">示例 3：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012641.png" style="width: 600px; height: 233px;padding: 10px; background: #fff; border-radius: .5rem;" /></p>

<pre>
<b>输入：</b>nums = [7,7,7,7,7,7], k = 3, edges = [[0,1],[0,2],[0,3],[0,4],[0,5]]
<b>输出：</b>42
<b>解释：</b>日增哥哥不需要执行任何操作，就可以得到最大价值之和 42 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= n == nums.length <= 2 * 10<sup>4</sup></code></li>
	<li><code>1 <= k <= 10<sup>9</sup></code></li>
	<li><code>0 <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>0 <= edges[i][0], edges[i][1] <= n - 1</code></li>
	<li>输入保证 <code>edges</code> 构成一棵合法的树。</li>
</ul>

### 思路

## 方法一：树形 DP

用「选或不选」思考。

对于以 $x$ 为根的子树，考虑 $x$ 和它的儿子 $y$ 之间的边是否操作。

- 定义 $f[x][0]$ 表示 $x$ 操作偶数次时，子树 $x$ 的除去 $x$ 的最大价值和。
- 定义 $f[x][1]$ 表示 $x$ 操作奇数次时，子树 $x$ 的除去 $x$ 的最大价值和。

初始化 $f[x][0]=0,\ f[x][1] = -\infty$。遍历并递归计算 $x$ 的所有儿子，设当前遍历到的儿子为 $y$，

- 情况一，不操作 $x$ 和 $y$：
  - 设 $r_0 = \max(f[y][0] + \textit{nums}[y], f[y][1] + (\textit{nums}[y] \oplus k))$。这是不操作 $x$ 和 $y$ 时，子树 $y$ 的最大价值和。
  - $f[x][0] = f[x][0] + r_0$。
  - $f[x][1] = f[x][1] + r_0$。
- 情况二，操作 $x$ 和 $y$：
  - 设 $r_1 = \max(f[y][1] + \textit{nums}[y], f[y][0] + (\textit{nums}[y] \oplus k))$。这是操作 $x$ 和 $y$ 时，子树 $y$ 的最大价值和。
  - $f[x][0] = f[x][1] + r_1$。注意操作后，$x$ 的操作次数的奇偶性变化了。
  - $f[x][1] = f[x][0] + r_1$。

两种情况取最大值，有

$$
\begin{aligned}
&f[x][0] = \max(f[x][0] + r_0, f[x][1] + r_1)\\&f[x][1] = \max(f[x][1] + r_0, f[x][0] + r_1) 
\end{aligned}

$$

注意这两个转移是同时发生的。最后答案为根节点对应的 $r_0$。

```go [sol]
func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		f0, f1 := 0, math.MinInt // f[x][0] 和 f[x][1]
		for _, y := range g[x] {
			if y != fa {
				r0, r1 := dfs(y, x)
				f0, f1 = max(f0+r0, f1+r1), max(f1+r0, f0+r1)
			}
		}
		return max(f0+nums[x], f1+(nums[x]^k)), max(f1+nums[x], f0+(nums[x]^k))
	}
	ans, _ := dfs(0, -1)
	return int64(ans)
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：结论 + 状态机 DP

#### 第一个观察

由于一个数异或两次（偶数次）$k$ 后保持不变，所以对于一条从 $x$ 到 $y$ 的简单路径，我们把路径上的所有边操作后，路径上除了 $x$ 和 $y$ 的其它节点都恰好操作两次，所以只有 $\textit{nums}[x]$ 和 $\textit{nums}[y]$ 都异或了 $k$，其余元素不变。
所以题目中的操作可以作用在**任意两个数**上。我们不需要建树，$\textit{edges}$ 是多余的。

#### 第二个观察

注意到，无论操作多少次，总是有**偶数个**元素异或了 $k$，其余元素不变，理由如下：

- 如果我们操作的两个数之前都没有异或过，那么操作后，异或 $k$ 的元素增加了 $2$。
- 如果我们操作的两个数之前都异或过，那么操作后，异或 $k$ 的元素减少了 $2$。
- 如果我们操作的两个数之前一个异或过，另一个没有异或过，那么操作后，异或 $k$ 的元素加一减一，不变。

结合这两个观察，问题变成：

- 选择 $\textit{nums}$ 中的**偶数**个元素，把这些数都异或 $k$，数组的最大元素和是多少？

这可以用状态机 DP 解决。

- 定义 $f[i][0]$ 表示选择 $\textit{nums}$ 的前 $i$ 数中的偶数个元素异或 $k$，得到的最大元素和。
- 定义 $f[i][1]$ 表示选择 $\textit{nums}$ 的前 $i$ 数中的奇数个元素异或 $k$，得到的最大元素和。

设 $x=\textit{nums}[i]$。

- 情况一，不操作 $x$：
  - $f[i+1][0] = f[i][0] + x$。
  - $f[i+1][1] = f[i][1] + x$。
- 情况二，操作 $x$：
  - $f[i+1][0] = f[i][1] + (x\oplus k)$。
  - $f[i+1][1] = f[i][0] + (x\oplus k)$。

两种情况取最大值，有

$$
\begin{aligned}
&f[i+1][0] = \max(f[i][0] + x, f[i][1] + (x\oplus k))\\&f[i+1][1] = \max(f[i][1] + x, f[i][0] + (x\oplus k))
\end{aligned}

$$

初始值 $f[0][0] = 0,\ f[0][1] = -\infty$。 答案为 $f[n][0]$。代码实现时，$f$ 数组的第一个维度可以压缩掉。

```go [sol]
func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	f0, f1 := 0, math.MinInt32
	for _, v := range nums {
		f0, f1 = max(v^k + f1, f0+v), max(v^k+f0, f1+v)
	}

	return int64(f0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
