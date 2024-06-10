#### 题目  

<p>现有 <code>n</code> 个机器人，编号从 <strong>1</strong> 开始，每个机器人包含在路线上的位置、健康度和移动方向。</p>

<p>给你下标从 <strong>0</strong> 开始的两个整数数组 <code>positions</code>、<code>healths</code> 和一个字符串 <code>directions</code>（<code>directions[i]</code> 为 <strong>&#39;L&#39;</strong> 表示 <strong>向左</strong> 或 <strong>&#39;R&#39;</strong> 表示 <strong>向右</strong>）。 <code>positions</code> 中的所有整数 <strong>互不相同</strong> 。</p>

<p>所有机器人以 <strong>相同速度</strong> <strong>同时</strong> 沿给定方向在路线上移动。如果两个机器人移动到相同位置，则会发生 <strong>碰撞</strong> 。</p>

<p>如果两个机器人发生碰撞，则将 <strong>健康度较低</strong> 的机器人从路线中 <strong>移除</strong> ，并且另一个机器人的健康度 <strong>减少 1</strong> 。幸存下来的机器人将会继续沿着与之前 <strong>相同</strong> 的方向前进。如果两个机器人的健康度相同，则将二者都从路线中移除。</p>

<p>请你确定全部碰撞后幸存下的所有机器人的 <strong>健康度</strong> ，并按照原来机器人编号的顺序排列。即机器人 1 （如果幸存）的最终健康度，机器人 2 （如果幸存）的最终健康度等。 如果不存在幸存的机器人，则返回空数组。</p>

<p>在不再发生任何碰撞后，请你以数组形式，返回所有剩余机器人的健康度（按机器人输入中的编号顺序）。</p>

<p><strong>注意：</strong>位置  <code>positions</code> 可能是乱序的。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img height="169" src="https://assets.leetcode.com/uploads/2023/05/15/image-20230516011718-12.png" width="808"/></p>

<pre><strong>输入：</strong>positions = [5,4,3,2,1], healths = [2,17,9,15,10], directions = &#34;RRRRR&#34;
<strong>输出：</strong>[2,17,9,15,10]
<strong>解释：</strong>在本例中不存在碰撞，因为所有机器人向同一方向移动。所以，从第一个机器人开始依序返回健康度，[2, 17, 9, 15, 10] 。
</pre>

<p><strong>示例 2：</strong></p>

<p><img height="176" src="https://assets.leetcode.com/uploads/2023/05/15/image-20230516004433-7.png" width="717"/></p>

<pre><strong>输入：</strong>positions = [3,5,2,6], healths = [10,10,15,12], directions = &#34;RLRL&#34;
<strong>输出：</strong>[14]
<strong>解释：</strong>本例中发生 2 次碰撞。首先，机器人 1 和机器人 2 将会碰撞，因为二者健康度相同，二者都将被从路线中移除。接下来，机器人 3 和机器人 4 将会发生碰撞，由于机器人 4 的健康度更小，则它会被移除，而机器人 3 的健康度变为 15 - 1 = 14 。仅剩机器人 3 ，所以返回 [14] 。
</pre>

<p><strong>示例 3：</strong></p>

<p><img height="172" src="https://assets.leetcode.com/uploads/2023/05/15/image-20230516005114-9.png" width="732"/></p>

<pre><strong>输入：</strong>positions = [1,2,5,6], healths = [10,10,11,11], directions = &#34;RLRL&#34;
<strong>输出：</strong>[]
<strong>解释：</strong>机器人 1 和机器人 2 将会碰撞，因为二者健康度相同，二者都将被从路线中移除。机器人 3 和机器人 4 将会碰撞，因为二者健康度相同，二者都将被从路线中移除。所以返回空数组 [] 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= positions.length == healths.length == directions.length == n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= positions[i], healths[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>directions[i] == &#39;L&#39;</code> 或 <code>directions[i] == &#39;R&#39;</code></li>
	<li><code>positions</code> 中的所有值互不相同</li>
</ul>
 
#### 思路  

本题思路和 [735. 行星碰撞](https://leetcode.cn/problems/asteroid-collision/) 是一样的。
用列表 $\textit{toLeft}$ 维护向左的机器人，用栈 $\textit{st}$ 维护向右的机器人。
按照 $\textit{positions}[i]$ 从小到大排序。遍历机器人，如果遇到一个向左的机器人 $p$，分类讨论：
- 如果 $p$ 的健康度小于栈顶，那么栈顶的健康度减一。
- 如果 $p$ 的健康度等于栈顶，那么移除栈顶。
- 如果 $p$ 的健康度大于栈顶，那么移除栈顶，将 $p$ 的健康度减一后加入 $\textit{toLeft}$。
- 请注意，如果健康度减一，那么在减一之前，健康度一定是大于 $1$ 的，**不存在健康度减为** $0$ **的情况**。
  最后剩余的就是 $\textit{toLeft}$ 和 $\textit{st}$ 中的机器人，合并，按照编号排序后，返回健康度列表。

```go  
func survivedRobotsHealths(positions []int, healths []int, directions string) (ans []int) {
	type pair struct {
		p, h, i int
		d       byte
	}
	const (
		Left  = 'L'
		Right = 'R'
	)

	pList := make([]*pair, 0, len(positions))
	for i, p := range positions {
		pList = append(pList, &pair{
			p: p,
			h: healths[i],
			i: i,
			d: directions[i],
		})
	}
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].p < pList[j].p
	})

	stack := []*pair{}
	for _, p := range pList {
		if len(stack) == 0 || stack[len(stack)-1].d == p.d {
			stack = append(stack, p)
			continue
		}

		if p.d == Left {
			for len(stack) > 0 && stack[len(stack)-1].d == Right {
				if stack[len(stack)-1].h == p.h {
					stack = stack[:len(stack)-1]
					p.h = 0
					break
				}
				if stack[len(stack)-1].h > p.h {
					stack[len(stack)-1].h--
					p.h = 0
					break
				} else {
					stack = stack[:len(stack)-1]
					p.h--
					if p.h == 0 {
						break
					}
				}
			}
		}
		if p.h > 0 {
			stack = append(stack, p)
		}
	}

	sort.Slice(stack, func(i, j int) bool {
		return stack[i].i < stack[j].i
	})
	for _, v := range stack {
		ans = append(ans, v.h)
	}

	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{positions}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。