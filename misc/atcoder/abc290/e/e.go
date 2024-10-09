package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// 错误的思路：区间 DP。复杂度太高，没有优化空间。
// 提示 1：一对数只需修改其中一个。

// 提示 2：正难则反，考虑有多少对数字无需修改。用所有数对个数，减去无需修改的数对个数，就是需要修改的数对个数。
// 所有数对个数 = sum(长为 i 的子数组个数 * floor(i/2))，其中【长为 i 的子数组个数】就是 n+1-i。
// 可以写个 for 循环计算，也可以用公式 O(1) 算。

// 提示 3：贡献法。对于 a[i]=a[j] 的这对数，考虑这对数能出现在多少个子数组中。
// 什么时候子数组的个数取决于 i，什么时候取决于 j？
// 答：设下标从 0 开始。如果 i+j<n，那么取决于 i，否则取决于 j。
//（想象成从 [i,j] 不断向外扩展，i 向左，j 向右）

// 提示 4：保存相同元素的下标列表 pos。

// 方法一：相向双指针。
// 设 p 为 pos 中的一个下标列表（下标从 0 开始）。
// 初始化 l=0，r=len(p)-1，循环直到 l >=r。
// 如果 p[l]+p[r]<n，那么有 p[l]+1 个子数组是包含 p[l] 和 p[r] 的。
// 此外，这意味着 p[l]+p[r-1] 也是 < n 的，那么也有 p[l]+1 个子数组是包含 p[l] 和 p[r-1] 的。
// 此外，这意味着 p[l]+p[r-2] 也是 < n 的，那么也有 p[l]+1 个子数组是包含 p[l] 和 p[r-2] 的。
// 依此类推，从 r'=r 到 r'=l+1，这样的 r' 一共有 r-l 个。
// 那么 p[l] 对答案的贡献为 (p[l]+1) * (r-l)。
// 如果 p[l]+p[r] >=n，那么同理可得 p[r] 对答案的贡献为 (n-p[r]) * (r-l)。
// 时间复杂度 O(n)。

// 方法二：二分查找。
// 如果你没有想到相向双指针，那么二分可能更适合你。
// 遍历 r。我们可以在 p 中二分找哪些 p[l] 满足 p[l]+p[r]<n，哪些 p[l] 满足 p[l]+p[r] >=n，分别统计。
// 时间复杂度 O(nlogn)。

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	// 所有数对个数 = sum(长为 i 的子数组个数 * floor(i/2))，其中【长为 i 的子数组个数】就是 n+1-i。
	// m := n / 2
	// ans := m * (m + 1) * (m*4 + n%2*6 - 1) / 6
	ans := 0
	for i := 1; i <= n; i++ {
		ans += (n + 1 - i) * (i / 2)
	}
	pos := make(map[int][]int)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		pos[v] = append(pos[v], i)
	}
	for _, p := range pos {
		l, r := 0, len(p)-1
		for l < r {
			if p[l]+p[r] < n {
				ans -= (p[l] + 1) * (r - l)
				l++
			} else {
				ans -= (n - p[r]) * (r - l)
				r--
			}
		}
	}
	Fprint(out, ans)
}

func runV2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	m := n / 2
	ans := m * (m + 1) * (m*4 + n%2*6 - 1) / 6
	pos := make(map[int][]int)
	sumP := make([][]int, n+1)
	for i := range sumP {
		sumP[i] = []int{0}
	}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		ps := pos[v]
		p := sort.SearchInts(ps, n-1-i)
		ans -= (len(ps) - p) * (n - i)
		ans -= sumP[v][p]
		pos[v] = append(pos[v], i)
		sumP[v] = append(sumP[v], sumP[v][len(sumP[v])-1]+i+1)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
