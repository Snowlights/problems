#### 题目  

<p>给你一个整数 <code>n</code> 。如果两个整数 <code>x</code> 和 <code>y</code> 满足下述条件，则认为二者形成一个质数对：</p>

<ul>
	<li><code>1 &lt;= x &lt;= y &lt;= n</code></li>
	<li><code>x + y == n</code></li>
	<li><code>x</code> 和 <code>y</code> 都是质数</li>
</ul>

<p>请你以二维有序列表的形式返回符合题目要求的所有 <code>[x<sub>i</sub>, y<sub>i</sub>]</code> ，列表需要按 <code>x<sub>i</sub></code> 的 <strong>非递减顺序</strong> 排序。如果不存在符合要求的质数对，则返回一个空数组。</p>

<p><strong>注意：</strong>质数是大于 <code>1</code> 的自然数，并且只有两个因子，即它本身和 <code>1</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>n = 10
<strong>输出：</strong>[[3,7],[5,5]]
<strong>解释：</strong>在这个例子中，存在满足条件的两个质数对。 
这两个质数对分别是 [3,7] 和 [5,5]，按照题面描述中的方式排序后返回。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>n = 2
<strong>输出：</strong>[]
<strong>解释：</strong>可以证明不存在和为 2 的质数对，所以返回一个空数组。 
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>6</sup></code></li>
</ul>
 
#### 思路  

暴力枚举质数 $x$ 和 $y=n-x$ 了，如果 $x\le y$ 且 $y$ 是质数，那么就把 $[x,y]$ 加入答案。

```go 
const mx = 1e6
var primes []int
var isP = [mx + 1]bool{}

func init() {
	for i := 2; i <= mx; i++ {
		isP[i] = true
	}
	for i := 2; i <= mx; i++ {
		if isP[i] {
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				isP[j] = false
			}
		}
	}
}

func findPrimePairs(n int) (ans [][]int) {
	for _, x := range primes {
		y := n - x
		if y < x {
			break
		}
		if isP[y] {
			ans = append(ans, []int{x, y})
		}
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}\left(\dfrac{n}{\log n}\right)$。$n$ 以内有 $\mathcal{O}\left(\dfrac{n}{\log n}\right)$ 个质数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。