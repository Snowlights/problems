#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个正整数 <code>k</code> 。</p>

<p>你可以对数组执行下述操作 <strong>任意次</strong> ：</p>

<ul>
	<li>从数组中选出长度为 <code>k</code> 的 <strong>任一</strong> 子数组，并将子数组中每个元素都 <strong>减去</strong> <code>1</code> 。</li>
</ul>

<p>如果你可以使数组中的所有元素都等于 <code>0</code> ，返回  <code>true</code><em> </em>；否则，返回<em> </em><code>false</code><em> </em>。</p>

<p><strong>子数组</strong> 是数组中的一个非空连续元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [2,2,3,1,1,0], k = 3
<strong>输出：</strong>true
<strong>解释：</strong>可以执行下述操作：
- 选出子数组 [2,2,3] ，执行操作后，数组变为 nums = [<em><strong>1</strong></em>,<em><strong>1</strong></em>,<em><strong>2</strong></em>,1,1,0] 。
- 选出子数组 [2,1,1] ，执行操作后，数组变为 nums = [1,1,<em><strong>1</strong></em>,<em><strong>0</strong></em>,<em><strong>0</strong></em>,0] 。
- 选出子数组 [1,1,1] ，执行操作后，数组变为 nums = [<em><strong>0</strong></em>,<em><strong>0</strong></em>,<em><strong>0</strong></em>,0,0,0] 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [1,3,1,1], k = 2
<strong>输出：</strong>false
<strong>解释：</strong>无法使数组中的所有元素等于 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>
 
#### 思路  

## 提示 1

想一想，如果 $\textit{nums}[0]>0$，那么必须要执行什么样的操作？

## 提示 2

对于 $\textit{nums}[0]>0$ 的情况，必须把 $\textit{nums}[0]$ 到 $\textit{nums}[k-1]$ 都减去 $\textit{nums}[0]$。  
然后思考 $\textit{nums}[1]$ 要怎么处理，依此类推。

## 提示 3

子数组同时加上/减去一个数，非常适合用**差分数组**来维护。  
设差分数组为 $d$。那么把 $\textit{nums}[i]$ 到 $\textit{nums}[i+k-1]$ 同时减去 $1$，等价于把 $d[i]$ 减 $1$，$d[i+k]$ 加 $1$。  
注意子数组长度必须恰好等于 $k$，所以当 $i+k\le n$ 时，才能执行上述操作。
遍历数组的同时，用变量 $\textit{sumD}$ 累加差分值。遍历到 $\textit{nums}[i]$ 时，$\textit{nums}[i]+\textit{sumD}$ 就是 $\textit{nums}[i]$ 的实际值了。  

分类讨论：  
- 如果 $\textit{nums}[i]<0$，由于无法让元素值增大，返回 `false`。
- 如果 $\textit{nums}[i]=0$，无需操作，遍历下一个数。
- 如果 $\textit{nums}[i]>0$：\r\n  - 如果 $i+k> n$，无法执行操作，所以 $\textit{nums}[i]$ 无法变成 $0$，返回 `false`。
- 如果 $i+k\le n$，按照上面说的执行操作，修改差分数组，遍历下一个数。 
  
如果遍历中途没有返回 `false`，那么最后返回 `true`。

```go 
func checkArray(nums []int, k int) bool {
	n := len(nums)
	d := make([]int, n+1)
	sumD := 0
	for i, x := range nums {
		sumD += d[i]
		x += sumD
		if x == 0 { // 无需操作
			continue
		}
		if x < 0 || i+k > n { // 无法操作
			return false
		}
		sumD -= x // 直接加到 sumD 中
		d[i+k] += x
	}
	return true
}

```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。