### 题目

<p>给你一个下标从 <strong>0</strong>&nbsp;开始只包含 <strong>正</strong>&nbsp;整数的数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>一开始，你可以将数组中 <strong>任意数量</strong> 元素增加 <strong>至多</strong> <code>1</code> 。</p>

<p>修改后，你可以从最终数组中选择 <strong>一个或者更多</strong>&nbsp;元素，并确保这些元素升序排序后是 <strong>连续</strong>&nbsp;的。比方说，<code>[3, 4, 5]</code> 是连续的，但是&nbsp;<code>[3, 4, 6]</code> 和&nbsp;<code>[1, 1, 2, 3]</code>&nbsp;不是连续的。<!-- notionvc: 312f8c5d-40d0-4cd1-96cc-9e96a846735b --></p>

<p>请你返回 <strong>最多</strong>&nbsp;可以选出的元素数目。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [2,1,5,1,1]
<b>输出：</b>3
<b>解释：</b>我们将下标 0 和 3 处的元素增加 1 ，得到结果数组 nums = [3,1,5,2,1] 。
我们选择元素 [<em><strong>3</strong></em>,<em><strong>1</strong></em>,5,<em><strong>2</strong></em>,1] 并将它们排序得到 [1,2,3] ，是连续元素。
最多可以得到 3 个连续元素。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [1,4,7,10]
<b>输出：</b>1
<b>解释：</b>我们可以选择的最多元素数目是 1 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>


### 思路

## 提示 1

把数组排序。   
为什么？设我们选的元素在排序后为 $b$，那么有 $b[i] + 1 = b[i+1]$，这意味着 $b$ 中元素在操作前，必然有 $b[i] \le b[i+1]$。反证：如果操作前 $b[i] > b[i+1]$，那么操作后 $b[i]$ 至多和 $b[i+1]$ 相等，不会出现 $b[i+1]$ 比 $b[i]$ 多 $1$ 的情况。所以可以排序。

## 提示 2

排序后，我们选的是 $\textit{nums}$ 中的一个**子序列**。  
定义 $f[x]$ 表示子序列的最后一个数是 $x$ 时，子序列的最大长度。  
从左到右遍历数组 $x = \textit{nums}[i]$：
- 如果操作，那么 $x+1$ 可以接在末尾为 $x$ 的子序列后面，即 $f[x+1] = f[x] + 1$。
- 如果不操作，那么 $x$ 可以接在末尾为 $x-1$ 的子序列后面，即 $f[x] = f[x-1] + 1$。

比如 $\textit{nums} = [1,2,2]$：
- 遍历到 $\textit{nums}[0]=1$ 时，$f[2]=1,\ f[1]=1$。
- 遍历到 $\textit{nums}[1]=2$ 时，$f[3]=f[2]+1=2,\ f[2]=f[1]+1=2$。注意要先计算 $f[x+1]$ 再计算 $f[x]$（不然这里会算出 $f[3]=3$）。此时 $f[1]$ 还是 $1$。
- 遍历到 $\textit{nums}[2]=2$ 时，$f[3]=f[2]+1=3,\ f[2]=f[1]+1=2$。此时 $f[1]$ 还是 $1$。

最后返回 $f[x]$ 的最大值。

```go [sol]

```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。
