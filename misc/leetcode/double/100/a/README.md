### 题目

<p>给你一个整数 <code>money</code> ，表示你总共有的钱数（单位为美元）和另一个整数 <code>children</code> ，表示你要将钱分配给多少个儿童。</p>

<p>你需要按照如下规则分配：</p>

<ul>
	<li>所有的钱都必须被分配。</li>
	<li>每个儿童至少获得 <code>1</code> 美元。</li>
	<li>没有人获得 <code>4</code> 美元。</li>
</ul>

<p>请你按照上述规则分配金钱，并返回 <strong>最多</strong> 有多少个儿童获得 <strong>恰好</strong><em> </em><code>8</code> 美元。如果没有任何分配方案，返回 <code>-1</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>money = 20, children = 3
<b>输出：</b>1
<b>解释：</b>
最多获得 8 美元的儿童数为 1 。一种分配方案为：
- 给第一个儿童分配 8 美元。
- 给第二个儿童分配 9 美元。
- 给第三个儿童分配 3 美元。
没有分配方案能让获得 8 美元的儿童数超过 1 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>money = 16, children = 2
<b>输出：</b>2
<b>解释：</b>每个儿童都可以获得 8 美元。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= money <= 200</code></li>
	<li><code>2 <= children <= 30</code></li>
</ul>

### 思路

首先，每人至少分配 $1$ 美元，把 $\textit{money}$ 减少 $\textit{children}$。
如果 $\textit{money}<0$，返回 $-1$。
然后不断给每个人 $7$ 美元（前面分配了 $1$ 美元），这样可以分给至多

$$
\textit{ans}=\min\left(\left\lfloor\dfrac{\textit{money}}{7}\right\rfloor,\textit{children}\right)

$$

个人。然后更新剩余 $\textit{money}$ 和剩余未分配的人数。

最后，分类讨论：

- 如果剩余 $0$ 人，且 $\textit{money}>0$，那么必须分给一个已经分到 $8$ 美元的人，$\textit{ans}$ 减一。
- 如果剩余 $1$ 人，且 $\textit{money}=3$，为避免分配 $4$ 美元，那么必须分给一个已经分到 $8$ 美元的人，$\textit{ans}$ 减一。（注意输入的 $\textit{children}\ge 2$）
- 其它情况全部给一个人，如果这个人分配到 $4$ 美元，他再给另一个人 $1$ 美元，这样 $\textit{ans}$ 不变。

```go  
func distMoney(money, children int) int {
	money -= children // 每人至少 1 美元
	if money < 0 {
		return -1
	}
	ans := min(money/7, children) // 初步分配，让尽量多的人分到 8 美元
	money -= ans * 7
	children -= ans
	if children == 0 && money > 0 || // 必须找一个前面分了 8 美元的人，分完剩余的钱
		children == 1 && money == 3 { // 不能有人恰好分到 4 美元
		ans--
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
