#### 题目

<p>给你一个下标从 <strong>0 </strong>开始的整数数组 <code>coins</code>，表示可用的硬币的面值，以及一个整数 <code>target</code> 。</p>

<p>如果存在某个 <code>coins</code> 的子序列总和为 <code>x</code>，那么整数 <code>x</code> 就是一个 <strong>可取得的金额 </strong>。</p>

<p>返回需要添加到数组中的<strong> 任意面值 </strong>硬币的 <strong>最小数量 </strong>，使范围 <code>[1, target]</code> 内的每个整数都属于 <strong>可取得的金额</strong> 。</p>

<p>数组的 <strong>子序列</strong> 是通过删除原始数组的一些（<strong>可能不删除</strong>）元素而形成的新的 <strong>非空</strong> 数组，删除过程不会改变剩余元素的相对位置。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>coins = [1,4,10], target = 19
<strong>输出：</strong>2
<strong>解释：</strong>需要添加面值为 2 和 8 的硬币各一枚，得到硬币数组 [1,2,4,8,10] 。
可以证明从 1 到 19 的所有整数都可由数组中的硬币组合得到，且需要添加到数组中的硬币数目最小为 2 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>coins = [1,4,10,5,7,19], target = 19
<strong>输出：</strong>1
<strong>解释：</strong>只需要添加一枚面值为 2 的硬币，得到硬币数组 [1,2,4,5,7,10,19] 。
可以证明从 1 到 19 的所有整数都可由数组中的硬币组合得到，且需要添加到数组中的硬币数目最小为 1 。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>coins = [1,1,1], target = 20
<strong>输出：</strong>3
<strong>解释：</strong>
需要添加面值为 4 、8 和 16 的硬币各一枚，得到硬币数组 [1,1,1,4,8,16] 。 
可以证明从 1 到 20 的所有整数都可由数组中的硬币组合得到，且需要添加到数组中的硬币数目最小为 3 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= target &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= coins.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= coins[i] &lt;= target</code></li>
</ul>

#### 思路

为方便描述，把 $0$ 也算作可以得到的数。

假设现在得到了 $[0,s-1]$ 内的所有整数，如果此时新发现了一个整数 $x$，那么把 $x$ 加到已得到的数字中，就得到了 $[x,s+x-1]$ 内的所有整数。

分类讨论：

- 如果 $x \le s$，那么合并这两个区间，我们可以得到 $[0,s+x-1]$ 内的所有整数。
- 如果 $x > s$，这意味着我们无法得到 $s$，那么就一定要把 $s$ 加到数组中（加一个比 $s$ 还小的数字就没法得到更大的数，不够贪），这样就可以得到 $[s,2s-1]$ 内的所有整数，再与 $[0,s-1]$ 合并，可以得到 $[0,2s-1]$ 内的所有整数。然后再重新考虑 $x$ 和 $s$ 的大小关系，继续分类讨论。

把 $\textit{coins}$ 排序，从小到大考虑 $x=\textit{coins}[i]$。按照上述分类讨论来看是否要添加数字。

```go  
func minimumAddedCoins(coins []int, target int) (ans int) {
	slices.Sort(coins)
	s, i := 1, 0
	for s <= target {
		if i < len(coins) && coins[i] <= s {
			s += coins[i]
			i++
		} else {
			s *= 2
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + \log \textit{target})$，其中 $n$ 为 $\textit{coins}$ 的长度。$s$ 至多翻倍 $\mathcal{O}(\log \textit{target})$ 次。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。
