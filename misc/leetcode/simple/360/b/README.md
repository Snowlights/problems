#### 题目

<p>给你两个正整数：<code>n</code> 和 <code>target</code> 。</p>

<p>如果数组 <code>nums</code> 满足下述条件，则称其为 <strong>美丽数组</strong> 。</p>

<ul>
	<li><code>nums.length == n</code>.</li>
	<li><code>nums</code> 由两两互不相同的正整数组成。</li>
	<li>在范围 <code>[0, n-1]</code> 内，<strong>不存在 </strong>两个 <strong>不同</strong> 下标 <code>i</code> 和 <code>j</code> ，使得 <code>nums[i] + nums[j] == target</code> 。</li>
</ul>

<p>返回符合条件的美丽数组所可能具备的 <strong>最小</strong> 和。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>n = 2, target = 3
<strong>输出：</strong>4
<strong>解释：</strong>nums = [1,3] 是美丽数组。
- nums 的长度为 n = 2 。
- nums 由两两互不相同的正整数组成。
- 不存在两个不同下标 i 和 j ，使得 nums[i] + nums[j] == 3 。
可以证明 4 是符合条件的美丽数组所可能具备的最小和。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>n = 3, target = 3
<strong>输出：</strong>8
<strong>解释：</strong>
nums = [1,3,4] 是美丽数组。 
- nums 的长度为 n = 3 。 
- nums 由两两互不相同的正整数组成。 
- 不存在两个不同下标 i 和 j ，使得 nums[i] + nums[j] == 3 。
可以证明 8 是符合条件的美丽数组所可能具备的最小和。</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>n = 1, target = 1
<strong>输出：</strong>1
<strong>解释：</strong>nums = [1] 是美丽数组。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n <= 10<sup>5</sup></code></li>
	<li><code>1 <= target <= 10<sup>5</sup></code></li>
</ul>

#### 思路

把 $\textit{target}$ 记作 $k$。对于 $[1,k-1]$ 内的数字：

- $1$ 和 $k-1$ 只能选其中一个；
- $2$ 和 $k-2$ 只能选其中一个；
- $3$ 和 $k-3$ 只能选其中一个；
- ……
- 一直到 $\left\lfloor\dfrac{k}{2}\right\rfloor$，无论 $k$ 是奇数还是偶数，它都可以选。

设 $m=\min\left(\left\lfloor\dfrac{k}{2}\right\rfloor, n\right)$，那么答案的第一段是从 $1$ 到 $m$，元素和为

$$
\dfrac{m(m+1)}{2}

$$

此时还剩下 $n-m$ 个数，只能从 $k$ 开始往后选，那么答案的第二段是从 $k$ 到 $k+n-m-1$，元素和为

$$
\dfrac{(2k+n-m-1)(n-m)}{2}

$$

所以答案为

$$
\dfrac{m(m+1) + (2k+n-m-1)(n-m)}{2}

$$

```go  
func minimumPossibleSum(n, k int) int64 {
	m := min(k/2, n)
	return int64((m*(m+1) + (k*2+n-m-1)*(n-m)) / 2)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$ 。
