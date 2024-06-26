#### 题目  

<p>公司里共有 <code>n</code> 名员工，按从 <code>0</code> 到 <code>n - 1</code> 编号。每个员工 <code>i</code> 已经在公司工作了 <code>hours[i]</code> 小时。</p>

<p>公司要求每位员工工作 <strong>至少</strong> <code>target</code> 小时。</p>

<p>给你一个下标从 <strong>0</strong> 开始、长度为 <code>n</code> 的非负整数数组 <code>hours</code> 和一个非负整数 <code>target</code> 。</p>

<p>请你用整数表示并返回工作至少 <code>target</code> 小时的员工数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>hours = [0,1,2,3,4], target = 2
<strong>输出：</strong>3
<strong>解释：</strong>公司要求每位员工工作至少 2 小时。
- 员工 0 工作 0 小时，不满足要求。
- 员工 1 工作 1 小时，不满足要求。
- 员工 2 工作 2 小时，满足要求。
- 员工 3 工作 3 小时，满足要求。
- 员工 4 工作 4 小时，满足要求。
共有 3 位满足要求的员工。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>hours = [5,1,4,2,2], target = 6
<strong>输出：</strong>0
<strong>解释：</strong>公司要求每位员工工作至少 6 小时。
共有 0 位满足要求的员工。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == hours.length &lt;= 50</code></li>
	<li><code>0 &lt;= hours[i], target &lt;= 10<sup>5</sup></code></li>
</ul>
 
#### 思路  

简单一次遍历

```go 
func numberOfEmployeesWhoMetTarget(hours []int, tar int) (ans int) {
	for _, v  := range hours {
		if v >= tar {
			ans++
		}
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$ 。