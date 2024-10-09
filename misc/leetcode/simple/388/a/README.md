#### 题目

<p>给你一个长度为 <code>n</code> 的数组 <code>apple</code> 和另一个长度为 <code>m</code> 的数组 <code>capacity</code> 。</p>

<p>一共有 <code>n</code> 个包裹，其中第 <code>i</code> 个包裹中装着 <code>apple[i]</code> 个苹果。同时，还有 <code>m</code> 个箱子，第 <code>i</code> 个箱子的容量为 <code>capacity[i]</code> 个苹果。</p>

<p>请你选择一些箱子来将这 <code>n</code> 个包裹中的苹果重新分装到箱子中，返回你需要选择的箱子的<strong> 最小</strong> 数量。</p>

<p><strong>注意</strong>，同一个包裹中的苹果可以分装到不同的箱子中。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>apple = [1,3,2], capacity = [4,3,1,5,2]
<strong>输出：</strong>2
<strong>解释：</strong>使用容量为 4 和 5 的箱子。
总容量大于或等于苹果的总数，所以可以完成重新分装。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>apple = [5,5,5], capacity = [2,4,2,7]
<strong>输出：</strong>4
<strong>解释：</strong>需要使用所有箱子。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n == apple.length <= 50</code></li>
	<li><code>1 <= m == capacity.length <= 50</code></li>
	<li><code>1 <= apple[i], capacity[i] <= 50</code></li>
	<li>输入数据保证可以将包裹中的苹果重新分装到箱子中。</li>
</ul>

#### 思路

题目说，同一个包裹中的苹果可以分装到不同的箱子中。
那么按照容量从大到小选择箱子装苹果，直到所有苹果均装入箱子为止。
注意题目保证可以将包裹中的苹果重新分装到箱子中。

```go [sol]
func minimumBoxes(apple []int, capacity []int) int {
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})
	sum, ans := 0, 0
	for _, v := range apple {
		sum += v
	}
	for _, v := range capacity {
		if sum >= v {
			ans++
			sum -= v
		} else {
			break
		}
	}
	if sum > 0 {
		ans++
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\log m)$，其中 $n$ 为 $\textit{apple}$ 的长度，$m$ 为 $\textit{capacity}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。
