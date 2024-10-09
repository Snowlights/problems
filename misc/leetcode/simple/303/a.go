package simple_303

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
	"math/bits"
)

// 1
func repeatedCharacter(s string) byte {

	h := make(map[byte]bool)
	for _, v := range s {
		if h[byte(v)] {
			return byte(v)
		}
		h[byte(v)] = true
	}
	return byte(' ')
}

// 2
func equalPairs(grid [][]int) int {
	n := len(grid)
	ans := 0
	for i := 0; i < n; i++ {
		row := grid[i]
		for j := 0; j < n; j++ {
			tmp := 0
			for tmp < n && row[tmp] == grid[tmp][j] {
				tmp++
			}
			if tmp == n {
				ans++
			}
		}

	}
	return ans
}

// 3
type pair struct {
	r int
	f string
}

type FoodRatings struct {
	fr map[string]int
	fc map[string]string
	ct map[string]*redblacktree.Tree
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	fr := map[string]int{}
	fc := map[string]string{}
	ct := map[string]*redblacktree.Tree{}
	for i, f := range foods {
		r, c := ratings[i], cuisines[i]
		fr[f] = r
		fc[f] = c
		if ct[c] == nil {
			ct[c] = redblacktree.NewWith(func(x, y interface{}) int {
				a, b := x.(pair), y.(pair)
				if a.r != b.r {
					return utils.IntComparator(b.r, a.r)
				}
				return utils.StringComparator(a.f, b.f)
			})
		}
		ct[c].Put(pair{r, f}, nil)
	}
	return FoodRatings{fr, fc, ct}
}

func (r FoodRatings) ChangeRating(f string, newRating int) {
	t := r.ct[r.fc[f]]
	t.Remove(pair{r.fr[f], f})     // 移除旧数据
	t.Put(pair{newRating, f}, nil) // 添加新数据
	r.fr[f] = newRating
}

func (r FoodRatings) HighestRated(cuisine string) string {
	return r.ct[cuisine].Left().Key.(pair).f
}

// 4
func countExcellentPairs(nums []int, k int) (ans int64) {
	vis := map[int]bool{}
	cnt := map[int]int{}
	for _, x := range nums {
		if !vis[x] {
			vis[x] = true
			cnt[bits.OnesCount(uint(x))]++
		}
	}
	for cx, ccx := range cnt {
		for cy, ccy := range cnt {
			if cx+cy >= k { // (x,y) 是优质数对
				ans += int64(ccx) * int64(ccy)
			}
		}
	}
	return
}
