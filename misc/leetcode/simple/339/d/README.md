#### 题目  

<p>给你一个整数 <code>n</code> 和一个在范围 <code>[0, n - 1]</code> 以内的整数 <code>p</code> ，它们表示一个长度为 <code>n</code> 且下标从 <strong>0</strong> 开始的数组 <code>arr</code> ，数组中除了下标为 <code>p</code> 处是 <code>1</code> 以外，其他所有数都是 <code>0</code> 。</p>

<p>同时给你一个整数数组 <code>banned</code> ，它包含数组中的一些位置。<code>banned</code> 中第 <strong>i</strong> 个位置表示 <code>arr[banned[i]] = 0</code> ，题目保证 <code>banned[i] != p</code> 。</p>

<p>你可以对 <code>arr</code> 进行 <strong>若干次</strong> 操作。一次操作中，你选择大小为 <code>k</code> 的一个 <strong>子数组</strong> ，并将它 <b>翻转</b> 。在任何一次翻转操作后，你都需要确保 <code>arr</code> 中唯一的 <code>1</code> 不会到达任何 <code>banned</code> 中的位置。换句话说，<code>arr[banned[i]]</code> 始终 <strong>保持</strong> <code>0</code> 。</p>

<p>请你返回一个数组 <code>ans</code> ，对于<em> </em><code>[0, n - 1]</code> 之间的任意下标 <code>i</code> ，<code>ans[i]</code> 是将 <code>1</code> 放到位置 <code>i</code> 处的 <strong>最少</strong> 翻转操作次数，如果无法放到位置 <code>i</code> 处，此数为 <code>-1</code> 。</p>

<ul>
	<li><strong>子数组</strong> 指的是一个数组里一段连续 <strong>非空</strong> 的元素序列。</li>
	<li>对于所有的 <code>i</code> ，<code>ans[i]</code> 相互之间独立计算。</li>
	<li>将一个数组中的元素 <strong>翻转</strong> 指的是将数组中的值变成 <strong>相反顺序</strong> 。</li>
</ul>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>n = 4, p = 0, banned = [1,2], k = 4
<b>输出：</b>[0,-1,-1,1]
<b>解释：</b><code>k = 4，所以只有一种可行的翻转操作，就是将整个数组翻转。一开始 </code>1<strong> </strong>在位置 0 处，所以将它翻转到位置 0 处需要的操作数为 0 。
我们不能将 1 翻转到 banned 中的位置，所以位置 1 和 2 处的答案都是 -1 。
通过一次翻转操作，可以将 1 放到位置 3 处，所以位置 3 的答案是 1 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>n = 5, p = 0, banned = [2,4], k = 3
<b>输出：</b>[0,-1,-1,-1,-1]
<b>解释：</b>这个例子中 1 一开始在位置 0 处，所以此下标的答案为 0 。
翻转的子数组长度为 k = 3 ，1 此时在位置 0 处，所以我们可以翻转子数组 [0, 2]，但翻转后的下标 2 在 banned 中，所以不能执行此操作。
由于 1 没法离开位置 0 ，所以其他位置的答案都是 -1 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>n = 4, p = 2, banned = [0,1,3], k = 1
<b>输出：</b>[-1,-1,0,-1]
<b>解释：</b>这个例子中，我们只能对长度为 1 的子数组执行翻转操作，所以 1 无法离开初始位置。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= p &lt;= n - 1</code></li>
	<li><code>0 &lt;= banned.length &lt;= n - 1</code></li>
	<li><code>0 &lt;= banned[i] &lt;= n - 1</code></li>
	<li><code>1 &lt;= k &lt;= n </code></li>
	<li><code>banned[i] != p</code></li>
	<li><code>banned</code> 中的值 <strong>互不相同</strong> 。</li>
</ul>
 
#### 思路  

#### 提示 1

对于子数组 $[L,R]$ 中的任意下标 $i$，翻转后的下标是 $L+R-i$（中心对称翻转，两个下标相加恒等于 $L+R$）。  
那么：
- 当子数组向右滑动时，$L$ 和 $R$ 都增加 $1$，所以翻转后的下标会增加 $2$
- 当子数组向左滑动时，$L$ 和 $R$ 都减少 $1$，所以翻转后的下标会减少 $2$

因此，**$i$ 翻转后的所有位置组成了一个公差为 $2$ 的等差数列**（不考虑 $\textit{banned}$）。  
如何求出这些位置的范围呢？注意当 $i$ 在数组边界 $0$ 或 $n-1$ 附近时，有些位置是无法翻转到的。

#### 提示 2

- 如果不考虑数组的边界，那么范围是 $[i-k+1, i+k-1]$。
- 如果 $i$ 在数组左边界 $0$ 附近，那么翻转时会受到数组左边界的约束，当子数组在最左边时，$L=0,R=k-1$，$i$ 翻转后是 $0+(k-1)-i=k-i-1$，所以小于 $k-i-1$ 的点是无法翻转到的；
- 如果 $i$ 在数组右边界 $n-1$ 附近，那么翻转时会受到数组右边界的约束，当子数组在最右边时，$L=n-k,R=n-1$，$i$ 翻转后是 $(n-k)+(n-1) - i=2n - k - i - 1$，所以大于 $2n - k - i - 1$ 的点是无法翻转到的。

所以实际范围为

$$
[\max(i-k+1,k-i-1), \min(i+k-1,2n - k - i - 1)]
$$

#### 提示 3

用两棵平衡树分别维护不等于 $p$ 也不在 $\textit{banned}$ 中的偶数下标和奇数下标。  
然后用 BFS **模拟**。  
在对应的平衡树上，一边遍历翻转后的所有位置，一边把平衡树上的下标删除，加到队列中。这样可以避免重复访问已经访问过的节点。

并查集的思路是，如果要删除一个元素，那么把它的下标 $j$ 和 $j+1$ 合并，这样后面删除的时候就会自动跳过已删除的元素。

```go 
type uf struct {
	fa []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}

func (u *uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	u.fa[x] = y
}

func minReverseOperations(n, p int, banned []int, k int) []int {
	ban := map[int]bool{p: true}
	for _, v := range banned {
		ban[v] = true
	}
	notBanned := [2][]int{}
	for i := 0; i < n; i++ {
		if !ban[i] {
			notBanned[i%2] = append(notBanned[i%2], i)
		}
	}
	notBanned[0] = append(notBanned[0], n)
	notBanned[1] = append(notBanned[1], n) // 哨兵
	ufs := [2]uf{newUnionFind(len(notBanned[0])), newUnionFind(len(notBanned[1]))}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	q := []int{p}
	for step := 0; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			ans[i] = step
			// 从 mn 到 mx 的所有位置都可以翻转到
			mn := max(i-k+1, k-i-1)
			mx := min(i+k-1, n*2-k-i-1)
			a, u := notBanned[mn%2], ufs[mn%2]
			for j := u.find(sort.SearchInts(a, mn)); a[j] <= mx; j = u.find(j + 1) {
				q = append(q, a[j])
				u.merge(j, j+1) // 删除 j
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

#### 复杂度分析  

- 时间复杂度：$O(n\log n)$。
- 空间复杂度：$O(n)$。