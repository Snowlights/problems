#### 题目  

<p>给你一个 <strong>正</strong> 整数 <code>n</code> 。</p>

<p>用 <code>even</code> 表示在 <code>n</code> 的二进制形式（下标从 <strong>0</strong> 开始）中值为 <code>1</code> 的偶数下标的个数。</p>

<p>用 <code>odd</code> 表示在 <code>n</code> 的二进制形式（下标从 <strong>0</strong> 开始）中值为 <code>1</code> 的奇数下标的个数。</p>

<p>返回整数数组<em> </em><code>answer</code><em> </em>，其中<em> </em><code>answer = [even, odd]</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>n = 17
<strong>输出：</strong>[2,0]
<strong>解释：</strong>17 的二进制形式是 10001 。 
下标 0 和 下标 4 对应的值为 1 。 
共有 2 个偶数下标，0 个奇数下标。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>n = 2
<strong>输出：</strong>[0,1]
<strong>解释：</strong>2 的二进制形式是 10 。 
下标 1 对应的值为 1 。 
共有 0 个偶数下标，1 个奇数下标。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 1000</code></li>
</ul>
 
#### 思路  

# 方法一：二进制基本操作

不断取最低位，然后右移，直到等于 $0$ 为止，这样可以取到每个比特位。

```go 
func evenOddBit(n int) []int {
	ans := make([]int, 2)
	for i := 0; n > 0; i ^= 1 {
		ans[i] += n & 1
		n >>= 1
	}
	return ans
}
```

#### 复杂度分析  

- 时间复杂度：$O(\log n)$。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

# 方法二：位掩码 + 库函数

利用位掩码 `0x55555555`（二进制的 $010101\cdots$），取出偶数下标比特和奇数下标比特，分别用库函数统计 $1$ 的个数。  
本题由于 $n$ 范围比较小，取 `0x5555` 作为位掩码。

```go  
func evenOddBit(n int) []int {
	const mask = 0x5555
	return []int{bits.OnesCount16(uint16(n & mask)), bits.OnesCount16(uint16(n & (mask >> 1)))}
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。仅用到若干额外变量。