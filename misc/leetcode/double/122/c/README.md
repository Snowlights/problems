### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> ，它只包含 <strong>正</strong> 整数。</p>

<p>你的任务是通过进行以下操作 <strong>任意次</strong> （可以是 0 次） <strong>最小化</strong> <code>nums</code> 的长度：</p>

<ul>
	<li>在 <code>nums</code> 中选择 <strong>两个不同</strong> 的下标 <code>i</code> 和 <code>j</code> ，满足 <code>nums[i] > 0</code> 且 <code>nums[j] > 0</code> 。</li>
	<li>将结果 <code>nums[i] % nums[j]</code> 插入 <code>nums</code> 的结尾。</li>
	<li>将 <code>nums</code> 中下标为 <code>i</code> 和 <code>j</code> 的元素删除。</li>
</ul>

<p>请你返回一个整数，它表示进行任意次操作以后<em> </em><code>nums</code> 的 <strong>最小长度</strong> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,4,3,1]
<b>输出：</b>1
<b>解释：</b>使数组长度最小的一种方法是：
操作 1 ：选择下标 2 和 1 ，插入 nums[2] % nums[1] 到数组末尾，得到 [1,4,3,1,3] ，然后删除下标为 2 和 1 的元素。
nums 变为 [1,1,3] 。
操作 2 ：选择下标 1 和 2 ，插入 nums[1] % nums[2] 到数组末尾，得到 [1,1,3,1] ，然后删除下标为 1 和 2 的元素。
nums 变为 [1,1] 。
操作 3 ：选择下标 1 和 0 ，插入 nums[1] % nums[0] 到数组末尾，得到 [1,1,0] ，然后删除下标为 1 和 0 的元素。
nums 变为 [0] 。
nums 的长度无法进一步减小，所以答案为 1 。
1 是可以得到的最小长度。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [5,5,5,10,5]
<b>输出：</b>2
<b>解释：</b>使数组长度最小的一种方法是：
操作 1 ：选择下标 0 和 3 ，插入 nums[0] % nums[3] 到数组末尾，得到 [5,5,5,10,5,5] ，然后删除下标为 0 和 3 的元素。
nums 变为 [5,5,5,5] 。
操作 2 ：选择下标 2 和 3 ，插入 nums[2] % nums[3] 到数组末尾，得到 [5,5,5,5,0] ，然后删除下标为 2 和 3 的元素。
nums 变为 [5,5,0] 。
操作 3 ：选择下标 0 和 1 ，插入 nums[0] % nums[1] 到数组末尾，得到 [5,5,0,0] ，然后删除下标为 0 和 1 的元素。
nums 变为 [0,0] 。
nums 的长度无法进一步减小，所以答案为 2 。
2 是可以得到的最小长度。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [2,3,4]
<b>输出：</b>1
<b>解释：</b>使数组长度最小的一种方法是：
操作 1 ：选择下标 1 和 2 ，插入 nums[1] % nums[2] 到数组末尾，得到 [2,3,4,3] ，然后删除下标为 1 和 2 的元素。
nums 变为 [2,3] 。
操作 2 ：选择下标 1 和 0 ，插入 nums[1] % nums[0] 到数组末尾，得到 [2,3,1] ，然后删除下标为 1 和 0 的元素。
nums 变为 [1] 。
nums 的长度无法进一步减小，所以答案为 1 。
1 是可以得到的最小长度。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

### 思路

考虑这个例子：$\textit{nums}=[2,3,4,5,6]$，每次操作都可以选择 $2$ 和另一个数字 $x$，
由于 $x>2$，所以 $2\bmod x = 2$，于是操作等价于：

- 移除 $x$。

所以最后必定只会剩下 $2$。

所以，如果数组中的最小值只有一个，我们可以操作成只剩下一个数，返回 $1$。
但如果最小值不止一个呢？如果能构造出一个小于 $m = \min(\textit{nums})$ 的正整数，那么也可以返回 $1$。

**结论**：
当且仅当 $\textit{nums}$ 中有不是 $m$ 的倍数的数，我们才能构造出一个小于 $m$ 的正整数。

**证明**：
如果有不是 $m$ 的倍数的数 $x$，那么 $0 < x\bmod m < m$，构造成功。如果所有数都是 $m$ 的倍数，
那么任意两个数的模都是 $m$ 的倍数，我们无法得到一个在 $[1,m-1]$ 内的数。
如果所有数都是 $m$ 的倍数，我们可以先用 $m$ 把大于 $m$ 的数都移除，然后剩下的 $\textit{cnt}$ 个 $m$ 两两一对消除，
最后剩下 $\left\lceil\dfrac{\textit{cnt}}{2}\right\rceil$ 个数。

```go [sol]
func minimumArrayLength(nums []int) int {
	mm, cnt := nums[0], 0
	for _, v := range nums {
		if v < mm {
			mm = v
		}
	}
	for _, v := range nums {
		if v % mm > 0 {
			return 1
		}
		if v == mm {
			cnt++
		}
	}
	return (cnt + 1) / 2
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
