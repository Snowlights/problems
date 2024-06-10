#### 题目

<p>给你一个下标从 <strong>0 </strong>开始的二维数组 <code>variables</code> ，其中 <code>variables[i] = [a<sub>i</sub>, b<sub>i</sub>, c<sub>i,</sub> m<sub>i</sub>]</code>，以及一个整数 <code>target</code> 。</p>

<p>如果满足以下公式，则下标 <code>i</code> 是 <strong>好下标</strong>：</p>

<ul>
	<li><code>0 <= i < variables.length</code></li>
	<li><code>((a<sub>i</sub><sup>b<sub>i</sub></sup> % 10)<sup>c<sub>i</sub></sup>) % m<sub>i</sub> == target</code></li>
</ul>

<p>返回一个由<strong> 好下标 </strong>组成的数组，<strong>顺序不限</strong> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>variables = [[2,3,3,10],[3,3,3,1],[6,1,1,4]], target = 2
<strong>输出：</strong>[0,2]
<strong>解释：</strong>对于 variables 数组中的每个下标 i ：
1) 对于下标 0 ，variables[0] = [2,3,3,10] ，(2<sup>3</sup> % 10)<sup>3</sup> % 10 = 2 。
2) 对于下标 1 ，variables[1] = [3,3,3,1] ，(3<sup>3</sup> % 10)<sup>3</sup> % 1 = 0 。
3) 对于下标 2 ，variables[2] = [6,1,1,4] ，(6<sup>1</sup> % 10)<sup>1</sup> % 4 = 2 。
因此，返回 [0,2] 作为答案。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>variables = [[39,3,1000,1000]], target = 17
<strong>输出：</strong>[]
<strong>解释：</strong>对于 variables 数组中的每个下标 i ：
1) 对于下标 0 ，variables[0] = [39,3,1000,1000] ，(39<sup>3</sup> % 10)<sup>1000</sup> % 1000 = 1 。
因此，返回 [] 作为答案。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= variables.length <= 100</code></li>
	<li><code>variables[i] == [a<sub>i</sub>, b<sub>i</sub>, c<sub>i</sub>, m<sub>i</sub>]</code></li>
	<li><code>1 <= a<sub>i</sub>, b<sub>i</sub>, c<sub>i</sub>, m<sub>i</sub> <= 10<sup>3</sup></code></li>
	<li><code><font face="monospace">0 <= target <= 10<sup>3</sup></font></code></li>
</ul>

#### 思路

模拟+快速幂

```go  
func getGoodIndices(variables [][]int, target int) []int {
	ans := make([]int, 0, len(variables))
	pow := func(x, n, m int) int {
		//x %= mod
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % m
			}
			x = x * x % m
		}
		return res
	}
	for i, v := range variables {
		if pow(pow(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{variables}$ 的长度，$U$ 为 $b_i$ 和 $c_i$ 的最大值，本题为 $10^3$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。
