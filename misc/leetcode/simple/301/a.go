package simple_301

import (
	"sort"
	"strings"
)

// 1
func fillCups(amount []int) int {
	sort.Ints(amount)
	ans := amount[2]
	if amount[0]+amount[1] <= ans {
		return ans
	}
	ans += (amount[0] + amount[1] - ans + 1) / 2

	return ans
}

// 2
type SmallestInfiniteSet struct {
	h map[int]bool
}

func Constructor() SmallestInfiniteSet {
	h := make(map[int]bool)
	for i := 0; i <= 1000; i++ {
		h[i] = true
	}
	return SmallestInfiniteSet{h}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := 1; ; i++ {
		if this.h[i] {
			this.h[i] = false
			return i
		}
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	this.h[num] = true
}

// 3
func canChange(start, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
		return false
	}
	j := 0
	for i, c := range start {
		if c != '_' {
			for target[j] == '_' {
				j++
			}
			if i != j && c == 'L' == (i < j) {
				return false
			}
			j++
		}
	}
	return true
}

//4

const mod, mx int = 1e9 + 7, 1e4 + 20

var ks [mx][]int
var F, invF [mx]int

func init() {
	for i := 2; i < mx; i++ {
		x := i
		for p := 2; p*p <= x; p++ {
			k := 0
			for ; x%p == 0; x /= p {
				k++
			}
			if k > 0 {
				ks[i] = append(ks[i], k)
			}
		}
		if x > 1 {
			ks[i] = append(ks[i], 1)
		}
	}
	F[0] = 1
	for i := 1; i < mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx-1] = pow(F[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func idealArrays(n, maxValue int) (ans int) {
	for m := 1; m <= maxValue; m++ {
		mul := 1
		for _, k := range ks[m] {
			comb := F[n+k-1] * invF[k] % mod * invF[n-1] % mod
			mul = mul * comb % mod
		}
		ans = (ans + mul) % mod
	}
	return ans
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
