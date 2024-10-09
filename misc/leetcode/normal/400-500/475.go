package _00_500

import (
	"math"
	"math/bits"
	"strconv"
)

// 483
func smallestGoodBase(n string) string {
	nVal, _ := strconv.Atoi(n)
	mMax := bits.Len(uint(nVal)) - 1
	for m := mMax; m > 1; m-- {
		k := int(math.Pow(float64(nVal), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == nVal {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(nVal - 1)
}

// 496
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))

	for i, v := range nums1 {
		j := 0
		for nums2[j] != v {
			j++
		}
		for j < len(nums2) {
			if nums2[j] > v {
				ans[i] = nums2[j]
			}
		}

		if j == len(nums2) {
			ans[i] = -1
		}
	}

	return ans
}
