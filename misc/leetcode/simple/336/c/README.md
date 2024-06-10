#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组<code>nums</code> 。每次操作中，你可以：</p>

<ul>
	<li>选择两个满足 <code>0 &lt;= i, j &lt; nums.length</code> 的不同下标 <code>i</code> 和 <code>j</code> 。</li>
	<li>选择一个非负整数 <code>k</code> ，满足 <code>nums[i]</code> 和 <code>nums[j]</code> 在二进制下的第 <code>k</code> 位（下标编号从 <strong>0</strong> 开始）是 <code>1</code> 。</li>
	<li>将 <code>nums[i]</code> 和 <code>nums[j]</code> 都减去 <code>2<sup>k</sup></code> 。</li>
</ul>

<p>如果一个子数组内执行上述操作若干次后，该子数组可以变成一个全为 <code>0</code> 的数组，那么我们称它是一个 <strong>美丽</strong> 的子数组。</p>

<p>请你返回数组 <code>nums</code> 中 <strong>美丽子数组</strong> 的数目。</p>

<p>子数组是一个数组中一段连续 <strong>非空</strong> 的元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [4,3,1,2,4]
<b>输出：</b>2
<b>解释：</b>nums 中有 2 个美丽子数组：[4,<em><strong>3,1,2</strong></em>,4] 和 [<em><strong>4,3,1,2,4</strong></em>] 。
- 按照下述步骤，我们可以将子数组 [3,1,2] 中所有元素变成 0 ：
  - 选择 [<em><strong>3</strong></em>, 1, <em><strong>2</strong></em>] 和 k = 1 。将 2 个数字都减去 2<sup>1</sup> ，子数组变成 [1, 1, 0] 。
  - 选择 [<em><strong>1</strong></em>, <em><strong>1</strong></em>, 0] 和 k = 0 。将 2 个数字都减去 2<sup>0</sup> ，子数组变成 [0, 0, 0] 。
- 按照下述步骤，我们可以将子数组 [4,3,1,2,4] 中所有元素变成 0 ：
  - 选择 [<em><strong>4</strong></em>, 3, 1, 2, <em><strong>4</strong></em>] 和 k = 2 。将 2 个数字都减去 2<sup>2</sup> ，子数组变成 [0, 3, 1, 2, 0] 。
  - 选择 [0, <em><strong>3</strong></em>, <em><strong>1</strong></em>, 2, 0] 和 k = 0 。将 2 个数字都减去 2<sup>0</sup> ，子数组变成 [0, 2, 0, 2, 0] 。
  - 选择 [0, <em><strong>2</strong></em>, 0, <em><strong>2</strong></em>, 0] 和 k = 1 。将 2 个数字都减去 2<sup>1</sup> ，子数组变成 [0, 0, 0, 0, 0] 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [1,10,4]
<b>输出：</b>0
<b>解释：</b>nums 中没有任何美丽子数组。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>
 
#### 思路  

下文中 $\oplus$ 表示异或运算  
对于数组 $\textit{nums}$，定义它的前缀异或和 $\textit{s}[0]=0$，$\textit{s}[i+1] = \bigoplus\limits_{j=0}^{i}\textit{nums}[j]$。  
根据这个定义，有 $s[i+1]=s[i]\oplus\textit{nums}[i]$。  
例如 $\textit{nums}=[4,3,1,2,4]$，对应的前缀异或和数组为 $s=[0,4,7,6,4,0]$。  
通过前缀异或和，我们可以把**子数组的异或和转换成两个前缀异或和的异或**，即

$$
\bigoplus_{j=\textit{left}}^{\textit{right}}\textit{nums}[j] = \bigoplus\limits_{j=0}^{\textit{right}}\textit{nums}[j] \oplus \bigoplus\limits_{j=0}^{\textit{left}-1}\textit{nums}[j] = \textit{s}[\textit{right}+1]\oplus \textit{s}[\textit{left}]
$$

例如 $\textit{nums}$ 的子数组 $[3,1,2]$ 的异或和就可以用 $s[4]\oplus s[1]=4\oplus 4=0$ 算出来。
> 注：为方便计算，常用左闭右开区间 $[\textit{left},\textit{right})$ 来表示从 $\textit{nums}[\textit{left}]$ 到 $\textit{nums}[\textit{right}-1]$ 的子数组，此时子数组的异或和为 $\textit{s}[\textit{right}] \oplus \textit{s}[\textit{left}]$。
> 注 2：$s[0]=0$ 表示一个空数组的异或和。为什么要额外定义它？想一想，如果要计算的子数组恰好是一个前缀（从 $\textit{nums}[0]$ 开始），你要用 $s[\textit{right}]$ 异或谁呢？通过定义 $s[0]=0$，任意子数组（包括前缀）都可以表示为两个前缀异或和的异或。

#### 提示 1

由于每次操作的都是同一个比特位，可以把每一位单独看。

#### 提示 2

每次都去掉两个 $1$，要是美丽子数组，需要子数组内这个比特位的 $1$ 的个数是偶数。

#### 提示 3

由于 $1\oplus 1=0$，把所有比特位合起来看，美丽子数组这等价于子数组的异或和等于 $0$。

#### 提示 4

利用前缀异或和 $s$，问题相当于在求 $s$ 中有多少对 $s[\textit{left}]$ 和 $s[\textit{right}]$ 满足 $s[\textit{right}]\oplus s[\textit{left}] = 0$，即 $s[\textit{right}]= s[\textit{left}]$，因为异或为 $0$ 的两个数字必然相等。

也就是说，我们实际计算的是 $s$ 中有多少对相同数字。  
我们可以一边遍历 $s$，一边用一个哈希表 $\textit{cnt}$ 统计 $s[i]$ 的出现次数，累加遍历中的 $\textit{cnt}[s[i]]$，即为答案。

```go 
func beautifulSubarrays(nums []int) (ans int64) {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] ^ x
	}
	cnt := map[int]int{}
	for _, x := range s {
		// 先计入答案再统计个数，如果反过来的话，就相当于把空子数组也计入答案了
		ans += int64(cnt[x])
		cnt[x]++
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。