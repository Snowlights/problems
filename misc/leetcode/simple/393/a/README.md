#### 题目

<p>给你一个字符串 <code>s</code>，表示一个 12 小时制的时间格式，其中一些数字（可能没有）被 <code>"?"</code> 替换。</p>

<p>12 小时制时间格式为 <code>"HH:MM"</code> ，其中 <code>HH</code> 的取值范围为 <code>00</code> 至 <code>11</code>，<code>MM</code> 的取值范围为 <code>00</code> 至 <code>59</code>。最早的时间为 <code>00:00</code>，最晚的时间为 <code>11:59</code>。</p>

<p>你需要将 <code>s</code> 中的<strong> 所有</strong> <code>"?"</code> 字符替换为数字，使得结果字符串代表的时间是一个<strong> 有效 </strong>的 12 小时制时间，并且是可能的 <strong>最晚 </strong>时间。</p>

<p>返回结果字符串。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "1?:?4"</span></p>

<p><strong>输出：</strong> <span class="example-io">"11:54"</span></p>

<p><strong>解释：</strong> 通过替换 <code>"?"</code> 字符，可以得到的最晚12小时制时间是 <code>"11:54"</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "0?:5?"</span></p>

<p><strong>输出：</strong> <span class="example-io">"09:59"</span></p>

<p><strong>解释：</strong> 通过替换 <code>"?"</code> 字符，可以得到的最晚12小时制时间是 <code>"09:59"</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>s.length == 5</code></li>
	<li><code>s[2]</code> 是字符 <code>":"</code></li>
	<li>除 <code>s[2]</code> 外，其他字符都是数字或 <code>"?"</code></li>
	<li>输入保证在替换 <code>"?"</code> 字符后至少存在一个介于 <code>"00:00"</code> 和 <code>"11:59"</code> 之间的时间。</li>
</ul>

#### 思路

模拟

``` go
func findLatestTime(s string) string {
	t := []byte(s)
	if t[0] == '?' {
		if t[1] == '?' || t[1] <= '1' {
			t[0] = '1'
		} else {
			t[0] = '0'
		}
	}
	if t[1] == '?' {
		if t[0] == '1' {
			t[1] = '1'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [1446. 连续字符](https://leetcode.cn/problems/consecutive-characters/) 1165
- [1869. 哪种连续子字符串更长](https://leetcode.cn/problems/longer-contiguous-segments-of-ones-than-zeros/) 1205
- [1957. 删除字符使字符串变好](https://leetcode.cn/problems/delete-characters-to-make-fancy-string/) 1358
- [978. 最长湍流子数组](https://leetcode.cn/problems/longest-turbulent-subarray/) 1393
- [2110. 股票平滑下跌阶段的数目](https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/) 1408
- [228. 汇总区间](https://leetcode.cn/problems/summary-ranges/)
- [2760. 最长奇偶子数组](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/) 1420
- [1887. 使数组元素相等的减少操作次数](https://leetcode.cn/problems/reduction-operations-to-make-the-array-elements-equal/) 1428
- [845. 数组中的最长山脉](https://leetcode.cn/problems/longest-mountain-in-array/) 1437
- [2038. 如果相邻两个颜色均相同则删除当前颜色](https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/) 1468
- [1759. 统计同质子字符串的数目](https://leetcode.cn/problems/count-number-of-homogenous-substrings/) 1491
- [3011. 判断一个数组是否可以变为有序](https://leetcode.cn/problems/find-if-array-can-be-sorted/) 1497
- [1578. 使绳子变成彩色的最短时间](https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/) 1574
- [1839. 所有元音按顺序排布的最长子字符串](https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order/) 1580
- [2765. 最长交替子序列](https://leetcode.cn/problems/longest-alternating-subarray/) 1581
- [467. 环绕字符串中唯一的子字符串](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/) ~1700
- [2948. 交换得到字典序最小的数组](https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements/) 2047
- [2393. 严格递增的子数组个数](https://leetcode.cn/problems/count-strictly-increasing-subarrays/)（会员题）
- [2436. 使子数组最大公约数大于一的最小分割数](https://leetcode.cn/problems/minimum-split-into-subarrays-with-gcd-greater-than-one/)（会员题）
- [2495. 乘积为偶数的子数组数](https://leetcode.cn/problems/number-of-subarrays-having-even-product/)（会员题）
- [3063. 链表频率](https://leetcode.cn/problems/linked-list-frequency/)（会员题）


## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)