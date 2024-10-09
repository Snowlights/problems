#### 题目

<p>有一个大型的 <code>(m - 1) x (n - 1)</code> 矩形田地，其两个对角分别是 <code>(1, 1)</code> 和 <code>(m, n)</code> ，田地内部有一些水平栅栏和垂直栅栏，分别由数组 <code>hFences</code> 和 <code>vFences</code> 给出。</p>

<p>水平栅栏为坐标 <code>(hFences[i], 1)</code> 到 <code>(hFences[i], n)</code>，垂直栅栏为坐标 <code>(1, vFences[i])</code> 到 <code>(m, vFences[i])</code> 。</p>

<p>返回通过<strong> 移除 </strong>一些栅栏（<strong>可能不移除</strong>）所能形成的最大面积的<strong> 正方形 </strong>田地的面积，或者如果无法形成正方形田地则返回 <code>-1</code>。</p>

<p>由于答案可能很大，所以请返回结果对 <code>10<sup>9</sup> + 7</code> <strong>取余</strong> 后的值。</p>

<p><strong>注意：</strong>田地外围两个水平栅栏（坐标 <code>(1, 1)</code> 到 <code>(1, n)</code> 和坐标 <code>(m, 1)</code> 到 <code>(m, n)</code> ）以及两个垂直栅栏（坐标 <code>(1, 1)</code> 到 <code>(m, 1)</code> 和坐标 <code>(1, n)</code> 到 <code>(m, n)</code> ）所包围。这些栅栏<strong> 不能</strong> 被移除。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/05/screenshot-from-2023-11-05-22-40-25.png" /></p>

<pre>
<strong>输入：</strong>m = 4, n = 3, hFences = [2,3], vFences = [2]
<strong>输出：</strong>4
<strong>解释：</strong>移除位于 2 的水平栅栏和位于 2 的垂直栅栏将得到一个面积为 4 的正方形田地。
</pre>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/22/maxsquareareaexample1.png" style="width: 285px; height: 242px;" /></p>

<pre>
<strong>输入：</strong>m = 6, n = 7, hFences = [2], vFences = [4]
<strong>输出：</strong>-1
<strong>解释：</strong>可以证明无法通过移除栅栏形成正方形田地。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= m, n &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= hFences.length, vFences.length &lt;= 600</code></li>
	<li><code>1 &lt; hFences[i] &lt; m</code></li>
	<li><code>1 &lt; vFences[i] &lt; n</code></li>
	<li><code>hFences</code> 和 <code>vFences</code> 中的元素是唯一的。</li>
</ul>

#### 思路

暴力枚举横向栅栏和纵向栅栏每两个栅栏之间的距离，计算最大值

```go  [sol]
func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	h, v := make(map[int]bool), make(map[int]bool)
	cal := func(arr []int, m map[int]bool) {
		for i, v := range arr {
			for _, vv := range arr[i+1:] {
				m[abs(vv-v)] = true
			}
		}
	}
	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)
	cal(hFences, h)
	cal(vFences, v)
	ans, mod := -1, 1_000_000_007
	for i := range h {
		if v[i] {
			ans = max(ans, i)
		}
	}
	if ans == -1 {
		return ans
	}
	return ans * ans % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(h^2 + v^2)$，其中 $h$ 为 $\textit{hFences}$ 的长度，其中 $v$ 为 $\textit{vFences}$ 的长度。
- 空间复杂度：$\mathcal{O}(h^2 + v^2)$。
