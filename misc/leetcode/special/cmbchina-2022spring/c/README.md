### 题目

有一个的木棒摆成的长方形网格，各个木棒，各个端点按照如下左图所示的方式编号。
![middle_img_v2_de308c49-ceb8-4946-9f24-953f94cbd74g.png](https://pic.leetcode-cn.com/1637557264-hILEsK-middle_img_v2_de308c49-ceb8-4946-9f24-953f94cbd74g.png)


在其中取走部分木棒，取走的木棒的编号记录于整数数组 `indices`。取完之后木棒网格一种可能的形式例如上右图所示。在 `0` 时刻，在网格中点燃某个端点。
- 一根木棒的一端被点燃，或者其他木棒燃烧到了相邻接的端点，则自此端点开始向另一端均匀地燃烧。
- 每根木棒从一端均匀地燃烧到另一端花费时间 `1`。

给定网格的高 `height`、宽 `width` 以及 `indices`。请问在 `0` 时刻点燃哪一个端点，可以使得在最短的时间内全体木棒都燃尽。返回一个升序排序的数组记录所有满足要求的端点编号。如果不存在这样的端点，返回空数组。


**注意：**
- 输入用例保证，`indices` 中的木棒编号都是合法的。


**示例 1：**
>输入：
> `height = 1, width = 2, indices = [3]`
>
>输出：`[0,1,2,3,4,5]`
>
>解释：如图所示，点燃所有端点，燃尽时间均为 `3`。
![p2_small.png](https://pic.leetcode-cn.com/1637557612-ePJqgt-p2_small.png)



**示例 2：**
>输入：
> `height = 2, width = 2, indices = [2,5,6,7,8,10,11]`
>
>输出：`[2]`
>
>解释：如下图所示。点燃端点 `[2]` 燃尽时间为 `2`，是最小燃尽时间。
![p3_small.png](https://pic.leetcode-cn.com/1637557626-nUTqoD-p3_small.png)



**示例 3：**
>输入：
> `height = 1, width = 1, indices = [0,3]`
>
>输出：`[]`
>
>解释：点燃任何一个端点，都无法燃尽所有木棒。因此返回空数组。


**提示：**
- `1 <= height, width <= 50`
- `0 <= indices.length <= (height+1)*width+height*(width+1)`
- `indices` 中的值两两不同
- `0 <= indices[i] < (height+1)*width+height*(width+1)`

### 思路

暴力枚举每个点开始时所需要的时间，进行一个`dfs`
最后判断所需要的时间和之前所有点的时间谁更小，如果相等加入答案

```go 
func lightSticks(height int, width int, indices []int) (ans []int) {

	// 火柴编码
	t, total, oneLine := math.MaxUint32, 2*height*width+height+width, 2*width+1
	point := func(x, y int) int {
		return (width+1)*x + y
	}
	top := func(x, y int) int {
		return (x-1)*oneLine + width + y
	}
	left := func(x, y int) int {
		return x*oneLine + y - 1
	}
	right := func(x, y int) int {
		return x*oneLine + y
	}
	down := func(x, y int) int {
		return x*oneLine + width + y
	}
	del := make([]bool, total)
	for _, v := range indices {
		del[v] = true
	}

	for i := 0; i <= height; i++ {
		for j := 0; j <= width; j++ {
			vis := append([]bool{}, del...)
			step, l := 0, total-len(indices)
			q := [][]int{{i, j}}

			for ; len(q) > 0; step++ {
				tmp := q
				q = nil

				for _, v := range tmp {
					x, y := v[0], v[1]
					if id := top(x, y); x > 0 && !vis[id] {
						vis[id] = true
						l--
						q = append(q, []int{x - 1, y})
					}
					if id := down(x, y); x < height && !vis[id] {
						vis[id] = true
						l--
						q = append(q, []int{x + 1, y})
					}
					if id := left(x, y); y > 0 && !vis[id] {
						vis[id] = true
						l--
						q = append(q, []int{x, y - 1})
					}
					if id := right(x, y); y < width && !vis[id] {
						vis[id] = true
						l--
						q = append(q, []int{x, y + 1})
					}
				}
			}
			if l > 0 {
				continue
			}

			if step < t {
				t = step
				ans = []int{point(i, j)}
			} else if step == t {
				ans = append(ans, point(i, j))
			}
		}
	}

	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2*m^2)$。$n$ 为 $height$ ，$m$ 为 $weight$ 
- 空间复杂度：$\mathcal{O}(n^2*m^2)$，
