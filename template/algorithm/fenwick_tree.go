package algorithm

// 树状数组（Fenwick Tree），二叉索引树（Binary Index Tree, BIT）
// https://en.wikipedia.org/wiki/Fenwick_tree
// 树状数组 tree 的基本用途是维护序列 a 的前缀和（tree 和 a 的下标都从 1 开始）
// tree[i] = a[i-lowbit(i)+1] + ... + a[i]
// 看图 https://oi-wiki.org/ds/fenwick/
// 更新 a[i] 的时候，会首先更新最下面的包含 a[i] 的 tree[i]，然后逐渐往上，更新包含这个元素的 tree[i]
// 这些节点（下标）都在 i 后面，所以更新的时候是从小往大算
// 计算某个前缀的时候，需要拆分区间，先把最右边的 arr[i-lowbit(i)+1] + ... + arr[i] 算出来，然后再去掉 i 最低位，算下一个区间
// 所以计算前缀是从大往小算
// 这里从小往大和从大往小说的是 i 的变化

// 树状数组,
// 模版题：https://loj.ac/p/131  solution: https://loj.ac/s/1589486
// 模版题：https://loj.ac/p/130  solution: https://loj.ac/s/1589521
// 结构体写法
// 单点查询时：a为原始数组，tree为更新之后的diff值(或者a使用前缀和进行优化，tree为diff值)、
type fenwick struct {
	a    []int64 // 原始数据，单点查询+区间更新传入
	tree []int64 // 树状数组
	diff []int64 // 差分数组, 用于区间加、区间求和
}

func newFenwickTree(n int64, a []int64) fenwick {
	return fenwick{a, make([]int64, n+1), make([]int64, n+1)}
}

// 位置 i 增加 val
// 1<=i<=n
func (f fenwick) add(i, val int64) {
	val1 := i * val
	for ; i < int64(len(f.tree)); i += i & -i {
		f.tree[i] += val
		f.diff[i] += val1
	}
}

// 求前缀和 [0,i]
// 0<=i<=n
func (f fenwick) sum(i int64) (res int64) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

func (f fenwick) sumDiff(i int64) (res int64) {
	for ; i > 0; i &= i - 1 {
		res += f.diff[i]
	}
	return
}

// 求区间和 [l,r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int64) int64 {
	return f.sum(r) - f.sum(l-1)
}

func (f fenwick) queryDiff(l, r int64) int64 {
	return (r+1)*f.sum(r) - l*f.sum(l-1) - (f.sumDiff(r) - f.sumDiff(l-1))
}

// 差分树状数组，可用于区间更新+单点查询 queryOne(i) = a[i] + sum(i) // a 从 1 开始
// r+1 即使超过 n 也没关系，因为不会用到
// i >= 1
func (f fenwick) queryOne(i int64) int64 { return f.a[i] + f.sum(i) }

// [l,r]
func (f fenwick) addRange(l, r, val int64) { f.add(l, val); f.add(r+1, -val) }

// fenwick tree with mod
// const mod = 1e9 + 7
//
//type fenwick struct {
//	a    []int64 // 原始数据，单点查询+区间更新传入
//	tree []int64 // 树状数组
//	diff []int64 // 差分数组, 用于区间加、区间求和
//}
//
//func newFenwickTree(n int64, a []int64) fenwick {
//	return fenwick{a, make([]int64, n+1), make([]int64, n+1)}
//}
//
//// 位置 i 增加 val
//// 1<=i<=n
//func (f fenwick) add(i, val int64) {
//	val1 := i * val
//	for ; i < int64(len(f.tree)); i += i & -i {
//		f.tree[i] += val
//		f.tree[i] %= mod
//		f.diff[i] += val1
//		f.diff[i] %= mod
//	}
//}
//
//// 求前缀和 [0,i]
//// 0<=i<=n
//func (f fenwick) sum(i int64) (res int64) {
//	for ; i > 0; i &= i - 1 {
//		res += f.tree[i]
//		res %= mod
//	}
//	return
//}
//
//func (f fenwick) sumDiff(i int64) (res int64) {
//	for ; i > 0; i &= i - 1 {
//		res += f.diff[i]
//		res %= mod
//	}
//	return
//}
//
//// 求区间和 [l,r]
//// 1<=l<=r<=n
//func (f fenwick) query(l, r int64) int64 {
//	ans := (f.sum(r) - f.sum(l-1) + mod) % mod
//	for ans < 0 {
//		ans += mod
//	}
//	return ans
//}
//
//func (f fenwick) queryDiff(l, r int64) int64 {
//	ans := (r+1)*f.sum(r)%mod - l*f.sum(l-1)%mod - (f.sumDiff(r)%mod - f.sumDiff(l-1)%mod) + mod
//	for ans < 0 {
//		ans += mod
//	}
//	return ans
//}
//
//// 差分树状数组，可用于区间更新+单点查询 queryOne(i) = a[i] + sum(i) // a 从 1 开始
//// r+1 即使超过 n 也没关系，因为不会用到
//// i >= 1
//func (f fenwick) queryOne(i int64) int64 { return f.a[i] + f.sum(i) }
//
//// [l,r]
//func (f fenwick) addRange(l, r, val int64) { f.add(l, val); f.add(r+1, -val) }
