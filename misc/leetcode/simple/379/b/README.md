#### 题目

<p>现有一个下标从 <strong>0</strong> 开始的 <code>8 x 8</code> 棋盘，上面有 <code>3</code> 枚棋子。</p>

<p>给你 <code>6</code> 个整数 <code>a</code> 、<code>b</code> 、<code>c</code> 、<code>d</code> 、<code>e</code> 和 <code>f</code> ，其中：</p>

<ul>
	<li><code>(a, b)</code> 表示白色车的位置。</li>
	<li><code>(c, d)</code> 表示白色象的位置。</li>
	<li><code>(e, f)</code> 表示黑皇后的位置。</li>
</ul>

<p>假定你只能移动白色棋子，返回捕获黑皇后所需的<strong>最少</strong>移动次数。</p>

<p><strong>请注意</strong>：</p>

<ul>
	<li>车可以向垂直或水平方向移动任意数量的格子，但不能跳过其他棋子。</li>
	<li>象可以沿对角线方向移动任意数量的格子，但不能跳过其他棋子。</li>
	<li>如果车或象能移向皇后所在的格子，则认为它们可以捕获皇后。</li>
	<li>皇后不能移动。</li>
</ul>

<p> </p>

<p><strong class="example">示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/21/ex1.png" style="width: 600px; height: 600px; padding: 10px; background: #fff; border-radius: .5rem;" />
<pre>
<strong>输入：</strong>a = 1, b = 1, c = 8, d = 8, e = 2, f = 3
<strong>输出：</strong>2
<strong>解释：</strong>将白色车先移动到 (1, 3) ，然后移动到 (2, 3) 来捕获黑皇后，共需移动 2 次。
由于起始时没有任何棋子正在攻击黑皇后，要想捕获黑皇后，移动次数不可能少于 2 次。
</pre>

<p><strong class="example">示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/21/ex2.png" style="width: 600px; height: 600px;padding: 10px; background: #fff; border-radius: .5rem;" />
<pre>
<strong>输入：</strong>a = 5, b = 3, c = 3, d = 4, e = 5, f = 2
<strong>输出：</strong>1
<strong>解释：</strong>可以通过以下任一方式移动 1 次捕获黑皇后：
- 将白色车移动到 (5, 2) 。
- 将白色象移动到 (5, 2) 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= a, b, c, d, e, f <= 8</code></li>
	<li>两枚棋子不会同时出现在同一个格子上。</li>
</ul>

#### 思路

分类讨论：
- 如果车能直接攻击到皇后，答案是 $1$。
- 如果象能直接攻击到皇后，答案是 $1$。
- 如果车被象挡住，那么移走象，车就可以攻击到皇后，答案是 $2$。小知识：这在国际象棋中称作「闪击」。
- 如果象被车挡住，那么移走车，象就可以攻击到皇后，答案是 $2$。
- 如果车不能直接攻击到皇后，那么车可以水平移动或者垂直移动，其中一种方式必定不会被象挡住，可以攻击到皇后，答案是 $2$。

对于车，如果和皇后在同一水平线或者同一竖直线，且中间没有象，那么就可以直接攻击到皇后。
对于象，如果和皇后在同一斜线，且中间没有车，那么就可以直接攻击到皇后。

```go [sol]
func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	notInMiddle := func(l, m, r int) bool {
		return m < min(l, r) || m > max(r, l)
	}
	if (a == e && (c != e || notInMiddle(b, d, f))) ||
		(b == f && (d != f || notInMiddle(a, c, e))) ||
		(c+d == e+f && (a+b != e+f || notInMiddle(c, a, e))) ||
		(c-d == e-f && (a-b != e-f || notInMiddle(c, a, e))) {
		return 1
	}
	return 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
