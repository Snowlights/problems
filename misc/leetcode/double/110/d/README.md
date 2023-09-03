### 题目

<p>给你两个长度相等下标从 <strong>0</strong> 开始的整数数组 <code>nums1</code> 和 <code>nums2</code> 。每一秒，对于所有下标 <code>0 <= i < nums1.length</code> ，<code>nums1[i]</code> 的值都增加 <code>nums2[i]</code> 。操作 <strong>完成后</strong> ，你可以进行如下操作：</p>

<ul>
	<li>选择任一满足 <code>0 <= i < nums1.length</code> 的下标 <code>i</code> ，并使 <code>nums1[i] = 0</code> 。</li>
</ul>

<p>同时给你一个整数 <code>x</code> 。</p>

<p>请你返回使 <code>nums1</code> 中所有元素之和 <strong>小于等于</strong> <code>x</code> 所需要的 <strong>最少</strong> 时间，如果无法实现，那么返回 <code>-1</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums1 = [1,2,3], nums2 = [1,2,3], x = 4
<b>输出：</b>3
<b>解释：</b>
第 1 秒，我们对 i = 0 进行操作，得到 nums1 = [0,2+2,3+3] = [0,4,6] 。
第 2 秒，我们对 i = 1 进行操作，得到 nums1 = [0+1,0,6+3] = [1,0,9] 。
第 3 秒，我们对 i = 2 进行操作，得到 nums1 = [1+1,0+2,0] = [2,2,0] 。
现在 nums1 的和为 4 。不存在更少次数的操作，所以我们返回 3 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums1 = [1,2,3], nums2 = [3,3,3], x = 4
<b>输出：</b>-1
<b>解释：</b>不管如何操作，nums1 的和总是会超过 x 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums1.length <= 10<sup>3</sup></code></li>
	<li><code>1 <= nums1[i] <= 10<sup>3</sup></code></li>
	<li><code>0 <= nums2[i] <= 10<sup>3</sup></code></li>
	<li><code>nums1.length == nums2.length</code></li>
	<li><code>0 <= x <= 10<sup>6</sup></code></li>
</ul>

### 思路

### 提示 1

每个下标 $i$ 至多操作一次。
因为操作多次的话，可以只保留最后一次，前面的操作是完全多余的（反而会让其它数字变得更大）。
所以至多操作 $n$ 次。
试试枚举答案。

### 提示 2

考虑第 $t$ 秒元素之和最小是多少。
如果从一开始到第 $t$ 秒都不做任何操作的话，元素之和等于

$$
s_1 + s_2\cdot t

$$

其中 $s_1$ 是 $\textit{nums}_1$ 的元素和，$s_2$ 是 $\textit{nums}_2$ 的元素和。
考虑如何**分配**这 $t$ 次操作，让数组元素分别**减少**多少。用 $s_1 + s_2\cdot t$ 减去这些元素减少量之和的最大值，就是第 $t$ 秒元素之和的最小值。

### 提示 3

假设已经选好了要操作的元素，那么 $\textit{num}_2[i]$ 越大，操作的时间就应该越靠后。
例如 $t=3$，假设选的三个数的下标分别是 $0,1,2$，且 $\textit{nums}_2[0]\le\textit{nums}_2[1]\le\textit{nums}_2[2]$。我们可以让这些数分别减少

$$
\textit{nums}_1[0] + \textit{nums}_2[0] \cdot x \\
\textit{nums}_1[1] + \textit{nums}_2[1] \cdot y \\
\textit{nums}_1[2] + \textit{nums}_2[2] \cdot z

$$

根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728)，上式中的 $x,y,z$ 应分别取 $1,2,3$，分别对应在第 $1,2,3$ 秒操作，能让第 $3$ 秒的 $s_1 + s_2\cdot t$ 减少多少。

### 提示 4

在第 $t$ 秒，$s_1 + s_2\cdot t$ 的**减少量的最大值**相当于求解如下问题：
按照 $\textit{nums}_2[i]$ 从小到大排序后，从 $\textit{nums}_1$ 中选择一个长为 $t$ 的子序列，子序列第 $j$ 个数的 $\textit{nums}_2[i]$ 的系数为 $j$，计算减少量的最大值。
设子序列第 $j$ 个数（$j$ 从 $1$ 开始）的下标为 $i$，那么它对减少量的贡献是

$$
\textit{nums}_1[i] + \textit{nums}_2[i] \cdot j

$$

由于上式中 $\textit{nums}_1$ 并不是有序的，无法贪心，考虑用动态规划求解。定义 $f[i+1][j]$ 表示从前 $i$ 个数中选出 $j$ 个数，减少量最大是多少。考虑第 $i$ 个数「选或不选」：

- 不选：$f[i+1][j] = f[i][j]$。
- 选：$f[i+1][j] = f[i][j-1] + \textit{nums}_1[i] + \textit{nums}_2[i] \cdot j$。
- 这两种情况取最大值，即

$$
f[i+1][j] = max(f[i][j], f[i][j-1] + \textit{nums}_1[i] + \textit{nums}_2[i] \cdot j)

$$

初始值 $f[0][j]=0$。
答案就是第一个满足

$$
s_1 + s_2\cdot t - f[t] \le x

$$

的 $t$。如果不存在，则返回 $-1$。
代码实现时，第一个维度可以优化掉

```go  
func minimumTime(a []int, b []int, x int) (ans int) {

	type pair struct {
		x, y int
	}
	pList := make([]*pair, 0, len(a))
	suma, sumb := 0, 0
	for i := range a {
		pList = append(pList, &pair{
			x: a[i],
			y: b[i],
		})
		suma += a[i]
		sumb += b[i]
	}
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].y < pList[j].y ||
			pList[i].y == pList[j].y && pList[i].x < pList[j].x
	})

	f := make([]int, len(a)+1)
	for _, p := range pList {
		for i := len(a) - 1; i >= 0; i-- {
			f[i+1] = max(f[i+1], f[i]+(i+1)*p.y+p.x)
		}
	}

	for i := 0; i <= len(a); i++ {
		if sumb*i+suma-f[i] <= x {
			return i
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
