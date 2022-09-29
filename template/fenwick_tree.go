package template

// 树状数组,
// 模版题：https://loj.ac/p/131  solution: https://loj.ac/s/1589486
// 模版题：https://loj.ac/p/130  solution: https://loj.ac/s/1589521
// 结构体写法
// 单点查询时：a为原始数组，tree为更新之后的diff值(或者a使用前缀和进行优化，tree为diff值)、
//
type fenwick struct {
	a, tree []int64
}

func newFenwickTree(n int, a []int64) fenwick {
	return fenwick{a, make([]int64, n+1)}
}

// 位置 i 增加 val
// 1<=i<=n
func (f fenwick) add(i int, val int64) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i] += val
	}
}

// 求前缀和 [0,i]
// 0<=i<=n
func (f fenwick) sum(i int) (res int64) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

// 求区间和 [l,r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int) int64 {
	return f.sum(r) - f.sum(l-1)
}

// i >= 1
func (f fenwick) queryOne(i int) int64 { return f.a[i] + f.sum(i) }

// [l,r]
func (f fenwick) addRange(l, r int, val int64) { f.add(l, val); f.add(r+1, -val) }
