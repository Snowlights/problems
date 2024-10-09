package double_88

import "sort"

// 1
func equalFrequency(word string) bool {
next:
	for i := range word {
		cnt := map[rune]int{}
		for _, c := range word[:i] + word[i+1:] {
			cnt[c]++
		}
		same := 0
		for _, c := range cnt {
			if same == 0 {
				same = c
			} else if c != same {
				continue next
			}
		}
		return true
	}
	return false
}

// 2
type LUPrefix struct {
	x int
	h map[int]bool
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		x: 1,
		h: make(map[int]bool),
	}
}

func (this *LUPrefix) Upload(video int) {
	this.h[video] = true
	for this.h[this.x] {
		this.x++
	}
}

func (this *LUPrefix) Longest() int {
	return this.x - 1
}

// 3
func xorAllNums(nums1, nums2 []int) (ans int) {
	if len(nums2)%2 > 0 {
		for _, x := range nums1 {
			ans ^= x
		}
	}
	if len(nums1)%2 > 0 {
		for _, x := range nums2 {
			ans ^= x
		}
	}
	return
}

// 4
func numberOfPairs(a, nums2 []int, diff int) (ans int64) {
	for i, x := range nums2 {
		a[i] -= x
	}
	b := append(sort.IntSlice{}, a...)
	b.Sort() // 配合下面的二分，离散化

	t := make(BIT, len(a)+1)
	for _, x := range a {
		ans += int64(t.query(b.Search(x + diff + 1)))
		t.add(b.Search(x) + 1)
	}
	return
}

type BIT []int

func (t BIT) add(x int) {
	for x < len(t) {
		t[x]++
		x += x & -x
	}
}

func (t BIT) query(x int) (res int) {
	for x > 0 {
		res += t[x]
		x &= x - 1
	}
	return
}

// func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
//
//	a, b := make([]int, 0), make([]int, 0)
//	ans := int64(0)
//	for i, v := range nums1 {
//		a = append(a, v-nums2[i])
//	}
//	b = append(b, a...)
//	sort.Ints(b)
//
//	t := newFenwickTree(int64(len(nums1))+1, nil)
//	for _, v := range a {
//		ans += int64(t.query(1, int64(sort.SearchInts(b, v+diff+1))))
//		t.add(int64(sort.SearchInts(b, v)) + 1, 1)
//	}
//
//	return ans
//}
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
//		f.diff[i] += val1
//	}
//}
//
//// 求前缀和 [0,i]
//// 0<=i<=n
//func (f fenwick) sum(i int64) (res int64) {
//	for ; i > 0; i &= i - 1 {
//		res += f.tree[i]
//	}
//	return
//}
//
//func (f fenwick) sumDiff(i int64) (res int64) {
//	for ; i > 0; i &= i - 1 {
//		res += f.diff[i]
//	}
//	return
//}
//
//// 求区间和 [l,r]
//// 1<=l<=r<=n
//func (f fenwick) query(l, r int64) int64 {
//	return f.sum(r) - f.sum(l-1)
//}
//
//func (f fenwick) queryDiff(l, r int64) int64 {
//	return (r+1)*f.sum(r) - l*f.sum(l-1) - (f.sumDiff(r) - f.sumDiff(l-1))
//}
//
//// 差分树状数组，可用于区间更新+单点查询 queryOne(i) = a[i] + sum(i) // a 从 1 开始
//// r+1 即使超过 n 也没关系，因为不会用到
//// i >= 1
//func (f fenwick) queryOne(i int64) int64 { return f.a[i] + f.sum(i) }
//
//// [l,r]
//func (f fenwick) addRange(l, r, val int64) { f.add(l, val); f.add(r+1, -val) }
