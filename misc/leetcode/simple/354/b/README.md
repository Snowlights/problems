#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个 <strong>非负</strong> 整数 <code>k</code> 。</p>

<p>在一步操作中，你可以执行下述指令：</p>

<ul>
	<li>在范围 <code>[0, nums.length - 1]</code> 中选择一个 <strong>此前没有选过</strong> 的下标 <code>i</code> 。</li>
	<li>将 <code>nums[i]</code> 替换为范围 <code>[nums[i] - k, nums[i] + k]</code> 内的任一整数。</li>
</ul>

<p>数组的 <strong>美丽值</strong> 定义为数组中由相等元素组成的最长子序列的长度。</p>

<p>对数组 <code>nums</code> 执行上述操作任意次后，返回数组可能取得的 <strong>最大</strong> 美丽值。</p>

<p><strong>注意：</strong>你 <strong>只</strong> 能对每个下标执行 <strong>一次</strong> 此操作。</p>

<p>数组的 <strong>子序列</strong> 定义是：经由原数组删除一些元素（也可能不删除）得到的一个新数组，且在此过程中剩余元素的顺序不发生改变。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [4,6,1,2], k = 2
<strong>输出：</strong>3
<strong>解释：</strong>在这个示例中，我们执行下述操作：
- 选择下标 1 ，将其替换为 4（从范围 [4,8] 中选出），此时 nums = [4,4,1,2] 。
- 选择下标 3 ，将其替换为 4（从范围 [0,4] 中选出），此时 nums = [4,4,1,4] 。
执行上述操作后，数组的美丽值是 3（子序列由下标 0 、1 、3 对应的元素组成）。
可以证明 3 是我们可以得到的由相等元素组成的最长子序列长度。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [1,1,1,1], k = 10
<strong>输出：</strong>4
<strong>解释：</strong>在这个示例中，我们无需执行任何操作。
数组 nums 的美丽值是 4（整个数组）。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>0 <= nums[i], k <= 10<sup>5</sup></code></li>
</ul>

#### 思路

由于选的是子序列，且子序列的元素都相等，所以**元素顺序对答案没有影响**，可以先对数组**排序**。
由于替换操作替换的是一个连续范围内的数，所以排序后，选出的子序列必然也是一段**连续子数组**。
那么问题变成：「找最长的连续子数组，其最大值减最小值不超过 $2k$」，只要子数组满足这个要求，其中的元素都可以变成同一个数。
这个问题可以用 [同向双指针] 解决。枚举 $\textit{nums}[\textit{right}]$ 作为子数组的最后一个数，一旦 $\textit{nums}[\textit{right}]-\textit{nums}[\textit{left}]>2k$，就移动左端点。
$\textit{right}-\textit{left}+1$ 是子数组的长度，取所有长度最大值，即为答案。

```go
func maximumBeauty(nums []int, k int) (ans int) {
	sort.Ints(nums)
	left := 0
	for right, x := range nums {
		for x-nums[left] > k*2 {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- =时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。
