### 题目

<p>给你一个  <code>n x 2</code> 的二维数组 <code>points</code> ，它表示二维平面上的一些点坐标，其中 <code>points[i] = [x<sub>i</sub>, y<sub>i</sub>]</code> 。</p>

<p>我们定义 x 轴的正方向为 <strong>右</strong> （<strong>x 轴递增的方向</strong>），x 轴的负方向为 <strong>左</strong> （<strong>x 轴递减的方向</strong>）。类似的，我们定义 y 轴的正方向为 <strong>上</strong> （<strong>y 轴递增的方向</strong>），y 轴的负方向为 <strong>下</strong> （<strong>y 轴递减的方向</strong>）。</p>

<p>你需要安排这 <code>n</code> 个人的站位，这 <code>n</code> 个人中包括 liupengsay 和小羊肖恩 。你需要确保每个点处 <strong>恰好</strong> 有 <strong>一个</strong> 人。同时，liupengsay 想跟小羊肖恩单独玩耍，所以 liupengsay 会以 liupengsay<b> </b>的坐标为 <strong>左上角</strong> ，小羊肖恩的坐标为 <strong>右下角</strong> 建立一个矩形的围栏（<strong>注意</strong>，围栏可能 <strong>不</strong> 包含任何区域，也就是说围栏可能是一条线段）。如果围栏的 <strong>内部</strong> 或者 <strong>边缘</strong> 上有任何其他人，liupengsay 都会难过。</p>

<p>请你在确保 liupengsay <strong>不会</strong> 难过的前提下，返回 liupengsay 和小羊肖恩可以选择的 <strong>点对</strong> 数目。</p>

<p><b>注意</b>，liupengsay 建立的围栏必须确保 liupengsay 的位置是矩形的左上角，小羊肖恩的位置是矩形的右下角。比方说，以 <code>(1, 1)</code> ，<code>(1, 3)</code> ，<code>(3, 1)</code> 和 <code>(3, 3)</code> 为矩形的四个角，给定下图的两个输入，liupengsay 都不能建立围栏，原因如下：</p>

<ul>
	<li>图一中，liupengsay 在 <code>(3, 3)</code> 且小羊肖恩在 <code>(1, 1)</code> ，liupengsay 的位置不是左上角且小羊肖恩的位置不是右下角。</li>
	<li>图二中，liupengsay 在 <code>(1, 3)</code> 且小羊肖恩在 <code>(1, 1)</code> ，小羊肖恩的位置不是在围栏的右下角。</li>
</ul>
<img alt="" src="https://assets.leetcode.com/uploads/2024/01/04/example0alicebob-1.png" style="width: 750px; height: 308px;padding: 10px; background: #fff; border-radius: .5rem;" />
<p> </p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/01/04/example1alicebob.png" style="width: 376px; height: 308px; padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem;" /></p>

<pre>
<b>输入：</b>points = [[1,1],[2,2],[3,3]]
<b>输出：</b>0
<strong>解释：</strong>没有办法可以让 liupengsay 的围栏以 liupengsay 的位置为左上角且小羊肖恩的位置为右下角。所以我们返回 0 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<p><strong class="example"><a href="https://pic.leetcode.cn/1706880313-YelabI-example2.jpeg"><img alt="" src="https://pic.leetcode.cn/1706880313-YelabI-example2.jpeg" style="width: 900px; height: 247px;" /></a></strong></p>

<pre>
<b>输入：</b>points = [[6,2],[4,4],[2,6]]
<b>输出：</b>2
<b>解释：</b>总共有 2 种方案安排 liupengsay 和小羊肖恩的位置，使得 liupengsay 不会难过：
- liupengsay 站在 (4, 4) ，小羊肖恩站在 (6, 2) 。
- liupengsay 站在 (2, 6) ，小羊肖恩站在 (4, 4) 。
不能安排 liupengsay 站在 (2, 6) 且小羊肖恩站在 (6, 2) ，因为站在 (4, 4) 的人处于围栏内。
</pre>

<p><strong class="example">示例 3：</strong></p>

<p><strong class="example"><a href="https://pic.leetcode.cn/1706880311-mtPGYC-example3.jpeg"><img alt="" src="https://pic.leetcode.cn/1706880311-mtPGYC-example3.jpeg" style="width: 900px; height: 247px;" /></a></strong></p>

<pre>
<b>输入：</b>points = [[3,1],[1,3],[1,1]]
<b>输出：</b>2
<b>解释：</b>总共有 2 种方案安排 liupengsay 和小羊肖恩的位置，使得 liupengsay 不会难过：
- liupengsay 站在 (1, 1) ，小羊肖恩站在 (3, 1) 。
- liupengsay 站在 (1, 3) ，小羊肖恩站在 (1, 1) 。
不能安排 liupengsay 站在 (1, 3) 且小羊肖恩站在 (3, 1) ，因为站在 (1, 1) 的人处于围栏内。
注意围栏是可以不包含任何面积的，上图中第一和第二个围栏都是合法的。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= n <= 50</code></li>
	<li><code>points[i].length == 2</code></li>
	<li><code>0 <= points[i][0], points[i][1] <= 50</code></li>
	<li><code>points[i]</code> 点对两两不同。</li>
</ul>

### 思路

将 $\textit{points}$ 按照横坐标**从小到大**排序，横坐标相同的，按照纵坐标**从大到小**排序。如此一来，在枚举 $\textit{points}[i]$ 和 $\textit{points}[j]$ 时（$i<j$），就只需要关心纵坐标的大小。固定 $\textit{points}[i]$，然后枚举 $\textit{points}[j]$：

- 如果 $\textit{points}[j][1]$ 比之前枚举的点的纵坐标都大，那么矩形内没有其它点，符合要求，答案加一。
- 如果 $\textit{points}[j][1]$ 小于等于之前枚举的某个点的纵坐标，那么矩形内有其它点，不符合要求。

所以在枚举 $\textit{points}[j]$ 的同时，需要维护纵坐标的最大值 $\textit{maxY}$。这也解释了为什么横坐标相同的，
按照纵坐标**从大到小**排序。这保证了横坐标相同时，我们总是优先枚举更靠上的点，不会误把包含其它点的矩形也当作符合要求的矩形。

```go [sol]
func numberOfPairs(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] > points[j][1]
	})
	for i, p := range points {
		y0 := p[1]
		maxY := math.MinInt
		for _, q := range points[i+1:] {
			y := q[1]
			if y <= y0 && y > maxY {
				maxY = y
				ans++
			}
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
