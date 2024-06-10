#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的二进制数组 <code>nums</code>，其长度为 <code>n</code> ；另给你一个 <strong>正整数 </strong><code>k</code> 以及一个 <strong>非负整数 </strong><code>maxChanges</code> 。</p>

<p>灵茶山艾府在玩一个游戏，游戏的目标是让灵茶山艾府使用 <strong>最少 </strong>数量的 <strong>行动 </strong>次数从 <code>nums</code> 中拾起 <code>k</code> 个 1 。游戏开始时，灵茶山艾府可以选择数组 <code>[0, n - 1]</code> 范围内的任何索引<code>index</code> 站立。如果 <code>nums[index] == 1</code> ，灵茶山艾府就会拾起一个 1 ，并且 <code>nums[index]</code> 变成<code>0</code>（这<strong> 不算 </strong>作一次行动）。之后，灵茶山艾府可以执行 <strong>任意数量</strong> 的 <strong>行动</strong>（<strong>包括</strong><strong>零次</strong>），在每次行动中灵茶山艾府必须 <strong>恰好 </strong>执行以下动作之一：</p>

<ul>
	<li>选择任意一个下标 <code>j != index</code> 且满足 <code>nums[j] == 0</code> ，然后将 <code>nums[j]</code> 设置为 <code>1</code> 。这个动作最多可以执行 <code>maxChanges</code> 次。</li>
	<li>选择任意两个相邻的下标 <code>x</code> 和 <code>y</code>（<code>|x - y| == 1</code>）且满足 <code>nums[x] == 1</code>, <code>nums[y] == 0</code> ，然后交换它们的值（将 <code>nums[y] = 1</code> 和 <code>nums[x] = 0</code>）。如果 <code>y == index</code>，在这次行动后灵茶山艾府拾起一个 1 ，并且 <code>nums[y]</code> 变成 <code>0</code> 。</li>
</ul>

<p>返回灵茶山艾府拾起 <strong>恰好 </strong><code>k</code> 个 1 所需的 <strong>最少 </strong>行动次数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">nums = [1,1,0,0,0,1,1,0,0,1], k = 3, maxChanges = 1</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">3</span></p>

<p><strong>解释：</strong>如果游戏开始时灵茶山艾府在 <code>index == 1</code> 的位置上，按照以下步骤执行每个动作，他可以利用 <code>3</code> 次行动拾取 <code>3</code> 个 1 ：</p>

<ul>
	<li>游戏开始时灵茶山艾府拾取了一个 1 ，<code>nums[1]</code> 变成了 <code>0</code>。此时 <code>nums</code> 变为 <code>[1,<strong><u>0</u></strong>,0,0,0,1,1,0,0,1]</code> 。</li>
	<li>选择 <code>j == 2</code> 并执行第一种类型的动作。<code>nums</code> 变为 <code>[1,<strong><u>0</u></strong>,1,0,0,1,1,0,0,1]</code></li>
	<li>选择 <code>x == 2</code> 和 <code>y == 1</code> ，并执行第二种类型的动作。<code>nums</code> 变为 <code>[1,<strong><u>1</u></strong>,0,0,0,1,1,0,0,1]</code> 。由于 <code>y == index</code>，灵茶山艾府拾取了一个 1 ，<code>nums</code> 变为&nbsp; <code>[1,<strong><u>0</u></strong>,0,0,0,1,1,0,0,1]</code> 。</li>
	<li>选择 <code>x == 0</code> 和 <code>y == 1</code> ，并执行第二种类型的动作。<code>nums</code> 变为 <code>[0,<strong><u>1</u></strong>,0,0,0,1,1,0,0,1]</code> 。由于 <code>y == index</code>，灵茶山艾府拾取了一个 1 ，<code>nums</code> 变为&nbsp; <code>[0,<strong><u>0</u></strong>,0,0,0,1,1,0,0,1]</code> 。</li>
</ul>

<p>请注意，灵茶山艾府也可能执行其他的 <code>3</code> 次行动序列达成拾取 <code>3</code> 个 1 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;"><!-- 以下是示例内容的中文翻译，同时保留了原有的HTML格式和注释 -->
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">nums = [0,0,0,0], k = 2, maxChanges = 3</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">4</span></p>

<p><strong>解释：</strong>如果游戏开始时灵茶山艾府在 <code>index == 0</code> 的位置上，按照以下步骤执行每个动作，他可以利用 <code>4</code> 次行动拾取 <code>2</code> 个 1 ：</p>

<ul>
	<li>选择 <code>j == 1</code> 并执行第一种类型的动作。<code>nums</code> 变为 <code>[<strong><u>0</u></strong>,1,0,0]</code> 。</li>
	<li>选择 <code>x == 1</code> 和 <code>y == 0</code> ，并执行第二种类型的动作。<code>nums</code> 变为 <code>[<strong><u>1</u></strong>,0,0,0]</code> 。由于 <code>y == index</code>，灵茶山艾府拾起了一个 1 ，<code>nums</code> 变为 <code>[<strong><u>0</u></strong>,0,0,0]</code> 。</li>
	<li>再次选择 <code>j == 1</code> 并执行第一种类型的动作。<code>nums</code> 变为 <code>[<strong><u>0</u></strong>,1,0,0]</code> 。</li>
	<li>再次选择 <code>x == 1</code> 和 <code>y == 0</code> ，并执行第二种类型的动作。<code>nums</code> 变为 <code>[<strong><u>1</u></strong>,0,0,0]</code> 。由于<code>y == index</code>，灵茶山艾府拾起了一个 1 ，<code>nums</code> 变为 <code>[<strong><u>0</u></strong>,0,0,0]</code> 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 1</code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= maxChanges &lt;= 10<sup>5</sup></code></li>
	<li><code>maxChanges + sum(nums) &gt;= k</code></li>
</ul>


#### 思路

## 提示 1

把 $0$ 看成「空位」。

第二种操作相当于把一个 $1$ 移动到它相邻的空位上，如果我们想得到一个下标在 $j$ 的 $1$，必须操作 $|\textit{index} - j|$ 次。

对于第一种操作，我们可以选择把和 $\textit{index}$ 相邻的 $0$ 变成 $1$（在此之前先移动相邻的 $1$），然后结合第二种操作，把相邻的 $1$ 移动到 $\textit{index}$，总共操作 $2$ 次就可以得到一个 $1$。

## 提示 2

看上去应当优先使用第一种操作+第二种操作（毕竟只需要操作 $2$ 次就能得到一个 $1$），在数组中有 $1$ 且 $\textit{maxChanges}$ 很大的情况下，答案会是 $(k-1)\cdot 2$ 吗？

不对，对于 $\textit{index}, \textit{index}-1, \textit{index}+1$ 这三个位置上的 $1$。我们可以用更少的操作得到：

- $\textit{index}$ 位置上的 $1$ 相当于操作 $0$ 次。
- $\textit{index}-1$ 和 $\textit{index}+1$ 位置上的 $1$ 只需操作 $1$ 次。

设 $c$ 为 $\textit{nums}$ 中的最长连续 $1$ 的长度（实际上只需要看是否有 $3$ 个或 $2$ 个连续的 $1$）。如果 $c>k$ 则 $c=k$。

如果 $\textit{maxChanges}\ge k-c$，我们可以先使用 $\max(c-1, 0)$ 次第二种操作，收集这连续的 $c$ 个 $1$，然后对于其余 $k-c$ 个 $1$，都可以用 $2$ 次操作得到，此时可以直接返回 $\max(c-1, 0) + (k-c)\cdot 2$。

接下来，要解决的就是 $\textit{maxChanges}$ 比较小的情况了。

想一想，如果 $\textit{maxChanges}=0$，也就是只能使用第二种操作，要如何计算答案呢？

## 提示 3

这是一个标准的「货仓选址」问题。我们首先算出所有 $1$ 的位置，记到数组 $\textit{pos}$ 中，例如示例 1 的 $\textit{nums} = [1,1,0,0,0,1,1,0,0,1]$，其 $\textit{pos}=[0, 1, 5, 6, 9]$。

如果 $\textit{maxChanges}=0$，我们可以枚举所有长为 $3$ 的子数组，例如 $[0,1,5]$，就好比在坐标轴上的 $0,1,5$ 位置上有三个货仓，把工厂建在哪里，可以使得货仓到工厂的距离和最小？

根据 [中位数贪心及其证明](https://leetcode.cn/problems/5TxKeK/solution/zhuan-huan-zhong-wei-shu-tan-xin-dui-din-7r9b/)，最优解是把工厂建在货仓的中位数上。例如 $[0,1,5]$ 中的 $1$，此时距离和等于 $|0-1|+|1-1|+|5-1| = 6$。

利用前缀和，可以 $\mathcal{O}(1)$ 算出子数组中的所有数到其中位数的距离之和，原理请看 [图解](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/)。

最后，如果 $\textit{maxChanges}>0$，我们可以先计算所有长为 $k - \textit{maxChanges}$ 的子数组的货仓选址问题，然后再通过 $\textit{maxChanges}\cdot 2$ 次操作得到 $\textit{maxChanges}$ 个 $1$。

注：也有不用前缀和的滑窗做法，具体见 2968 题的 [题解方法二](https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score/solution/hua-dong-chuang-kou-zhong-wei-shu-tan-xi-nuvr/)。

```go [sol-Go]
func minimumMoves(nums []int, k, maxChanges int) int64 {
	pos := []int{}
	c := 0 // nums 中连续的 1 长度
	for i, x := range nums {
		if x == 0 {
			continue
		}
		pos = append(pos, i) // 记录 1 的位置
		c = max(c, 1)
		if i > 0 && nums[i-1] == 1 {
			if i > 1 && nums[i-2] == 1 {
				c = 3 // 有 3 个连续的 1
			} else {
				c = max(c, 2) // 有 2 个连续的 1
			}
		}
	}

	c = min(c, k)
	if maxChanges >= k-c {
		// 其余 k-c 个 1 可以全部用两次操作得到
		return int64(max(c-1, 0) + (k-c)*2)
	}

	n := len(pos)
	sum := make([]int, n+1)
	for i, x := range pos {
		sum[i+1] = sum[i] + x
	}

	ans := math.MaxInt
	// 除了 maxChanges 个数可以用两次操作得到，其余的 1 只能一步步移动到 pos[i]
	size := k - maxChanges
	for right := size; right <= n; right++ {
		// s1+s2 是 j 在 [left, right) 中的所有 pos[j] 到 pos[(left+right)/2] 的距离之和
		left := right - size
		i := left + size/2
		s1 := pos[i]*(i-left) - (sum[i] - sum[left])
		s2 := sum[right] - sum[i] - pos[i]*(right-i)
		ans = min(ans, s1+s2)
	}
	return int64(ans + maxChanges*2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 题单：中位数贪心

- [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- [2033. 获取单值网格的最小操作数](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/) 1672
- [2448. 使数组相等的最小开销](https://leetcode.cn/problems/minimum-cost-to-make-array-equal/) 2005
- [2607. 使子数组元素和相等](https://leetcode.cn/problems/make-k-subarray-sums-equal/) 2071
- [2967. 使数组成为等数数组的最小代价](https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/) 2116
- [1478. 安排邮筒](https://leetcode.cn/problems/allocate-mailboxes/) 2190
- [2968. 执行操作使频率分数最大](https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score/) 2444
- [1703. 得到连续 K 个 1 的最少相邻交换次数](https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/) 2467
- [LCP 24. 数字游戏](https://leetcode.cn/problems/5TxKeK/)
- [296. 最佳的碰头地点](https://leetcode.cn/problems/best-meeting-point/) 二维的情况（会员题）
