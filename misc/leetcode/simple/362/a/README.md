#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的二维整数数组 <code>nums</code> 表示汽车停放在数轴上的坐标。对于任意下标 <code>i</code>，<code>nums[i] = [start<sub>i</sub>, end<sub>i</sub>]</code> ，其中 <code>start<sub>i</sub></code> 是第 <code>i</code> 辆车的起点，<code>end<sub>i</sub></code> 是第 <code>i</code> 辆车的终点。</p>
<p>返回数轴上被车 <strong>任意部分</strong> 覆盖的整数点的数目。</p>
<p> </p>
<p><strong class="example">示例 1：</strong></p>
<pre>
<strong>输入：</strong>nums = [[3,6],[1,5],[4,7]]
<strong>输出：</strong>7
<strong>解释：</strong>从 1 到 7 的所有点都至少与一辆车相交，因此答案为 7 。
</pre>
<p><strong class="example">示例 2：</strong></p>
<pre>
<strong>输入：</strong>nums = [[1,3],[5,8]]
<strong>输出：</strong>7
<strong>解释：</strong>1、2、3、5、6、7、8 共计 7 个点满足至少与一辆车相交，因此答案为 7 。
</pre>
<p> </p>
<p><strong>提示：</strong></p>
<ul>
<li><code>1 <= nums.length <= 100</code></li>
<li><code>nums[i].length == 2</code></li>
<li><code><font face="monospace">1 <= start<sub>i</sub> <= end<sub>i</sub> <= 100</font></code></li>
</ul>

#### 思路

之前在力扣上写过一篇文章：[【算法小课堂】差分数组](https://leetcode.cn/circle/discuss/FfMCgb/)。  
根据这篇文章，可以用查封数组求出每个点被覆盖多少次。  
答案就是覆盖次数大于 $0$ 的点的个数。

```go  
func numberOfPoints(nums [][]int) (ans int) {
	diff := [102]int{}
	for _, p := range nums {
		diff[p[0]]++
		diff[p[1]+1]--
	}
	s := 0
	for _, d := range diff {
		s += d
		if s > 0 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+\max{\textit{end}_i})$。其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\max{\textit{end}_i})$。