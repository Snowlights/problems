### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个正整数 <code>k</code> 。</p>

<p>你可以对数组执行以下操作 <strong>任意次</strong> ：</p>

<ul>
	<li>选择数组里的 <strong>任意</strong> 一个元素，并将它的 <strong>二进制</strong> 表示 <strong>翻转</strong> 一个数位，翻转数位表示将 <code>0</code> 变成 <code>1</code> 或者将 <code>1</code> 变成 <code>0</code> 。</li>
</ul>

<p>你的目标是让数组里 <strong>所有</strong> 元素的按位异或和得到 <code>k</code> ，请你返回达成这一目标的 <strong>最少 </strong>操作次数。</p>

<p><strong>注意</strong>，你也可以将一个数的前导 0 翻转。比方说，数字 <code>(101)<sub>2</sub></code> 翻转第四个数位，得到 <code>(1101)<sub>2</sub></code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [2,1,3,4], k = 1
<b>输出：</b>2
<b>解释：</b>我们可以执行以下操作：
- 选择下标为 2 的元素，也就是 3 == (011)<sub>2</sub> ，我们翻转第一个数位得到 (010)<sub>2</sub> == 2 。数组变为 [2,1,2,4] 。
- 选择下标为 0 的元素，也就是 2 == (010)<sub>2</sub> ，我们翻转第三个数位得到 (110)<sub>2</sub> == 6 。数组变为 [6,1,2,4] 。
最终数组的所有元素异或和为 (6 XOR 1 XOR 2 XOR 4) == 1 == k 。
无法用少于 2 次操作得到异或和等于 k 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,0,2,0], k = 0
<b>输出：</b>0
<strong>解释：</strong>数组所有元素的异或和为 (2 XOR 0 XOR 2 XOR 0) == 0 == k 。所以不需要进行任何操作。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>0 <= nums[i] <= 10<sup>6</sup></code></li>
	<li><code>0 <= k <= 10<sup>6</sup></code></li>
</ul>

### 思路

设 $\textit{nums}$ 的异或和为 $s$。
$s=k$ 等价于 $s\oplus k = 0$，其中 $\oplus$ 表示异或。
设 $x = s\oplus k$，我们把 $\textit{nums}$ 中的任意数字的某个比特位翻转，那么 $x$ 的这个比特位也会翻转。
要让 $x=0$，就必须把 $x$ 中的每个 $1$ 都翻转，所以 $x$ 中的 $1$ 的个数就是我们的操作次数。

```go [sol]
func minOperations(nums []int, k int) int {
	for _, v := range nums {
		k ^= v
	}
	return bits.OnesCount(uint(k))
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。
