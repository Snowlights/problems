#### 题目

<p>给你两个下标从 <strong>1</strong>&nbsp;开始的整数数组&nbsp;<code>nums</code> 和&nbsp;<code>changeIndices</code>&nbsp;，数组的长度分别为&nbsp;<code>n</code> 和&nbsp;<code>m</code>&nbsp;。</p>

<p>一开始，<code>nums</code>&nbsp;中所有下标都是未标记的，你的任务是标记 <code>nums</code>&nbsp;中 <strong>所有</strong>&nbsp;下标。</p>

<p>从第 <code>1</code>&nbsp;秒到第 <code>m</code>&nbsp;秒（<b>包括&nbsp;</b>第&nbsp;<code>m</code>&nbsp;秒），对于每一秒 <code>s</code>&nbsp;，你可以执行以下操作 <strong>之一</strong>&nbsp;：</p>

<ul>
	<li>选择范围&nbsp;<code>[1, n]</code>&nbsp;中的一个下标 <code>i</code>&nbsp;，并且将&nbsp;<code>nums[i]</code> <strong>减少</strong>&nbsp;<code>1</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[changeIndices[s]]</code>&nbsp;设置成任意的 <strong>非负</strong>&nbsp;整数。</li>
	<li>选择范围&nbsp;<code>[1, n]</code>&nbsp;中的一个下标&nbsp;<code>i</code>&nbsp;， 满足&nbsp;<code>nums[i]</code> <strong>等于</strong> <code>0</code>, 并 <strong>标记</strong>&nbsp;下标&nbsp;<code>i</code> 。</li>
	<li>什么也不做。</li>
</ul>

<p>请你返回范围 <code>[1, m]</code>&nbsp;中的一个整数，表示最优操作下，标记&nbsp;<code>nums</code>&nbsp;中 <strong>所有</strong>&nbsp;下标的 <strong>最早秒数</strong>&nbsp;，如果无法标记所有下标，返回 <code>-1</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,2,3], changeIndices = [1,3,2,2,2,2,3]
<b>输出：</b>6
<b>解释：</b>这个例子中，我们总共有 7 秒。按照以下操作标记所有下标：
第 1 秒：将 nums[changeIndices[1]] 变为 0 。nums 变为 [0,2,3] 。
第 2 秒：将 nums[changeIndices[2]] 变为 0 。nums 变为 [0,2,0] 。
第 3 秒：将 nums[changeIndices[3]] 变为 0 。nums 变为 [0,0,0] 。
第 4 秒：标记下标 1 ，因为 nums[1] 等于 0 。
第 5 秒：标记下标 2 ，因为 nums[2] 等于 0 。
第 6 秒：标记下标 3 ，因为 nums[3] 等于 0 。
现在所有下标已被标记。
最早可以在第 6 秒标记所有下标。
所以答案是 6 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [0,0,1,2], changeIndices = [1,2,1,2,1,2,1,2]
<b>输出：</b>7
<b>解释：</b>这个例子中，我们总共有 8 秒。按照以下操作标记所有下标：
第 1 秒：标记下标 1 ，因为 nums[1] 等于 0 。
第 2 秒：标记下标 2 ，因为 nums[2] 等于 0 。
第 3 秒：将 nums[4] 减少 1 。nums 变为 [0,0,1,1] 。
第 4 秒：将 nums[4] 减少 1 。nums 变为 [0,0,1,0] 。
第 5 秒：将 nums[3] 减少 1 。nums 变为 [0,0,0,0] 。
第 6 秒：标记下标 3 ，因为 nums[3] 等于 0 。
第 7 秒：标记下标 4 ，因为 nums[4] 等于 0 。
现在所有下标已被标记。
最早可以在第 7 秒标记所有下标。
所以答案是 7 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,3], changeIndices = [1,2,3]
<b>输出：</b>-1
<strong>解释：</strong>这个例子中，无法标记所有下标，因为我们没有足够的秒数。
所以答案是 -1 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == nums.length &lt;= 5000</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= m == changeIndices.length &lt;= 5000</code></li>
	<li><code>1 &lt;= changeIndices[i] &lt;= n</code></li>
</ul>

#### 思路

这题的题意很抽象，给大家形象地解释一下：  
你有 $n$ 门课程需要考试，第 $i$ 门课程需要用 $\textit{nums}[i]$ 天复习。同一天只能复习一门课程（**慢速**复习）。  
在第 $i$ 天，你可以**快速**搞定第 $\textit{changeIndices}[i]$ 门课程的复习。  
你可以在任意一天完成一门课程的考试（前提是复习完成）。考试这一天不能复习。  
搞定所有课程的复习+考试，至少要多少天？

## 提示 1

答案越大，越能够搞定所有课程，反之越不能。 有单调性，可以**二分答案**。

## 提示 2

如果决定在第 $i$ 天快速复习第 $\textit{changeIndices}[i]$ 门课程，那么在第 $i$ 天前慢速复习这门课程是没有意义的。同理，如果决定慢速复习某门课程，那么也没必要对这门课程使用快速复习。  
在 $\textit{nums}[i] > 1$ 的情况下，如果用快速复习+考试，只需要花费 $2$ 天。这比只用慢速复习+考试要更好。  
但是！如果一门课程对应的 $\textit{changeIndices}$ 很靠后，可能没有时间快速复习这门课程并完成考试。比如只剩下 $2$ 天，但是还有 $3$ 门课程没有考。  
此外，如果一门课程的复习时间很长（$\textit{nums}[i]$ 很大），当我们把后续时间都用在快速复习其它复习时间比较小的课程上，可能就没有时间快速复习 $\textit{nums}[i]$ 很大的课程了（还要留一天来考试）。  
如何权衡把哪些课程快速复习，把哪些课程慢速复习呢？

## 提示 3

设二分的答案为 $\textit{mx}$。我们倒着遍历 $\textit{changeIndices}$ 的前 $\textit{mx}$ 个数，和 [周赛第三题](https://leetcode.cn/problems/earliest-second-to-mark-indices-i/) 一样，尽量选择靠左的时间来快速复习，这样右边就有更充足的时间用来考试。  
用一个数组 $\textit{firstT}$ 记录 $1$ 到 $n$（代码中是 $0$ 到 $n-1$）在 $\textit{changeIndices}$ 中首次出现的下标。初始化可用天数 $\textit{cnt}=0$。
- 设当前天数为 $t$，设 $i = \textit{changeIndices}[t] - 1$。
- 如果 $i$ 不是在 $\textit{changeIndices}$ 中首次出现的数，或者 $\textit{nums}[i]\le 1$，那么把时间留给左边再决定做什么，$\textit{cnt}$ 加一。
- 否则如果 $\textit{cnt}>0$，我们直接快速复习第 $i$ 门课程，并消耗一天用来考试，把 $\textit{cnt}$ 减一。然后把 $(\textit{nums}[i],i)$ 加到一个**小根堆**中。
- 否则如果 $\textit{cnt}=0$，那么尝试在小根堆中「反悔」一个复习时间比 $\textit{nums}[i]$ 小的数。如果堆为空或者堆顶大于等于 $\textit{nums}[i]$ 就不反悔，否则弹出堆顶并把 $\textit{cnt}$ 加二（一天快速复习，一天考试），然后做法同上述 $\textit{cnt}>0$ 的情况。这里从堆中弹出的课程，相当于用更靠左的时间去慢速复习+考试。

遍历结束后，对于每个未快速复习的课程，全部使用慢速复习+考试，将 $\textit{cnt}$ 减去这些课程对应的 $\textit{nums}[i]+1$。如果最终 $\textit{cnt}\ge 0$ 则说明可以在 $\textit{mx}$ 天内搞定所有课程的复习+考试。

```go [sol]
func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}

	total := n
	for _, v := range nums {
		total += v // 慢速复习+考试所需天数
	}

	firstT := make([]int, n)
	for t := m - 1; t >= 0; t-- {
		firstT[changeIndices[t]-1] = t + 1
	}

	h := hp{}
	ans := n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		cnt, slow := 0, total
		h = h[:0]
		for t := mx - 1; t >= 0; t-- {
			i := changeIndices[t] - 1
			v := nums[i]
			if v <= 1 || t != firstT[i]-1 {
				cnt++ // 留给左边，用来快速复习/考试
				continue
			}
			if cnt == 0 {
				if len(h) == 0 || v <= h[0].v {
					cnt++ // 留给左边，用来快速复习/考试
					continue
				}
				slow += heap.Pop(&h).(pair).v + 1
				cnt += 2 // 反悔：一天快速复习，一天考试
			}
			slow -= v + 1
			cnt-- // 快速复习，然后消耗一天来考试
			heap.Push(&h, pair{v, i})
		}
		return cnt >= slow // 剩余天数不能慢速复习+考试
	})
	if ans > m {
		return -1
	}
	return ans
}

type pair struct{ v, i int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].v < h[j].v }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log (mn))$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{changeIndices}$ 的长度。二分的时候保证 $n\le m$，时间复杂度以 $m$ 为主。注意堆中至多有 $n$ 个元素。
- 空间复杂度：$\mathcal{O}(n)$。

