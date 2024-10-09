#### 题目

<p>给你两个下标从 <strong>1</strong> 开始的整数数组 <code>nums</code> 和 <code>changeIndices</code> ，数组的长度分别为 <code>n</code> 和 <code>m</code> 。</p>

<p>一开始，<code>nums</code> 中所有下标都是未标记的，你的任务是标记 <code>nums</code> 中 <strong>所有</strong> 下标。</p>

<p>从第 <code>1</code> 秒到第 <code>m</code> 秒（<b>包括 </b>第 <code>m</code> 秒），对于每一秒 <code>s</code> ，你可以执行以下操作 <strong>之一</strong> ：</p>

<ul>
	<li>选择范围 <code>[1, n]</code> 中的一个下标 <code>i</code> ，并且将 <code>nums[i]</code> <strong>减少</strong> <code>1</code> 。</li>
	<li>如果 <code>nums[changeIndices[s]]</code> <strong>等于</strong> <code>0</code> ，<strong>标记</strong> 下标 <code>changeIndices[s]</code> 。</li>
	<li>什么也不做。</li>
</ul>

<p>请你返回范围 <code>[1, m]</code> 中的一个整数，表示最优操作下，标记 <code>nums</code> 中 <strong>所有</strong> 下标的 <strong>最早秒数</strong> ，如果无法标记所有下标，返回 <code>-1</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [2,2,0], changeIndices = [2,2,2,2,3,2,2,1]
<b>输出：</b>8
<b>解释：</b>这个例子中，我们总共有 8 秒。按照以下操作标记所有下标：
第 1 秒：选择下标 1 ，将 nums[1] 减少 1 。nums 变为 [1,2,0] 。
第 2 秒：选择下标 1 ，将 nums[1] 减少 1 。nums 变为 [0,2,0] 。
第 3 秒：选择下标 2 ，将 nums[2] 减少 1 。nums 变为 [0,1,0] 。
第 4 秒：选择下标 2 ，将 nums[2] 减少 1 。nums 变为 [0,0,0] 。
第 5 秒，标记 changeIndices[5] ，也就是标记下标 3 ，因为 nums[3] 等于 0 。
第 6 秒，标记 changeIndices[6] ，也就是标记下标 2 ，因为 nums[2] 等于 0 。
第 7 秒，什么也不做。
第 8 秒，标记 changeIndices[8] ，也就是标记下标 1 ，因为 nums[1] 等于 0 。
现在所有下标已被标记。
最早可以在第 8 秒标记所有下标。
所以答案是 8 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [1,3], changeIndices = [1,1,1,2,1,1,1]
<b>输出：</b>6
<b>解释：</b>这个例子中，我们总共有 7 秒。按照以下操作标记所有下标：
第 1 秒：选择下标 2 ，将 nums[2] 减少 1 。nums 变为 [1,2] 。
第 2 秒：选择下标 2 ，将 nums[2] 减少 1 。nums 变为 [1,1] 。
第 3 秒：选择下标 2 ，将 nums[2] 减少 1 。nums 变为 [1,0] 。
第 4 秒：标记 changeIndices[4] ，也就是标记下标 2 ，因为 nums[2] 等于 0 。
第 5 秒：选择下标 1 ，将 nums[1] 减少 1 。nums 变为 [0,0] 。
第 6 秒：标记 changeIndices[6] ，也就是标记下标 1 ，因为 nums[1] 等于 0 。
现在所有下标已被标记。
最早可以在第 6 秒标记所有下标。
所以答案是 6 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>Input:</strong> nums = [0,1], changeIndices = [2,2,2]
<strong>Output:</strong> -1
<strong>Explanation:</strong> 这个例子中，无法标记所有下标，因为下标 1 不在 changeIndices 中。
所以答案是 -1 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n == nums.length <= 2000</code></li>
	<li><code>0 <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>1 <= m == changeIndices.length <= 2000</code></li>
	<li><code>1 <= changeIndices[i] <= n</code></li>
</ul>

#### 思路

这题的题意很抽象，给大家形象地解释一下：

你有 $n$ 门课程需要考试，第 $i$ 门课程需要用 $\textit{nums}[i]$ 天复习。同一天只能复习一门课程。
在第 $i$ 天，你可以选择参加第 $\textit{changeIndices}[i]$ 门课程的考试。考试这一天不能复习。

搞定所有课程的复习+考试，至少要多少天？

## 提示 1

答案越大，越能够搞定所有课程，反之越不能。有单调性，可以**二分答案**。

## 提示 2

考试的时间越晚越好，这样我们能有更充足的时间复习。设二分的答案为 $\textit{mx}$。在 $\textit{mx}$ 天内，设 $\textit{i}$ 在 $\textit{changeIndices}$ 中出现的最后下标为 $\textit{lastT}[\textit{i}]$，
即第 $i$ 门课程的最晚考试时间。如果 $i$ 不在 $\textit{changeIndices}$ 的前 $\textit{mx}$ 个数中，二分返回 false。

- 初始化 $\textit{cnt}=0$，遍历 $\textit{changeIndices}$ 的前 $\textit{mx}$ 个数。
- 如果 $i\ne \textit{lastT}[i]$，这一天只能用来复习（或者什么也不做），但还不知道要复习哪一门课程，所以暂时记录一下，把 $\textit{cnt}$ 加一。
- 如果 $i= \textit{lastT}[i]$，先从 $\textit{cnt}$ 中消耗 $\textit{nums}[i]$ 天用来复习，然后考试。
- 如果 $\textit{cnt}<\textit{nums}[i]$ 则无法完成考试，二分返回 false。

```go [sol]

func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}

	lastT := make([]int, n)
	ans := n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		clear(lastT)
		for t, idx := range changeIndices[:mx] {
			lastT[idx-1] = t + 1
		}
		if slices.Contains(lastT, 0) { // 有课程没有考试时间
			return false
		}

		cnt := 0
		for i, idx := range changeIndices[:mx] {
			idx--
			if i == lastT[idx]-1 { // 考试
				if nums[idx] > cnt { // 没时间复习
					return false
				}
				cnt -= nums[idx] // 复习这门课程
			} else {
				cnt++ // 留着后面用
			}
		}
		return true
	})
	if ans > m {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m$ 为 $\textit{changeIndices}$ 的长度。二分的时候保证 $n\le m$，时间复杂度以 $m$ 为主。
- 空间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度
