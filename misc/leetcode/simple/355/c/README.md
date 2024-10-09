#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始、长度为 <code>n</code> 的数组 <code>usageLimits</code> 。</p>

<p>你的任务是使用从 <code>0</code> 到 <code>n - 1</code> 的数字创建若干组，并确保每个数字 <code>i</code> 在 <strong>所有组</strong> 中使用的次数总共不超过 <code>usageLimits[i]</code> 次。此外，还必须满足以下条件：</p>

<ul>
	<li>每个组必须由 <strong>不同</strong> 的数字组成，也就是说，单个组内不能存在重复的数字。</li>
	<li>每个组（除了第一个）的长度必须 <strong>严格大于</strong> 前一个组。</li>
</ul>

<p>在满足所有条件的情况下，以整数形式返回可以创建的最大组数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><code><strong>输入：</strong>usageLimits</code> = [1,2,5]
<strong>输出：</strong>3
<strong>解释：</strong>在这个示例中，我们可以使用 0 至多一次，使用 1 至多 2 次，使用 2 至多 5 次。
一种既能满足所有条件，又能创建最多组的方式是： 
组 1 包含数字 [2] 。
组 2 包含数字 [1,2] 。
组 3 包含数字 [0,1,2] 。 
可以证明能够创建的最大组数是 3 。 
所以，输出是 3 。 </pre>

<p><strong>示例 2：</strong></p>

<pre><code><strong>输入：</strong></code><code>usageLimits</code> = [2,1,2]
<strong>输出：</strong>2
<strong>解释：</strong>在这个示例中，我们可以使用 0 至多 2 次，使用 1 至多 1 次，使用 2 至多 2 次。
一种既能满足所有条件，又能创建最多组的方式是： 
组 1 包含数字 [0] 。 
组 2 包含数字 [1,2] 。
可以证明能够创建的最大组数是 2 。 
所以，输出是 2 。 
</pre>

<p><strong>示例 3：</strong></p>

<pre><code><strong>输入：</strong></code><code>usageLimits</code> = [1,1]
<strong>输出：</strong>1
<strong>解释：</strong>在这个示例中，我们可以使用 0 和 1 至多 1 次。 
一种既能满足所有条件，又能创建最多组的方式是：
组 1 包含数字 [0] 。
可以证明能够创建的最大组数是 1 。 
所以，输出是 1 。 
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= usageLimits.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= usageLimits[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

排序，再判断能否递增
先排序，因为题目只关心最大组数
- 有点类似动态规划思路：x(i) = 使用排序后的数组usageLimits2[:i+1]能得到的最大数组
- 假设已经计算好了x(i-1), 想计算x(i)时，主要需要考虑 usageLimits2[i] 以及  usageLimits2[:i]的总和, 来判断最大数组能否增长
    - x(i) = x(i-1)+1的情况
        - 如果usageLimits2[i] > x(i), 则能生成的最大数组+1. 我们可以直接把i单独作为1个数组，以及之前的数组后加上i。如从[1,2] --> [1,2,4]；原来对应的数组为[1],[0,1]; 新所对应的数组为[2],[1,2],[0,1,2]
        - 如果总和满足数量约束,  综和大于等于 (x[i-1]+1)*(x[i-1]+2)/2。 如从[2,2] --> [2,2,2]. 原来的数组为 [0],[0,1]； 新数组为[0],[1,2],[0,1,2]
    - 否则为 x(i) = x(i-1)
- 这里注意每次只能递增1，即使数量更多。 如[1,1] （能创建1组） --> [1,1,4] 时，虽然有6个数，从数字上看能创建3组（1+2+3 = 6），但其实依然只有2组满足条件的数。


```go 
func maxIncreasingGroups(usageLimits []int) (ans int) {
	sort.Ints(usageLimits)
	ans, cur := 0, 0
	for _, v := range usageLimits {
		cur += v
		if cur > ans {
			ans++
			cur -= ans
		}
	}
	return ans
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(n)$ 。