### 题目

公园规划小组为了让公园景观层次丰富，决定按以下方案对各花坛内的植物进行统一规划：
+ 一条小路两端的花坛不能栽种同一种花
+ 与同一花坛相连的两个花坛也不能栽种同一种花。

已知公园内有编号为 `0 ~ num-1`的若干花坛，任意两个花坛均可通过小路直接或间接到达。
公园中共有 `num-1` 条双向小路连接花坛，`roads[i] = [x, y]` 表示花坛 `x` 和花坛 `y` 之间存在小路将二者相连。

请返回这些花坛最少需要几种花。

**示例 1：**
>输入：`roads = [[0,1],[1,3],[1,2]]`
>
>输出：`4`
>
>解释：4 个花坛中编号为 1 的花坛与其余花坛均相连，因此至少需要栽种 4 种的花。

**示例 2：**
>输入：`roads = [[0,1],[0,2],[1,3],[2,5],[3,6],[5,4]]`
>
>输出：`3`

**提示：**
- `1 <= roads.length <= 10^5`
- `0 <= roads[i][0],roads[i][1] <= roads.length`

### 思路
看相邻的图最多有几条边，再算上自身求最大值

```go 
func numFlowers(roads [][]int) (ans int) {

	n := len(roads)
	g := make([][]int, n+1)
	for _, r := range roads {
		from, to := r[0], r[1]
		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}

	for i := range g {
		ans = max(ans, len(g[i]))
	}
	ans++
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。 $n$ 为 $\textit{roads}$ 的长度
- 空间复杂度：$\mathcal{O}(n)$。
