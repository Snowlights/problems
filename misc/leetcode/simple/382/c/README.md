#### 题目

<p>Alice 和 Bob 在一个长满鲜花的环形草地玩一个回合制游戏。环形的草地上有一些鲜花，Alice 到 Bob 之间顺时针有 <code>x</code> 朵鲜花，逆时针有 <code>y</code> 朵鲜花。</p>

<p>游戏过程如下：</p>

<ol>
	<li>Alice 先行动。</li>
	<li>每一次行动中，当前玩家必须选择顺时针或者逆时针，然后在这个方向上摘一朵鲜花。</li>
	<li>一次行动结束后，如果所有鲜花都被摘完了，那么 <strong>当前</strong> 玩家抓住对手并赢得游戏的胜利。</li>
</ol>

<p>给你两个整数 <code>n</code> 和 <code>m</code> ，你的任务是求出满足以下条件的所有 <code>(x, y)</code> 对：</p>

<ul>
	<li>按照上述规则，Alice 必须赢得游戏。</li>
	<li>Alice 顺时针方向上的鲜花数目 <code>x</code> 必须在区间 <code>[1,n]</code> 之间。</li>
	<li>Alice 逆时针方向上的鲜花数目 <code>y</code> 必须在区间 <code>[1,m]</code> 之间。</li>
</ul>

<p>请你返回满足题目描述的数对 <code>(x, y)</code> 的数目。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>n = 3, m = 2
<b>输出：</b>3
<b>解释：</b>以下数对满足题目要求：(1,2) ，(3,2) ，(2,1) 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>n = 1, m = 1
<b>输出：</b>0
<b>解释：</b>没有数对满足题目要求。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n, m <= 10<sup>5</sup></code></li>
</ul>

#### 思路

由于每回合玩家恰好摘掉一朵花，所以当且仅当 $x+y$ 是奇数，Alice 必胜。

- 如果 $x$ 是奇数，那么 $y$ 必须是偶数。
- 如果 $x$ 是偶数，那么 $y$ 必须是奇数。

$[1,n]$ 中有 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个偶数，$\left\lceil\dfrac{n}{2}\right\rceil$ 个奇数。
对于 $[1,m]$ 也同理。所以答案为

$$
\left\lfloor\dfrac{n}{2}\right\rfloor\cdot \left\lceil\dfrac{m}{2}\right\rceil + \left\lfloor\dfrac{m}{2}\right\rfloor\cdot \left\lceil\dfrac{n}{2}\right\rceil

$$

进一步优化：

- 当 $y=1,3,5,cdots$ 时，可以选 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个 $x$。
- 当 $y=2,4,6,cdots$ 时，可以选 $\left\lceil\dfrac{n}{2}\right\rceil$ 个 $x$。

两个两个一组，把 $y=1$ 和 $y=2$ 加起来，可以选 $\left\lfloor\dfrac{n}{2}\right\rfloor + \left\lceil\dfrac{n}{2}\right\rceil = n$ 个 $x$。
依此类推。

- 如果 $m$ 是偶数，那么答案就是 $\dfrac{nm}{2}$。
- 如果 $m$ 是奇数，当 $y=m$ 时，可以选 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个 $x$，所以答案是 $\left\lfloor\dfrac{nm}{2}\right\rfloor$。

综合这两种情况，答案是

$$
\left\lfloor\dfrac{nm}{2}\right\rfloor

$$

```go [sol]
func flowerGame(n int, m int) int64 {
	return int64(m * n) / 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$O(1)$。
