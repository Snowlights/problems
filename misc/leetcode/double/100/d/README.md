### 题目  

<p>给你一个整数数组 <code>ranks</code> ，表示一些机械工的 <strong>能力值</strong> 。<code>ranks<sub>i</sub></code> 是第 <code>i</code> 位机械工的能力值。能力值为 <code>r</code> 的机械工可以在 <code>r * n<sup>2</sup></code> 分钟内修好 <code>n</code> 辆车。</p>

<p>同时给你一个整数 <code>cars</code> ，表示总共需要修理的汽车数目。</p>

<p>请你返回修理所有汽车 <strong>最少</strong> 需要多少时间。</p>

<p><strong>注意：</strong>所有机械工可以同时修理汽车。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>ranks = [4,2,3,1], cars = 10
<b>输出：</b>16
<b>解释：</b>
- 第一位机械工修 2 辆车，需要 4 * 2 * 2 = 16 分钟。
- 第二位机械工修 2 辆车，需要 2 * 2 * 2 = 8 分钟。
- 第三位机械工修 2 辆车，需要 3 * 2 * 2 = 12 分钟。
- 第四位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
16 分钟是修理完所有车需要的最少时间。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>ranks = [5,1,8], cars = 6
<b>输出：</b>16
<b>解释：</b>
- 第一位机械工修 1 辆车，需要 5 * 1 * 1 = 5 分钟。
- 第二位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
- 第三位机械工修 1 辆车，需要 8 * 1 * 1 = 8 分钟。
16 分钟时修理完所有车需要的最少时间。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= ranks.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= ranks[i] &lt;= 100</code></li>
	<li><code>1 &lt;= cars &lt;= 10<sup>6</sup></code></li>
</ul>

### 思路

如果无法用 $t$ 分钟修完所有车，那么同样无法用小于 $t$ 分钟的时间修完所有车。有单调性，可以二分答案。
根据题意，需要满足

$$
rn^2 \le t

$$

解得

$$
n\le \sqrt\dfrac{t}{r}

$$

所以能力值为 $r$ 的机械工最多可以修

$$
\left\lfloor\sqrt\dfrac{t}{r}\right\rfloor

$$

辆车。
如果

$$
\sum_{i=0}^{i=n-1}\left\lfloor\sqrt\dfrac{t}{\textit{ranks}[i]}\right\rfloor \ge \textit{cars}

$$

则说明可以在 $t$ 分钟修理完所有汽车，根据这个公式来二分答案。
二分上界为 $\min(\textit{ranks}) \cdot \textit{cars}^2$，即让能力值最低的机械工修理所有汽车。

### 答疑

n**问**：开方直接取整的做法是否会有精度误差？
**答**：代码中对整数开方，只要整数转浮点没有丢失精度（在 $2^{53}-1$ 内），开方出来的整数部分就是正确的。具体可以参考 IEEE 754。

```go  
func repairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	cnt := [101]int{}
	for _, r := range ranks {
		if r < minR {
			minR = r
		}
		cnt[r]++
	}
	return int64(sort.Search(minR*cars*cars, func(t int) bool {
		s := 0
		for r := minR; r <= 100; r++ {
			s += int(math.Sqrt(float64(t/r))) * cnt[r]
		}
		return s >= cars
	}))
}
```

### 复杂度分析

- 时间复杂度：$O(n + U\log(mc^2))$，其中 $n$ 为 $\textit{ranks}$ 的长度，$U=\max(\textit{ranks})$，$m=\min(\textit{ranks})$，$c=\textit{cars}$。
- 空间复杂度：$O(U)$。