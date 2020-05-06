package three_sum_closest

import (
	"math"
	"sort"
)

// my solution
// Runtime 12ms, Memory 2.7MB
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	ret := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if target == sum {
					return target
				}
				// find the absolute value
				x, y := sum-target, ret-target
				if x < 0 {
					x *= -1
				}
				if y < 0 {
					y *= -1
				}
				if x <= y {
					ret = sum
				}
			}
		}
	}
	return ret
}

// best answer from community
// faster than 100%
func best_threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	var res int = nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			cur := nums[i] + nums[start] + nums[end]
			if cur == target {
				return target
			} else if cur < target {
				start++
			} else {
				end--
			}
			d1 := math.Abs(float64(cur - target))
			d2 := math.Abs(float64(res - target))
			if d1 < d2 {
				res = cur
			}
		}
	}
	return res
}
