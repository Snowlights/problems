#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个整数 <code>k</code> 。</p>

<p>一次操作中，你可以选择 <code>nums</code> 中满足 <code>0 <= i < nums.length - 1</code> 的一个下标 <code>i</code> ，并将 <code>nums[i]</code> 和 <code>nums[i + 1]</code> 替换为数字 <code>nums[i] & nums[i + 1]</code> ，其中 <code>&</code> 表示按位 <code>AND</code> 操作。</p>

<p>请你返回 <strong>至多</strong> <code>k</code> 次操作以内，使 <code>nums</code> 中所有剩余元素按位 <code>OR</code> 结果的 <strong>最小值</strong> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,5,3,2,7], k = 2
<b>输出：</b>3
<b>解释：</b>执行以下操作：
1. 将 nums[0] 和 nums[1] 替换为 (nums[0] & nums[1]) ，得到 nums 为 [1,3,2,7] 。
2. 将 nums[2] 和 nums[3] 替换为 (nums[2] & nums[3]) ，得到 nums 为 [1,3,2] 。
最终数组的按位或值为 3 。
3 是 k 次操作以内，可以得到的剩余元素的最小按位或值。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [7,3,15,14,2,8], k = 4
<b>输出：</b>2
<b>解释：</b>执行以下操作：
1. 将 nums[0] 和 nums[1] 替换为 (nums[0] & nums[1]) ，得到 nums 为 [3,15,14,2,8] 。
2. 将 nums[0] 和 nums[1] 替换为 (nums[0] & nums[1]) ，得到 nums 为 [3,14,2,8] 。
3. 将 nums[0] 和 nums[1] 替换为 (nums[0] & nums[1]) ，得到 nums 为 [2,2,8] 。
4. 将 nums[1] 和 nums[2] 替换为 (nums[1] & nums[2]) ，得到 nums 为 [2,0] 。
最终数组的按位或值为 2 。
2 是 k 次操作以内，可以得到的剩余元素的最小按位或值。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [10,7,10,3,9,14,9,4], k = 1
<b>输出：</b>15
<b>解释：</b>不执行任何操作，nums 的按位或值为 15 。
15 是 k 次操作以内，可以得到的剩余元素的最小按位或值。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>0 <= nums[i] < 2<sup>30</sup></code></li>
	<li><code>0 <= k < nums.length</code></li>
</ul>

#### 思路

## 提示 1

从高到低考虑：答案在这一位能不能是 $0$？

## 提示 2

一次操作，用 AND 合并两个相邻数字。
多次操作，相当于把一段**连续子数组**合并成 $0$。
尝试从左到右合并（忽略低位和必须是 $1$ 的位），如果合并出 $0$，就开始合并下一段。

以 $[5,2,3,6]$ 为例说明，这四个数的二进制表示如下：

$$
\begin{aligned}
&101\\
&010\\
&011\\
&110
\end{aligned}

$$

设 $k=2$。从高到低考虑：答案在这一位能不能是 $0$？

- 最高位有两个 $1$，合并掉这两个 $1$ 需要操作 $2$ 次（$\le k$），所以答案的最高位可以是 $0$。
- 对于次高位，我们需要通过一连串的合并，让合并结果的最高位和次高位都是 $0$。数组前两个数（只看最高位和次高位）可以合并成 $0$，操作 $1$ 次。数组后两个数（只看最高位和次高位）无法合并成 $0$，那么用前两个数合并出来的 $0$，与后两个数合并，操作 $2$ 次，得到 $0$。所以一共要操作 $3$ 次才能让最高位和次高位都是 $0$，无法做到，所以答案的次高位一定是 $1$。
- 对于最低位，我们需要通过一连串的合并，让合并结果的最高位和最低位都是 $0$。注意我们**无需考虑次高位**，因为前面已经确定答案这一位是 $1$ 了。数组前两个数（只看最高位和最低位）可以合并成 $0$，操作 $1$ 次。数组后两个数（只看最高位和最低位）也可以合并成 $0$，操作 $1$ 次。一共操作 $2$ 次，所以答案的最低位可以是 $0$。
- 综上所述，答案的二进制表示为 $010$，即十进制 $2$。

注意，如果整个数组都无法合并成 $0$，那么代码计算出来的操作次数是 $n$，题目保证了这是大于 $k$ 的，无需特判这种情况。

```go [sol]
func minOrAfterOperations(nums []int, k int) (ans int) {
	mask := 0
	for b := 29; b >= 0; b-- {
		mask |= 1 << b
		cnt := 0  // 操作次数
		and := -1 // -1 的二进制全为 1
		for _, x := range nums {
			and &= x & mask
			if and != 0 {
				cnt++ // 合并 x，操作次数加一
			} else {
				and = -1 // 准备合并下一段
			}
		}
		if cnt > k {
			ans |= 1 << b  // 答案的这个比特位必须是 1
			mask ^= 1 << b // 后面不考虑这个比特位
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。
