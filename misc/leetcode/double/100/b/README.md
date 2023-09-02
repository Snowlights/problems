### 题目  

<p>给你一个下标从 0 开始的整数数组 <code>nums</code> 。你需要将 <code>nums</code> 重新排列成一个新的数组 <code>perm</code> 。</p>

<p>定义 <code>nums</code> 的 <strong>伟大值</strong> 为满足 <code>0 &lt;= i &lt; nums.length</code> 且 <code>perm[i] &gt; nums[i]</code> 的下标数目。</p>

<p>请你返回重新排列 <code>nums</code> 后的 <strong>最大</strong> 伟大值。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [1,3,5,2,1,3,1]
<b>输出：</b>4
<b>解释：</b>一个最优安排方案为 perm = [2,5,1,3,3,1,1] 。
在下标为 0, 1, 3 和 4 处，都有 perm[i] &gt; nums[i] 。因此我们返回 4 。</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [1,2,3,4]
<b>输出：</b>3
<b>解释：</b>最优排列为 [2,3,4,1] 。
在下标为 0, 1 和 2 处，都有 perm[i] &gt; nums[i] 。因此我们返回 3 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
### 思路  

# 方法一：贪心 + 排序 + 双指针

### 提示 1

田忌赛马。

### 提示 2-1

想一想，$\textit{nums}$ 的最小值是否要参与贡献伟大值？要和谁匹配？

### 提示 2-2

$\textit{nums}$ 的最小值要参与匹配，否则更大的数字更难匹配上。  
$\textit{nums}$ 的最小值要与次小值匹配，这样后面的数字才能取匹配更大的数。

### 提示 3

为了方便实现，对 $\textit{nums}$ 从小到大排序。（为什么可以排序？因为只在乎匹配关系，与下标无关。）  
例如示例 1 排序后为 $[1,1,1,2,3,3,5]$。那么前三个 $1$ 分别与 $2,3,3$ 匹配，$2$ 与 $5$ 匹配，后面就没有数字能匹配了。

```go 
func maximizeGreatness(nums []int) (i int) {
	sort.Ints(nums)
	for _, x := range nums {
		if x > nums[i] {
			i++
		}
	}
	return
}
```

### 复杂度分析  

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序时的栈开销，仅用到若干额外变量。

# 方法二：利用两个指针的距离

考虑无法匹配的个数 $m$，答案为 $\textit{nums}$ 的长度减去 $m$。  
再来看方法一的双指针，设另一个指针为 $j$，即 $x=\textit{nums}[j]$。  
每次 `x > nums[i]` **不成立**时，两个指针错开的距离 $j-i$ 就增加 $1$。  
那么循环结束后，两个指针最终错开的距离 $j-i$，就是 $m$。  
以 $[1,1,2,2,2,2,3,3]$ 为例，初始 $i=0,j=0$：  
- $j=0$，无法匹配，$j$ 加一。
- $j=1$，无法匹配，$j$ 加一。
- $j=2$，可以匹配，$i,j$ 都加一。
- $j=3$，可以匹配，$i,j$ 都加一。
- $j=4$，无法匹配，$j$ 加一。注意此时 $i$ 指向第一个 $2$。
- $j=5$，无法匹配，$j$ 加一。
- $j=6$，可以匹配，$i,j$ 都加一。注意此时 $j$ 指向最后一个 $2$ 的右侧相邻元素。
- $j=7$，可以匹配，$i,j$ 都加一。循环结束。  

可以发现，当 $i$ 指向出现次数最多的数 $p$ 的时候，$j$ 要一直移动到下一个不等于 $p$ 的数，此时错开的距离是最大的。  
由于后面的数出现次数不会超过 $p$，所以不会出现无法匹配的情况。（可以用反证法证明，如果出现了，说明这个数的出现次数大于 $p$，矛盾。）  
所以 $m$ 就是 $\textit{nums}$ 中元素出现次数的最大值。用哈希表统计即可，无需排序。

```go 
func maximizeGreatnessV(nums []int) int {
	maxCnt := 0
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
		maxCnt = max(maxCnt, cnt[v])
	}
	return len(nums) - maxCnt
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。