#### 题目  

<p>卡车有两个油箱。给你两个整数，<code>mainTank</code> 表示主油箱中的燃料（以升为单位），<code>additionalTank</code> 表示副油箱中的燃料（以升为单位）。</p>

<p>该卡车每耗费 <code>1</code> 升燃料都可以行驶 <code>10</code> km。每当主油箱使用了 <code>5</code> 升燃料时，如果副油箱至少有 <code>1</code> 升燃料，则会将 <code>1</code> 升燃料从副油箱转移到主油箱。</p>

<p>返回卡车可以行驶的最大距离。</p>

<p>注意：从副油箱向主油箱注入燃料不是连续行为。这一事件会在每消耗 <code>5</code> 升燃料时突然且立即发生。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>mainTank = 5, additionalTank = 10
<strong>输出：</strong>60
<strong>解释：</strong>
在用掉 5 升燃料后，主油箱中燃料还剩下 (5 - 5 + 1) = 1 升，行驶距离为 50km 。
在用掉剩下的 1 升燃料后，没有新的燃料注入到主油箱中，主油箱变为空。
总行驶距离为 60km 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>mainTank = 1, additionalTank = 2
<strong>输出：</strong>10
<strong>解释：</strong>
在用掉 1 升燃料后，主油箱变为空。
总行驶距离为 10km 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= mainTank, additionalTank &lt;= 100</code></li>
</ul>
 
#### 思路  

如果 $\textit{mainTank}$ 有 $10^9$，那么方法一会超时。有没有更快的做法呢？

把方法一的减法改成除法，统计 `-=5` 发生了 $t$ 次。然后再一次性地把 $t$ 升注入主油箱。注意 $t$ 不能超过 $\textit{additionalTank}$。

```go 
func distanceTraveled(mainTank int, additionalTank int) (ans int) {

	for mainTank >= 5 {
		run := mainTank / 5
		ans += 5 * 10 * run
		mainTank -= 5 * run
		add := min(run, additionalTank)
		additionalTank -= add
		mainTank += add
	}

	return ans + mainTank*10
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(\log\textit{mainTank})$。每次循环 $\textit{mainTank}$ 至少减为原来的 $\dfrac{1}{4}$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
