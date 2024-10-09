#### 题目

<p>给你一个下标从 <strong>1</strong> 开始、长度为 <code>n</code> 的整数数组 <code>nums</code> 。</p>

<p>现定义函数 <code>greaterCount</code> ，使得 <code>greaterCount(arr, val)</code> 返回数组 <code>arr</code> 中<strong> 严格大于</strong> <code>val</code> 的元素数量。</p>

<p>你需要使用 <code>n</code> 次操作，将 <code>nums</code> 的所有元素分配到两个数组 <code>arr1</code> 和 <code>arr2</code> 中。在第一次操作中，将 <code>nums[1]</code> 追加到 <code>arr1</code> 。在第二次操作中，将 <code>nums[2]</code> 追加到 <code>arr2</code> 。之后，在第 <code>i</code> 次操作中：</p>

<ul>
	<li>如果 <code>greaterCount(arr1, nums[i]) > greaterCount(arr2, nums[i])</code> ，将 <code>nums[i]</code> 追加到 <code>arr1</code> 。</li>
	<li>如果 <code>greaterCount(arr1, nums[i]) < greaterCount(arr2, nums[i])</code> ，将 <code>nums[i]</code> 追加到 <code>arr2</code> 。</li>
	<li>如果 <code>greaterCount(arr1, nums[i]) == greaterCount(arr2, nums[i])</code> ，将 <code>nums[i]</code> 追加到元素数量较少的数组中。</li>
	<li>如果仍然相等，那么将 <code>nums[i]</code> 追加到 <code>arr1</code> 。</li>
</ul>

<p>连接数组 <code>arr1</code> 和 <code>arr2</code> 形成数组 <code>result</code> 。例如，如果 <code>arr1 == [1,2,3]</code> 且 <code>arr2 == [4,5,6]</code> ，那么 <code>result = [1,2,3,4,5,6]</code> 。</p>

<p>返回整数数组 <code>result</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,1,3,3]
<strong>输出：</strong>[2,3,1,3]
<strong>解释：</strong>在前两次操作后，arr1 = [2] ，arr2 = [1] 。
在第 3 次操作中，两个数组中大于 3 的元素数量都是零，并且长度相等，因此，将 nums[3] 追加到 arr1 。
在第 4 次操作中，两个数组中大于 3 的元素数量都是零，但 arr2 的长度较小，因此，将 nums[4] 追加到 arr2 。
在 4 次操作后，arr1 = [2,3] ，arr2 = [1,3] 。
因此，连接形成的数组 result 是 [2,3,1,3] 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [5,14,3,1,2]
<strong>输出：</strong>[5,3,1,2,14]
<strong>解释：</strong>在前两次操作后，arr1 = [5] ，arr2 = [14] 。
在第 3 次操作中，两个数组中大于 3 的元素数量都是一，并且长度相等，因此，将 nums[3] 追加到 arr1 。
在第 4 次操作中，arr1 中大于 1 的元素数量大于 arr2 中的数量（2 > 1），因此，将 nums[4] 追加到 arr1 。
在第 5 次操作中，arr1 中大于 2 的元素数量大于 arr2 中的数量（2 > 1），因此，将 nums[5] 追加到 arr1 。
在 5 次操作后，arr1 = [5,3,1,2] ，arr2 = [14] 。
因此，连接形成的数组 result 是 [5,3,1,2,14] 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [3,3,3,3]
<strong>输出：</strong>[3,3,3,3]
<strong>解释：</strong>在 4 次操作后，arr1 = [3,3] ，arr2 = [3,3] 。
因此，连接形成的数组 result 是 [3,3,3,3] 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 <= n <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

#### 思路

将元素**离散化**成 $[1,m]$ 中的元素，其中 $m$ 为 $\textit{nums}$ 中的不同元素个数。  
这可以对 $\textit{nums}$ 排序去重后，在数组中二分查找得到。  
记 $\textit{arr}_1$ 为 $a$，记 $\textit{arr}_2$ 为 $b$。用两棵树状数组分别维护 $a$ 和 $b$ 中的每个元素的出现次数，即可快速计算 $\texttt{greaterCount}$。然后按照题目要求模拟即可。

```go [sol]
func resultArray(nums []int) []int {
	a, h := make([]int, 0, len(nums)), make(map[int]bool)
	for _, v := range nums {
		if h[v] {
			continue
		}
		a = append(a, v)
		h[v] = true
	}
	sort.Ints(a)
	n := len(a)
	at, bt := newFenwickTree(n, nil), newFenwickTree(n, nil)
	r1, r2 := []int{nums[0]}, []int{nums[1]}
	at.add(sort.SearchInts(a, nums[0])+1, 1)
	bt.add(sort.SearchInts(a, nums[1])+1, 1)
	for _, v := range nums[2:] {
		idx := sort.SearchInts(a, v) + 1
		v1, v2 := len(r1)-at.query(1, idx), len(r2)-bt.query(1, idx)
		if v1 > v2 || v1 == v2 && len(r1) <= len(r2) {
			r1 = append(r1, v)
			at.add(idx, 1)
		} else {
			r2 = append(r2, v)
			bt.add(idx, 1)
		}
	}
	return append(r1, r2...)
}

type fenwick struct {
	a    []int // 原始数据，单点查询+区间更新传入
	tree []int // 树状数组
	diff []int // 差分数组, 用于区间加、区间求和
}

func newFenwickTree(n int, a []int) fenwick {
	return fenwick{a, make([]int, n+1), make([]int, n+1)}
}

// 位置 i 增加 val
// 1<=i<=n
func (f fenwick) add(i, val int) {
	val1 := i * val
	for ; i < int(len(f.tree)); i += i & -i {
		f.tree[i] += val
		f.diff[i] += val1
	}
}

// 求前缀和 [0,i]
// 0<=i<=n
func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

func (f fenwick) sumDiff(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.diff[i]
	}
	return
}

// 求区间和 [l,r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int) int {
	return f.sum(r) - f.sum(l-1)
}

func (f fenwick) queryDiff(l, r int) int {
	return (r+1)*f.sum(r) - l*f.sum(l-1) - (f.sumDiff(r) - f.sumDiff(l-1))
}

// 差分树状数组，可用于区间更新+单点查询 queryOne(i) = a[i] + sum(i) // a 从 1 开始
// r+1 即使超过 n 也没关系，因为不会用到
// i >= 1
func (f fenwick) queryOne(i int) int { return f.a[i] + f.sum(i) }

// [l,r]
func (f fenwick) addRange(l, r, val int) { f.add(l, val); f.add(r+1, -val) }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
