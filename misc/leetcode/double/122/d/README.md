### 题目

<p>给你一个下标从 <strong>0</strong> 开始长度为 <code>n</code> 的整数数组 <code>nums</code> 和两个 <strong>正</strong> 整数 <code>k</code> 和 <code>dist</code> 。</p>

<p>一个数组的 <strong>代价</strong> 是数组中的 <strong>第一个</strong> 元素。比方说，<code>[1,2,3]</code> 的代价为 <code>1</code> ，<code>[3,4,1]</code> 的代价为 <code>3</code> 。</p>

<p>你需要将 <code>nums</code> 分割成 <code>k</code> 个 <strong>连续且互不相交</strong> 的子数组，满足 <strong>第二</strong> 个子数组与第 <code>k</code> 个子数组中第一个元素的下标距离 <strong>不超过</strong> <code>dist</code> 。换句话说，如果你将 <code>nums</code> 分割成子数组 <code>nums[0..(i<sub>1</sub> - 1)], nums[i<sub>1</sub>..(i<sub>2</sub> - 1)], ..., nums[i<sub>k-1</sub>..(n - 1)]</code> ，那么它需要满足 <code>i<sub>k-1</sub> - i<sub>1</sub> <= dist</code> 。</p>

<p>请你返回这些子数组的 <strong>最小</strong> 总代价。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,3,2,6,4,2], k = 3, dist = 3
<b>输出：</b>5
<b>解释：</b>将数组分割成 3 个子数组的最优方案是：[1,3] ，[2,6,4] 和 [2] 。这是一个合法分割，因为 i<sub>k-1</sub> - i<sub>1</sub> 等于 5 - 2 = 3 ，等于 dist 。总代价为 nums[0] + nums[2] + nums[5] ，也就是 1 + 2 + 2 = 5 。
5 是分割成 3 个子数组的最小总代价。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [10,1,2,2,2,1], k = 4, dist = 3
<b>输出：</b>15
<b>解释：</b>将数组分割成 4 个子数组的最优方案是：[10] ，[1] ，[2] 和 [2,2,1] 。这是一个合法分割，因为 i<sub>k-1</sub> - i<sub>1</sub> 等于 3 - 1 = 2 ，小于 dist 。总代价为 nums[0] + nums[1] + nums[2] + nums[3] ，也就是 10 + 1 + 2 + 2 = 15 。
分割 [10] ，[1] ，[2,2,2] 和 [1] 不是一个合法分割，因为 i<sub>k-1</sub> 和 i<sub>1</sub> 的差为 5 - 1 = 4 ，大于 dist 。
15 是分割成 4 个子数组的最小总代价。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [10,8,18,9], k = 3, dist = 1
<b>输出：</b>36
<b>解释：</b>将数组分割成 4 个子数组的最优方案是：[10] ，[8] 和 [18,9] 。这是一个合法分割，因为 i<sub>k-1</sub> - i<sub>1</sub> 等于 2 - 1 = 1 ，等于 dist 。总代价为 nums[0] + nums[1] + nums[2] ，也就是 10 + 8 + 18 = 36 。
分割 [10] ，[8,18] 和 [9] 不是一个合法分割，因为 i<sub>k-1</sub> 和 i<sub>1</sub> 的差为 3 - 1 = 2 ，大于 dist 。
36 是分割成 3 个子数组的最小总代价。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 <= n <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>3 <= k <= n</code></li>
	<li><code>k - 2 <= dist <= n - 2</code></li>
</ul>

### 思路

第一段的第一个数是确定的，即 $\textit{nums}[0]$。
如果知道了第二段的第一个数的位置（记作 $p$），第三段的第一个数的位置，……，第 $k$ 段的第一个数的位置（记作 $q$），那么这个划分方案也就确定了。
考虑到 $q-p \le \textit{dist}$，本题相当于在一个大小固定为 $\textit{dist}+1$ 的滑动窗口内，求前 $k-1$ 小的元素和。仿照 [480. 滑动窗口中位数](https://leetcode.cn/problems/sliding-window-median/)，这可以用两个有序集合来做：

1. 初始化两个有序集合 $L$ 和 $R$。注意：为方便计算，把 $k$ 减一。
2. 把 $\textit{nums}[1]$ 到 $\textit{nums}[\textit{dist}+1]$ 加到 $L$ 中。
3. 保留 $L$ 最小的 $k$ 个数，把其余数丢到 $R$ 中。
4. 从 $i=\textit{dist}+2$ 开始滑窗。
5. 先把 $\textit{out} = \textit{nums}[i-\textit{dist}-1]$ 移出窗口：如果 $\textit{out}$ 在 $L$ 中，就从 $L$ 中移除，否则从 $R$ 中移除。
6. 然后把 $\textit{in} = \textit{nums}[i]$ 移入窗口：如果 $\textit{in}$ 小于 $L$ 中的最大元素，则加入 $L$，否则加入 $R$。
7. 上面两步做完后，如果 $L$ 中的元素个数小于 $k$（等于 $k-1$），则从 $R$ 中取一个最小元素加入 $L$；反之，如果 $L$ 中的元素个数大于 $k$（等于 $k+1$），则从 $L$ 中取一个最大元素加入 $R$。

上述过程维护 $L$ 中元素之和 $\textit{sumL}$，取 $\textit{sumL}$ 的最小值，即为答案。

```go [sol]
func minimumCost(nums []int, k, dist int) int64 {
	k--
	L := redblacktree.NewWithIntComparator()
	R := redblacktree.NewWithIntComparator()
	add := func(t *redblacktree.Tree, x int) {
		c, ok := t.Get(x)
		if ok {
			t.Put(x, c.(int)+1)
		} else {
			t.Put(x, 1)
		}
	}
	del := func(t *redblacktree.Tree, x int) {
		c, _ := t.Get(x)
		if c.(int) > 1 {
			t.Put(x, c.(int)-1)
		} else {
			t.Remove(x)
		}
	}

	sumL := nums[0]
	for _, x := range nums[1 : dist+2] {
		sumL += x
		add(L, x)
	}
	sizeL := dist + 1

	l2r := func() {
		x := L.Right().Key.(int)
		sumL -= x
		sizeL--
		del(L, x)
		add(R, x)
	}
	r2l := func() {
		x := R.Left().Key.(int)
		sumL += x
		sizeL++
		del(R, x)
		add(L, x)
	}
	for sizeL > k {
		l2r()
	}

	ans := sumL
	for i := dist + 2; i < len(nums); i++ {
		// 移除 out
		out := nums[i-dist-1]
		if _, ok := L.Get(out); ok {
			sumL -= out
			sizeL--
			del(L, out)
		} else {
			del(R, out)
		}

		// 添加 in
		in := nums[i]
		if in < L.Right().Key.(int) {
			sumL += in
			sizeL++
			add(L, in)
		} else {
			add(R, in)
		}

		// 维护大小
		if sizeL == k-1 {
			r2l()
		} else if sizeL == k+1 {
			l2r()
		}

		ans = min(ans, sumL)
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log \textit{dist})$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\textit{dist})$。
