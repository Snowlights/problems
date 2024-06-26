#### 题目

<p>给你一个下标从 <strong>0</strong> 开始、大小为 <code>m x n</code> 的网格 <code>image</code> ，表示一个灰度图像，其中 <code>image[i][j]</code> 表示在范围 <code>[0..255]</code> 内的某个像素强度。另给你一个<strong> 非负 </strong>整数 <code>threshold</code> 。</p>

<p>如果 <code>image[a][b]</code> 和 <code>image[c][d]</code> 满足 <code>|a - c| + |b - d| == 1</code> ，则称这两个像素是<strong> 相邻像素</strong> 。</p>

<p><strong>区域 </strong>是一个 <code>3 x 3</code> 的子网格，且满足区域中任意两个 <strong>相邻</strong> 像素之间，像素强度的<strong> 绝对差 </strong><strong> 小于或等于 </strong><code>threshold</code> 。</p>

<p><strong>区域</strong> 内的所有像素都认为属于该区域，而一个像素 <strong>可以 </strong>属于 <strong>多个</strong> 区域。</p>

<p>你需要计算一个下标从 <strong>0</strong> 开始、大小为 <code>m x n</code> 的网格 <code>result</code> ，其中 <code>result[i][j]</code> 是 <code>image[i][j]</code> 所属区域的 <strong>平均 </strong>强度，<strong>向下取整 </strong>到最接近的整数。如果 <code>image[i][j]</code> 属于多个区域，<code>result[i][j]</code> 是这些区域的<strong> </strong><strong>“取整后的平均强度”</strong> 的<strong> 平均值</strong>，也 <strong>向下取整 </strong>到最接近的整数。如果 <code>image[i][j]</code> 不属于任何区域，则 <code>result[i][j]</code><strong> 等于 </strong><code>image[i][j]</code> 。</p>

<p>返回网格 <code>result</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/21/example0corrected.png" style="width: 832px; height: 275px;" />
<pre>
<strong>输入：</strong>image = [[5,6,7,10],[8,9,10,10],[11,12,13,10]], threshold = 3
<strong>输出：</strong>[[9,9,9,9],[9,9,9,9],[9,9,9,9]]
<strong>解释：</strong>图像中存在两个区域，如图片中的阴影区域所示。第一个区域的平均强度为 9 ，而第二个区域的平均强度为 9.67 ，向下取整为 9 。两个区域的平均强度为 (9 + 9) / 2 = 9 。由于所有像素都属于区域 1 、区域 2 或两者，因此 result 中每个像素的强度都为 9 。
注意，在计算多个区域的平均值时使用了向下取整的值，因此使用区域 2 的平均强度 9 来进行计算，而不是 9.67 。
</pre>

<p><strong class="example">示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/21/example1corrected.png" style="width: 805px; height: 377px;" />
<pre>
<strong>输入：</strong>image = [[10,20,30],[15,25,35],[20,30,40],[25,35,45]], threshold = 12
<strong>输出：</strong>[[25,25,25],[27,27,27],[27,27,27],[30,30,30]]
<strong>解释：</strong>图像中存在两个区域，如图片中的阴影区域所示。第一个区域的平均强度为 25 ，而第二个区域的平均强度为 30 。两个区域的平均强度为 (25 + 30) / 2 = 27.5 ，向下取整为 27 。图像中第 0 行的所有像素属于区域 1 ，因此 result 中第 0 行的所有像素为 25 。同理，result 中第 3 行的所有像素为 30 。图像中第 1 行和第 2 行的像素属于区域 1 和区域 2 ，因此它们在 result 中的值为 27 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>image = [[5,6,7],[8,9,10],[11,12,13]], threshold = 1
<strong>输出：</strong>[[5,6,7],[8,9,10],[11,12,13]]
<strong>解释：</strong>图像中不存在任何区域，因此对于所有像素，result[i][j] == image[i][j] 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 <= n, m <= 500</code></li>
	<li><code>0 <= image[i][j] <= 255</code></li>
	<li><code>0 <= threshold <= 255</code></li>
</ul>

#### 思路

1. 遍历所有 $3 \times 3$ 的子网格。
2. 遍历网格内的所有左右相邻格子和上下相邻格子，如果存在差值超过 $\textit{threshold}$ 的情况，则枚举下一个子网格。
3. 如果合法，计算子网格的平均值 $\textit{avg}$，等于子网格的元素和除以 $9$ 下取整。
4. 更新子网格内的 $\textit{result}[i][j]$，由于需要计算平均值，我们先把 $\textit{avg}$ 加到 $\textit{result}[i][j]$ 中，同时用一个 $\textit{cnt}$ 矩阵统计 $(i,j)$ 在多少个合法子网格内。
5. 最后返回答案。如果 $\textit{cnt}[i][j]=0$ 则 $\textit{result}[i][j] = \textit{image}[i][j]$，否则 $\textit{result}[i][j] = \left\lfloor\dfrac{\textit{result}[i][j]}{\textit{cnt}[i][j]}\right\rfloor$。

```go [sol]
func resultGrid(a [][]int, threshold int) [][]int {
	m, n := len(a), len(a[0])
	result := make([][]int, m)
	cnt := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		cnt[i] = make([]int, n)
	}
	for i := 2; i < m; i++ {
	next:
		for j := 2; j < n; j++ {
			// 检查左右相邻格子
			for _, row := range a[i-2 : i+1] {
				if abs(row[j-2]-row[j-1]) > threshold || abs(row[j-1]-row[j]) > threshold {
					continue next // 不合法，下一个
				}
			}

			// 检查上下相邻格子
			for y := j - 2; y <= j; y++ {
				if abs(a[i-2][y]-a[i-1][y]) > threshold || abs(a[i-1][y]-a[i][y]) > threshold {
					continue next // 不合法，下一个
				}
			}

			// 合法，计算 3x3 子网格的平均值
			avg := 0
			for x := i - 2; x <= i; x++ {
				for y := j - 2; y <= j; y++ {
					avg += a[x][y]
				}
			}
			avg /= 9

			// 更新 3x3 子网格内的 result
			for x := i - 2; x <= i; x++ {
				for y := j - 2; y <= j; y++ {
					result[x][y] += avg // 先累加，最后再求平均值
					cnt[x][y]++
				}
			}
		}
	}

	for i, row := range cnt {
		for j, c := range row {
			if c == 0 { // (i,j) 不属于任何子网格
				result[i][j] = a[i][j]
			} else {
				result[i][j] /= c // 求平均值
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(Amn)$，其中 m 和 n 分别为 $\textit{a}$ 的行数和列数，A=9 表示子网格大小。
- 空间复杂度：$O(mn)$。
