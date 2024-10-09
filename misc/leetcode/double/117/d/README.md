### 题目

<p>给你一个下标从 <strong>0</strong> 开始大小为 <code>m * n</code> 的整数矩阵 <code>values</code> ，表示 <code>m</code> 个不同商店里 <code>m * n</code> 件不同的物品。每个商店有 <code>n</code> 件物品，第 <code>i</code> 个商店的第 <code>j</code> 件物品的价值为 <code>values[i][j]</code> 。除此以外，第 <code>i</code> 个商店的物品已经按照价值非递增排好序了，也就是说对于所有 <code>0 <= j < n - 1</code> 都有 <code>values[i][j] >= values[i][j + 1]</code> 。</p>

<p>每一天，你可以在一个商店里购买一件物品。具体来说，在第 <code>d</code> 天，你可以：</p>

<ul>
	<li>选择商店 <code>i</code> 。</li>
	<li>购买数组中最右边的物品 <code>j</code> ，开销为 <code>values[i][j] * d</code> 。换句话说，选择该商店中还没购买过的物品中最大的下标 <code>j</code> ，并且花费 <code>values[i][j] * d</code> 去购买。</li>
</ul>

<p><strong>注意</strong>，所有物品都视为不同的物品。比方说如果你已经从商店 <code>1</code> 购买了物品 <code>0</code> ，你还可以在别的商店里购买其他商店的物品 <code>0</code> 。</p>

<p>请你返回购买所有 <code>m * n</code> 件物品需要的 <strong>最大开销</strong> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>values = [[8,5,2],[6,4,1],[9,7,3]]
<b>输出：</b>285
<b>解释：</b>第一天，从商店 1 购买物品 2 ，开销为 values[1][2] * 1 = 1 。
第二天，从商店 0 购买物品 2 ，开销为 values[0][2] * 2 = 4 。
第三天，从商店 2 购买物品 2 ，开销为 values[2][2] * 3 = 9 。
第四天，从商店 1 购买物品 1 ，开销为 values[1][1] * 4 = 16 。
第五天，从商店 0 购买物品 1 ，开销为 values[0][1] * 5 = 25 。
第六天，从商店 1 购买物品 0 ，开销为 values[1][0] * 6 = 36 。
第七天，从商店 2 购买物品 1 ，开销为 values[2][1] * 7 = 49 。
第八天，从商店 0 购买物品 0 ，开销为 values[0][0] * 8 = 64 。
第九天，从商店 2 购买物品 0 ，开销为 values[2][0] * 9 = 81 。
所以总开销为 285 。
285 是购买所有 m * n 件物品的最大总开销。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>values = [[10,8,6,4,2],[9,7,5,3,2]]
<b>输出：</b>386
<b>解释：</b>第一天，从商店 0 购买物品 4 ，开销为 values[0][4] * 1 = 2 。
第二天，从商店 1 购买物品 4 ，开销为 values[1][4] * 2 = 4 。
第三天，从商店 1 购买物品 3 ，开销为 values[1][3] * 3 = 9 。
第四天，从商店 0 购买物品 3 ，开销为 values[0][3] * 4 = 16 。
第五天，从商店 1 购买物品 2 ，开销为 values[1][2] * 5 = 25 。
第六天，从商店 0 购买物品 2 ，开销为 values[0][2] * 6 = 36 。
第七天，从商店 1 购买物品 1 ，开销为 values[1][1] * 7 = 49 。
第八天，从商店 0 购买物品 1 ，开销为 values[0][1] * 8 = 64 。
第九天，从商店 1 购买物品 0 ，开销为 values[1][0] * 9 = 81 。
第十天，从商店 0 购买物品 0 ，开销为 values[0][0] * 10 = 100 。
所以总开销为 386 。
386 是购买所有 m * n 件物品的最大总开销。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= m == values.length <= 10</code></li>
	<li><code>1 <= n == values[i].length <= 10<sup>4</sup></code></li>
	<li><code>1 <= values[i][j] <= 10<sup>6</sup></code></li>
	<li><code>values[i]</code> 按照非递增顺序排序。</li>
</ul>

### 思路

## 提示 1

根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728)，$\textit{values}[i][j]$ 越小的数，应该越早购买。

## 提示 2

$\textit{values}$ 的最后一列一定包含最小的数，因为其余列的数都不会比最后一列的这 $m$ 个数小。

次小的数呢？

## 提示 3

考虑每行当前最后一个数，这 $m$ 个数中一定包含次小的数，因为和上面一样，其余数字都不会比这 $m$ 个数小。

## 提示 4

根据数学归纳法，我们可以按照 $\textit{values}[i][j]$ 从小到大的顺序取到所有元素。

所以，把所有数合并在一起排序即可。

也可以用最小堆实现，那样空间更小。

```go  
func maxSpending(values [][]int) int64 {
	m, n := len(values), len(values[0])
	a := make([]int, 0, m*n)
	for _, row := range values {
		a = append(a, row...)
	}
	sort.Ints(a)
	ans := 0
	for i, x := range a {
		ans += x * (i + 1)
	}
	return int64(ans)
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log (mn))$，其中 $m$ 和 $n$ 分别为 $\textit{values}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。
