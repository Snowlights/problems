#### 题目

<p>给你一个下标从 <strong>0</strong> 开始、长度为 <code>n</code> 的整数数组 <code>nums</code> ，其中 <code>n</code> 是班级中学生的总数。班主任希望能够在让所有学生保持开心的情况下选出一组学生：</p>

<p>如果能够满足下述两个条件之一，则认为第 <code>i</code> 位学生将会保持开心：</p>

<ul>
	<li>这位学生被选中，并且被选中的学生人数 <strong>严格大于</strong> <code>nums[i]</code> 。</li>
	<li>这位学生没有被选中，并且被选中的学生人数 <strong>严格小于</strong> <code>nums[i]</code> 。</li>
</ul>

<p>返回能够满足让所有学生保持开心的分组方法的数目。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,1]
<strong>输出：</strong>2
<strong>解释：</strong>
有两种可行的方法：
班主任没有选中学生。
班主任选中所有学生形成一组。 
如果班主任仅选中一个学生来完成分组，那么两个学生都无法保持开心。因此，仅存在两种可行的方法。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [6,0,3,3,6,7,2,7]
<strong>输出：</strong>3
<strong>解释：</strong>
存在三种可行的方法：
班主任选中下标为 1 的学生形成一组。
班主任选中下标为 1、2、3、6 的学生形成一组。
班主任选中所有学生形成一组。 
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt; nums.length</code></li>
</ul>


#### 思路

为了方便判断，把 $\textit{nums}$ 从小到大排序。
如果 $\textit{nums}[0] > 0$，那么所有 $\textit{nums}[i]$ 都是大于 $0$ 的，我们可以一个学生都不选。
如果 $\textit{nums}[i] < i+1 < \textit{nums}[i+1]$，这意味着选择 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 这一共 $i+1$ 个学生，是满足要求的：由于数组已经排好序，$\textit{nums}[0]$ 到 $\textit{nums}[i]$ 都是小于 $i+1$ 的，而 $\textit{nums}[i+1]$ 到 $\textit{nums}[n-1]$ 都是大于 $i+1$ 的。
特别地，如果 $i=n-1$，我们只需要判断是否满足 $\textit{nums}[i] < n$。

```go 
func countWays(nums []int) (ans int) {
	sort.Ints(nums)
	if nums[0] > 0 {
		ans++
	}
	for i, x := range nums {
		if x < i+1 && (i == len(nums)-1 || i+1 < nums[i+1]) {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。