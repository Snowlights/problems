### 题目

<p>给你一个数组 <code>nums</code> ，它包含若干正整数。</p>

<p>一开始分数 <code>score = 0</code> ，请你按照下面算法求出最后分数：</p>

<ul>
	<li>从数组中选择最小且没有被标记的整数。如果有相等元素，选择下标最小的一个。</li>
	<li>将选中的整数加到 <code>score</code> 中。</li>
	<li>标记 <strong>被选中元素</strong>，如果有相邻元素，则同时标记 <strong>与它相邻的两个元素</strong> 。</li>
	<li>重复此过程直到数组中所有元素都被标记。</li>
</ul>

<p>请你返回执行上述算法后最后的分数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [2,1,3,4,5,2]
<b>输出：</b>7
<b>解释：</b>我们按照如下步骤标记元素：
- 1 是最小未标记元素，所以标记它和相邻两个元素：[<em><strong>2</strong></em>,<em><strong>1</strong></em>,<em><strong>3</strong></em>,4,5,2] 。
- 2 是最小未标记元素，所以标记它和左边相邻元素：[<em><strong>2</strong></em>,<em><strong>1</strong></em>,<em><strong>3</strong></em>,4,<em><strong>5</strong></em>,<em><strong>2</strong></em>] 。
- 4 是仅剩唯一未标记的元素，所以我们标记它：[<em><strong>2</strong></em>,<em><strong>1</strong></em>,<em><strong>3</strong></em>,<em><strong>4</strong></em>,<em><strong>5</strong></em>,<em><strong>2</strong></em>] 。
总得分为 1 + 2 + 4 = 7 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [2,3,5,1,3,2]
<b>输出：</b>5
<b>解释：</b>我们按照如下步骤标记元素：
- 1 是最小未标记元素，所以标记它和相邻两个元素：[2,3,<em><strong>5</strong></em>,<em><strong>1</strong></em>,<em><strong>3</strong></em>,2] 。
- 2 是最小未标记元素，由于有两个 2 ，我们选择最左边的一个 2 ，也就是下标为 0 处的 2 ，以及它右边相邻的元素：[<em><strong>2</strong></em>,<em><strong>3</strong></em>,<em><strong>5</strong></em>,<em><strong>1</strong></em>,<em><strong>3</strong></em>,2] 。
- 2 是仅剩唯一未标记的元素，所以我们标记它：[<em><strong>2</strong></em>,<em><strong>3</strong></em>,<em><strong>5</strong></em>,<em><strong>1</strong></em>,<em><strong>3</strong></em>,<em><strong>2</strong></em>] 。
总得分为 1 + 2 + 2 = 5 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>6</sup></code></li>
</ul>

### 思路

# 方法一：带着下标去排序

技巧题，把 $\textit{nums}[i]$ 及其下标绑定后，按照元素值从小到大排序，元素值相同的按照下标排序。  
然后按照题目模拟，用一个 $\textit{vis}$ 数组来实现标记。  
也可以生成一个下标数组，对下标排序。

```go  
func findScore(nums []int) (ans int64) {
	type pair struct{ v, i int }
	a := make([]pair, len(nums))
	for i, x := range nums {
		a[i] = pair{x, i + 1} // +1 保证下面 for 循环下标不越界
	}
	sort.Slice(a, func(i, j int) bool {
		a, b := a[i], a[j]
		return a.v < b.v || a.v == b.v && a.i < b.i
	})
	vis := make([]bool, len(nums)+2) // 保证下标不越界
	for _, p := range a {
		if !vis[p.i] {
			vis[p.i-1] = true
			vis[p.i+1] = true // 标记相邻的两个元素
			ans += int64(p.v)
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

# 方法二：转换 + 分组循环

把 $\textit{nums}$ 视作由若干严格递减子段组成的数组。  
例如示例 1 可以看成 $[2,1]+[3]+[4]+[5,2]$， 示例 2 可以看成 $[2]+[3]+[5,1]+[3,2]$。
从左到右遍历 $\textit{nums}$，严格递减子段的最小值 $\textit{nums}[i]$ 一定可以选，
因为它比 $\textit{nums}[i-1]$ 小，且不会超过 $\textit{nums}[i+1]$。
如果等于 $\textit{nums}[i+1]$，由于我们是从左到右遍历的，下标也是最小的。  
$\textit{nums}[i]$ 选了，那么这一段左侧的 $\textit{nums}[i-2],\textit{nums}[i-4],\cdots$ 就可以一起选了（因为不能选相邻的），
且 $\textit{nums}[i+1]$ 不能选。  
于是遍历 $\textit{nums}$ 就可以算出答案了。

```go  
func findScore(nums []int) (ans int64) {
	for i, n := 0, len(nums); i < n; i += 2 { // i 选了 i+1 不能选
		i0 := i
		for i+1 < n && nums[i] > nums[i+1] { // 找到下坡的坡底
			i++
		}
		for j := i; j >= i0; j -= 2 { // 从坡底 i 到坡顶 i0，每隔一个累加
			ans += int64(nums[j])
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n + U\log(mc^2))$，其中 $n$ 为 $\textit{ranks}$ 的长度，$U=\max(\textit{ranks})$，$m=\min(\textit{ranks})$，$c=\textit{cars}$。
- 空间复杂度：$O(U)$。
