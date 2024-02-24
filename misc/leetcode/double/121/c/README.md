### 题目

<p>给你两个正整数 <code>x</code> 和 <code>y</code> 。</p>

<p>一次操作中，你可以执行以下四种操作之一：</p>

<ol>
	<li>如果 <code>x</code> 是 <code>11</code> 的倍数，将 <code>x</code> 除以 <code>11</code> 。</li>
	<li>如果 <code>x</code> 是 <code>5</code> 的倍数，将 <code>x</code> 除以 <code>5</code> 。</li>
	<li>将 <code>x</code> 减 <code>1</code> 。</li>
	<li>将 <code>x</code> 加 <code>1</code> 。</li>
</ol>

<p>请你返回让 <code>x</code> 和 <code>y</code> 相等的 <strong>最少</strong> 操作次数。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>x = 26, y = 1
<b>输出：</b>3
<b>解释</b><strong>：</strong>我们可以通过以下操作将 26 变为 1 ：
1. 将 x 减 1
2. 将 x 除以 5
3. 将 x 除以 5
将 26 变为 1 最少需要 3 次操作。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>x = 54, y = 2
<b>输出：</b>4
<b>解释：</b>我们可以通过以下操作将 54 变为 2 ：
1. 将 x 加 1
2. 将 x 除以 11
3. 将 x 除以 5
4. 将 x 加 1
将 54 变为 2 最少需要 4 次操作。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>x = 25, y = 30
<b>输出：</b>5
<b>解释：</b>我们可以通过以下操作将 25 变为 30 ：
1. 将 x 加 1
2. 将 x 加 1
3. 将 x 加 1
4. 将 x 加 1
5. 将 x 加 1
将 25 变为 30 最少需要 5 次操作。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= x, y <= 10<sup>4</sup></code></li>
</ul>

### 思路

每个操作都可以理解成：从 $x$ 向操作后的数连边。
在这张图上跑 BFS，求出从 $x$ 到 $y$ 的最短路，即为答案。
注意，如果 $x<y$ 那么只能用加一操作，此时可以直接算出操作次数。

```go [sol]
func minimumOperationsToMakeEqual(x, y int) int {
	if x <= y {
		return y - x
	}

	q, vis := []int{x}, make(map[int]bool)
	vis[x] = true
	for step := 0; ; step++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == y {
				return step
			}
			if v%11 == 0 && !vis[v/11]{
				q = append(q, v/11)
				vis[v/11] = true
			}
			if v%5 == 0 && !vis[v/5]{
				q = append(q, v/5)
				vis[v/5] = true
			}
			if !vis[v+1] {
				q = append(q, v+1)
				vis[v+1] = true
			}
			if !vis[v-1] {
				q = append(q, v-1)
				vis[v-1] = true
			}
		}
	}
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{x}$ 。
- 空间复杂度：$\mathcal{O}(n)$。
